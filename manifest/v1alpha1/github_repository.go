package v1alpha1

import (
	"fmt"
	"strings"
)

const KindGitHubRepository = "GitHubRepository"

type GitHubRepositorySpec struct {
	// Owner and Name form the stable repository identity. Reconciliation does not
	// treat transfer or rename as an ordinary in-place update.
	Owner string `yaml:"owner"`
	Name  string `yaml:"name"`
	// Nil means unmanaged. Non-nil means reconcile the provided value, including
	// explicit empty strings where supported by the downstream API.
	Visibility  *string `yaml:"visibility,omitempty"`
	Description *string `yaml:"description,omitempty"`
	Homepage    *string `yaml:"homepage,omitempty"`
	// AutoInit is a create-time-only input honored when creating a repository.
	AutoInit      bool    `yaml:"autoInit,omitempty"`
	DefaultBranch *string `yaml:"defaultBranch,omitempty"`
	// Omitted means unmanaged; a present empty list means manage and clear.
	Topics      []string                         `yaml:"topics,omitempty"`
	Features    *GitHubRepositoryFeaturesSpec    `yaml:"features,omitempty"`
	MergePolicy *GitHubRepositoryMergePolicySpec `yaml:"mergePolicy,omitempty"`

	// Narrowed out of the active v1alpha1 schema for now:
	// Archived *bool `yaml:"archived,omitempty"`
	// Initialization *GitHubRepositoryInitializationSpec `yaml:"initialization,omitempty"`
	// SecurityAndAnalysis *GitHubRepositorySecurityAndAnalysisSpec `yaml:"securityAndAnalysis,omitempty"`
	// Pages *GitHubRepositoryPagesSpec `yaml:"pages,omitempty"`
	// CustomProperties []GitHubRepositoryCustomPropertySpec `yaml:"customProperties,omitempty"`
	// Branches []GitHubRepositoryBranchSpec `yaml:"branches,omitempty"`
}

type GitHubRepositoryFeaturesSpec struct {
	HasIssues   *bool `yaml:"hasIssues,omitempty"`
	HasProjects *bool `yaml:"hasProjects,omitempty"`
	HasWiki     *bool `yaml:"hasWiki,omitempty"`

	// Narrowed out of the active v1alpha1 schema for now:
	// HasDownloads *bool `yaml:"hasDownloads,omitempty"`
}

type GitHubRepositoryMergePolicySpec struct {
	AllowSquashMerge    *bool `yaml:"allowSquashMerge,omitempty"`
	AllowMergeCommit    *bool `yaml:"allowMergeCommit,omitempty"`
	AllowRebaseMerge    *bool `yaml:"allowRebaseMerge,omitempty"`
	AllowAutoMerge      *bool `yaml:"allowAutoMerge,omitempty"`
	AllowUpdateBranch   *bool `yaml:"allowUpdateBranch,omitempty"`
	DeleteBranchOnMerge *bool `yaml:"deleteBranchOnMerge,omitempty"`

	// Narrowed out of the active v1alpha1 schema for now:
	// SquashMergeCommitTitle *string `yaml:"squashMergeCommitTitle,omitempty"`
	// SquashMergeCommitMessage *string `yaml:"squashMergeCommitMessage,omitempty"`
	// MergeCommitTitle *string `yaml:"mergeCommitTitle,omitempty"`
	// MergeCommitMessage *string `yaml:"mergeCommitMessage,omitempty"`
}

