package v1alpha1

import (
	"fmt"
	"strings"
)

const KindGitHubRepository = "GitHubRepository"

type GitHubRepositorySpec struct {
	Owner               string                                   `yaml:"owner"`
	Name                string                                   `yaml:"name"`
	Visibility          string                                   `yaml:"visibility,omitempty"`
	Description         string                                   `yaml:"description,omitempty"`
	Homepage            string                                   `yaml:"homepage,omitempty"`
	AutoInit            bool                                     `yaml:"autoInit,omitempty"`
	DefaultBranch       string                                   `yaml:"defaultBranch,omitempty"`
	Topics              []string                                 `yaml:"topics,omitempty"`
	Archived            *bool                                    `yaml:"archived,omitempty"`
	Features            *GitHubRepositoryFeaturesSpec            `yaml:"features,omitempty"`
	Initialization      *GitHubRepositoryInitializationSpec      `yaml:"initialization,omitempty"`
	MergePolicy         *GitHubRepositoryMergePolicySpec         `yaml:"mergePolicy,omitempty"`
	SecurityAndAnalysis *GitHubRepositorySecurityAndAnalysisSpec `yaml:"securityAndAnalysis,omitempty"`
	Pages               *GitHubRepositoryPagesSpec               `yaml:"pages,omitempty"`
	CustomProperties    []GitHubRepositoryCustomPropertySpec     `yaml:"customProperties,omitempty"`
	Branches            []GitHubRepositoryBranchSpec             `yaml:"branches,omitempty"`
}

type GitHubRepositoryFeaturesSpec struct {
	HasIssues    *bool `yaml:"hasIssues,omitempty"`
	HasProjects  *bool `yaml:"hasProjects,omitempty"`
	HasWiki      *bool `yaml:"hasWiki,omitempty"`
	HasDownloads *bool `yaml:"hasDownloads,omitempty"`
}

type GitHubRepositoryInitializationSpec struct {
	GitignoreTemplate string `yaml:"gitignoreTemplate,omitempty"`
	LicenseTemplate   string `yaml:"licenseTemplate,omitempty"`
	IsTemplate        *bool  `yaml:"isTemplate,omitempty"`
}

type GitHubRepositoryMergePolicySpec struct {
	AllowSquashMerge         *bool  `yaml:"allowSquashMerge,omitempty"`
	AllowMergeCommit         *bool  `yaml:"allowMergeCommit,omitempty"`
	AllowRebaseMerge         *bool  `yaml:"allowRebaseMerge,omitempty"`
	AllowAutoMerge           *bool  `yaml:"allowAutoMerge,omitempty"`
	AllowUpdateBranch        *bool  `yaml:"allowUpdateBranch,omitempty"`
	DeleteBranchOnMerge      *bool  `yaml:"deleteBranchOnMerge,omitempty"`
	SquashMergeCommitTitle   string `yaml:"squashMergeCommitTitle,omitempty"`
	SquashMergeCommitMessage string `yaml:"squashMergeCommitMessage,omitempty"`
	MergeCommitTitle         string `yaml:"mergeCommitTitle,omitempty"`
	MergeCommitMessage       string `yaml:"mergeCommitMessage,omitempty"`
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
	BuildType     string                           `yaml:"buildType,omitempty"`
	CNAME         string                           `yaml:"cname,omitempty"`
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
	if s.Visibility != "" && !contains([]string{"public", "private", "internal"}, s.Visibility) {
		return fmt.Errorf("unsupported spec.visibility %q", s.Visibility)
	}

	if err := validateUniqueNonEmptyStrings("spec.topics", s.Topics); err != nil {
		return err
	}

	if err := validateFeaturesSpec(s.Features); err != nil {
		return err
	}

	if err := validateInitializationSpec(s.Initialization); err != nil {
		return err
	}

	if err := validateMergePolicySpec(s.MergePolicy); err != nil {
		return err
	}

	if err := validateSecurityAndAnalysisSpec(s.SecurityAndAnalysis); err != nil {
		return err
	}

	if err := validatePagesSpec(s.Pages); err != nil {
		return err
	}

	if err := validateCustomPropertiesSpec(s.CustomProperties); err != nil {
		return err
	}

	if err := validateBranchesSpec(s.Branches); err != nil {
		return err
	}

	if s.MergePolicy != nil && explicitlyTrue(branchesRequireLinearHistory(s.Branches)) &&
		explicitlyFalse(s.MergePolicy.AllowSquashMerge) && explicitlyFalse(s.MergePolicy.AllowRebaseMerge) {
		return fmt.Errorf("spec.mergePolicy must allow squash or rebase merges when branch protection requires linear history")
	}

	return nil
}

