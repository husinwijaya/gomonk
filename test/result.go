package test

import "gomonk/test/structs"

type IFace1Monk struct {
	OnA func()
	OnB func(b int) int
	OnC func(c structs.Struct1) structs.Struct1
	OnD func([]structs.Struct2) *[]structs.Struct2
	OnE func([]*structs.Struct1) ([]*structs.Struct1, error)
}

func (m *IFace1Monk) A() {
	m.OnA()
}

func (m *IFace1Monk) B(b int) int {
	return m.OnB(b)
}

func (m *IFace1Monk) C(c structs.Struct1) structs.Struct1 {
	return m.OnC(c)
}

func (m *IFace1Monk) D(struct2s []structs.Struct2) *[]structs.Struct2 {
	return m.OnD(struct2s)
}

func (m *IFace1Monk) E(struct1s []*structs.Struct1) ([]*structs.Struct1, error) {
	return m.OnE(struct1s)
}