/*
Previous broader GitHubRepository-adjacent shapes retained as commented reference:

type GitHubRepositoryInitializationSpec struct {
	GitignoreTemplate string `yaml:"gitignoreTemplate,omitempty"`
	LicenseTemplate   string `yaml:"licenseTemplate,omitempty"`
	IsTemplate        *bool  `yaml:"isTemplate,omitempty"`
}

type GitHubRepositorySecurityAndAnalysisSpec struct {
	AdvancedSecurity                      *GitHubRepositorySecuritySettingSpec `yaml:"advancedSecurity,omitempty"`
	CodeSecurity                          *GitHubRepositorySecuritySettingSpec `yaml:"codeSecurity,omitempty"`
	SecretScanning                        *GitHubRepositorySecuritySettingSpec `yaml:"secretScanning,omitempty"`
	SecretScanningPushProtection          *GitHubRepositorySecuritySettingSpec `yaml:"secretScanningPushProtection,omitempty"`
	SecretScanningAIDetection             *GitHubRepositorySecuritySettingSpec `yaml:"secretScanningAIDetection,omitempty"`
	SecretScanningNonProviderPatterns     *GitHubRepositorySecuritySettingSpec `yaml:"secretScanningNonProviderPatterns,omitempty"`
	SecretScanningDelegatedAlertDismissal *GitHubRepositorySecuritySettingSpec `yaml:"secretScanningDelegatedAlertDismissal,omitempty"`
	SecretScanningDelegatedBypass         *GitHubRepositorySecuritySettingSpec `yaml:"secretScanningDelegatedBypass,omitempty"`
}

type GitHubRepositorySecuritySettingSpec struct {
	Status string `yaml:"status,omitempty"`
}

type GitHubRepositoryPagesSpec struct {
	BuildType     *string                          `yaml:"buildType,omitempty"`
	CNAME         *string                          `yaml:"cname,omitempty"`
	HTTPSEnforced *bool                            `yaml:"httpsEnforced,omitempty"`
	Source        *GitHubRepositoryPagesSourceSpec `yaml:"source,omitempty"`
}

type GitHubRepositoryPagesSourceSpec struct {
	Branch string `yaml:"branch"`
	Path   string `yaml:"path"`
}

type GitHubRepositoryCustomPropertySpec struct {
	Name  string `yaml:"name"`
	Value any    `yaml:"value"`
}

type GitHubRepositoryBranchSpec struct {
	Name       string                                `yaml:"name"`
	Protection *GitHubRepositoryBranchProtectionSpec `yaml:"protection,omitempty"`
}

type GitHubRepositoryBranchProtectionSpec struct {
	RequiredStatusChecks        *GitHubRequiredStatusChecksSpec `yaml:"requiredStatusChecks,omitempty"`
	EnforceAdmins               *bool                           `yaml:"enforceAdmins,omitempty"`
	PullRequestReviews          *GitHubPullRequestReviewsSpec   `yaml:"pullRequestReviews,omitempty"`
	Restrictions                *GitHubActorAllowanceSpec       `yaml:"restrictions,omitempty"`
	BypassPullRequestAllowances *GitHubActorAllowanceSpec       `yaml:"bypassPullRequestAllowances,omitempty"`
	RequiredLinearHistory       *bool                           `yaml:"requiredLinearHistory,omitempty"`
}

type GitHubRequiredStatusChecksSpec struct {
	Strict bool                            `yaml:"strict"`
	Checks []GitHubRequiredStatusCheckSpec `yaml:"checks,omitempty"`
}

type GitHubRequiredStatusCheckSpec struct {
	Context string `yaml:"context"`
	AppID   *int64 `yaml:"appID,omitempty"`
}

type GitHubPullRequestReviewsSpec struct {
	DismissalRestrictions        *GitHubActorAllowanceSpec `yaml:"dismissalRestrictions,omitempty"`
	DismissStaleReviews          *bool                     `yaml:"dismissStaleReviews,omitempty"`
	RequireCodeOwnerReviews      *bool                     `yaml:"requireCodeOwnerReviews,omitempty"`
	RequiredApprovingReviewCount *int                      `yaml:"requiredApprovingReviewCount,omitempty"`
	RequireLastPushApproval      *bool                     `yaml:"requireLastPushApproval,omitempty"`
}

type GitHubActorAllowanceSpec struct {
	Users []string `yaml:"users,omitempty"`
	Teams []string `yaml:"teams,omitempty"`
	Apps  []string `yaml:"apps,omitempty"`
}
*/

type GitHubRepositoryManifest struct {
	APIVersion string               `yaml:"apiVersion"`
	Kind       string               `yaml:"kind"`
	Metadata   Metadata             `yaml:"metadata"`
	Spec       GitHubRepositorySpec `yaml:"spec"`
}

