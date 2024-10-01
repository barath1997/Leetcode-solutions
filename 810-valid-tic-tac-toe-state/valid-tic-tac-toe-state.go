/*Types of invalid states:

Number of O > Number of X
Number of X - Number of O > 1
greater than 2 Win states for X or O
Both X and O have a win state
X wins but Number of X == Number of O
O wins but Number of O < Number of X
*/

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

    // if all are spaces  ["   ","   ","   "]
	if spaces == 9 {
		return true
	}
    
    // move row by row and then compare elements in each row
	for i := 0; i < 9; i += 3 {
		if (full[i] == full[i+1]) && (full[i+1] == full[i+2]) {
			if full[i] == 'X' {
				xWins++
			} else if full[i] == 'O' {
				oWins++
			}
		}
	}
    
    // move column by column and then compare elements in each column
	for j := 0; j < 3; j++ {
		if (full[j] == full[j+3]) && (full[j+3] == full[j+6]) {
			if full[j] == 'X' {
				xWins++
			} else if full[j] == 'O' {
				oWins++
			}
		}
	}
    
    // compare elements in leading diagonal
	if (full[0] == full[4]) && (full[4] == full[8]) {
		if full[0] == 'X' {
			xWins++
		} else if full[0] == 'O' {
			oWins++
		}
	}
    
    // comapre elements in the secondary diagonal
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

    // validating conditions inputs like ["O  ","   ","   "] :starting with O or generally always X >= O because we start with X.  
	if oCount > xCount {
		return false
	}

    // if the player puts X consecutively without giving chance to O
	if xCount-oCount > 1 {
		return false
	}
    
    // we can have at max 2 win states for both X and O in a game
	if xWins > 2 || oWins > 2 {
		return false
	}

    // both X and O cant win in one game, there can be only one winner, then game stops
	if xWins > 0 && oWins > 0 {
		return false
	}

    // if X won then his X count should be more than O because we start with X if both are equal then there should not be a win state.
	if xWins > 0 && xCount == oCount {
		return false
	}

    // if O wins , then O count should be at max equal to X count, if O wins then game finishes then O count becomes equal to X count , because we start with X.
	if oWins > 0 && oCount < xCount {
		return false
	}
    
    // if all conditions fails , return true
	return true
}