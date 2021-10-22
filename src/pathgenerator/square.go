package pathgenerator

import "fmt"

type Square struct {
	rank int8
	file int8
	dist int
}

// Equals compares given square with the receiver one.
func (s *Square) Equals(comparedS *Square) bool {
	if (s.file == comparedS.file) && (s.rank == comparedS.rank) && (s.dist == comparedS.dist) {
		return true
	}

	return false
}

// Formats the square on letters and ranks format (A1,C3,F5 etc.)
func (s *Square) String() string {
	return fmt.Sprintf("%s%d", fileLettersMapping[s.file], s.rank+1)
}

func NewSquare(file, rank int8) (sq Square) {
	sq.file = file
	sq.rank = rank

	return sq
}

// NewSquareFromNotation generates a square struct from strings with square and letters format.
func NewSquareFromNotation(squareNotation string) (sq Square, err error) {
	file, rank, err := TextToSquare(squareNotation)
	if err != nil {
		return sq, err
	}

	sq.file = file
	sq.rank = rank

	return sq, nil
}

// transforming file numbers to letter representations.
var fileLettersMapping = map[int8]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
}
