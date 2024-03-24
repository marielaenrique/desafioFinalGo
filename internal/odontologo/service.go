package odontologo

import "desafioFinalGo/internal/domain"

type IService interface {
	GetByID(id int) (*domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) (domain.Odontologo, error)
}

type Service struct {
	r IRepository
}

func NewService(r IRepository) IService {
	return &Service{r}
}

func (s *Service) GetByID(id int) (*domain.Odontologo, error) {
	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return odontologo, nil
}

func (s *Service) Create(odontologo domain.Odontologo) (domain.Odontologo, error) {
	newOdontologo, err := s.r.Create(odontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return newOdontologo, nil
}

func (s *Service) Update(id int, odontologo domain.Odontologo) (domain.Odontologo, error) {
	o, err := s.r.Update(id, odontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return o, nil
}
