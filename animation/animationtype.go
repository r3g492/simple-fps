package animation

type Type int

const (
	Idle Type = iota
	Move
	Attack
)

var typeName = map[Type]string{
	Idle:   "idle",
	Move:   "move",
	Attack: "attack",
}

func (at Type) String() string {
	return typeName[at]
}
