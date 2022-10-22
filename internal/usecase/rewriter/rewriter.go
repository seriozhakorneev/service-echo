package rewriter

import (
	"service-echo/config"
)

// Rewriter -.
type Rewriter struct {
	name, value, new string
}

// New -.
func New(rules config.Rules) *Rewriter {
	return &Rewriter{
		rules.Name,
		rules.Value,
		rules.New,
	}
}

// Rewrite -.
func (r *Rewriter) Rewrite(m map[string]any) {
	for key, value := range m {
		switch vv := value.(type) {
		case string:
			if key == r.name &&
				vv == r.value {
				m[key] = r.new
			}
		case map[string]any:
			r.Rewrite(vv)
		case []any:
			for i, v := range vv {
				switch vvv := v.(type) {
				case string:
					if key == r.name &&
						vvv == r.value {
						vv[i] = r.new
					}
				case map[string]any:
					r.Rewrite(vvv)
				}
			}
		}
	}
}
