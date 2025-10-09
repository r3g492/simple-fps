package enemy

type AnimationType int

const (
	Idle AnimationType = iota
	Move
	Attack
)

var stateName = map[AnimationType]string{
	Idle:   "idle",
	Move:   "move",
	Attack: "attack",
}

func (ss AnimationType) String() string {
	return stateName[ss]
}
