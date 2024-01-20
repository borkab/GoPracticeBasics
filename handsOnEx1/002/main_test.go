package main

import (
	"testing"
)

func TestSpeak(t *testing.T) {

	t.Run("Person", func(t *testing.T) {
		per := Person{
			"Joe",
			"Blake",
		}

		got := per.Speak()
		want := "Joe Blake"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("SecretAgent", func(t *testing.T) {
		sec := SecretAgent{
			true,
			Person{
				"Julianna",
				"Crain",
			},
		}

		got := sec.Speak()
		want := "Julianna Crain has  a licence to kill"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

		// if licence != true {
		// 	fmt.Println("You do not have a Licence to Kill")
		// }

	})
}