func NewGitHubRepositoryManifest(metadata Metadata, spec GitHubRepositorySpec) GitHubRepositoryManifest {
	return GitHubRepositoryManifest{
		APIVersion: APIVersion,
		Kind:       KindGitHubRepository,
		Metadata:   metadata,
		Spec:       spec,
	}
}

func (m GitHubRepositoryManifest) Validate() error {
	if m.APIVersion == "" {
		return fmt.Errorf("missing apiVersion")
	}

	if m.Kind == "" {
		return fmt.Errorf("missing kind")
	}

	if m.Kind != KindGitHubRepository {
		return fmt.Errorf("unsupported kind %q", m.Kind)
	}

	if m.Metadata.Name == "" {
		return fmt.Errorf("missing metadata.name")
	}

	if m.Spec.Owner == "" {
		return fmt.Errorf("missing spec.owner")
	}

	if m.Spec.Name == "" {
		return fmt.Errorf("missing spec.name")
	}

	if err := m.Spec.Validate(); err != nil {
		return err
	}

	return nil
}

func (s GitHubRepositorySpec) Validate() error {
	if s.Visibility != nil && *s.Visibility != "" && !contains([]string{"public", "private", "internal"}, *s.Visibility) {
		return fmt.Errorf("unsupported spec.visibility %q", *s.Visibility)
	}

	if err := validateUniqueNonEmptyStrings("spec.topics", s.Topics); err != nil {
		return err
	}

	// Previously active validation hooks kept here as commented reference:
	// if err := validateInitializationSpec(s.Initialization); err != nil { return err }
	// if err := validateMergePolicySpec(s.MergePolicy); err != nil { return err }
	// if err := validateSecurityAndAnalysisSpec(s.SecurityAndAnalysis); err != nil { return err }
	// if err := validatePagesSpec(s.Pages); err != nil { return err }
	// if err := validateCustomPropertiesSpec(s.CustomProperties); err != nil { return err }
	// if err := validateBranchesSpec(s.Branches); err != nil { return err }
	// if s.MergePolicy != nil && explicitlyTrue(branchesRequireLinearHistory(s.Branches)) &&
	// 	explicitlyFalse(s.MergePolicy.AllowSquashMerge) && explicitlyFalse(s.MergePolicy.AllowRebaseMerge) {
	// 	return fmt.Errorf("spec.mergePolicy must allow squash or rebase merges when branch protection requires linear history")
	// }

	return nil
}

func validateUniqueNonEmptyStrings(field string, values []string) error {
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return fmt.Errorf("%s must not contain blank values", field)
		}

		key := strings.ToLower(trimmed)
		if _, ok := seen[key]; ok {
			return fmt.Errorf("%s contains duplicate value %q", field, value)
		}

		seen[key] = struct{}{}
	}

	return nil
}

func contains(values []string, candidate string) bool {
	for _, value := range values {
		if value == candidate {
			return true
		}
	}

	return false
}

/*
Previous helper set retained as commented reference:

func validateInitializationSpec(spec *GitHubRepositoryInitializationSpec) error
func validateMergePolicySpec(spec *GitHubRepositoryMergePolicySpec) error
func validateSecurityAndAnalysisSpec(spec *GitHubRepositorySecurityAndAnalysisSpec) error
func validatePagesSpec(spec *GitHubRepositoryPagesSpec) error
func validateCustomPropertiesSpec(specs []GitHubRepositoryCustomPropertySpec) error
func validateBranchesSpec(specs []GitHubRepositoryBranchSpec) error
func validateBranchProtectionSpec(branchName string, spec *GitHubRepositoryBranchProtectionSpec) error
func validateActorAllowanceSpec(branchName, field string, spec *GitHubActorAllowanceSpec) error
func branchesRequireLinearHistory(branches []GitHubRepositoryBranchSpec) *bool
func explicitlyTrue(value *bool) bool
func explicitlyFalse(value *bool) bool
*/
