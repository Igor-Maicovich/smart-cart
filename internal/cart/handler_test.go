package cart

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddItem(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{
		AddFn: func(ctx context.Context, item Item) error {
			if item.Name == "" {
				return errors.New("name required")
			}
			return nil
		},
	}

	handler := NewHandler(mockService)

	router := gin.Default()
	router.POST("/cart", handler.AddItem)

	body := []byte(`{"name":"Apple","price":1.5,"quantity":2}`)
	req, _ := http.NewRequest("POST", "/cart", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "item added")
}

func TestGetCart(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{
		GetAllFn: func() ([]Item, error) {
			return []Item{
				{ID: 1, Name: "Apple", Price: 1.5, Quantity: 2},
				{ID: 2, Name: "Banana", Price: 0.9, Quantity: 3},
			}, nil
		},
	}

	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/cart", handler.GetCart)

	req, _ := http.NewRequest("GET", "/cart", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Apple")
	assert.Contains(t, w.Body.String(), "Banana")
}

func TestDeleteItem(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{
		DeleteFn: func(id int) error {
			if id == 0 {
				return errors.New("invalid id")
			}
			return nil
		},
	}

	handler := NewHandler(mockService)

	router := gin.Default()
	router.DELETE("/cart/:id", handler.DeleteItem)

	req, _ := http.NewRequest("DELETE", "/cart/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"deleted_id":1`)
}
