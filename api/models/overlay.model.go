package models

type Alert struct {
	ID              int `gorm:"type:integer; primaryKey"`
	EnableGif       bool
	MinAmountNotify uint
	MinAmountGIF    uint
	Sound           string
	WordFilter      string
	UserID          uint `gorm:"unique"`
	AlertDesign     *AlertDesign
}

type AlertDesign struct {
	ID              uint `gorm:"type:integer; primaryKey"`
	AlertID         uint `gorm:"unique"`
	BackgroundColor string
	HighlightColor  string
	TextColor       string
	TextTemplate    string
	Border          bool
	TextTickness    uint
	Duration        uint
	Font            string
}
