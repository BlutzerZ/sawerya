package models

type Alert struct {
	ID              int   `gorm:"type:integer; primaryKey"`
	EnableGif       uint8 `gorm:"type:tinyint(1)"`
	MinAmountNotify uint
	MinAmountGIF    uint
	Sound           string
	WordFilter      string
	UserID          uint `gorm:"unique"`
	AlertDesign     *AlertDesign
}
