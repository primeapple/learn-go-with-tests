package walk

import (
	"reflect"
	"testing"
)

func TestWalker(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Toni"},
			[]string{"Toni"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
                City string
			}{"Toni", "Berlin"},
			[]string{"Toni", "Berlin"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, expected %v", got, test.ExpectedCalls)
			}
		})
	}
}
