package pathgenerator

import "testing"

func TestText2SquareParse(t *testing.T) {
	m := map[string][2]int8{ // test dataset from
		"A1": {0, 0},
		"B2": {1, 1},
		"C3": {2, 2},
		"D4": {3, 3},
		"E5": {4, 4},
		"F6": {5, 5},
		"G7": {6, 6},
		"H8": {7, 7},
	}

	for k, v := range m {
		file, rank, err := TextToSquare(k)
		if err != nil {
			t.Errorf("TextToSquare parsing encountered with an error : %s ", err.Error())
		}

		if file != v[0] || rank != v[1] {
			t.Errorf("TextToSquare parsing returns wrong result : Expected %v , but got %v%v .", v, file, rank)
		}
	}

	if _, _, err := TextToSquare("好的"); err == nil {
		t.Errorf("TextToSquare not returning the error for longer inputs")
	}

	if _, _, err := TextToSquare("1a"); err == nil {
		t.Errorf("TextToSquare not returning the error for wrong inputs")
	}
}
