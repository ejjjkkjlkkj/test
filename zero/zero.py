"""ZERO: laboratoire expérimental sans modèle d'IA externe.

Aucune dépendance externe. Le système manipule des traces symboliques
et expérimente des transformations plutôt que de prédire une réponse.
"""

from __future__ import annotations

import sys
import json
import time
import random
from pathlib import Path


SEED = 20261006


class Trace:
    def __init__(self, parts: list[str], history: list[str] | None = None):
        self.parts = parts
        self.history = history or []

    def text(self) -> str:
        return " | ".join(self.parts)

    def copy(self) -> "Trace":
        return Trace(self.parts[:], self.history[:])


def perceive(problem: str) -> Trace:
    words = [w.strip(".,;:!?()[]{}\"'").lower()
             for w in problem.split() if w.strip()]
    return Trace(words, ["perception"])


def split_world(t: Trace) -> list[Trace]:
    out = [t.copy()]
    if len(t.parts) > 2:
        middle = len(t.parts) // 2
        out.append(Trace(
            t.parts[:middle] + ["<gap>"] + t.parts[middle:],
            t.history + ["insert-gap"],
        ))
    unique = []
    for p in t.parts:
        if p not in unique:
            unique.append(p)
    if unique != t.parts:
        out.append(Trace(unique, t.history + ["remove-repetition"]))
    return out


def recombine(a: Trace, b: Trace) -> Trace:
    left = a.parts[:]
    right = b.parts[:]
    cut_a = len(left) // 2
    cut_b = len(right) // 2
    parts = left[:cut_a] + right[cut_b:]
    return Trace(parts, a.history + ["cross-recombine"])


def perturb(t: Trace, rng: random.Random) -> Trace:
    p = t.parts[:]
    if len(p) >= 2:
        i = rng.randrange(len(p))
        j = rng.randrange(len(p))
        p[i], p[j] = p[j], p[i]
    elif p:
        p.append("<unknown>")
    return Trace(p, t.history + ["perturb"])


def consequence(t: Trace, target: str) -> tuple[int, str]:
    text = t.text()
    target_words = set(target.lower().split())
    own_words = set(t.parts)
    overlap = len(target_words & own_words)

    # Ce n'est pas un score de vérité. C'est un signal expérimental :
    # une transformation est intéressante si elle conserve du sens
    # tout en produisant une structure différente.
    novelty = len(set(t.history))
    coherence = sum(1 for a, b in zip(t.parts, t.parts[1:]) if a != b)
    signal = overlap * 3 + min(novelty, 5) + min(coherence, 5)
    return signal, f"overlap={overlap}; novelty={novelty}; structure={coherence}"


def experiment(problem: str, rounds: int = 12) -> list[dict]:
    rng = random.Random(SEED)
    population = [perceive(problem)]
    observations: list[dict] = []

    for step in range(rounds):
        candidates: list[Trace] = []
        for item in population:
            candidates.extend(split_world(item))
            candidates.append(perturb(item, rng))

        if len(population) >= 2:
            candidates.append(recombine(population[0], population[-1]))

        ranked = []
        for candidate in candidates:
            signal, detail = consequence(candidate, problem)
            ranked.append((signal, candidate, detail))

        ranked.sort(key=lambda x: (x[0], len(x[1].history)), reverse=True)
        population = [x[1] for x in ranked[:4]]

        best_signal, best, detail = ranked[0]
        observations.append({
            "step": step,
            "signal": best_signal,
            "trace": best.text(),
            "history": best.history,
            "detail": detail,
        })

    return observations


def main() -> None:
    problem = " ".join(sys.argv[1:]).strip()
    if not problem:
        problem = "comprendre quelque chose que la procédure ne décrit pas"

    result = experiment(problem)

    for row in result:
        print(
            f"[{row['step']:02d}] signal={row['signal']:02d} "
            f"{row['trace']}  ({row['detail']})"
        )

    Path("runs").mkdir(exist_ok=True)
    stamp = time.strftime("%Y%m%d-%H%M%S")
    path = Path("runs") / f"{stamp}.json"
    path.write_text(
        json.dumps(
            {"problem": problem, "seed": SEED, "observations": result},
            ensure_ascii=False,
            indent=2,
        ),
        encoding="utf-8",
    )
    print(f"\nTRACE: {path}")


if __name__ == "__main__":
    main()
