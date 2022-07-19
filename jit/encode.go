package main

import (
	"fmt"

	zserio "github.com/woven-planet/go-zserio"
	"github.com/woven-planet/go-zserio/internal/generator"
	"github.com/woven-planet/go-zserio/internal/model"
)

type M struct {
	Model *model.Model

	Pkg  string
	Name string
	Vals map[string]any
}

// MarshalZserio implements the zserio.Marshaler interface.
func (v *M) MarshalZserio(w zserio.Writer) error {
	// var err error

	// if err = ztype.WriteString(w, v.Street); err != nil {
	// 	return err
	// }

	// if err = ztype.WriteBool(w, v.Number != nil); err != nil {
	// 	return err
	// }
	// if v.Number != nil {

	// 	if err = ztype.WriteUint8(w, (*v.Number)); err != nil {
	// 		return err
	// 	}
	// }

	// _ = err // to avoid "declared but not used" warning
	// return nil

	pkg := v.Model.Packages["contacts"]

	addrStruct := pkg.Structs["Address"]

	for _, f := range addrStruct.Fields {
		nativeType, err := generator.GoNativeType(pkg, f.Type)
		if err != nil {
			return err
		}
		fmt.Println(nativeType)
	}

	return nil
}

// UnmarshalZserio implements the zserio.Unmarshaler interface.
func (v *M) UnmarshalZserio(r zserio.Reader) error {
	return nil
}

func (v *M) ZserioBitSize(bitPosition int) (int, error) {
	return 0, nil
}
