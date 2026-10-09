use serde::{Deserialize, Serialize};
use sha2::{Digest, Sha256};
use std::collections::{BTreeMap, BTreeSet};
use std::fs::{self, File};
use std::io::{self, BufRead, BufReader, Read, Write};
use std::path::{Path, PathBuf};
use std::time::{Duration, Instant, SystemTime, UNIX_EPOCH};

#[derive(Clone, Debug, Serialize, Deserialize, PartialEq, Eq, PartialOrd, Ord)]
pub struct Origin {
    pub repository: String,
    pub branch: String,
    pub commit_sha: String,
    pub path: String,
    pub blob_sha: String,
    pub size_bytes: Option<u64>,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct FileRecord {
    pub origin: Origin,
    pub indexed_sha256: Option<String>,
    pub indexed_bytes: Option<u64>,
    pub status: String,
    pub note: Option<String>,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct IndexSnapshot {
    pub schema_version: u32,
    pub generated_unix_ms: u128,
    pub source_manifests: Vec<String>,
    pub records: Vec<FileRecord>,
    pub unique_blob_count: usize,
    pub duplicate_blob_count: usize,
    pub incomplete_count: usize,
    pub status: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct LookupResult {
    pub status: String,
    pub query: String,
    pub matches: Vec<Origin>,
    pub cache_age_ms: u128,
    pub freshness: String,
}

#[derive(Default)]
pub struct MemoryIndex {
    by_path: BTreeMap<String, Vec<Origin>>,
    by_blob: BTreeMap<String, Vec<Origin>>,
    built_unix_ms: u128,
}

impl MemoryIndex {
    pub fn from_snapshot(snapshot: &IndexSnapshot) -> Self {
        let mut idx = Self { built_unix_ms: snapshot.generated_unix_ms, ..Self::default() };
        for record in &snapshot.records {
            let key = format!("{}\0{}\0{}", record.origin.repository, record.origin.branch, record.origin.path);
            idx.by_path.entry(key).or_default().push(record.origin.clone());
            idx.by_blob.entry(record.origin.blob_sha.clone()).or_default().push(record.origin.clone());
        }
        idx
    }

    pub fn lookup_path(&self, repository: &str, branch: &str, path: &str) -> LookupResult {
        let key = format!("{repository}\0{branch}\0{path}");
        let matches = self.by_path.get(&key).cloned().unwrap_or_default();
        self.result(format!("{repository}:{branch}:{path}"), matches)
    }

    pub fn lookup_blob(&self, blob_sha: &str) -> LookupResult {
        let matches = self.by_blob.get(blob_sha).cloned().unwrap_or_default();
        self.result(blob_sha.to_owned(), matches)
    }

    fn result(&self, query: String, matches: Vec<Origin>) -> LookupResult {
        let now = unix_ms();
        LookupResult {
            status: if matches.is_empty() { "NOT_FOUND".into() } else { "INDEX_HIT".into() },
            query,
            matches,
            cache_age_ms: now.saturating_sub(self.built_unix_ms),
            freshness: "UNVERIFIED_REQUIRES_HEAD_RECONCILIATION".into(),
        }
    }
}

pub fn build_index(manifest_paths: &[PathBuf], source_root: Option<&Path>) -> io::Result<IndexSnapshot> {
    let mut records = Vec::new();
    let mut seen = BTreeSet::new();
    let mut unique_blobs = BTreeSet::new();
    let mut duplicate_count = 0usize;
    let mut incomplete = 0usize;

    for manifest_path in manifest_paths {
        let file = File::open(manifest_path)?;
        for (line_no, line) in BufReader::new(file).lines().enumerate() {
            let line = line?;
            if line.trim().is_empty() || line.starts_with('#') { continue; }
            let parsed: serde_json::Value = match serde_json::from_str(&line) {
                Ok(v) => v,
                Err(e) => {
                    incomplete += 1;
                    records.push(FileRecord {
                        origin: Origin { repository: String::new(), branch: String::new(), commit_sha: String::new(), path: format!("{}:line:{}", manifest_path.display(), line_no + 1), blob_sha: String::new(), size_bytes: None },
                        indexed_sha256: None, indexed_bytes: None, status: "INVALID_MANIFEST_ROW".into(), note: Some(e.to_string()),
                    });
                    continue;
                }
            };
            let Some(path) = parsed.get("path").and_then(|v| v.as_str()) else { continue };
            let Some(blob_sha) = parsed.get("object_sha").or_else(|| parsed.get("sha")).and_then(|v| v.as_str()) else { continue };
            let repo = parsed.get("repository").and_then(|v| v.as_str()).unwrap_or("").to_owned();
            let branch = parsed.get("branch").and_then(|v| v.as_str()).unwrap_or("").to_owned();
            let commit = parsed.get("commit_sha").and_then(|v| v.as_str()).unwrap_or("").to_owned();
            let size = parsed.get("size_bytes").or_else(|| parsed.get("size")).and_then(|v| v.as_u64());
            let typ = parsed.get("type").and_then(|v| v.as_str()).unwrap_or("blob");
            if typ != "blob" { continue; }
            let origin = Origin { repository: repo, branch, commit_sha: commit, path: path.to_owned(), blob_sha: blob_sha.to_owned(), size_bytes: size };
            if !seen.insert(origin.clone()) { continue; }
            if !unique_blobs.insert(blob_sha.to_owned()) { duplicate_count += 1; }

            let mut record = FileRecord { origin, indexed_sha256: None, indexed_bytes: None, status: "INVENTORIED_NOT_CONTENT_VERIFIED".into(), note: None };
            if let Some(root) = source_root {
                let source_file = root.join(&record.origin.repository.replace('/', "__")).join(&record.origin.branch).join(path);
                match hash_file(&source_file) {
                    Ok((digest, bytes)) => {
                        record.indexed_sha256 = Some(digest);
                        record.indexed_bytes = Some(bytes);
                        record.status = "CONTENT_HASHED_LOCAL_SNAPSHOT".into();
                    }
                    Err(e) => {
                        record.status = "SOURCE_BYTES_UNAVAILABLE".into();
                        record.note = Some(e.to_string());
                        incomplete += 1;
                    }
                }
            } else {
                incomplete += 1;
            }
            records.push(record);
        }
    }

    Ok(IndexSnapshot {
        schema_version: 1,
        generated_unix_ms: unix_ms(),
        source_manifests: manifest_paths.iter().map(|p| p.display().to_string()).collect(),
        unique_blob_count: unique_blobs.len(),
        duplicate_blob_count: duplicate_count,
        incomplete_count: incomplete,
        status: if incomplete == 0 { "INDEXED_WITH_LOCAL_CONTENT_EVIDENCE".into() } else { "PARTIAL_INVENTORY_OR_CONTENT_UNVERIFIED".into() },
        records,
    })
}

pub fn write_snapshot(snapshot: &IndexSnapshot, output: &Path) -> io::Result<()> {
    let bytes = serde_json::to_vec_pretty(snapshot).map_err(io::Error::other)?;
    let temp = output.with_extension("tmp");
    if let Some(parent) = output.parent() { fs::create_dir_all(parent)?; }
    {
        let mut f = File::create(&temp)?;
        f.write_all(&bytes)?;
        f.sync_all()?;
    }
    fs::rename(temp, output)?;
    Ok(())
}

fn hash_file(path: &Path) -> io::Result<(String, u64)> {
    let mut file = File::open(path)?;
    let mut hash = Sha256::new();
    let mut bytes = 0u64;
    let mut buf = [0u8; 64 * 1024];
    loop {
        let n = file.read(&mut buf)?;
        if n == 0 { break; }
        hash.update(&buf[..n]);
        bytes += n as u64;
    }
    Ok((format!("{:x}", hash.finalize()), bytes))
}

fn unix_ms() -> u128 {
    SystemTime::now().duration_since(UNIX_EPOCH).unwrap_or(Duration::ZERO).as_millis()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn origin(repo: &str, branch: &str, path: &str, sha: &str) -> Origin {
        Origin { repository: repo.into(), branch: branch.into(), commit_sha: "commit1".into(), path: path.into(), blob_sha: sha.into(), size_bytes: Some(12) }
    }

    #[test]
    fn in_memory_lookup_preserves_origin() {
        let snapshot = IndexSnapshot {
            schema_version: 1, generated_unix_ms: unix_ms(), source_manifests: vec![],
            records: vec![FileRecord { origin: origin("owner/repo", "main", "src/lib.rs", "blob1"), indexed_sha256: None, indexed_bytes: None, status: "INVENTORIED".into(), note: None }],
            unique_blob_count: 1, duplicate_blob_count: 0, incomplete_count: 0, status: "TEST".into()
        };
        let index = MemoryIndex::from_snapshot(&snapshot);
        let found = index.lookup_path("owner/repo", "main", "src/lib.rs");
        assert_eq!(found.status, "INDEX_HIT");
        assert_eq!(found.matches.len(), 1);
        assert_eq!(found.matches[0].blob_sha, "blob1");
    }

    #[test]
    fn lookup_missing_path_is_explicit() {
        let snapshot = IndexSnapshot { schema_version: 1, generated_unix_ms: unix_ms(), source_manifests: vec![], records: vec![], unique_blob_count: 0, duplicate_blob_count: 0, incomplete_count: 0, status: "TEST".into() };
        assert_eq!(MemoryIndex::from_snapshot(&snapshot).lookup_blob("missing").status, "NOT_FOUND");
    }
}
