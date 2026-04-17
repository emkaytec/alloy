package v1alpha1

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestEnvelopeUnmarshal(t *testing.T) {
	t.Parallel()

	var envelope Envelope
	data := []byte(`apiVersion: anvil.example.io/v1alpha1
kind: GitHubRepository
metadata:
  name: example-repo
spec:
  owner: example-org
  name: example-repo
`)

	if err := yaml.Unmarshal(data, &envelope); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}

	if envelope.APIVersion != APIVersion {
		t.Fatalf("expected apiVersion %q, got %q", APIVersion, envelope.APIVersion)
	}

	if envelope.Kind != KindGitHubRepository {
		t.Fatalf("expected kind %q, got %q", KindGitHubRepository, envelope.Kind)
	}

	if envelope.Metadata.Name != "example-repo" {
		t.Fatalf("expected metadata.name to be example-repo, got %q", envelope.Metadata.Name)
	}

	var spec GitHubRepositorySpec
	if err := envelope.Spec.Decode(&spec); err != nil {
		t.Fatalf("Decode returned error: %v", err)
	}

	if spec.Owner != "example-org" {
		t.Fatalf("expected spec.owner to be example-org, got %q", spec.Owner)
	}

	if spec.Name != "example-repo" {
		t.Fatalf("expected spec.name to be example-repo, got %q", spec.Name)
	}
}

func boolPtr(v bool) *bool {
	return &v
}

func intPtr(v int) *int {
	return &v
}

func stringPtr(v string) *string {
	return &v
}
