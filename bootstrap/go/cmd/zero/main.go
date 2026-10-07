package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	zero "zero-bootstrap"
)

type output struct {
	Canonical string              `json:"canonical"`
	Authorization zero.Authorization `json:"authorization"`
	Result       zero.Result      `json:"result"`
	Previous     zero.SemanticState `json:"previous"`
	Next         zero.SemanticState `json:"next"`
	Transitioned bool              `json:"transitioned"`
}

func main() {
	actor := flag.String("actor", "", "execution actor")
	operation := flag.String("operation", "OBSERVE", "ZERO operation")
	flag.Parse()

	if *actor == "" {
		fmt.Fprintln(os.Stderr, "error: --actor is required")
		os.Exit(2)
	}

	in := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for in.Scan() {
		line := in.Text()
		receipt, err := zero.ExecuteRecord(line, *actor, *operation, zero.SemanticState{})
		if err != nil {
			_ = enc.Encode(output{
				Canonical:     receipt.Canonical,
				Authorization: receipt.Authorization,
				Result:        receipt.Result,
				Previous:      receipt.Previous,
				Next:          receipt.Next,
				Transitioned:  receipt.Transitioned,
			})
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		if err := enc.Encode(output{
			Canonical:     receipt.Canonical,
			Authorization: receipt.Authorization,
			Result:        receipt.Result,
			Previous:      receipt.Previous,
			Next:          receipt.Next,
			Transitioned:  receipt.Transitioned,
		}); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
