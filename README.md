# Alloy

Alloy holds shared manifest schema code used by Anvil tooling.

It is intended to provide a small, stable Go module for versioned manifest types and schema-oriented validation shared by `anvil` and `smith`.

Example `GitHubRepository` manifest:

```yaml
apiVersion: anvil.example.io/v1alpha1
kind: GitHubRepository
metadata:
  name: alloy
spec:
  owner: emkaytec
  name: alloy
  visibility: public
  description: Shared manifest schema code for Anvil
  autoInit: true
  defaultBranch: main
  topics:
    - anvil
    - manifests
  features:
    hasIssues: true
    hasProjects: false
    hasWiki: false
    # hasDownloads: true
  mergePolicy:
    allowSquashMerge: true
    allowMergeCommit: false
    allowRebaseMerge: true
    allowAutoMerge: true
    deleteBranchOnMerge: true
    # squashMergeCommitTitle: PR_TITLE
    # squashMergeCommitMessage: PR_BODY
    # mergeCommitTitle: PR_TITLE
    # mergeCommitMessage: PR_BODY
  # archived: false
  # initialization:
  #   gitignoreTemplate: Go
  #   licenseTemplate: mit
  #   isTemplate: false
  # securityAndAnalysis:
  #   advancedSecurity:
  #     status: enabled
  #   secretScanning:
  #     status: enabled
  #   secretScanningPushProtection:
  #     status: enabled
  # pages:
  #   buildType: legacy
  #   source:
  #     branch: main
  #     path: /docs
  # customProperties:
  #   - name: service
  #     value: alloy
  # branches:
  #   - name: main
  #     protection:
  #       enforceAdmins: true
  #       requiredLinearHistory: true
  #       requiredStatusChecks:
  #         strict: true
  #         checks:
  #           - context: ci/test
  #       pullRequestReviews:
  #         dismissStaleReviews: true
  #         requireCodeOwnerReviews: true
  #         requiredApprovingReviewCount: 1
```

Example `HCPTerraformWorkspace` manifest:

```yaml
apiVersion: anvil.example.io/v1alpha1
kind: HCPTerraformWorkspace
metadata:
  name: platform-workspace
spec:
  organization: emkaytec
  name: emkaytec-hcp-terraform
  projectID: prj-abc123
  description: Control plane workspace for shared HCP Terraform infrastructure
  terraformVersion: ~> 1.14.8
  workingDirectory: terraform
  executionMode: agent
  agentPoolID: apool-abc123
  allowDestroyPlan: true
  assessmentsEnabled: true
  autoApply: false
  autoApplyRunTrigger: true
  autoDestroyActivityDuration: 14d
  fileTriggersEnabled: true
  globalRemoteState: false
  queueAllRuns: false
  sourceName: anvil
  sourceURL: https://github.com/emkaytec/anvil
  speculativeEnabled: true
  sshKeyID: sshkey-abc123
  settingOverwrites:
    executionMode: true
    terraformVersion: true
  tags:
    - platform
    - hcp
  tagBindings:
    - key: env
      value: prod
    - key: service
      value: platform
  triggerPatterns:
    - terraform/**/*.tf
    - modules/**/*.tf
  triggerPrefixes:
    - terraform/
    - modules/
  remoteStateConsumerIDs:
    - ws-consumer-1
    - ws-consumer-2
  vcsRepo:
    identifier: emkaytec/hcp-terraform
    branch: main
    ingressSubmodules: false
    oauthTokenID: ot-abc123
    tagsRegex: ^v[0-9]+\.[0-9]+\.[0-9]+$
  variables:
    - key: AWS_REGION
      category: env
      value: us-east-1
      description: Default AWS region for remote runs
    - key: account_id
      category: terraform
      value: '"123456789012"'
      hcl: true
  variableSetIDs:
    - varset-abc123
  runTriggers:
    - sourceWorkspaceID: ws-network
  teamAccess:
    - teamName: platform
      access: admin
    - teamName: operators
      access: custom
      runs: apply
      variables: write
      stateVersions: read
      workspaceLocking: true
  notifications:
    - name: run-events
      destinationType: generic
      enabled: true
      url: https://hooks.example.com/tfc
      token: shared-secret
      triggers:
        - run:created
        - run:completed
        - run:errored
```