func validateFeaturesSpec(spec *GitHubRepositoryFeaturesSpec) error {
	if spec == nil {
		return nil
	}

	return nil
}

func validateInitializationSpec(spec *GitHubRepositoryInitializationSpec) error {
	if spec == nil {
		return nil
	}

	if strings.TrimSpace(spec.GitignoreTemplate) == "" && spec.GitignoreTemplate != "" {
		return fmt.Errorf("spec.initialization.gitignoreTemplate must not be blank")
	}

	if strings.TrimSpace(spec.LicenseTemplate) == "" && spec.LicenseTemplate != "" {
		return fmt.Errorf("spec.initialization.licenseTemplate must not be blank")
	}

	return nil
}

func validateMergePolicySpec(spec *GitHubRepositoryMergePolicySpec) error {
	if spec == nil {
		return nil
	}

	if spec.SquashMergeCommitTitle != "" && !contains([]string{"PR_TITLE", "COMMIT_OR_PR_TITLE"}, spec.SquashMergeCommitTitle) {
		return fmt.Errorf("unsupported spec.mergePolicy.squashMergeCommitTitle %q", spec.SquashMergeCommitTitle)
	}

	if spec.SquashMergeCommitMessage != "" && !contains([]string{"PR_BODY", "COMMIT_MESSAGES", "BLANK"}, spec.SquashMergeCommitMessage) {
		return fmt.Errorf("unsupported spec.mergePolicy.squashMergeCommitMessage %q", spec.SquashMergeCommitMessage)
	}

	if spec.SquashMergeCommitMessage != "" && spec.SquashMergeCommitTitle == "" {
		return fmt.Errorf("spec.mergePolicy.squashMergeCommitTitle is required when spec.mergePolicy.squashMergeCommitMessage is set")
	}

	if spec.MergeCommitTitle != "" && !contains([]string{"PR_TITLE", "MERGE_MESSAGE"}, spec.MergeCommitTitle) {
		return fmt.Errorf("unsupported spec.mergePolicy.mergeCommitTitle %q", spec.MergeCommitTitle)
	}

	if spec.MergeCommitMessage != "" && !contains([]string{"PR_BODY", "PR_TITLE", "BLANK"}, spec.MergeCommitMessage) {
		return fmt.Errorf("unsupported spec.mergePolicy.mergeCommitMessage %q", spec.MergeCommitMessage)
	}

	if spec.MergeCommitMessage != "" && spec.MergeCommitTitle == "" {
		return fmt.Errorf("spec.mergePolicy.mergeCommitTitle is required when spec.mergePolicy.mergeCommitMessage is set")
	}

	return nil
}

func validateSecurityAndAnalysisSpec(spec *GitHubRepositorySecurityAndAnalysisSpec) error {
	if spec == nil {
		return nil
	}

	settings := map[string]*GitHubRepositorySecuritySettingSpec{
		"spec.securityAndAnalysis.advancedSecurity":                      spec.AdvancedSecurity,
		"spec.securityAndAnalysis.codeSecurity":                          spec.CodeSecurity,
		"spec.securityAndAnalysis.secretScanning":                        spec.SecretScanning,
		"spec.securityAndAnalysis.secretScanningPushProtection":          spec.SecretScanningPushProtection,
		"spec.securityAndAnalysis.secretScanningAIDetection":             spec.SecretScanningAIDetection,
		"spec.securityAndAnalysis.secretScanningNonProviderPatterns":     spec.SecretScanningNonProviderPatterns,
		"spec.securityAndAnalysis.secretScanningDelegatedAlertDismissal": spec.SecretScanningDelegatedAlertDismissal,
		"spec.securityAndAnalysis.secretScanningDelegatedBypass":         spec.SecretScanningDelegatedBypass,
	}

	for field, setting := range settings {
		if setting == nil {
			continue
		}

		if setting.Status == "" {
			return fmt.Errorf("missing %s.status", field)
		}

		if !contains([]string{"enabled", "disabled"}, setting.Status) {
			return fmt.Errorf("unsupported %s.status %q", field, setting.Status)
		}
	}

	return nil
}

