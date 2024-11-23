package model

type AttackType int
type DefenceType int

const (
	NormalAT AttackType = iota
	Explosive
	Piercing
	Mystic
	Siege
)

const (
	NormalDT DefenceType = iota
	Light
	Heavy
	Special
	Structure
)

const (
	resist    = 0.5
	normal    = 1
	effective = 1.5
	weak      = 2
)

var damageTable = [5][5]float64{
	// Normal, Light, Heavy, Special, Structure
	{normal, normal, normal, normal, normal},  // Normal
	{normal, weak, normal, resist, resist},    // Explosive
	{normal, resist, weak, normal, normal},    // Piercing
	{normal, normal, resist, weak, normal},    // Mystic
	{normal, normal, resist, effective, weak}, // Siege
}

func ArmourMatchUp(attack AttackType, defence DefenceType) float64 {
	return damageTable[attack][defence]
}
