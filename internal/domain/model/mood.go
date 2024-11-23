package model

type Terrain string

const (
	City   Terrain = "City"
	Desert Terrain = "Desert"
	Indoor Terrain = "Indoor"
)

type TerrainSynergy map[Terrain]Mood

func NewTerrainSynergy(city, desert, indoor Mood) TerrainSynergy {
	return TerrainSynergy{
		City:   city,
		Desert: desert,
		Indoor: indoor,
	}
}

type Mood struct {
	Evaluation           string
	DamageMultiplier     float64
	BlockRateBehindCover float64
	ChanceToIgnoreBlock  float64
}

var (
	MoodSS = Mood{"SS", 1.3, 0.75, 0.75}
	MoodS  = Mood{"S", 1.2, 0.6, 0.6}
	MoodA  = Mood{"A", 1.1, 0.45, 0.45}
	MoodB  = Mood{"B", 1, 0.3, 0.3}
	MoodC  = Mood{"C", 0.9, 0.15, 0.15}
	MoodD  = Mood{"D", 0.8, 0, 0}
)
