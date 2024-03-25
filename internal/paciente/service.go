package paciente

import "desafioFinalGo/internal/domain"

type IService interface {
	GetByID(id int) (*domain.Paciente, error)
	Create(paciente domain.Paciente) (domain.Paciente, error)
	Update(id int, paciente domain.Paciente) (domain.Paciente, error)
	Patch(id int, paciente domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type Service struct {
	r IRepository
}

func NewService(r IRepository) IService {
	return &Service{r}
}

func (s *Service) GetByID(id int) (*domain.Paciente, error) {
	paciente, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return paciente, nil
}

func (s *Service) Create(paciente domain.Paciente) (domain.Paciente, error) {
	newPaciente, err := s.r.Create(paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return newPaciente, nil
}

func (s *Service) Update(id int, paciente domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Update(id, paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *Service) Patch(id int, paciente domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Patch(id, paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *Service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
