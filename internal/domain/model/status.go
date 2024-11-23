package model

import "fmt"

type Status map[StatusName]int

type StatusName string

const (
	MaxHP               StatusName = "MaxHP"
	ATK                 StatusName = "ATK"
	DEF                 StatusName = "DEF"
	Healing             StatusName = "Healing"
	Accuracy            StatusName = "Accuracy"
	Evasion             StatusName = "Evasion"
	Crit                StatusName = "Crit"       // CriticalRate
	CritRES             StatusName = "CritRES"    // CriticalResist
	CritDMG             StatusName = "CritDMG"    // CriticalDamage
	CritDMGRES          StatusName = "CritDMGRES" // CriticalDamageResist
	ArmourPenetration   StatusName = "ArmourPenetration"
	Stability           StatusName = "Stability"
	NormalAttackRange   StatusName = "NormalAttackRange"
	CCPower             StatusName = "CCPower" // CrowdControlPower
	CCRES               StatusName = "CCRES"   // CrowdControlResist
	CostRecovery        StatusName = "CostRecovery"
	ExplosiveEfficiency StatusName = "ExplosiveEfficiency"
	PiercingEfficiency  StatusName = "PiercingEfficiency"
	MysticEfficiency    StatusName = "MysticEfficiency"
	SiegeEfficiency     StatusName = "SiegeEfficiency"
)

var AllStatus = [20]StatusName{
	MaxHP, ATK, DEF, Healing, Accuracy, Evasion, Crit, CritRES, CritDMG, CritDMGRES,
	ArmourPenetration, Stability, NormalAttackRange, CCPower, CCRES, CostRecovery,
	ExplosiveEfficiency, PiercingEfficiency, MysticEfficiency, SiegeEfficiency,
}

var LevelDependentStatus = [4]StatusName{MaxHP, ATK, DEF, Healing}

type StatusInfo map[StatusName]StatusInfoItem

// StatusInfoItem is used to calculate the status of the character.
// Some statuses (MaxHP, ATK, DEF, and Healing) depend on the level and stars.
// Other statuses use their InitValue as is.
type StatusInfoItem struct {
	Name      StatusName
	InitValue int
	RateValue float64
}

func (si StatusInfoItem) Calc(level, star int) int {
	switch si.Name {
	case MaxHP, ATK, DEF, Healing:
		x := si.InitValue + int(float64(si.InitValue)*si.RateValue)*(level-1)
		if si.Name == DEF {
			return x
		}
		return int(float64(x) * starIncreaseRate[si.Name][star-1])
	default:
		panic(fmt.Sprintf("Status %s are not level-dependent.", si.Name))
	}
}

var starIncreaseRate = map[StatusName][5]float64{
	MaxHP:   {1, 1.05, 1.12, 1.21, 1.35},
	ATK:     {1, 1.1, 1.22, 1.36, 1.53},
	Healing: {1, 1.075, 1.175, 1.295, 1.445},
}

type StatusBuff struct {
	Name     StatusName
	AddValue int
	MulValue float64
}
