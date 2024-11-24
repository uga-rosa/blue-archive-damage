package model

// UniqueItem is a struct that represents a unique item.
// At tier 1, it has only status.
// At tier 2, BasicSkill improves.
type UniqueItem struct {
	Name       string
	Tier       int
	Status     Status
	BasicSkill BasicSkill
}
