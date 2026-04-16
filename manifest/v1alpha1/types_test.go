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
			Initialization: &GitHubRepositoryInitializationSpec{
				GitignoreTemplate: "Go",
				LicenseTemplate:   "mit",
				IsTemplate:        boolPtr(false),
			},
			MergePolicy: &GitHubRepositoryMergePolicySpec{
				AllowSquashMerge:         boolPtr(true),
				AllowMergeCommit:         boolPtr(false),
				AllowRebaseMerge:         boolPtr(true),
				AllowAutoMerge:           boolPtr(true),
				DeleteBranchOnMerge:      boolPtr(true),
				SquashMergeCommitTitle:   stringPtr("PR_TITLE"),
				SquashMergeCommitMessage: stringPtr("PR_BODY"),
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

func TestGitHubRepositoryManifestValidateRequiresSquashTitleWhenMessageConfigured(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
			Name:  "example-repo",
			MergePolicy: &GitHubRepositoryMergePolicySpec{
				SquashMergeCommitMessage: stringPtr("PR_BODY"),
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.mergePolicy.squashMergeCommitTitle is required when spec.mergePolicy.squashMergeCommitMessage is set" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateRejectsInvalidSecurityStatus(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
			Name:  "example-repo",
			SecurityAndAnalysis: &GitHubRepositorySecurityAndAnalysisSpec{
				SecretScanning: &GitHubRepositorySecuritySettingSpec{Status: "maybe"},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `unsupported spec.securityAndAnalysis.secretScanning.status "maybe"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateRejectsInvalidPagesSourcePath(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
			Name:  "example-repo",
			Pages: &GitHubRepositoryPagesSpec{
				Source: &GitHubRepositoryPagesSourceSpec{
					Branch: "main",
					Path:   "/site",
				},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `unsupported spec.pages.source.path "/site"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateRejectsLinearHistoryWithoutCompatibleMergeStrategy(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
			Name:  "example-repo",
			MergePolicy: &GitHubRepositoryMergePolicySpec{
				AllowSquashMerge: boolPtr(false),
				AllowRebaseMerge: boolPtr(false),
			},
			Branches: []GitHubRepositoryBranchSpec{
				{
					Name: "main",
					Protection: &GitHubRepositoryBranchProtectionSpec{
						RequiredLinearHistory: boolPtr(true),
					},
				},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.mergePolicy must allow squash or rebase merges when branch protection requires linear history" {
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
			MergePolicy: &GitHubRepositoryMergePolicySpec{
				SquashMergeCommitTitle:   stringPtr(""),
				SquashMergeCommitMessage: stringPtr(""),
				MergeCommitTitle:         stringPtr(""),
				MergeCommitMessage:       stringPtr(""),
			},
			Pages: &GitHubRepositoryPagesSpec{
				BuildType: stringPtr(""),
				CNAME:     stringPtr(""),
			},
		},
	)

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

func TestGitHubRepositoryManifestValidateRejectsDeprecatedHasDownloads(t *testing.T) {
	t.Parallel()

	manifest := NewGitHubRepositoryManifest(
		Metadata{Name: "example-repo"},
		GitHubRepositorySpec{
			Owner: "example-org",
			Name:  "example-repo",
			Features: &GitHubRepositoryFeaturesSpec{
				HasDownloads: boolPtr(true),
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.features.hasDownloads is deprecated and unsupported" {
		t.Fatalf("unexpected error: %v", err)
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
