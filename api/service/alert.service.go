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

func (s *AlertService) GetAlertByUserID(ID int) (models.Alert, error) {
	alert, err := s.repository.FindAlertByUserID(ID)

	return alert, err
}

func (s *AlertService) UpdateAlert(req *dto.UpdateAlertRequest) error {
	ID := 2 // soon need update with session
	alert := models.Alert{
		EnableGif:       req.EnableGif,
		MinAmountNotify: req.MinAmountNotify,
		MinAmountGIF:    req.MinAmountGIF,
		Sound:           req.Sound,
		WordFilter:      req.WordFilter,
	}

	err := s.repository.UpdateAlert(uint(ID), &alert)

	return err
}

func (s *AlertService) UpdateAlertDesign(req *dto.UpdateAlertDesignRequest) error {
	ID := 2 // soon need update with session
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
