package repodata

import "gopkg.in/yaml.v3"

type ModuleMDDefaults struct {
	Document string `yaml:"document"`
	Version  int    `yaml:"version"`
	Data     struct {
		Module   string              `yaml:"module"`
		Stream   string              `yaml:"stream"`
		Profiles map[string][]string `yaml:"profiles"`
	} `yaml:"data"`
}

type ModuleMD struct {
	Document string `yaml:"document"`
	Version  int    `yaml:"version"`
	Data     struct {
		Name        string `yaml:"name"`
		Stream      string `yaml:"stream"`
		Version     int64  `yaml:"version"`
		Context     string `yaml:"context"`
		Arch        string `yaml:"arch"`
		Summary     string `yaml:"summary"`
		Description string `yaml:"description"`
		License     struct {
			Module  []string `yaml:"module"`
			Content []string `yaml:"content"`
		} `yaml:"license"`

		Dependencies []struct {
			Buildrequires struct {
				Platform []string `yaml:"platform"`
			} `yaml:"buildrequires"`
			Requires struct {
				Platform []string `yaml:"platform"`
			} `yaml:"requires"`
		} `yaml:"dependencies"`

		Profiles map[string]struct {
			Rpms []string `yaml:"rpms"`
		} `yaml:"profiles"`

		Components struct {
			Rpms map[string]struct {
				Rationale string   `yaml:"rationale"`
				Ref       string   `yaml:"ref"`
				Arches    []string `yaml:"arches"`
			} `yaml:"rpms"`
		} `yaml:"components"`
		Artifacts struct {
			Rpms []string `yaml:"rpms"`
		} `yaml:"artifacts"`
	} `yaml:"data"`
}

type ModuleItem struct {
	ModuleMD *ModuleMD
	Defaults *ModuleMDDefaults
}

func (m *ModuleItem) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&m.Defaults); err == nil {
		return nil
	}
	return value.Decode(&m.ModuleMD)
}
