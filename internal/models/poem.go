package models

type Poem struct {
	ID             int `gorm:"primary_key"`
	Poem           string
	Interpretation string
	ImageURL       string
	AudioURL       string
}
