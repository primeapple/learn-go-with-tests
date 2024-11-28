package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

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
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Toni", 29},
			[]string{"Toni"},
		},
		{
			"nested struct",
			Person{"Toni", Profile{"London", 33}},
			[]string{"Toni", "London"},
		},
		{
			"pointers",
			&Person{"Toni", Profile{"London", 33}},
			[]string{"Toni", "London"},
		},
		{
			"slices",
			[]Profile{
				{"London", 3},
				{"Berlin", 23},
			},
			[]string{"London", "Berlin"},
		},
		{
			"arrays",
			[2]Profile{
				{"London", 3},
				{"Berlin", 23},
			},
			[]string{"London", "Berlin"},
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
