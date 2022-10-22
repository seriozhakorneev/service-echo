package rewriter

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRewriter_Rewrite(t *testing.T) {
	t.Parallel()

	r := Rewriter{
		name:  "pokemon",
		value: "pikachu",
		new:   "bulbasaur",
	}

	tests := []struct {
		input, expected []byte
	}{
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
	}

	for _, test := range tests {
		var data map[string]any

		err := json.Unmarshal(test.input, &data)
		if err != nil {
			t.Fatal("Unmarshal failed:", err)
		}

		r.Rewrite(data)

		test.input, err = json.Marshal(data)
		if err != nil {
			t.Fatal("Marshal failed:", err)
		}

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Fatalf("Expected object: %s\n Got: %s", test.expected, test.input)
		}
	}
}
