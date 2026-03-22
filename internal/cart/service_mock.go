package cart

import (
	"context"
)

type MockService struct {
	AddFn    func(ctx context.Context, item Item) error
	GetAllFn func() ([]Item, error)
	UpdateFn func(id int, input Item) (Item, error)
	DeleteFn func(id int) error
}

func (m *MockService) AddItem(ctx context.Context, item Item) error {
	if m.AddFn != nil {
		return m.AddFn(ctx, item)
	}
	return nil
}

func (m *MockService) GetAll() ([]Item, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn()
	}
	return nil, nil
}

func (m *MockService) Update(id int, input Item) (Item, error) {
	if m.UpdateFn != nil {
		return m.UpdateFn(id, input)
	}
	return input, nil
}

func (m *MockService) Delete(id int) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(id)
	}
	return nil
}
