package lib

type Int interface {
    int | int8 | int16 | int32 | int64
}

type Uint interface {
    uint | uint8 | uint16 | uint32 | uint64
}

type Integers interface {
    Int | Uint
}

type Float interface {
    float32 | float64
}

type Complex interface {
    complex64 | complex128
}

type DecimalsT interface {
    Float | Complex
}

type NumbersT interface {
    Float | Complex | Int | Uint
}
