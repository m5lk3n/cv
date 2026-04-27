package resume

import "testing"

func TestLoadEmbeddedResume(t *testing.T) {
	r, err := Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if r.Basics.Name == "" {
		t.Error("expected embedded resume to have a basics.name")
	}
}

func TestLoadFromBytes(t *testing.T) {
	in := []byte(`{"basics":{"name":"Alice","label":"Tester"}}`)
	r, err := LoadFromBytes(in)
	if err != nil {
		t.Fatalf("LoadFromBytes error: %v", err)
	}
	if r.Basics.Name != "Alice" || r.Basics.Label != "Tester" {
		t.Errorf("unexpected resume: %+v", r.Basics)
	}
}

func TestLoadFromBytesInvalid(t *testing.T) {
	if _, err := LoadFromBytes([]byte("not json")); err == nil {
		t.Error("expected error for invalid JSON")
	}
}
