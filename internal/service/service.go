package service

import "github.com/ShekleinAleksey/survey/internal/repository"

type Service struct {
	SurveyService *SurveyService
}

func NewService(r repository.SurveyRepository) *Service {
	return &Service{
		SurveyService: NewSurveyService(r),
	}
}
