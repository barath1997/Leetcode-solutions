func validTicTacToe(board []string) bool {

	full := ""
	xWins, oWins := 0, 0
	spaces := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			full += string(board[i][j])
			if string(board[i][j]) == " " {
                spaces++
            }
		}
	}

	if spaces == 9 {
		return true
	}

	for i := 0; i < 9; i += 3 {
		if (full[i] == full[i+1]) && (full[i+1] == full[i+2]) {
			if full[i] == 'X' {
				xWins++
			} else if full[i] == 'O' {
				oWins++
			}
		}
	}

	for j := 0; j < 3; j++ {
		if (full[j] == full[j+3]) && (full[j+3] == full[j+6]) {
			if full[j] == 'X' {
				xWins++
			} else if full[j] == 'O' {
				oWins++
			}
		}
	}

	if (full[0] == full[4]) && (full[4] == full[8]) {
		if full[0] == 'X' {
			xWins++
		} else if full[0] == 'O' {
			oWins++
		}
	}

	if (full[2] == full[4]) && (full[4] == full[6]) {
		if full[2] == 'X' {
			xWins++
		} else if full[2] == 'O' {
			oWins++
		}
	}

	xCount, oCount := 0,0

	for i := 8; i >= 0; i-- {
		if full[i] == 'X' {
			xCount++
		}
		if full[i] == 'O' {
			oCount++
		}
	}

	if oCount > xCount {
		return false
	}
	if xCount-oCount > 1 {
		return false
	}

	if xWins > 2 || oWins > 2 {
		return false
	}
	if xWins > 0 && oWins > 0 {
		return false
	}
	if xWins > 0 && xCount == oCount {
		return false
	}
	if oWins > 0 && oCount < xCount {
		return false
	}

	return true
}