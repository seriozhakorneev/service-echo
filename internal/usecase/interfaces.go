// Package usecase implements application business logic. Each logic group in own file.
package usecase

type (
	// Echo -.
	Echo interface {
		Reflect(any) any
	}
)
