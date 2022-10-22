// Package usecase implements application business logic. Each logic group in own file.
package usecase

type (
	// Echo -.
	Echo interface {
		Rewrite(any) any
	}

	// Rewriter -.
	Rewriter interface {
		Rewrite(m map[string]interface{})
	}
)
