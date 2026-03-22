package cart

import "context"

type ServiceInterface interface {
	AddItem(ctx context.Context, item Item) error
	GetAll() ([]Item, error)
	Update(id int, input Item) (Item, error)
	Delete(id int) error
}

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}
func (s *Service) AddItem(ctx context.Context, item Item) error {
	return s.repo.AddItem(ctx, item)
}

func (s *Service) GetAll() ([]Item, error) {
	return s.repo.GetAll()
}

func (s *Service) Update(id int, input Item) (Item, error) {
	return s.repo.Update(id, input)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
