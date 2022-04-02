package test

import "gomonk/test/structs"

type IFace1 interface {
	A()
	B(b int) int
	C(c structs.Struct1) structs.Struct1
	D([]structs.Struct2) *[]structs.Struct2
	E([]*structs.Struct1) ([]*structs.Struct1, error)
}
