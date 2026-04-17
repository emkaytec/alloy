package v1alpha1

import (
	"fmt"
	"regexp"
	"strings"
)

const KindHCPTerraformWorkspace = "HCPTerraformWorkspace"

var autoDestroyActivityDurationPattern = regexp.MustCompile(`^[1-9][0-9]{0,3}[dh]$`)

type HCPTerraformWorkspaceSpec struct {
	// Organization and Name form the stable workspace identity. Reconciliation
	// does not treat moving the workspace to a different organization or renaming
	// it as an ordinary in-place update.
	Organization string `yaml:"organization"`
	Name         string `yaml:"name"`
	// Nil means unmanaged. Non-nil means reconcile the provided value, including
	// explicit empty strings where supported by the downstream API.
	ProjectID                   *string                                 `yaml:"projectID,omitempty"`
	Description                 *string                                 `yaml:"description,omitempty"`
	TerraformVersion            *string                                 `yaml:"terraformVersion,omitempty"`
	WorkingDirectory            *string                                 `yaml:"workingDirectory,omitempty"`
	ExecutionMode               *string                                 `yaml:"executionMode,omitempty"`
	AgentPoolID                 *string                                 `yaml:"agentPoolID,omitempty"`
	AllowDestroyPlan            *bool                                   `yaml:"allowDestroyPlan,omitempty"`
	AssessmentsEnabled          *bool                                   `yaml:"assessmentsEnabled,omitempty"`
	AutoApply                   *bool                                   `yaml:"autoApply,omitempty"`
	AutoApplyRunTrigger         *bool                                   `yaml:"autoApplyRunTrigger,omitempty"`
	AutoDestroyAt               *string                                 `yaml:"autoDestroyAt,omitempty"`
	AutoDestroyActivityDuration *string                                 `yaml:"autoDestroyActivityDuration,omitempty"`
	FileTriggersEnabled         *bool                                   `yaml:"fileTriggersEnabled,omitempty"`
	GlobalRemoteState           *bool                                   `yaml:"globalRemoteState,omitempty"`
	ProjectRemoteState          *bool                                   `yaml:"projectRemoteState,omitempty"`
	QueueAllRuns                *bool                                   `yaml:"queueAllRuns,omitempty"`
	SourceName                  *string                                 `yaml:"sourceName,omitempty"`
	SourceURL                   *string                                 `yaml:"sourceURL,omitempty"`
	SpeculativeEnabled          *bool                                   `yaml:"speculativeEnabled,omitempty"`
	SSHKeyID                    *string                                 `yaml:"sshKeyID,omitempty"`
	SettingOverwrites           *HCPTerraformWorkspaceSettingOverwrites `yaml:"settingOverwrites,omitempty"`
	VCSRepo                     *HCPTerraformWorkspaceVCSRepoSpec       `yaml:"vcsRepo,omitempty"`
	// Omitted means unmanaged; a present empty list means manage and clear.
	Tags                   []string                                `yaml:"tags,omitempty"`
	TagBindings            []HCPTerraformWorkspaceTagBindingSpec   `yaml:"tagBindings,omitempty"`
	TriggerPatterns        []string                                `yaml:"triggerPatterns,omitempty"`
	TriggerPrefixes        []string                                `yaml:"triggerPrefixes,omitempty"`
	RemoteStateConsumerIDs []string                                `yaml:"remoteStateConsumerIDs,omitempty"`
	Variables              []HCPTerraformWorkspaceVariableSpec     `yaml:"variables,omitempty"`
	VariableSetIDs         []string                                `yaml:"variableSetIDs,omitempty"`
	RunTriggers            []HCPTerraformWorkspaceRunTriggerSpec   `yaml:"runTriggers,omitempty"`
	TeamAccess             []HCPTerraformWorkspaceTeamAccessSpec   `yaml:"teamAccess,omitempty"`
	Notifications          []HCPTerraformWorkspaceNotificationSpec `yaml:"notifications,omitempty"`

	// Left out of the active schema for now:
	// Operations *bool `yaml:"operations,omitempty"`
	// HYOKEnabled *bool `yaml:"hyokEnabled,omitempty"`
	// DataRetentionPolicyID *string `yaml:"dataRetentionPolicyID,omitempty"`
}

