package sdl

type v2StorageAttributes v2ResourceAttributes

type v2ResourceStorage struct {
	Size       byteQuantity        `yaml:"size"`
	Attributes v2StorageAttributes `yaml:"attributes,omitempty"`
}
