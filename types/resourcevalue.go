package types

import (
	"fmt"
	"math/big"

	"github.com/pkg/errors"
)

var (
	errOverflow  = errors.Errorf("resource value overflow")
	errCannotSub = errors.Errorf("cannot sub resources when lhs does not have same units as rhs")
)

/*
ResourceValue the big point of this small change is to ensure math operations on resources
not resulting with negative value which panic on unsigned types as well as overflow which leads to panic too
instead reasonable error is returned.
Each resource using this type as value can take extra advantage of it to check upper bounds
For example in SDL v1 CPU units were handled as uint32 and operation like math.MaxUint32 + 2
would cause application to panic. But nowadays
	const CPULimit = math.MaxUint32

	func (c *CPU) add(rhs CPU) error {
		res, err := c.Units.add(rhs.Units)
		if err != nil {
			return err
		}

		if res.Units.Value() > CPULimit {
			return ErrOverflow
		}

		c.Units = res

		return nil
	}
*/
type ResourceValue struct {
	val big.Int
}

func NewResourceValue(val uint64) ResourceValue {
	res := ResourceValue{}
	res.val.SetUint64(val)

	return res
}

func (v ResourceValue) Value() uint64 {
	return v.val.Uint64()
}

func (v ResourceValue) equals(rhs ResourceValue) bool {
	return v.val.Cmp(&rhs.val) == 0
}

func (v ResourceValue) le(rhs ResourceValue) bool {
	res := v.val.Cmp(&rhs.val)
	return res == 0 || res == -1
}

func (v ResourceValue) lt(rhs ResourceValue) bool {
	res := v.val.Cmp(&rhs.val)
	return res == -1
}

func (v ResourceValue) ge(rhs ResourceValue) bool {
	res := v.val.Cmp(&rhs.val)
	return res == 0 || res == 1
}

func (v ResourceValue) gt(rhs ResourceValue) bool {
	res := v.val.Cmp(&rhs.val)
	return res == 1
}

func (v ResourceValue) add(rhs ResourceValue) (ResourceValue, error) {
	var res big.Int

	_ = res.Add(&v.val, &rhs.val)

	if res.Sign() == -1 {
		return ResourceValue{}, errOverflow
	}

	return ResourceValue{res}, nil
}

func (v ResourceValue) sub(rhs ResourceValue) (ResourceValue, error) {
	var res big.Int

	_ = res.Sub(&v.val, &rhs.val)

	if res.Sign() == -1 {
		return ResourceValue{}, errCannotSub
	}

	return ResourceValue{res}, nil
}

func (v ResourceValue) MarshalJSON() ([]byte, error) {
	return []byte(v.val.String()), nil
}

func (v *ResourceValue) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}

	var z big.Int
	_, ok := z.SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}

	v.val = z

	return nil
}
