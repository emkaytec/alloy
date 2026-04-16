package v1alpha1

import "gopkg.in/yaml.v3"

const APIVersion = "anvil.example.io/v1alpha1"

type Metadata struct {
	Name string `yaml:"name"`
}

type Envelope struct {
	APIVersion string    `yaml:"apiVersion"`
	Kind       string    `yaml:"kind"`
	Metadata   Metadata  `yaml:"metadata"`
	Spec       yaml.Node `yaml:"spec"`
}
