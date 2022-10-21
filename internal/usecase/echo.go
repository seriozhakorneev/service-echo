package usecase

// EchoUseCase -.
type EchoUseCase struct {
	// TODO: rewriter might be here
}

// New -.
func New() *EchoUseCase {
	return &EchoUseCase{}
}

// Reflect - returns data back unchanged or rewritten if rewrite rules active.
func (uc *EchoUseCase) Reflect(data any) any {
	return data
}
