package model

// UniqueWeapon is a struct that represents a unique weapon.
// At 1 star, it has only status.
// At 2 star, EnhancedSkill improves.
// At 3 star, Terrain synergy in one area improves.
type UniqueWeapon struct {
	Name            string
	Star            int
	Level           int
	Status          Status
	EnhancedSkill   EnhancedSkill
	SynergyOverride SynergyOverride
}

type SynergyOverride struct {
	Terrain Terrain
	Mood    Mood
}
