package pathgenerator

import "errors"

const chessMatrixSize = 8

var ErrNoRouteFound = errors.New("no route found between given squares")

func GenerateKnightPath(srcNotation, dstNotation string) ([]string, error) {
	src, err := NewSquareFromNotation(srcNotation)
	if err != nil {
		return nil, err
	}

	dst, err := NewSquareFromNotation(dstNotation)
	if err != nil {
		return nil, err
	}

	// generating path with the different attempts leading the destination.
	pathSteps := knightPathGenerator(src, dst, chessMatrixSize)

	maxLevel := len(pathSteps)

	// if we don't have a one top level that hit the destination , that means its not possible to move between this squares .
	// (despite its not very likely)
	if tops := pathSteps[maxLevel]; len(tops) != 1 {
		return nil, ErrNoRouteFound
	}

	for i := maxLevel; i != 0; i-- {
		validSquares := make([]Square, 0)

		for _, topLevelSquare := range pathSteps[i] {
			for _, v := range pathSteps[i-1] {
				// validating the moves on backwards if knight can actually perform this move for eliminate unused paths
				// caused by tryouts of path generator
				if validateKnightMove(topLevelSquare, v) {
					// it is very likely that we will have more than one path , but we returning only one of them (for now),
					// once we found one of them , we are breaking the loop . we only need the single path for each step .
					validSquares = append(validSquares, v)

					break // we can remove this break if we decide to support multiple routes in the future .
				}
			}
		}

		pathSteps[i-1] = validSquares
	}

	resp := make([]string, 0)

	for i := 0; i < len(pathSteps); i++ {
		for _, step := range pathSteps[i] {
			resp = append(resp, step.String())
		}
	}

	return resp, nil
}

func knightPathGenerator(src, dest Square, dimensionLength int8) map[int][]Square {
	visitedSquares := make(map[Square]struct{})
	levels := make(map[int][]Square)

	q := make([]Square, 0)
	q = append(q, src)

	for len(q) != 0 {
		sqr := q[0]
		q = q[1:] // we are removing the received square from the list

		if _, ok := visitedSquares[sqr]; !ok {
			visitedSquares[sqr] = struct{}{} // marking current square as visited

			// checking for the possible coordinates knight can move
			for _, v := range knightMoves {
				rank1 := sqr.rank + v.Rank
				file1 := sqr.file + v.File

				// verifying that possible coordinates are exists in the given matrix.
				if validateCoordinates(rank1, file1, dimensionLength) {
					q = append(q, Square{rank1, file1, sqr.dist + 1})

					levels[sqr.dist+1] = append(levels[sqr.dist+1], Square{rank1, file1, sqr.dist + 1})

					// we hit the destination on this level
					if (rank1 == dest.rank) && (file1 == dest.file) {
						// this level already found the destination , thats why lets drop the other remaining attempts.
						levels[sqr.dist+1] = []Square{{rank1, file1, sqr.dist + 1}}

						return levels
					}
				}
			}
		}
	}

	return levels
}
