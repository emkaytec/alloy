package v1alpha1

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	APIVersion           = "anvil.example.io/v1alpha1"
	KindGitHubRepository = "GitHubRepository"
)

type Metadata struct {
	Name string `yaml:"name"`
}

type Envelope struct {
	APIVersion string    `yaml:"apiVersion"`
	Kind       string    `yaml:"kind"`
	Metadata   Metadata  `yaml:"metadata"`
	Spec       yaml.Node `yaml:"spec"`
}

type GitHubRepositorySpec struct {
	Owner       string `yaml:"owner"`
	Name        string `yaml:"name"`
	Visibility  string `yaml:"visibility"`
	Description string `yaml:"description"`
	AutoInit    bool   `yaml:"autoInit"`
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

	return nil
}
