package model

type Equipments [3]Equipment

type Equipment struct {
	Name           string
	Category       EquipmentCategory
	Tier           int
	StatusIncrease map[StatusName]StatusBuff
}

type EquipmentCategory string

// First slot
const (
	Hat    EquipmentCategory = "Hat"
	Gloves EquipmentCategory = "Gloves"
	Shoes  EquipmentCategory = "Shoes"
)

// Second slot
const (
	Bag     EquipmentCategory = "Bag"
	Badge   EquipmentCategory = "Badge"
	Hairpin EquipmentCategory = "Hairpin"
)

// Third slot
const (
	Charm    EquipmentCategory = "Charm"
	Watch    EquipmentCategory = "Watch"
	Necklace EquipmentCategory = "Necklace"
)
