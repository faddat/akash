package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestV2ResourceCPU_Valid(t *testing.T) {
	var stream = `
units: 0.1
attributes:
  arch:
    - amd64
`
	var p v2ResourceCPU

	arch := []string{
		"amd64",
	}

	err := yaml.Unmarshal([]byte(stream), &p)
	require.NoError(t, err)
	require.Equal(t, cpuQuantity(100), p.Units)
	require.NotNil(t, p.Attributes["arch"])
	require.Equal(t, arch, p.Attributes["arch"])
}

func TestV2ResourceCPU_ArchDuplicates(t *testing.T) {
	var stream = `
units: 0.1
attributes:
  arch:
    - amd64
    - amd64
`
	var p v2ResourceCPU

	arch := []string{
		"amd64",
	}

	err := yaml.Unmarshal([]byte(stream), &p)
	require.NoError(t, err)
	require.Equal(t, cpuQuantity(100), p.Units)
	require.NotNil(t, p.Attributes["arch"])
	require.Equal(t, arch, p.Attributes["arch"])
}
