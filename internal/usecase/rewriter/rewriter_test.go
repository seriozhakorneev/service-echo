package rewriter

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"service-echo/config"
	"service-echo/internal/usecase"
)

type rewriterTest struct{ input, expected []byte }

var (
	rewriter Rewriter
	tests    []rewriterTest
)

func TestMain(m *testing.M) {

	rewriter = Rewriter{
		name:  "pokemon",
		value: "pikachu",
		new:   "bulbasaur",
	}

	tests = []rewriterTest{
		{
			[]byte(`{"pokemon":"pikachu"}`),
			[]byte(`{"pokemon":"bulbasaur"}`),
		},
		{
			[]byte(`{"field":"pikachu"}`),
			[]byte(`{"field":"pikachu"}`),
		},
		{
			[]byte(`{"pokemon":"value"}`),
			[]byte(`{"pokemon":"value"}`),
		},
		{
			[]byte(`{"field":{"pokemon":"pikachu"}}`),
			[]byte(`{"field":{"pokemon":"bulbasaur"}}`),
		},
		{
			[]byte(`{"field":[{"pokemon":"pikachu"}]}`),
			[]byte(`{"field":[{"pokemon":"bulbasaur"}]}`),
		},
		{
			[]byte(`{"field":[{"pokemon":"pikachu"},{"pokemon":"pikachu"}]}`),
			[]byte(`{"field":[{"pokemon":"bulbasaur"},{"pokemon":"bulbasaur"}]}`),
		},
		{
			[]byte(`{"field":[[[{"pokemon":"pikachu"}],[{"pokemon":"pikachu"}]]]}`),
			[]byte(`{"field":[[[{"pokemon":"bulbasaur"}],[{"pokemon":"bulbasaur"}]]]}`),
		},
	}

	code := m.Run()
	os.Exit(code)
}

func TestRewriter_New(t *testing.T) {
	t.Parallel()

	expNilInterface := usecase.Rewriter(nil)
	rInterface := New(config.Rewriter{Active: false})

	if rInterface != expNilInterface {
		t.Fatalf("Expected: %v, got: %v", expNilInterface, rInterface)
	}

	rInterface = New(config.Rewriter{Active: true})
	if rInterface == nil {
		t.Fatalf("Expected: %v, got: nil", rInterface)
	}
}

func TestRewriter_Rewrite(t *testing.T) {
	t.Parallel()

	for _, test := range tests {
		var data map[string]any

		err := json.Unmarshal(test.input, &data)
		if err != nil {
			t.Fatal("Unmarshal failed:", err)
		}

		rewriter.Rewrite(data)

		test.input, err = json.Marshal(data)
		if err != nil {
			t.Fatal("Marshal failed:", err)
		}

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Fatalf("Expected object: %s\n Got: %s", test.expected, test.input)
		}
	}
}

func BenchmarkRewriterRewrite(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("input %s", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var data map[string]any

				err := json.Unmarshal(test.input, &data)
				if err != nil {
					b.Fatal("Unmarshal failed:", err)
				}

				rewriter.Rewrite(data)
			}
		})
	}
}
