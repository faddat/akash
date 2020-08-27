package sdl

import (
	"sort"

	"github.com/pkg/errors"
	"github.com/xtgo/set"
	"gopkg.in/yaml.v3"
)

type v2CPUAttributes v2ResourceAttributes

type v2ResourceCPU struct {
	Units      cpuQuantity     `yaml:"units"`
	Attributes v2CPUAttributes `yaml:"attributes,omitempty"`
}

func (sdl *v2CPUAttributes) UnmarshalYAML(node *yaml.Node) error {
	attr := make(v2CPUAttributes)

	for i := 0; i+1 < len(node.Content); i += 2 {
		var value interface{}
		switch node.Content[i].Value {
		case "arch":
			var archSlice []string

			if err := node.Content[i+1].Decode(&archSlice); err != nil {
				return err
			}

			// remove duplicates if any
			archSlice = set.Strings(archSlice)

			// keep ordering stable
			sort.Strings(archSlice)

			value = archSlice
		default:
			return errors.Errorf("unsupported cpu attribute \"%s\"", node.Content[i].Value)
		}

		attr[node.Content[i].Value] = value
	}

	*sdl = attr

	return nil
}
