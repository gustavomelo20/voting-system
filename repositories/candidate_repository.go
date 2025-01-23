package repositories

import (
	"voting-system/models"

	"gorm.io/gorm"
)

type CandidateRepository interface {
	FindAll() ([]models.Candidate, error)
	FindByID(id uint) (*models.Candidate, error)
	AddVote(id uint) error
	SumVotes(id uint) (int, error)
}

type candidateRepository struct {
	db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) CandidateRepository {
	return &candidateRepository{db: db}
}

func (r *candidateRepository) FindAll() ([]models.Candidate, error) {
	var candidates []models.Candidate
	err := r.db.Find(&candidates).Error
	return candidates, err
}

func (r *candidateRepository) FindByID(id uint) (*models.Candidate, error) {
	var candidate models.Candidate
	err := r.db.First(&candidate, id).Error
	return &candidate, err
}

func (r *candidateRepository) AddVote(candidateID uint) error {
	vote := models.Votes{
		Candidate: int(candidateID),
	}

	if err := r.db.Create(&vote).Error; err != nil {
		return err
	}
	return nil
}

func (r *candidateRepository) SumVotes(id uint) (int, error) {
	var totalVotes int
	err := r.db.Model(&models.Votes{}).
		Where("candidate = ?", id).
		Select("COUNT(*)").
		Scan(&totalVotes).Error
	return totalVotes, err
}
