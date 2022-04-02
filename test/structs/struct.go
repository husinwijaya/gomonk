package structs

type Struct1 struct {
}

type Struct2 struct {
}

type AnInt int64

type AnotherInterface interface {
	DoSomething(Struct1) Struct2
}