type HCPTerraformWorkspaceSettingOverwrites struct {
	ExecutionMode       *bool `yaml:"executionMode,omitempty"`
	AgentPoolID         *bool `yaml:"agentPoolID,omitempty"`
	AutoApply           *bool `yaml:"autoApply,omitempty"`
	FileTriggersEnabled *bool `yaml:"fileTriggersEnabled,omitempty"`
	GlobalRemoteState   *bool `yaml:"globalRemoteState,omitempty"`
	QueueAllRuns        *bool `yaml:"queueAllRuns,omitempty"`
	SpeculativeEnabled  *bool `yaml:"speculativeEnabled,omitempty"`
	TerraformVersion    *bool `yaml:"terraformVersion,omitempty"`
	TriggerPatterns     *bool `yaml:"triggerPatterns,omitempty"`
	TriggerPrefixes     *bool `yaml:"triggerPrefixes,omitempty"`
	WorkingDirectory    *bool `yaml:"workingDirectory,omitempty"`
}

type HCPTerraformWorkspaceVCSRepoSpec struct {
	Identifier *string `yaml:"identifier,omitempty"`
	Branch     *string `yaml:"branch,omitempty"`
	// Nil means unmanaged. Explicit false means disable submodule fetching.
	IngressSubmodules *bool   `yaml:"ingressSubmodules,omitempty"`
	OAuthTokenID      *string `yaml:"oauthTokenID,omitempty"`
	TagsRegex         *string `yaml:"tagsRegex,omitempty"`

	// Left out of the active schema for now:
	// GithubAppInstallationID *string `yaml:"githubAppInstallationID,omitempty"`
}

type HCPTerraformWorkspaceTagBindingSpec struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type HCPTerraformWorkspaceVariableSpec struct {
	Key      string `yaml:"key"`
	Category string `yaml:"category"`
	// Value is intentionally required when a variable is declared in the manifest.
	// Sensitive variables should still provide the desired value here so the
	// downstream reconciler can create or update them.
	Value       string  `yaml:"value"`
	Description *string `yaml:"description,omitempty"`
	Sensitive   *bool   `yaml:"sensitive,omitempty"`
	HCL         *bool   `yaml:"hcl,omitempty"`
}

type HCPTerraformWorkspaceRunTriggerSpec struct {
	SourceWorkspaceID string `yaml:"sourceWorkspaceID"`
}

type HCPTerraformWorkspaceTeamAccessSpec struct {
	TeamName         string  `yaml:"teamName"`
	Access           string  `yaml:"access"`
	Runs             *string `yaml:"runs,omitempty"`
	Variables        *string `yaml:"variables,omitempty"`
	StateVersions    *string `yaml:"stateVersions,omitempty"`
	SentinelMocks    *string `yaml:"sentinelMocks,omitempty"`
	WorkspaceLocking *bool   `yaml:"workspaceLocking,omitempty"`
	RunTasks         *bool   `yaml:"runTasks,omitempty"`
}

type HCPTerraformWorkspaceNotificationSpec struct {
	Name            string   `yaml:"name"`
	DestinationType string   `yaml:"destinationType"`
	Enabled         *bool    `yaml:"enabled,omitempty"`
	URL             *string  `yaml:"url,omitempty"`
	Token           *string  `yaml:"token,omitempty"`
	Triggers        []string `yaml:"triggers,omitempty"`
}

type HCPTerraformWorkspaceManifest struct {
	APIVersion string                    `yaml:"apiVersion"`
	Kind       string                    `yaml:"kind"`
	Metadata   Metadata                  `yaml:"metadata"`
	Spec       HCPTerraformWorkspaceSpec `yaml:"spec"`
}

