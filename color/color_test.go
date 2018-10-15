package color

import (
	"testing"
)

func TestColoringBlack(t *testing.T) {

	color := Brack
	message := Coloring(color, color.String())

	if message != "[30mBrack[0m" {
		t.Fatalf("error: %s", color.String())
	}

	t.Log(message)
	t.Log(message)
}

func TestColoringRed(t *testing.T) {

	color := Red
	message := Coloring(color, color.String())

	if message != "[31mRed[0m" {
		t.Fatalf("error: %s", color.String())
	}

	t.Log(message)
}

func TestColoringGreen(t *testing.T) {

	color := Green
	message := Coloring(color, color.String())

	if message != "[32mGreen[0m" {
		t.Fatalf("error: %s", color.String())
	}

	t.Log(message)
}

func TestColoringYellow(t *testing.T) {

	color := Yellow
	message := Coloring(color, color.String())

	if message != "[33mYellow[0m" {
		t.Fatalf("error: %s", color.String())
	}

	t.Log(message)
}

func TestColoringBlue(t *testing.T) {

	color := Blue
	message := Coloring(color, color.String())

	if message != "[34mBlue[0m" {
		t.Fatalf("error: %s", color.String())
	}

	t.Log(message)
}

func TestColoringMagenta(t *testing.T) {

	color := Magenta
	message := Coloring(color, color.String())

	if message != "[35mMagenta[0m" {
		t.Fatalf("error: %s", color.String())
	}

	t.Log(message)
}

func TestColoringCyan(t *testing.T) {

	color := Cyan
	message := Coloring(color, color.String())

	if message != "[36mCyan[0m" {
		t.Fatalf("error: %s", color.String())
	}

	t.Log(message)
}
