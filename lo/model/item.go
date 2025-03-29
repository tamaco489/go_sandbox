package model

type Item struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

func NewItem(id uint32, name string) *Item {
	return &Item{
		ID:   id,
		Name: name,
	}
}

type ItemList struct {
	Items []Item `json:"items"`
}

func NewItemList(items []Item) *ItemList {
	return &ItemList{
		Items: items,
	}
}
