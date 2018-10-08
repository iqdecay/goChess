package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ChessPieceType struct {
	name string
	// if the vertical variation of a move is i, and the horizontal one is j,
	// then if i*j is not in "moves", then it is not a legal move for the piece
	moves     []int
	asciiRep  string // used for board representation in CLI
	originalX []int  // the original x positions of a ChessPiece
	originalY []int  // the original y positions of a ChessPiece
}

type ChessPiece struct {
	chesspiecetype ChessPieceType
	colour         string
	id             int // each piece is identified by its unique id
	number         int // over the number of piece of a certain type
}

type Colour struct {
	colour string
}

func Change(c Colour) Colour {
	if c.colour == "white" {
		c.colour = "black"
	} else {
		c.colour = "white"
	}
	return c
}

func TranslateMove(userMove string, lettersToInt map[byte]int) [4]int {
	a := lettersToInt[userMove[0]]
	b, _ := strconv.Atoi(string(userMove[1]))
	c := lettersToInt[userMove[2]]
	d, _ := strconv.Atoi(string(userMove[3]))
	return [4]int{a, b, c, d}
}

func GetUserInput(c Colour) (userMove string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the next %s move, then press Enter:\n", c.colour)
	userMove, err := reader.ReadString('\n')
	if err != nil {
		return GetUserInput(c)
	} else {
		return userMove
	}
}

//func playMove(move [4]int,board [8][8]int,boardRep [8][8]string, chessGame []ChessPiece) {

func main() {
	//--------------------- GAME INITIALIZATION------------------------------
	// We create the board that will hold the position of the pieces
	board := [8][8]int{}
	// We create its string representation
	boardRep := [8][8]string{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			boardRep[i][j] = "_"

		}
	}
	//Creating the complex moves
	var squareList = []int{}
	for i := 1; i < 9; i++ {
		squareList = append(squareList, i*i, -i*i)
	}
	queenList := append(squareList, 0)

	//Create the mapping between moves and coordinates
	lettersToInt := make(map[byte]int)
	letters := [8]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	for index, letter := range letters {
		lettersToInt[letter] = index
	}

	//Creating the different types of pieces
	bishop := ChessPieceType{"bishop", squareList, "B", []int{2, 5}, []int{0, 7}}
	knight := ChessPieceType{"knight", []int{-2, 2}, "H", []int{1, 6}, []int{0, 7}}
	king := ChessPieceType{"king", []int{-1, 0, 1}, "K", []int{4}, []int{0, 7}}
	queen := ChessPieceType{"queen", queenList, "Q", []int{3}, []int{0, 7}}
	tower := ChessPieceType{"tower", []int{0}, "T", []int{0, 7}, []int{0, 7}}
	pawn := ChessPieceType{"pawn", []int{-1, 0, 1}, "P", []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{1, 6}}

	//Each pieces appears a certain number of time
	chessSet := make(map[int][]ChessPieceType)
	chessSet[1] = []ChessPieceType{queen, king}
	chessSet[2] = []ChessPieceType{bishop, knight, tower}
	chessSet[8] = []ChessPieceType{pawn}

	//We create the game with the previously defined variables
	chessGame := make(map[int]ChessPiece)
	for number, pieceSlice := range chessSet {
		for i := 0; i < len(pieceSlice); i++ {
			pieceType := pieceSlice[i]
			for j := 0; j < number; j++ {
				id_white := len(chessGame) + 1
				id_black := id_white + 1
				//We create each piece for one colour
				chessGame[id_white] = ChessPiece{pieceType, "white", id_white, j}
				chessGame[id_black] = ChessPiece{pieceType, "black", id_black, j}
			}
		}
	}
	// After building the set, we must place it on the board
	for _, piece := range chessGame {
		id := piece.id
		number := piece.number
		pieceType := piece.chesspiecetype
		colour := piece.colour
		// Depending on its number, we place it differently
		j := pieceType.originalX[number]
		// Depending  on the colour, it is either on the top or bottom row
		var colourId int
		if colour == "white" {
			colourId = 1
		} else {
			colourId = 0
		}
		i := pieceType.originalY[colourId]
		boardRep[i][j] = pieceType.asciiRep
		// if the piece is white, put it in lowercase
		if colour == "white" {
			boardRep[i][j] = "w" + strings.ToLower(pieceType.asciiRep)
		}
		board[i][j] = id
	}

	//--------------------------- BEGINNING THE ACTUAL GAME ---------------------------------------

	continueGame := true          // Will be false whenever there is a checkmate
	turnColour := Colour{"white"} // Holds the color of the next player to play a move
	for continueGame {
		/*
			The structure will be as follow :
				°play move
					if move illegal:  // for now a move is legal iff
					it can perform it, if it places you in a check
					situation, it is not mentioned
					The legality of the move should be checked all at once
						return to °
					if move is eating a piece :
						eat the piece (remove eaten piece)
					update position due to move
				if move now is a check
					say "check"
					if checkmate :
						continueGame = false
				turnColour changes color
		*/
		userMove := GetUserInput(turnColour)
		coordinateMove := TranslateMove(userMove, lettersToInt)
		playMove(coordinateMove)
		//if ThereIsCheckmate(board) {
		//continueGame = false
		//fmt.Println("%s won the game !", turnColour)
		//}
		turnColour = Change(turnColour)

	}
}
