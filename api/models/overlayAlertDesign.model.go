package models

type AlertDesign struct {
	ID              uint `gorm:"type:integer; primaryKey"`
	AlertID         uint `gorm:"unique"`
	BackgroundColor string
	HighlightColor  string
	TextColor       string
	TextTemplate    string
	Border          uint8 `gorm:"type:tinyint(1)"`
	TextTickness    uint
	Duration        uint
	Font            string
}
