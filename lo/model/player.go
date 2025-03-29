package model

import "time"

type Player struct {
	ID              uint64          `json:"id"`
	ProviderUID     string          `json:"provider_uid"`
	Level           uint32          `json:"level"`
	Experience      uint32          `json:"experience"`
	Balance         uint32          `json:"balance"`
	AchievementList AchievementList `json:"achievement_list"`
	ItemList        ItemList        `json:"item_list"`
	Settings        map[string]bool `json:"settings"`
	Titles          map[string]bool `json:"titles"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

func NewPlayer(
	id uint64,
	providerUID string,
	level, experience, balance uint32,
	achievements []Achievement,
	items []Item,
	settings map[string]bool,
	titles map[string]bool,
) *Player {
	if settings == nil {
		settings = make(map[string]bool)
	}
	if titles == nil {
		titles = make(map[string]bool)
	}
	achievementList := NewAchievementList(achievements)
	itemList := NewItemList(items)
	now := time.Now()
	return &Player{
		ID:              id,
		ProviderUID:     providerUID,
		Level:           level,
		Experience:      experience,
		Balance:         balance,
		AchievementList: *achievementList,
		ItemList:        *itemList,
		Settings:        settings,
		Titles:          titles,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