func NewHCPTerraformWorkspaceManifest(metadata Metadata, spec HCPTerraformWorkspaceSpec) HCPTerraformWorkspaceManifest {
	return HCPTerraformWorkspaceManifest{
		APIVersion: APIVersion,
		Kind:       KindHCPTerraformWorkspace,
		Metadata:   metadata,
		Spec:       spec,
	}
}

func (m HCPTerraformWorkspaceManifest) Validate() error {
	if m.APIVersion == "" {
		return fmt.Errorf("missing apiVersion")
	}

	if m.Kind == "" {
		return fmt.Errorf("missing kind")
	}

	if m.Kind != KindHCPTerraformWorkspace {
		return fmt.Errorf("unsupported kind %q", m.Kind)
	}

	if m.Metadata.Name == "" {
		return fmt.Errorf("missing metadata.name")
	}

	if m.Spec.Organization == "" {
		return fmt.Errorf("missing spec.organization")
	}

	if m.Spec.Name == "" {
		return fmt.Errorf("missing spec.name")
	}

	if err := m.Spec.Validate(); err != nil {
		return err
	}

	return nil
}

func (s HCPTerraformWorkspaceSpec) Validate() error {
	if s.ExecutionMode != nil && *s.ExecutionMode != "" && !contains([]string{"remote", "local", "agent"}, *s.ExecutionMode) {
		return fmt.Errorf("unsupported spec.executionMode %q", *s.ExecutionMode)
	}

	if s.ExecutionMode != nil && *s.ExecutionMode == "agent" {
		if s.AgentPoolID == nil || strings.TrimSpace(*s.AgentPoolID) == "" {
			return fmt.Errorf("spec.agentPoolID is required when spec.executionMode is \"agent\"")
		}
	}

	if s.AgentPoolID != nil && strings.TrimSpace(*s.AgentPoolID) != "" && s.ExecutionMode != nil && contains([]string{"remote", "local"}, *s.ExecutionMode) {
		return fmt.Errorf("spec.agentPoolID is only supported when spec.executionMode is \"agent\"")
	}

	if s.AutoDestroyActivityDuration != nil && *s.AutoDestroyActivityDuration != "" && !autoDestroyActivityDurationPattern.MatchString(*s.AutoDestroyActivityDuration) {
		return fmt.Errorf("unsupported spec.autoDestroyActivityDuration %q", *s.AutoDestroyActivityDuration)
	}

	if explicitlyTrue(s.GlobalRemoteState) && explicitlyTrue(s.ProjectRemoteState) {
		return fmt.Errorf("spec.globalRemoteState and spec.projectRemoteState cannot both be true")
	}

	if explicitlyTrue(s.GlobalRemoteState) && len(s.RemoteStateConsumerIDs) > 0 {
		return fmt.Errorf("spec.remoteStateConsumerIDs cannot be set when spec.globalRemoteState is true")
	}

	if err := validateUniqueNonEmptyStrings("spec.tags", s.Tags); err != nil {
		return err
	}

	if err := validateTagBindingsSpec(s.TagBindings); err != nil {
		return err
	}

	if err := validateUniqueNonEmptyStrings("spec.triggerPatterns", s.TriggerPatterns); err != nil {
		return err
	}

	if err := validateUniqueNonEmptyStrings("spec.triggerPrefixes", s.TriggerPrefixes); err != nil {
		return err
	}

	if err := validateUniqueNonEmptyStrings("spec.remoteStateConsumerIDs", s.RemoteStateConsumerIDs); err != nil {
		return err
	}

	if err := validateUniqueNonEmptyStrings("spec.variableSetIDs", s.VariableSetIDs); err != nil {
		return err
	}

	if err := validateVCSRepoSpec(s.VCSRepo); err != nil {
		return err
	}

	if err := validateWorkspaceVariablesSpec(s.Variables); err != nil {
		return err
	}

	if err := validateRunTriggersSpec(s.RunTriggers); err != nil {
		return err
	}

	if err := validateTeamAccessSpec(s.TeamAccess); err != nil {
		return err
	}

	if err := validateNotificationsSpec(s.Notifications); err != nil {
		return err
	}

	if s.ExecutionMode != nil && *s.ExecutionMode == "local" {
		if len(s.Variables) > 0 {
			return fmt.Errorf("spec.variables is not supported when spec.executionMode is \"local\"")
		}

		if len(s.VariableSetIDs) > 0 {
			return fmt.Errorf("spec.variableSetIDs is not supported when spec.executionMode is \"local\"")
		}

		if len(s.Notifications) > 0 {
			return fmt.Errorf("spec.notifications is not supported when spec.executionMode is \"local\"")
		}
	}

	return nil
}

