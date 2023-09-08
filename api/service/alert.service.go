package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/api/repository"
)

type AlertService struct {
	repository *repository.AlertRepository
}

func NewAlertService() *AlertService {
	repo := repository.NewAlertRepository()
	return &AlertService{
		repository: repo,
	}
}

func (s *AlertService) GetAlertByUserID(ID uint) (models.Alert, error) {
	alert, err := s.repository.FindAlertByUserID(ID)

	return alert, err
}

func (s *AlertService) UpdateAlert(ID uint, req *dto.UpdateAlertRequest) error {
	alert := models.Alert{
		EnableGif:       req.EnableGif,
		MinAmountNotify: req.MinAmountNotify,
		MinAmountGIF:    req.MinAmountGIF,
		Sound:           req.Sound,
		WordFilter:      req.WordFilter,
	}

	err := s.repository.UpdateAlert(ID, &alert)

	return err
}

func (s *AlertService) UpdateAlertDesign(ID uint, req *dto.UpdateAlertDesignRequest) error {
	alert, err := s.repository.FindAlertByUserID(ID)
	if err != nil {
		return err
	}

	// convert
	alert.AlertDesign.BackgroundColor = req.BackgroundColor
	alert.AlertDesign.HighlightColor = req.HighlightColor
	alert.AlertDesign.TextColor = req.TextColor
	alert.AlertDesign.TextTemplate = req.TextTemplate
	alert.AlertDesign.Border = req.Border
	alert.AlertDesign.TextTickness = req.TextTickness
	alert.AlertDesign.Duration = req.TextTickness
	alert.AlertDesign.Duration = req.Duration
	alert.AlertDesign.Font = req.Font

	s.repository.UpdateAlertDesign(&alert)

	return err
}
