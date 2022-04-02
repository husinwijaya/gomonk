# gomonk

gomonk is a go code generator that generate implementation stub of an interface that you can easily override.

# Usage

```
gomonk [go source file] [interface name] [outout file] <custom package name>
```

for example, the following input file:

```go
package test

import "gomonk/test/structs"

type IFace2 interface {
	D([]structs.Struct2, []structs.Struct2, []structs.Struct2) *[]structs.Struct2
	A()
	B(b int) int
	C(c1 structs.Struct1, c2 structs.AnInt, c3 structs.AnotherInterface) structs.Struct1
	E(_ []*structs.Struct1, _ []*structs.Struct1) ([]*structs.Struct1, error)
}
```

with following command:

```bash
gomonk test/sample.go IFace2 test/out/output.go custompackagename
```

will output:

```go
package custompackagename

import "gomonk/test/structs"

type IFace2Impl struct {
	OnD func([]structs.Struct2, []structs.Struct2, []structs.Struct2) *[]structs.Struct2
	OnA func()
	OnB func(b int) int
	OnC func(c1 structs.Struct1, c2 structs.AnInt, c3 structs.AnotherInterface) structs.Struct1
	OnE func(_ []*structs.Struct1, _ []*structs.Struct1) ([]*structs.Struct1, error)
}

func (i *IFace2Impl) D(p0 []structs.Struct2, p1 []structs.Struct2, p2 []structs.Struct2) *[]structs.Struct2 {
	return i.OnD(p0, p1, p2)
}

func (i *IFace2Impl) A() {
	i.OnA()
}

func (i *IFace2Impl) B(b int) int {
	return i.OnB(b)
}

func (i *IFace2Impl) C(c1 structs.Struct1, c2 structs.AnInt, c3 structs.AnotherInterface) structs.Struct1 {
	return i.OnC(c1, c2, c3)
}

func (i *IFace2Impl) E(p0 []*structs.Struct1, p1 []*structs.Struct1) ([]*structs.Struct1, error) {
	return i.OnE(p0, p1)
}
```
