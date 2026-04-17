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
			Owner:         "example-org",
			Name:          "example-repo",
			Visibility:    stringPtr("internal"),
			Description:   stringPtr("Repository reconciled through Anvil"),
			Homepage:      stringPtr("https://example.com"),
			AutoInit:      true,
			DefaultBranch: stringPtr("main"),
			Topics:        []string{"platform", "anvil"},
			Features: &GitHubRepositoryFeaturesSpec{
				HasIssues:   boolPtr(true),
				HasProjects: boolPtr(false),
				HasWiki:     boolPtr(false),
			},
			MergePolicy: &GitHubRepositoryMergePolicySpec{
				AllowSquashMerge:    boolPtr(true),
				AllowMergeCommit:    boolPtr(false),
				AllowRebaseMerge:    boolPtr(true),
				AllowAutoMerge:      boolPtr(true),
				DeleteBranchOnMerge: boolPtr(true),
			},
		},
	)

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

/*
Legacy broader-shape fixture retained as commented reference:

Initialization: &GitHubRepositoryInitializationSpec{
	GitignoreTemplate: "Go",
	LicenseTemplate:   "mit",
	IsTemplate:        boolPtr(false),
},
SecurityAndAnalysis: &GitHubRepositorySecurityAndAnalysisSpec{
	AdvancedSecurity:             &GitHubRepositorySecuritySettingSpec{Status: "enabled"},
	SecretScanning:               &GitHubRepositorySecuritySettingSpec{Status: "enabled"},
	SecretScanningPushProtection: &GitHubRepositorySecuritySettingSpec{Status: "enabled"},
},
Pages: &GitHubRepositoryPagesSpec{
	BuildType:     stringPtr("legacy"),
	HTTPSEnforced: boolPtr(true),
	Source: &GitHubRepositoryPagesSourceSpec{
		Branch: "main",
		Path:   "/docs",
	},
},
CustomProperties: []GitHubRepositoryCustomPropertySpec{
	{Name: "service", Value: "alloy"},
	{Name: "tier", Value: "platform"},
},
Branches: []GitHubRepositoryBranchSpec{
	{
		Name: "main",
		Protection: &GitHubRepositoryBranchProtectionSpec{
			RequiredStatusChecks: &GitHubRequiredStatusChecksSpec{
				Strict: true,
				Checks: []GitHubRequiredStatusCheckSpec{
					{Context: "ci/build"},
					{Context: "ci/test"},
				},
			},
			EnforceAdmins:         boolPtr(true),
			RequiredLinearHistory: boolPtr(true),
			PullRequestReviews: &GitHubPullRequestReviewsSpec{
				DismissStaleReviews:          boolPtr(true),
				RequireCodeOwnerReviews:      boolPtr(true),
				RequiredApprovingReviewCount: intPtr(2),
			},
		},
	},
},
*/

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

func TestGitHubRepositoryManifestValidateRejectsInvalidVisibility(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner:      "example-org",
			Name:       "example-repo",
			Visibility: stringPtr("secret"),
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `unsupported spec.visibility "secret"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateRejectsBlankTopics(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner:  "example-org",
			Name:   "example-repo",
			Topics: []string{"platform", " "},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.topics must not contain blank values" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateRejectsDuplicateTopics(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner:  "example-org",
			Name:   "example-repo",
			Topics: []string{"platform", "Platform"},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.topics contains duplicate value "Platform"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateAllowsExplicitEmptyOptionalStrings(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner:         "example-org",
			Name:          "example-repo",
			Visibility:    stringPtr(""),
			Description:   stringPtr(""),
			Homepage:      stringPtr(""),
			DefaultBranch: stringPtr(""),
		},
	)

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

/*
Legacy removed tests retained as commented reference:

func TestGitHubRepositoryManifestValidateRequiresSquashTitleWhenMessageConfigured(t *testing.T)
func TestGitHubRepositoryManifestValidateRejectsInvalidSecurityStatus(t *testing.T)
func TestGitHubRepositoryManifestValidateRejectsInvalidPagesSourcePath(t *testing.T)
func TestGitHubRepositoryManifestValidateRejectsLinearHistoryWithoutCompatibleMergeStrategy(t *testing.T)
func TestGitHubRepositoryManifestValidateRejectsDeprecatedHasDownloads(t *testing.T)
*/

func boolPtr(v bool) *bool {
	return &v
}

func intPtr(v int) *int {
	return &v
}

func stringPtr(v string) *string {
	return &v
}
