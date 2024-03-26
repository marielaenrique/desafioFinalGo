package turno

import "desafioFinalGo/internal/domain"

type IService interface {
	GetByID(id int) (*domain.Turno, error)
	Create(turno domain.Turno) (domain.Turno, error)
	Update(id int, turno domain.Turno) (domain.Turno, error)
	Patch(id int, turno domain.Turno) (domain.Turno, error)
}

type Service struct {
	r IRepository
}

func NewService(r IRepository) IService {
	return &Service{r}
}

func (s *Service) GetByID(id int) (*domain.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) Create(turno domain.Turno) (domain.Turno, error) {

	newTurno, err := s.r.Create(turno)
	if err != nil {
		return domain.Turno{}, err
	}
	return newTurno, nil
}

func (s *Service) Update(id int, turno domain.Turno) (domain.Turno, error) {
	t, err := s.r.Update(id, turno)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (s *Service) Patch(id int, turno domain.Turno) (domain.Turno, error) {
	t, err := s.r.Patch(id, turno)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}
