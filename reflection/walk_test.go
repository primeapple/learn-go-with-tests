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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	containsNeedle := false
	for _, x := range haystack {
		if x == needle {
			containsNeedle = true
		}
	}

	if !containsNeedle {
		t.Errorf("expected %v to contain %q but didn't", haystack, needle)
	}
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
		{
			"maps",
			map[string]string{
				"Cow":   "Emma",
				"Sheep": "Bertha",
			},
			[]string{"Emma", "Bertha"},
		},
        {
            "function",
            func() (Profile, Profile) {
                return Profile{"London", 100}, Profile{"Berlin", 200}
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

	t.Run("maps", func(t *testing.T) {
		input := map[string]string{
			"Cow":   "Emma",
			"Sheep": "Bertha",
		}
		var got []string

		walk(input, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bertha")
		assertContains(t, got, "Emma")
	})

	t.Run("channels", func(t *testing.T) {
		inputChannel := make(chan Profile)
		go func() {
			inputChannel <- Profile{"Paris", 1000}
			inputChannel <- Profile{"Rome", 1001}
			close(inputChannel)
		}()

		var got []string
		want := []string{"Paris", "Rome"}

		walk(inputChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	})
}
