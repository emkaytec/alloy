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
