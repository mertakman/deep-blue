package pathgenerator

// validateKnightMove checks if move from one square to another is actually valid.
func validateKnightMove(from, to Square) bool {
	for i := 0; i < len(knightMoves); i++ {
		rank1 := from.rank + knightMoves[i].Rank
		file1 := from.file + knightMoves[i].File

		if rank1 == to.rank && file1 == to.file {
			return true
		}
	}

	return false
}

// validateCoordinates checks if given coordinate couldbe actually exists on the matrix as a square.
func validateCoordinates(rank, file int8, dimensionLength int8) bool {
	return (rank >= 0 && rank < dimensionLength) && (file >= 0 && file < dimensionLength)
}

// knightMoves are 8 possible directions from knights current position .
var knightMoves = [8]knightMove{
	{Rank: 2, File: -1},
	{Rank: 2, File: 1},
	{Rank: -2, File: 1},
	{Rank: -2, File: -1},
	{Rank: 1, File: 2},
	{Rank: 1, File: -2},
	{Rank: -1, File: 2},
	{Rank: -1, File: -2},
}

type knightMove struct {
	Rank, File int8
}
