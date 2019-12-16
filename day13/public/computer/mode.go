package computer

type ParameterMode int

const (
	Position  ParameterMode = 0
	Immediate ParameterMode = 1
	Relative  ParameterMode = 2
)

func (m ParameterMode) String() string {
	switch m {
	case Position:
		return "position"
	case Immediate:
		return "immediate"
	case Relative:
		return "relative"
	default:
		return "wat"
	}
}
