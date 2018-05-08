//Solves the English peg solitare board game
package main

import "fmt"

const N = 11 + 1 // length of a row(+1 for \n)

//The board must be surrounded by 2 illegal fileds in each direction
//so that the move() doesn't need to check the board boundaries
//periods represent illegeal fields,● are pegs, and ○ are holes.

var board = []rune(
	`...........
...........
....●●●....
....●●●....
..●●●●●●●..
..●●●○●●●..
..●●●●●●●..
....●●●....
....●●●....
...........
...........
`)

// center is the position of the center hole if there is a single one
//otherwise it is -1
var center int

func init() {
	n := 0
	for pos, field := range board {
		if field == '○' {
			center = pos
			n++
		}
	}
	if n != 1 {
		center = -1 //no single hole
	}
}

var moves int //number of times the move is called

//move tests if there is a peg at position pos that
//can jump over another peg direction dir. if the
//move is valid, it iss executed and move returns true.
//otherwise move returns false
func move(pos, dir int) bool {
	moves++
	if board[pos] == '●' && board[pos+dir] == '●' && board[pos+2*dir] == '○' {
		board[pos] = '○'
		board[pos+dir] = '○'
		board[pos+2*dir] = '●'
		return true
	}
	return false
}

//unmove reverts a previous executed valid move
func unmove(pos, dir int) {
	board[pos] = '●'
	board[pos+dir] = '●'
	board[pos+2*dir] = '○'
}

//solves tries to find a sequence of moves such that there is only one peg left at the end
//if the center is >=0 that last peg must be in the center position
//if a solution is found,solce prints the board after each move in a backward fashion
//ie the last board position is printed first, all the way back to the starting board position
func solve() bool {
	var last, n int
	for pos, field := range board {
		//try each board position
		if field == '●' {
			//found a peg
			for _, dir := range [...]int{-1, -N, +1, +N} {
				//try each direction
				if move(pos, dir) {
					//a valid move was found and executed,
					//see if this new board has a solution
					if solve() {
						unmove(pos, dir)
						fmt.Println(string(board))
						return true
					}
					unmove(pos, dir)
				}
			}
			last = pos
			n++
		}
	}
	//tried each possible move
	if n == 1 && (center < 0 || last == center) {
		//there's only one peg left
		fmt.Println(string(board))
		return true
	}
	//no solution found for this board
	return false
}

func main() {
	if !solve() {
		fmt.Println("no solution found")
	}
	fmt.Println(moves, "moves tried")
}
