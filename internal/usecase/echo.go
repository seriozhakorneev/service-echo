package usecase

// EchoUseCase -.
type EchoUseCase struct {
	rewriter Rewriter
}

// New -.
func New(r Rewriter) *EchoUseCase {
	return &EchoUseCase{rewriter: r}
}

// Rewrite - rewrites data with rewrite rules.
func (uc *EchoUseCase) Rewrite(data any) any {
	return data
}
