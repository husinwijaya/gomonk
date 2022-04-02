package test

import "gomonk/test/structs"

type IFace2 interface {
	D([]structs.Struct2, []structs.Struct2, []structs.Struct2) *[]structs.Struct2
	A()
	B(b int) int
	C(c1 structs.Struct1, c2 structs.AnInt, c3 structs.AnotherInterface) structs.Struct1
	E(_ []*structs.Struct1, _ []*structs.Struct1) ([]*structs.Struct1, error)
}
