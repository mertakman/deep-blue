package pathgenerator

import "testing"

func TestValidateKnightMove(t *testing.T) {
	route1, _ := NewSquareFromNotation("A8")
	route2, _ := NewSquareFromNotation("H3")

	if validateKnightMove(route1, route2) {
		t.Error("validateKnightMove validating the invalid moves of knight . please check the rules")
	}

	route1, _ = NewSquareFromNotation("C7")
	route2, _ = NewSquareFromNotation("B5")

	if !validateKnightMove(route1, route2) {
		t.Error("validateKnightMove returning false on valid events ,please check the rules")
	}
}

func TestValidateCoordinates(t *testing.T) {
	if validateCoordinates(7, 7, 5) {
		t.Error("Validator returning true for coordinates that outside of matrix.")
	}

	if !validateCoordinates(7, 7, 8) {
		t.Error("Validator returning false for coordinates that inside of matrix.")
	}
}
