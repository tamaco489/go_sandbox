package model

type Achievement struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

func NewAchievement(id uint32, name string) *Achievement {
	return &Achievement{
		ID:   id,
		Name: name,
	}
}

type AchievementList struct {
	Achievements []Achievement `json:"achievements"`
}

func NewAchievementList(achievements []Achievement) *AchievementList {
	return &AchievementList{
		Achievements: achievements,
	}
}
