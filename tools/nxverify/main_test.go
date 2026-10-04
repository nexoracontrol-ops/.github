package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	e := Evidence{
		EvidenceID:    "NX-EVID-TEST-001",
		SchemaVersion: "NX-EVIDENCE-001/v0.1",
		Subject:       "test",
		Claim:         "test claim",
		Source:        Source{Kind: "git", Locator: "github://example"},
		ObservedAt:    "2026-10-04T00:00:00Z",
		Status:        "ASSERTED",
	}
	if err := validate(e); err != nil {
		t.Fatal(err)
	}
}