func validatePagesSpec(spec *GitHubRepositoryPagesSpec) error {
	if spec == nil {
		return nil
	}

	if spec.BuildType != "" && !contains([]string{"legacy", "workflow"}, spec.BuildType) {
		return fmt.Errorf("unsupported spec.pages.buildType %q", spec.BuildType)
	}

	if spec.Source == nil {
		return nil
	}

	if spec.Source.Branch == "" {
		return fmt.Errorf("missing spec.pages.source.branch")
	}

	if !contains([]string{"/", "/docs"}, spec.Source.Path) {
		return fmt.Errorf("unsupported spec.pages.source.path %q", spec.Source.Path)
	}

	return nil
}

func validateCustomPropertiesSpec(specs []GitHubRepositoryCustomPropertySpec) error {
	seen := make(map[string]struct{}, len(specs))
	for _, spec := range specs {
		if spec.Name == "" {
			return fmt.Errorf("missing spec.customProperties.name")
		}

		key := strings.ToLower(spec.Name)
		if _, ok := seen[key]; ok {
			return fmt.Errorf("duplicate spec.customProperties name %q", spec.Name)
		}

		seen[key] = struct{}{}
		if spec.Value == nil {
			return fmt.Errorf("missing spec.customProperties[%s].value", spec.Name)
		}
	}

	return nil
}

func validateBranchesSpec(specs []GitHubRepositoryBranchSpec) error {
	seen := make(map[string]struct{}, len(specs))
	for _, spec := range specs {
		if spec.Name == "" {
			return fmt.Errorf("missing spec.branches.name")
		}

		if strings.Contains(spec.Name, "*") {
			return fmt.Errorf("spec.branches[%s].name must not contain wildcard characters", spec.Name)
		}

		key := strings.ToLower(spec.Name)
		if _, ok := seen[key]; ok {
			return fmt.Errorf("duplicate spec.branches name %q", spec.Name)
		}

		seen[key] = struct{}{}
		if err := validateBranchProtectionSpec(spec.Name, spec.Protection); err != nil {
			return err
		}
	}

	return nil
}

func validateBranchProtectionSpec(branchName string, spec *GitHubRepositoryBranchProtectionSpec) error {
	if spec == nil {
		return nil
	}

	if spec.RequiredStatusChecks != nil {
		for _, check := range spec.RequiredStatusChecks.Checks {
			if check.Context == "" {
				return fmt.Errorf("missing spec.branches[%s].protection.requiredStatusChecks.checks.context", branchName)
			}
		}
	}

	if spec.PullRequestReviews != nil && spec.PullRequestReviews.RequiredApprovingReviewCount != nil {
		count := *spec.PullRequestReviews.RequiredApprovingReviewCount
		if count < 0 || count > 6 {
			return fmt.Errorf("spec.branches[%s].protection.pullRequestReviews.requiredApprovingReviewCount must be between 0 and 6", branchName)
		}
	}

	if err := validateActorAllowanceSpec(branchName, "restrictions", spec.Restrictions); err != nil {
		return err
	}

	if err := validateActorAllowanceSpec(branchName, "bypassPullRequestAllowances", spec.BypassPullRequestAllowances); err != nil {
		return err
	}

	if spec.PullRequestReviews != nil {
		if err := validateActorAllowanceSpec(branchName, "pullRequestReviews.dismissalRestrictions", spec.PullRequestReviews.DismissalRestrictions); err != nil {
			return err
		}
	}

	return nil
}

func validateActorAllowanceSpec(branchName, field string, spec *GitHubActorAllowanceSpec) error {
	if spec == nil {
		return nil
	}

	if err := validateUniqueNonEmptyStrings(fmt.Sprintf("spec.branches[%s].protection.%s.users", branchName, field), spec.Users); err != nil {
		return err
	}

	if err := validateUniqueNonEmptyStrings(fmt.Sprintf("spec.branches[%s].protection.%s.teams", branchName, field), spec.Teams); err != nil {
		return err
	}

	if err := validateUniqueNonEmptyStrings(fmt.Sprintf("spec.branches[%s].protection.%s.apps", branchName, field), spec.Apps); err != nil {
		return err
	}

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

func branchesRequireLinearHistory(branches []GitHubRepositoryBranchSpec) *bool {
	for _, branch := range branches {
		if branch.Protection == nil || branch.Protection.RequiredLinearHistory == nil {
			continue
		}

		if *branch.Protection.RequiredLinearHistory {
			return branch.Protection.RequiredLinearHistory
		}
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

func explicitlyTrue(value *bool) bool {
	return value != nil && *value
}

func explicitlyFalse(value *bool) bool {
	return value != nil && !*value
}
