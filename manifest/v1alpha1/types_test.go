package v1alpha1

import "testing"

func TestNewGitHubRepositoryManifestUsesCurrentVersionAndKind(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
			Name:  "example-repo",
		},
	)

	if manifest.APIVersion != APIVersion {
		t.Fatalf("expected apiVersion %q, got %q", APIVersion, manifest.APIVersion)
	}

	if manifest.Kind != KindGitHubRepository {
		t.Fatalf("expected kind %q, got %q", KindGitHubRepository, manifest.Kind)
	}
}

func TestGitHubRepositoryManifestValidate(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
			Name:  "example-repo",
		},
	)

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateRequiresName(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "missing spec.name" {
		t.Fatalf("expected missing spec.name error, got %v", err)
	}
}
