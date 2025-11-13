package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	SurveyRepository *SurveyRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SurveyRepository: NewSurveyRepository(db),
	}
}
