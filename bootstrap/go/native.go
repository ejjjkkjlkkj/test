package zero

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

var nativeMagic = [4]byte{'Z', 'E', 'R', '0'}
const nativeVersion byte = 1

func putNativeString(b *bytes.Buffer, s string) error {
	if len(s) > int(^uint32(0)) {
		return errors.New("native string too large")
	}
	if err := binary.Write(b, binary.BigEndian, uint32(len(s))); err != nil {
		return err
	}
	_, err := b.WriteString(s)
	return err
}

func getNativeString(data []byte, off *int) (string, error) {
	if *off+4 > len(data) {
		return "", errors.New("native truncated length")
	}
	n := binary.BigEndian.Uint32(data[*off : *off+4])
	*off += 4
	end := uint64(*off) + uint64(n)
	if end > uint64(len(data)) {
		return "", errors.New("native truncated field")
	}
	s := string(data[*off:end])
	*off = int(end)
	return s, nil
}

func EncodeNative(r Record) ([]byte, error) {
	if err := Validate(r); err != nil {
		return nil, err
	}
	var b bytes.Buffer
	b.Write(nativeMagic[:])
	b.WriteByte(nativeVersion)
	for _, s := range []string{r.Version, r.Type, r.Identity} {
		if err := putNativeString(&b, s); err != nil { return nil, err }
	}
	if err := binary.Write(&b, binary.BigEndian, r.Sequence); err != nil {
		return nil, err
	}
	for _, s := range []string{r.Time, r.Source, r.Target, r.Payload, r.Proof} {
		if err := putNativeString(&b, s); err != nil { return nil, err }
	}
	return b.Bytes(), nil
}

func DecodeNative(data []byte) (Record, error) {
	if len(data) < 5 || !bytes.Equal(data[:4], nativeMagic[:]) || data[4] != nativeVersion {
		return Record{}, errors.New("invalid native header")
	}
	off := 5
	version, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	typ, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	identity, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	if off+8 > len(data) { return Record{}, errors.New("native truncated sequence") }
	seq := binary.BigEndian.Uint64(data[off:off+8]); off += 8
	timeValue, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	source, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	target, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	payload, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	proof, err := getNativeString(data, &off); if err != nil { return Record{}, err }
	if off != len(data) { return Record{}, fmt.Errorf("native trailing bytes: %d", len(data)-off) }
	r := Record{Version:version, Type:typ, Identity:identity, Sequence:seq, Time:timeValue, Source:source, Target:target, Payload:payload, Proof:proof}
	if err := Validate(r); err != nil { return Record{}, err }
	return r, nil
}
