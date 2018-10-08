package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
type ChessPieceType struct {
	name string
	// if the vertical variation of a move is i, and the horizontal one is j,
	// then if i*j is not in "moves", then it is not a legal move for the piece
	moves      []int
	ascii_rep  string // used for board representation in CLI
	original_x []int  // the original x positions of a ChessPiece
	original_y []int  // the original y positions of a ChessPiece
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

func Change(c Colour)  {
	if c.colour == "white" {
		c.colour = "black"
	} else{
		c.colour = "white"
	}
}

func GetUserInput(c Colour) (userMove string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the next %s move, then press Enter:\n",c.colour)
	userMove, err = reader.ReadString('\n')
	if err != nil {
		return GetUserInput(c.colour)
	}else {
		return userMove
	}
}





func main() {
	//--------------------- GAME INITIALIZATION------------------------------
	// We create the board that will hold the position of the pieces
	board := [8][8]int{}
	// We create its string representation
	board_rep := [8][8]string{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board_rep[i][j] = "_"

		}
	}
	//Creating the complex moves
	var square_list = []int{}
	for i := 1; i < 9; i++ {
			square_list = append(square_list, i*i, -i*i)
		}
	queen_list := append(square_list, 0)

	//Create the mapping between moves and coordinates
	lettersToInt := make(map[byte]int])
	letters := [8]byte{'a','b','c','d','e','f','g','h'}
	for index, letter := range letters {
		lettersToInt[letter] = index
	}

	

	//Creating the different types of pieces
	bishop := ChessPieceType{"bishop", square_list, "B", []int{2, 5}, []int{0, 7}}
	knight := ChessPieceType{"knight", []int{-2, 2}, "H", []int{1, 6}, []int{0, 7}}
	king := ChessPieceType{"king", []int{-1, 0, 1}, "K", []int{4}, []int{0, 7}}
	queen := ChessPieceType{"queen", queen_list, "Q", []int{3}, []int{0, 7}}
	tower := ChessPieceType{"tower", []int{0}, "T", []int{0, 7}, []int{0, 7}}
	pawn := ChessPieceType{"pawn", []int{-1, 0, 1}, "P", []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{1, 6}}

	//Each pieces appears a certain number of time
	chess_set := make(map[int][]ChessPieceType)
	chess_set[1] = []ChessPieceType{queen, king}
	chess_set[2] = []ChessPieceType{bishop, knight, tower}
	chess_set[8] = []ChessPieceType{pawn}

	//We create the game with the previously defined variables
	chess_game := make(map[int]ChessPiece)
	for number, piece_slice := range chess_set {
		for i := 0; i < len(piece_slice); i++ {
			piece_type := piece_slice[i]
			for j := 0; j < number; j++ {
				id_white := len(chess_game) + 1
				id_black := id_white + 1
				//We create each piece for one colour
				chess_game[id_white] = ChessPiece{piece_type, "white", id_white, j}
				chess_game[id_black] = ChessPiece{piece_type, "black", id_black, j}
			}
		}
	}
	// After building the set, we must place it on the board
	for _, piece := range chess_game {
		id := piece.id
		number := piece.number
		piece_type := piece.chesspiecetype
		colour := piece.colour
		// Depending on its number, we place it differently
		j := piece_type.original_x[number]
		// Depending  on the colour, it is either on the top or bottom row
		var colour_id int
		if colour == "white" {
			colour_id = 1
		} else {
			colour_id = 0
		}
		i := piece_type.original_y[colour_id]
		board_rep[i][j] = piece_type.ascii_rep
		// if the piece is white, put it in lowercase
		if colour == "white" {
			board_rep[i][j] = "w"+strings.ToLower(piece_type.ascii_rep)
		}
		board[i][j] = id
	}

	//--------------------------- BEGINNING THE ACTUAL GAME ---------------------------------------
	continueGame := true // Will be false whenever there is a checkmate
	turnColour := Colour{"white"}// Holds the color of the next player to play a move
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
		userMove = GetUserInput(turnColour)
		playMove(userMove)
		if ThereIsCheckmate(board) {
			continueGame = false
			fmt.Println("%s won the game !", turnColour)
		}
		Change(turnColour)
		



	}
}
