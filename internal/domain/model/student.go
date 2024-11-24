package model

type Student struct {
	Level int
	Star  int

	AttackType     AttackType
	DefenceType    DefenceType
	TerrainSynergy TerrainSynergy

	IsStriker     bool
	SupportStatus Status

	StatusInfo  StatusInfo
	BasicStatus Status

	Equipments   Equipments
	UniqueWeapon UniqueWeapon
	UniqueItem   UniqueItem

	ExSkill       ExSkill
	BasicSkill    BasicSkill
	EnhancedSkill EnhancedSkill
	SubSkill      SubSkill

	Bond         Bond
	OtherCostume []Student

	Buffs []StatusBuff
}

func (s *Student) SetLevel(level int) {
	s.Level = level
	for _, status := range LevelDependentStatus {
		s.BasicStatus[status] = s.StatusInfo[status].Calc(s.Level, s.Star)
	}
}

func (s *Student) SetStar(star int) {
	s.Star = star
	for _, status := range StarDependentStatus {
		s.BasicStatus[status] = s.StatusInfo[status].Calc(s.Level, s.Star)
	}
}

func (s *Student) GetStatus() Status {
	status := make(Status)
	buffMap := make(map[StatusName]StatusBuff)
	for _, buff := range s.Buffs {
		old := buffMap[buff.Name]
		old.AddValue += buff.AddValue
		old.MulValue += buff.MulValue
		buffMap[buff.Name] = old
	}
	for _, name := range AllStatus {
		x := s.BasicStatus[name] + s.Bond.Status[name] + s.UniqueWeapon.Status[name] + s.UniqueItem.Status[name]
		m := 1.0
		for _, equip := range s.Equipments {
			x += equip.StatusIncrease[name].AddValue
			m += equip.StatusIncrease[name].MulValue
		}
		for _, otherCostume := range s.OtherCostume {
			x += otherCostume.Bond.Status[name]
		}
		x += buffMap[name].AddValue
		m += buffMap[name].MulValue
		status[name] = int(float64(x) * m)
	}
	return status
}
