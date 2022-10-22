package usecase_test

import (
	"testing"

	"service-echo/internal/usecase"
)

type data any

func TestEchoReflect(t *testing.T) {
	t.Parallel()

	tests := []data{
		"string",
		27,
		27.0,
		false,
		`{"a": 10,"b": 2 }`,
	}

	uc := usecase.EchoUseCase{}
	for _, expectedData := range tests {
		result := uc.Reflect(expectedData)

		if result != expectedData {
			t.Fatalf("Expected data: %v, Got: %v", expectedData, result)
		}
	}
}
