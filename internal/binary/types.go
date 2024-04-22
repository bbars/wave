package binary

type (
	Int24  int32
	Uint24 uint32
)

type number interface {
	int8 | uint8 | int16 | uint16 | Int24 | Uint24 | int32 | uint32 | int64 | uint64 | float32 | float64
}
