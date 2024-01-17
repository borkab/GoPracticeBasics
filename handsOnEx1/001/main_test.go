package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Area of a square", func(t *testing.T) {
		sq := Square{
			5.0,
		}

		got := sq.Area()
		want := 25.0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Area of a circle", func(t *testing.T) {
		cir := Circle{
			4,
		}

		got := cir.Area()
		want := 50.26548245743669

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestInfo(t *testing.T) {

	checkInfo := func(t testing.TB, sha Shape, want float64) {
		t.Helper()
		got := Info(sha)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Info of a square", func(t *testing.T) {
		sq := Square{
			5,
		}
		checkInfo(t, sq, 25.0)
	})

	t.Run("Info of a circle", func(t *testing.T) {
		cir := Circle{
			4,
		}
		checkInfo(t, cir, 50.26548245743669)
	})
}
