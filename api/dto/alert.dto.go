package dto

type UpdateAlertRequest struct {
	EnableGif       uint8  `json:"enableGIF" binding:"min=0,max=1"`
	MinAmountNotify uint   `json:"minAmountNotify" binding:"required,min=4"`
	MinAmountGIF    uint   `json:"minAmountGIF" binding:"required,min=4"`
	Sound           string `json:"sound" binding:"required"`
	WordFilter      string `json:"wordFilter" binding:"required"`
}

type UpdateAlertDesignRequest struct {
	BackgroundColor string `json:"backgroundColor" binding:"required,min=7,max=7"`
	HighlightColor  string `json:"highlightColor" binding:"required,min=7,max=7"`
	TextColor       string `json:"textColor" binding:"required,min=7,max=7"`
	TextTemplate    string `json:"textTemplate" binding:"required,min=1"`
	Border          uint8  `json:"border" binding:"required,min=1"`
	TextTickness    uint   `json:"textTickness" binding:"required,min=1"`
	Duration        uint   `json:"duration" binding:"required,min=1"`
	Font            string `json:"font" binding:"required,min=3"`
}
