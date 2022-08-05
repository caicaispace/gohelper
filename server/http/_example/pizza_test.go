package _example_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/caicaispace/gohelper/server/http/_example"
)

func TestPizzasHandler(t *testing.T) {
	var pizzas = make(_example.Pizzas, 0)
	pizzas = append(pizzas, _example.Pizza{ID: 1, Name: "Pizza 1", Price: 10})
	tt := []struct {
		name       string
		method     string
		input      *_example.Pizzas
		want       string
		statusCode int
	}{
		{
			name:       "without pizzas",
			method:     http.MethodGet,
			input:      &_example.Pizzas{},
			want:       "Error: No pizzas found",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "with pizzas",
			method:     http.MethodGet,
			input:      &pizzas,
			want:       `[{"id":1,"name":"Foo","price":10}]`,
			statusCode: http.StatusOK,
		},
		{
			name:       "with bad method",
			method:     http.MethodPost,
			input:      &_example.Pizzas{},
			want:       "Method not allowed",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/orders", nil)
			responseRecorder := httptest.NewRecorder()

			_example.PizzasHandler{tc.input}.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
