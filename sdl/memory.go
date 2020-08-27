package sdl

type v2MemoryAttributes v2ResourceAttributes

type v2ResourceMemory struct {
	Size       byteQuantity       `yaml:"size"`
	Attributes v2MemoryAttributes `yaml:"-"`
}