func validateVCSRepoSpec(spec *HCPTerraformWorkspaceVCSRepoSpec) error {
	if spec == nil {
		return nil
	}

	if spec.Identifier == nil || strings.TrimSpace(*spec.Identifier) == "" {
		return fmt.Errorf("spec.vcsRepo.identifier is required")
	}

	if spec.OAuthTokenID == nil || strings.TrimSpace(*spec.OAuthTokenID) == "" {
		return fmt.Errorf("spec.vcsRepo.oauthTokenID is required")
	}

	if spec.TagsRegex != nil && strings.TrimSpace(*spec.TagsRegex) == "" {
		return fmt.Errorf("spec.vcsRepo.tagsRegex must not be blank")
	}

	return nil
}

func validateTagBindingsSpec(specs []HCPTerraformWorkspaceTagBindingSpec) error {
	seen := make(map[string]struct{}, len(specs))
	for _, spec := range specs {
		if strings.TrimSpace(spec.Key) == "" {
			return fmt.Errorf("spec.tagBindings.key must not be blank")
		}

		if strings.TrimSpace(spec.Value) == "" {
			return fmt.Errorf("spec.tagBindings.value must not be blank")
		}

		key := spec.Key + "\x00" + spec.Value
		if _, ok := seen[key]; ok {
			return fmt.Errorf("spec.tagBindings contains duplicate binding %q=%q", spec.Key, spec.Value)
		}

		seen[key] = struct{}{}
	}

	return nil
}

func validateWorkspaceVariablesSpec(specs []HCPTerraformWorkspaceVariableSpec) error {
	seen := make(map[string]struct{}, len(specs))
	for _, spec := range specs {
		if strings.TrimSpace(spec.Key) == "" {
			return fmt.Errorf("spec.variables.key must not be blank")
		}

		if strings.TrimSpace(spec.Category) == "" {
			return fmt.Errorf("spec.variables.category must not be blank")
		}

		if !contains([]string{"env", "terraform"}, spec.Category) {
			return fmt.Errorf("unsupported spec.variables.category %q", spec.Category)
		}

		if explicitlyTrue(spec.HCL) && spec.Category == "env" {
			return fmt.Errorf("spec.variables.hcl is only supported for terraform variables")
		}

		key := spec.Category + ":" + spec.Key
		if _, ok := seen[key]; ok {
			return fmt.Errorf("spec.variables contains duplicate variable %q", key)
		}

		seen[key] = struct{}{}
	}

	return nil
}

func validateRunTriggersSpec(specs []HCPTerraformWorkspaceRunTriggerSpec) error {
	seen := make(map[string]struct{}, len(specs))
	for _, spec := range specs {
		if strings.TrimSpace(spec.SourceWorkspaceID) == "" {
			return fmt.Errorf("spec.runTriggers.sourceWorkspaceID must not be blank")
		}

		if _, ok := seen[spec.SourceWorkspaceID]; ok {
			return fmt.Errorf("spec.runTriggers contains duplicate sourceWorkspaceID %q", spec.SourceWorkspaceID)
		}

		seen[spec.SourceWorkspaceID] = struct{}{}
	}

	return nil
}

