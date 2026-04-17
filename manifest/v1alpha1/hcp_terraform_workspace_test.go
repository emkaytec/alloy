package v1alpha1

import "testing"

func TestNewHCPTerraformWorkspaceManifestUsesCurrentVersionAndKind(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
		},
	)

	if manifest.APIVersion != APIVersion {
		t.Fatalf("expected apiVersion %q, got %q", APIVersion, manifest.APIVersion)
	}

	if manifest.Kind != KindHCPTerraformWorkspace {
		t.Fatalf("expected kind %q, got %q", KindHCPTerraformWorkspace, manifest.Kind)
	}
}

func TestHCPTerraformWorkspaceManifestValidate(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:                "emkaytec",
			Name:                        "emkaytec-platform",
			ProjectID:                   stringPtr("prj-abc123"),
			Description:                 stringPtr("Control plane for shared HCP Terraform infrastructure"),
			TerraformVersion:            stringPtr("~> 1.14.8"),
			WorkingDirectory:            stringPtr("terraform"),
			ExecutionMode:               stringPtr("agent"),
			AgentPoolID:                 stringPtr("apool-abc123"),
			AllowDestroyPlan:            boolPtr(true),
			AssessmentsEnabled:          boolPtr(true),
			AutoApply:                   boolPtr(false),
			AutoApplyRunTrigger:         boolPtr(true),
			AutoDestroyActivityDuration: stringPtr("14d"),
			FileTriggersEnabled:         boolPtr(true),
			GlobalRemoteState:           boolPtr(false),
			QueueAllRuns:                boolPtr(false),
			SourceName:                  stringPtr("anvil"),
			SourceURL:                   stringPtr("https://github.com/emkaytec/anvil"),
			SpeculativeEnabled:          boolPtr(true),
			SSHKeyID:                    stringPtr("sshkey-abc123"),
			SettingOverwrites:           &HCPTerraformWorkspaceSettingOverwrites{ExecutionMode: boolPtr(true), TerraformVersion: boolPtr(true)},
			Tags:                        []string{"platform", "hcp"},
			TagBindings:                 []HCPTerraformWorkspaceTagBindingSpec{{Key: "env", Value: "prod"}, {Key: "service", Value: "platform"}},
			TriggerPatterns:             []string{"terraform/**/*.tf", "modules/**/*.tf"},
			TriggerPrefixes:             []string{"terraform/", "modules/"},
			RemoteStateConsumerIDs:      []string{"ws-consumer-1", "ws-consumer-2"},
			VariableSetIDs:              []string{"varset-abc123"},
			VCSRepo:                     &HCPTerraformWorkspaceVCSRepoSpec{Identifier: stringPtr("emkaytec/hcp-terraform"), OAuthTokenID: stringPtr("ot-abc123"), Branch: stringPtr("main"), IngressSubmodules: boolPtr(false), TagsRegex: stringPtr("^v[0-9]+\\.[0-9]+\\.[0-9]+$")},
			Variables:                   []HCPTerraformWorkspaceVariableSpec{{Key: "AWS_REGION", Category: "env", Value: "us-east-1"}, {Key: "account_id", Category: "terraform", Value: "\"123456789012\"", HCL: boolPtr(true)}},
			RunTriggers:                 []HCPTerraformWorkspaceRunTriggerSpec{{SourceWorkspaceID: "ws-upstream"}},
			TeamAccess:                  []HCPTerraformWorkspaceTeamAccessSpec{{TeamName: "platform", Access: "admin"}, {TeamName: "operators", Access: "custom", Runs: stringPtr("apply"), Variables: stringPtr("write"), StateVersions: stringPtr("read"), WorkspaceLocking: boolPtr(true)}},
			Notifications:               []HCPTerraformWorkspaceNotificationSpec{{Name: "run-events", DestinationType: "generic", Enabled: boolPtr(true), URL: stringPtr("https://hooks.example.com/tfc"), Token: stringPtr("shared-secret"), Triggers: []string{"run:created", "run:completed", "run:errored"}}},
		},
	)

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRequiresName(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
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

func TestHCPTerraformWorkspaceManifestValidateRejectsInvalidExecutionMode(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:  "emkaytec",
			Name:          "emkaytec-platform",
			ExecutionMode: stringPtr("serverless"),
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `unsupported spec.executionMode "serverless"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRequiresAgentPoolForAgentMode(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:  "emkaytec",
			Name:          "emkaytec-platform",
			ExecutionMode: stringPtr("agent"),
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.agentPoolID is required when spec.executionMode is "agent"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsAgentPoolOutsideAgentMode(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:  "emkaytec",
			Name:          "emkaytec-platform",
			ExecutionMode: stringPtr("remote"),
			AgentPoolID:   stringPtr("apool-abc123"),
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.agentPoolID is only supported when spec.executionMode is "agent"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsInvalidAutoDestroyActivityDuration(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:                "emkaytec",
			Name:                        "emkaytec-platform",
			AutoDestroyActivityDuration: stringPtr("two-weeks"),
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `unsupported spec.autoDestroyActivityDuration "two-weeks"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsConflictingRemoteStateModes(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:       "emkaytec",
			Name:               "emkaytec-platform",
			GlobalRemoteState:  boolPtr(true),
			ProjectRemoteState: boolPtr(true),
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.globalRemoteState and spec.projectRemoteState cannot both be true" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsRemoteStateConsumersWhenGlobalEnabled(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:           "emkaytec",
			Name:                   "emkaytec-platform",
			GlobalRemoteState:      boolPtr(true),
			RemoteStateConsumerIDs: []string{"ws-consumer"},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.remoteStateConsumerIDs cannot be set when spec.globalRemoteState is true" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRequiresVCSRepoIdentifierAndOAuthToken(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			VCSRepo:      &HCPTerraformWorkspaceVCSRepoSpec{},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.vcsRepo.identifier is required" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateAllowsEmptyVCSBranch(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			VCSRepo: &HCPTerraformWorkspaceVCSRepoSpec{
				Identifier:   stringPtr("emkaytec/hcp-terraform"),
				OAuthTokenID: stringPtr("ot-abc123"),
				Branch:       stringPtr(""),
			},
		},
	)

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsBlankTags(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			Tags:         []string{"platform", " "},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.tags must not contain blank values" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsDuplicateTagBindings(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			TagBindings: []HCPTerraformWorkspaceTagBindingSpec{
				{Key: "env", Value: "prod"},
				{Key: "env", Value: "prod"},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.tagBindings contains duplicate binding "env"="prod"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsEnvVariableWithHCL(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			Variables: []HCPTerraformWorkspaceVariableSpec{
				{
					Key:      "AWS_REGION",
					Category: "env",
					Value:    "us-east-1",
					HCL:      boolPtr(true),
				},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "spec.variables.hcl is only supported for terraform variables" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsDuplicateVariables(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			Variables: []HCPTerraformWorkspaceVariableSpec{
				{Key: "AWS_REGION", Category: "env", Value: "us-east-1"},
				{Key: "AWS_REGION", Category: "env", Value: "us-west-2"},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.variables contains duplicate variable "env:AWS_REGION"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsDuplicateRunTriggers(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			RunTriggers: []HCPTerraformWorkspaceRunTriggerSpec{
				{SourceWorkspaceID: "ws-upstream"},
				{SourceWorkspaceID: "ws-upstream"},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.runTriggers contains duplicate sourceWorkspaceID "ws-upstream"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsCustomPermissionsOutsideCustomAccess(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			TeamAccess: []HCPTerraformWorkspaceTeamAccessSpec{
				{
					TeamName: "operators",
					Access:   "write",
					Runs:     stringPtr("apply"),
				},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.teamAccess custom permissions require spec.teamAccess.access to be "custom"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsInvalidNotificationTrigger(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization: "emkaytec",
			Name:         "emkaytec-platform",
			Notifications: []HCPTerraformWorkspaceNotificationSpec{
				{
					Name:            "run-events",
					DestinationType: "generic",
					URL:             stringPtr("https://hooks.example.com/tfc"),
					Triggers:        []string{"run:created", "run:paused"},
				},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `unsupported spec.notifications.trigger "run:paused"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateRejectsNotificationsForLocalMode(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:  "emkaytec",
			Name:          "emkaytec-platform",
			ExecutionMode: stringPtr("local"),
			Notifications: []HCPTerraformWorkspaceNotificationSpec{
				{
					Name:            "run-events",
					DestinationType: "generic",
					URL:             stringPtr("https://hooks.example.com/tfc"),
				},
			},
		},
	)

	err := manifest.Validate()
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `spec.notifications is not supported when spec.executionMode is "local"` {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHCPTerraformWorkspaceManifestValidateAllowsExplicitEmptyOptionalStrings(t *testing.T) {
	t.Parallel()

	manifest := NewHCPTerraformWorkspaceManifest(
		Metadata{Name: "platform-workspace"},
		HCPTerraformWorkspaceSpec{
			Organization:                "emkaytec",
			Name:                        "emkaytec-platform",
			ProjectID:                   stringPtr(""),
			Description:                 stringPtr(""),
			TerraformVersion:            stringPtr(""),
			WorkingDirectory:            stringPtr(""),
			ExecutionMode:               stringPtr(""),
			AutoDestroyAt:               stringPtr(""),
			AutoDestroyActivityDuration: stringPtr(""),
			SourceName:                  stringPtr(""),
			SourceURL:                   stringPtr(""),
			SSHKeyID:                    stringPtr(""),
		},
	)

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}
