package rewriter

import (
	"service-echo/config"
	"service-echo/internal/usecase"
)

// Rewriter -.
type Rewriter struct {
	name, value, new string
}

// New -.
func New(cfg config.Rewriter) usecase.Rewriter {
	if cfg.Active {
		return &Rewriter{
			cfg.Rules.Name,
			cfg.Rules.Value,
			cfg.Rules.New,
		}
	}

	return nil
}

// Rewrite - rewrites data with rewrite rules
func (r Rewriter) Rewrite(m map[string]any) {
	// start recursion-rewriter
	r.rec(m)
}

func (r Rewriter) arrReveal(a []any) {
	for _, el := range a {
		switch vv := el.(type) {
		case []any:
			r.arrReveal(vv)
		case map[string]any:
			r.rec(vv)
		}
	}
}

func (r Rewriter) rec(m map[string]any) {
	for key, value := range m {
		switch vv := value.(type) {
		case string:
			if key == r.name && vv == r.value {
				m[key] = r.new
			}
		case map[string]any:
			r.rec(vv)
		case []any:
			r.arrReveal(vv)
		}
	}
}
