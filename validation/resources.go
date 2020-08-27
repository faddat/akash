package validation

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	"github.com/ovrclk/akash/types"
)

var (
	ErrNoGroupsPresent = errors.New("validation: no groups present")
	ErrGroupEmptyName  = errors.New("validation: group has empty name")
)

// ValidateResourceList does basic validation for resources list
func ValidateResourceList(rlist types.ResourceGroup) error {
	return validateResourceList(defaultConfig, rlist)
}

func validateResourceLists(config ValConfig, rlists []types.ResourceGroup) error {

	if len(rlists) == 0 {
		return ErrNoGroupsPresent
	}

	if count := len(rlists); count > config.MaxGroupCount {
		return errors.Errorf("error: too many groups (%v > %v)", count, config.MaxGroupCount)
	}

	names := make(map[string]bool)

	for _, rlist := range rlists {

		if ok := names[rlist.GetName()]; ok {
			return errors.Errorf("error: duplicate name (%v)", rlist.GetName())
		}
		names[rlist.GetName()] = true

		if err := validateResourceList(config, rlist); err != nil {
			return err
		}
	}
	return nil
}

type resourceLimits struct {
	cpu     sdk.Uint
	memory  sdk.Uint
	storage sdk.Uint
}

func newLimits() resourceLimits {
	return resourceLimits{
		cpu:     sdk.ZeroUint(),
		memory:  sdk.ZeroUint(),
		storage: sdk.ZeroUint(),
	}
}

func (u *resourceLimits) add(rhs resourceLimits) {
	u.cpu.Add(rhs.cpu)
	u.memory.Add(rhs.memory)
	u.storage.Add(rhs.storage)
}

func (u *resourceLimits) mul(count uint32) {
	u.cpu.MulUint64(uint64(count))
	u.memory.MulUint64(uint64(count))
	u.storage.MulUint64(uint64(count))
}

func validateResourceList(config ValConfig, rlist types.ResourceGroup) error {
	if rlist.GetName() == "" {
		return ErrGroupEmptyName
	}

	units := rlist.GetResources()

	if count := len(units); count > config.MaxGroupUnits {
		return errors.Errorf("group %v: too many units (%v > %v)", rlist.GetName(), count, config.MaxGroupUnits)
	}

	limits := newLimits()

	for _, resource := range units {
		gLimits, err := validateResourceGroup(config, resource)
		if err != nil {
			return fmt.Errorf("group %v: %w", rlist.GetName(), err)
		}

		gLimits.mul(resource.Count)

		limits.add(gLimits)

		// TODO: validate pricing
		// if idx == 0 {
		// 	price = resource.Price
		// } else {
		// 	if resource.Price.Denom != price.Denom {
		// 		return fmt.Errorf("mixed denominations: (%v != %v)", price.Denom, resource.Price.Denom)
		// 	}
		// }
	}

	if limits.cpu.GT(sdk.NewUint(uint64(config.MaxGroupCPU))) || limits.cpu.LTE(sdk.ZeroUint()) {
		return errors.Errorf("group %v: invalid total cpu (%v > %v > %v fails)",
			rlist.GetName(), config.MaxGroupCPU, limits.cpu, 0)
	}

	if limits.memory.GT(sdk.NewUint(uint64(config.MaxGroupMemory))) || limits.memory.LTE(sdk.ZeroUint()) {
		return errors.Errorf("group %v: invalid total memory (%v > %v > %v fails)",
			rlist.GetName(), config.MaxGroupMemory, limits.memory, 0)
	}

	if limits.storage.GT(sdk.NewUint(uint64(config.MaxGroupStorage))) || limits.storage.LTE(sdk.ZeroUint()) {
		return errors.Errorf("group %v: invalid total storage (%v > %v > %v fails)",
			rlist.GetName(), config.MaxGroupStorage, limits.storage, 0)
	}

	return nil
}

func validateResourceGroup(config ValConfig, rg types.Resources) (resourceLimits, error) {
	limits, err := validateResourceUnit(config, rg.Resources)
	if err != nil {
		return resourceLimits{}, nil
	}

	if rg.Count > uint32(config.MaxUnitCount) || rg.Count < uint32(config.MinUnitCount) {
		return resourceLimits{}, errors.Errorf("error: invalid unit count (%v > %v > %v fails)",
			config.MaxUnitCount, rg.Count, config.MinUnitCount)
	}

	// TODO: validate pricing
	// if !rg.Price.IsPositive() {
	// 	return fmt.Errorf("error: invalid unit price (not positive fails)")
	// }

	return limits, nil
}

func validateResourceUnit(config ValConfig, units types.ResourceUnits) (resourceLimits, error) {
	limits := newLimits()

	if u := units.CPU; u != nil {
		if (u.Units.Value() > uint64(config.MaxUnitCPU)) || (u.Units.Value() < uint64(config.MinUnitCPU)) {
			return resourceLimits{}, errors.Errorf("error: invalid unit cpu (%v > %v > %v fails)",
				config.MaxUnitCPU, u.Units, config.MinUnitCPU)
		}
		limits.cpu.Add(sdk.NewUint(u.Units.Value()))
	}

	if u := units.Memory; u != nil {
		if (u.Size.Value() > uint64(config.MaxUnitMemory)) || (u.Size.Value() < uint64(config.MinUnitMemory)) {
			return resourceLimits{}, errors.Errorf("error: invalid unit memory (%v > %v > %v fails)",
				config.MaxUnitMemory, u.Size, config.MinUnitMemory)
		}
		limits.memory.Add(sdk.NewUint(u.Size.Value()))
	}

	if u := units.Storage; u != nil {
		if (u.Size.Value() > uint64(config.MaxUnitStorage)) || (u.Size.Value() < uint64(config.MinUnitStorage)) {
			return resourceLimits{}, errors.Errorf("error: invalid unit storage (%v > %v > %v fails)",
				config.MaxUnitStorage, u.Size, config.MinUnitStorage)
		}
		limits.storage.Add(sdk.NewUint(u.Size.Value()))
	}

	return limits, nil
}
