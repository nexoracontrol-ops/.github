package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Source struct {
	Kind     string `json:"kind"`
	Locator  string `json:"locator"`
}

type Evidence struct {
	EvidenceID      string  `json:"evidence_id"`
	SchemaVersion   string  `json:"schema_version"`
	Subject         string  `json:"subject"`
	Claim           string  `json:"claim"`
	Source          Source  `json:"source"`
	ObservedAt      string  `json:"observed_at"`
	ObservedBy      string  `json:"observed_by,omitempty"`
	ContentHash     string  `json:"content_hash,omitempty"`
	ArtifactVersion string  `json:"artifact_version,omitempty"`
	AuthorityRef    string  `json:"authority_ref,omitempty"`
	Status          string  `json:"status"`
	Notes           string  `json:"notes,omitempty"`
}

func validate(e Evidence) error {
	if !strings.HasPrefix(e.EvidenceID, "NX-EVID-") {
		return errors.New("invalid evidence_id")
	}
	if e.SchemaVersion != "NX-EVIDENCE-001/v0.1" {
		return errors.New("unsupported schema_version")
	}
	if e.Subject == "" || e.Claim == "" || e.Source.Kind == "" || e.Source.Locator == "" {
		return errors.New("missing required field")
	}
	if e.Status != "ASSERTED" && e.Status != "SUPPORTED" && e.Status != "VERIFIED" &&
		e.Status != "REJECTED" && e.Status != "SUPERSEDED" {
		return errors.New("invalid status")
	}
	if _, err := time.Parse(time.RFC3339, e.ObservedAt); err != nil {
		return fmt.Errorf("invalid observed_at: %w", err)
	}
	if e.ContentHash != "" {
		const prefix = "sha256:"
		if !strings.HasPrefix(e.ContentHash, prefix) || len(e.ContentHash) != len(prefix)+64 {
			return errors.New("invalid content_hash")
		}
		if _, err := hex.DecodeString(strings.TrimPrefix(e.ContentHash, prefix)); err != nil {
			return errors.New("invalid content_hash hex")
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: nxverify <evidence.json>")
		os.Exit(2)
	}

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	var e Evidence
	if err := json.Unmarshal(raw, &e); err != nil {
		fmt.Fprintln(os.Stderr, "invalid JSON:", err)
		os.Exit(1)
	}

	if err := validate(e); err != nil {
		fmt.Fprintln(os.Stderr, "REJECTED:", err)
		os.Exit(1)
	}

	digest := sha256.Sum256(raw)
	fmt.Printf("ACCEPTED\nevidence_id=%s\nraw_sha256=sha256:%x\n", e.EvidenceID, digest)
}