func validateTeamAccessSpec(specs []HCPTerraformWorkspaceTeamAccessSpec) error {
	seen := make(map[string]struct{}, len(specs))
	for _, spec := range specs {
		if strings.TrimSpace(spec.TeamName) == "" {
			return fmt.Errorf("spec.teamAccess.teamName must not be blank")
		}

		if !contains([]string{"read", "plan", "write", "admin", "custom"}, spec.Access) {
			return fmt.Errorf("unsupported spec.teamAccess.access %q", spec.Access)
		}

		key := strings.ToLower(spec.TeamName)
		if _, ok := seen[key]; ok {
			return fmt.Errorf("spec.teamAccess contains duplicate teamName %q", spec.TeamName)
		}

		seen[key] = struct{}{}

		if spec.Access != "custom" {
			if spec.Runs != nil || spec.Variables != nil || spec.StateVersions != nil || spec.SentinelMocks != nil || spec.WorkspaceLocking != nil || spec.RunTasks != nil {
				return fmt.Errorf("spec.teamAccess custom permissions require spec.teamAccess.access to be \"custom\"")
			}

			continue
		}

		if spec.Runs != nil && !contains([]string{"read", "plan", "apply"}, *spec.Runs) {
			return fmt.Errorf("unsupported spec.teamAccess.runs %q", *spec.Runs)
		}

		if spec.Variables != nil && !contains([]string{"none", "read", "write"}, *spec.Variables) {
			return fmt.Errorf("unsupported spec.teamAccess.variables %q", *spec.Variables)
		}

		if spec.StateVersions != nil && !contains([]string{"none", "read-outputs", "read", "write"}, *spec.StateVersions) {
			return fmt.Errorf("unsupported spec.teamAccess.stateVersions %q", *spec.StateVersions)
		}

		if spec.SentinelMocks != nil && !contains([]string{"none", "read"}, *spec.SentinelMocks) {
			return fmt.Errorf("unsupported spec.teamAccess.sentinelMocks %q", *spec.SentinelMocks)
		}
	}

	return nil
}

func validateNotificationsSpec(specs []HCPTerraformWorkspaceNotificationSpec) error {
	allowedTriggers := []string{
		"run:created",
		"run:planning",
		"run:needs_attention",
		"run:applying",
		"run:completed",
		"run:errored",
		"assessment:drifted",
		"assessment:check_failure",
		"assessment:failed",
		"workspace:auto_destroy_reminder",
		"workspace:auto_destroy_run_results",
	}

	seen := make(map[string]struct{}, len(specs))
	for _, spec := range specs {
		if strings.TrimSpace(spec.Name) == "" {
			return fmt.Errorf("spec.notifications.name must not be blank")
		}

		if !contains([]string{"generic", "email", "slack", "microsoft-teams"}, spec.DestinationType) {
			return fmt.Errorf("unsupported spec.notifications.destinationType %q", spec.DestinationType)
		}

		key := strings.ToLower(spec.Name)
		if _, ok := seen[key]; ok {
			return fmt.Errorf("spec.notifications contains duplicate name %q", spec.Name)
		}

		seen[key] = struct{}{}

		if spec.DestinationType != "email" && (spec.URL == nil || strings.TrimSpace(*spec.URL) == "") {
			return fmt.Errorf("spec.notifications.url is required when spec.notifications.destinationType is %q", spec.DestinationType)
		}

		if spec.Token != nil && strings.TrimSpace(*spec.Token) == "" {
			return fmt.Errorf("spec.notifications.token must not be blank")
		}

		if err := validateUniqueNonEmptyStrings("spec.notifications.triggers", spec.Triggers); err != nil {
			return err
		}

		for _, trigger := range spec.Triggers {
			if !contains(allowedTriggers, trigger) {
				return fmt.Errorf("unsupported spec.notifications.trigger %q", trigger)
			}
		}
	}

	return nil
}

func explicitlyTrue(value *bool) bool {
	return value != nil && *value
}
