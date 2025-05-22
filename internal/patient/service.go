package patient

import "context"

type Service interface {
	GetPatients(c context.Context, input Patient) ([]Patient, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetPatients(c context.Context, input Patient) ([]Patient, error) {
	patients, err := s.repo.Get(c, input)
	if err != nil {
		return nil, err
	}
	return patients, nil
}
