package services

import (
	"errors"
	"voting-system/models"
	"voting-system/repositories"
)

type CandidateService interface {
	GetAllCandidates() ([]models.Candidate, error)
	VoteForCandidate(id uint) error
	GetCandidateVotes(id uint) (int, error)
}

type candidateService struct {
	repo repositories.CandidateRepository
}

func NewCandidateService(repo repositories.CandidateRepository) CandidateService {
	return &candidateService{repo: repo}
}

func (s *candidateService) GetAllCandidates() ([]models.Candidate, error) {
	return s.repo.FindAll()
}

func (s *candidateService) VoteForCandidate(id uint) error {
	candidate, err := s.repo.FindByID(id)
	if err != nil || candidate == nil {
		return errors.New("candidato não encontrado")
	}

	go func() {
		err := s.repo.AddVote(id)
		if err != nil {
			errors.New("Erro ao adicionar voto para o candidato")
		}
	}()

	return nil
}

func (s *candidateService) GetCandidateVotes(id uint) (int, error) {
	return s.repo.SumVotes(id)
}
