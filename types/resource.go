package types

import (
	"fmt"
	"reflect"
)

type UnitType int

type Attribute struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value,omitempty"`
}

type Unit interface {
	String() string
	equals(Unit) bool
	add(Unit) error
	sub(Unit) error
	le(Unit) bool
}

type ResUnit interface {
	Equals(ResUnit) bool
	Add(unit ResUnit) bool
}

// ResourceUnits describes all available resources types for deployment/node etc
// if field is nil resource is not present in the given data-structure
type ResourceUnits struct {
	CPU     *CPU
	Memory  *Memory
	Storage *Storage
}

// Resources stores Unit details and Count value
type Resources struct {
	Resources ResourceUnits `json:"resources"`
	Count     uint32        `json:"count"`
}

// ResourceGroup is the interface that wraps GetName and GetResources methods
type ResourceGroup interface {
	GetName() string
	GetResources() []Resources
}

type CPU struct {
	Units      ResourceValue          `json:"units"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Memory struct {
	Size       ResourceValue          `json:"size"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Storage struct {
	Size       ResourceValue          `json:"size"`
	Attributes map[string]interface{} `json:"attributes"`
}

var _ Unit = (*CPU)(nil)
var _ Unit = (*Memory)(nil)
var _ Unit = (*Storage)(nil)

// AddUnit it rather searches for existing entry of the same type and sums values
// if type not found it appends
func (u ResourceUnits) Add(rhs ResourceUnits) (res ResourceUnits, err error) {
	res = u

	defer func() {
		if err != nil {
			res = ResourceUnits{}
		}
	}()

	if res.CPU != nil {
		if err = res.CPU.add(rhs.CPU); err != nil {
			return
		}
	} else {
		res.CPU = rhs.CPU
	}

	if res.Memory != nil {
		if err = res.Memory.add(rhs.Memory); err != nil {
			return
		}
	} else {
		res.Memory = rhs.Memory
	}

	if res.Storage != nil {
		if err = res.Storage.add(rhs.Storage); err != nil {
			return
		}
	} else {
		res.Storage = rhs.Storage
	}

	return
}

// Sub tbd
func (u ResourceUnits) Sub(rhs ResourceUnits) (res ResourceUnits, err error) {
	defer func() {
		if err != nil {
			res = ResourceUnits{}
		}
	}()

	if (u.CPU == nil && rhs.CPU != nil) ||
		(u.Memory == nil && rhs.Memory != nil) ||
		(u.Storage == nil && rhs.Storage != nil) {
		err = errCannotSub
		return
	}

	res = u

	if err = res.CPU.sub(rhs.CPU); err != nil {
		return
	}
	if err = res.Memory.sub(rhs.Memory); err != nil {
		return
	}
	if err = res.Storage.sub(rhs.Storage); err != nil {
		return
	}

	return
}

func (u ResourceUnits) Equals(rhs ResourceUnits) bool {
	return reflect.DeepEqual(u, rhs)
}

func (u CPU) String() string {
	return fmt.Sprintf("%v", u.Units)
}

func (u CPU) equals(other Unit) bool {
	rhs, valid := other.(CPU)
	if !valid {
		return false
	}

	if !u.Units.equals(rhs.Units) || len(u.Attributes) != len(rhs.Attributes) {
		return false
	}

	return reflect.DeepEqual(u.Attributes, rhs.Attributes)
}

func (u CPU) le(other Unit) bool {
	rhs, valid := other.(CPU)
	if !valid {
		return false
	}

	return u.Units.le(rhs.Units)
}

func (u CPU) add(other Unit) error {
	rhs, valid := other.(CPU)
	if !valid {
		return nil
	}

	res, err := u.Units.add(rhs.Units)
	if err != nil {
		return err
	}

	u.Units = res

	return nil
}

func (u CPU) sub(other Unit) error {
	rhs, valid := other.(CPU)
	if !valid {
		return nil
	}

	res, err := u.Units.sub(rhs.Units)
	if err != nil {
		return err
	}

	u.Units = res

	return nil
}

func (u Memory) String() string {
	return fmt.Sprintf("%v", u.Size)
}

func (u Memory) equals(other Unit) bool {
	rhs, valid := other.(Memory)
	if !valid {
		return false
	}

	if !u.Size.equals(rhs.Size) || len(u.Attributes) != len(rhs.Attributes) {
		return false
	}

	return reflect.DeepEqual(u.Attributes, rhs.Attributes)
}

func (u Memory) le(other Unit) bool {
	rhs, valid := other.(Memory)
	if !valid {
		return false
	}

	return u.Size.le(rhs.Size)
}

func (u Memory) add(other Unit) error {
	rhs, valid := other.(Memory)
	if !valid {
		return nil
	}

	res, err := u.Size.add(rhs.Size)
	if err != nil {
		return err
	}

	u.Size = res

	return nil
}

func (u Memory) sub(other Unit) error {
	rhs, valid := other.(Memory)
	if !valid {
		return nil
	}

	res, err := u.Size.sub(rhs.Size)
	if err != nil {
		return err
	}

	u.Size = res

	return nil
}

func (u Storage) String() string {
	return fmt.Sprintf("%v", u.Size)
}

func (u Storage) equals(other Unit) bool {
	rhs, valid := other.(Storage)
	if !valid {
		return false
	}

	if !u.Size.equals(rhs.Size) || len(u.Attributes) != len(rhs.Attributes) {
		return false
	}

	return reflect.DeepEqual(u.Attributes, rhs.Attributes)
}

func (u Storage) le(other Unit) bool {
	rhs, valid := other.(Storage)
	if !valid {
		return false
	}

	return u.Size.le(rhs.Size)
}

func (u Storage) add(other Unit) error {
	rhs, valid := other.(Storage)
	if !valid {
		return nil
	}

	res, err := u.Size.add(rhs.Size)
	if err != nil {
		return err
	}

	u.Size = res

	return nil
}

func (u Storage) sub(other Unit) error {
	rhs, valid := other.(Storage)
	if !valid {
		return nil
	}

	res, err := u.Size.sub(rhs.Size)
	if err != nil {
		return err
	}

	u.Size = res

	return nil
}
