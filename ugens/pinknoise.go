package ugens

// PinkNoise
type PinkNoise struct {
}

func (self PinkNoise) Rate(rate int8) *BaseNode {
	checkRate(rate)
	return NewNode("PinkNoise", rate, 0)
}
