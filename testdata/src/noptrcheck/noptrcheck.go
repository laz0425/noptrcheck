package noptrcheck

var LiteralKey map[int]int // OK

type LiteralStruct struct {
	Name    string
	Comment string
}

var NoPtrFieldKey map[LiteralStruct]bool // OK

var PtrKey map[*int]int // want `a pointer type cannot be used as a map key`

type MyNumber *int

var AliasPtrKey map[MyNumber]int // want `a pointer type cannot be used as a map key`

type MyStruct struct {
	Name    string
	Comment *string
}

var WithPtrField map[MyStruct]bool // want `MyStruct has pointer fields and cannot be used as a map key`

type SomeStruct struct {
	ID       int
	MyStruct MyStruct
}

var NestedWithPtrField map[SomeStruct]bool // want `SomeStruct has pointer fields and cannot be used as a map key`
