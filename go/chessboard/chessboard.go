package chessboard

// File represents vertical columns of squares and stores if a square is occupied by a piece
type File []bool

// Chessboard represents horizontal rows of squares and contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0

	for _, v := range cb[file] {
		if v {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	index := rank - 1

	if rank < 1 || rank > 8 {
		return count
	}

	for _, r := range cb {
		if r[index] {
			count++
		}

	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for _, v := range cb {
		for range v {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for _, v := range cb {
		for _, v := range v {
			if v {
				count++
			}
		}
	}

	return count
}
