package main

import (
	"fmt"
	"strconv"
)

type Color = int

const (
	White = iota
	Black = iota
)

type Game struct {
	board  Board
	pieces map[int]Piece
}

type PieceKind struct {
	symbol string // used for board representation in CLI
}

type Piece struct {
	kind   PieceKind
	colour Color
	id     int // each piece is identified by its unique id
}

type Board struct {
	Positions [8][8]int
}

func PrintBoard(b [8][8]string) {
	for _, line := range b {
		fmt.Println(line)
	}
}
func TranslateMove(userMove string, lettersToInt map[byte]int) ([4]int, bool) {
	a := lettersToInt[userMove[0]]
	a--
	b, _ := strconv.Atoi(string(userMove[1]))
	b--
	c := lettersToInt[userMove[2]]
	c--
	d, _ := strconv.Atoi(string(userMove[3]))
	d--
	translatedMove := [4]int{a, b, c, d}
	for _, i := range translatedMove {
		if i < 0 || i > 7 {
			return translatedMove, true
		}
	}
	return translatedMove, false
}

//func GetUserInput(c Colour) (userMove string, wrongFormat bool) {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Printf("Enter the next %s move, under lnln form, where l is a letter and n a number, then press Enter:\n", c.colour)
//	userMove, err := reader.ReadString('\n')
//	if err == nil && len(userMove) == 5 { // The input takes the newline into account
//		return userMove, false
//	} else {
//		return "", true
//	}
//}

//func playMove(move [4]int, board [8][8]int, boardRep [8][8]string, chessGame map[int]ChessPiece) ([8][8]int, [8][8]string) {
//	originLetter := move[0]
//	originNumber := move[1]
//	originNumber = 7 - originNumber
//	targetLetter := move[2]
//	targetNumber := move[3]
//	targetNumber = 7 - targetNumber
//	movedPieceId := board[originNumber][originLetter]
//	movedPiece := chessGame[movedPieceId]
//	if movedPiece.colour == "white" {
//		boardRep[targetNumber][targetLetter] = strings.ToLower(movedPiece.chessPieceType.asciiRep)
//	} else {
//		boardRep[targetNumber][targetLetter] = movedPiece.chessPieceType.asciiRep
//	}
//	board[originNumber][originLetter] = 0
//	boardRep[originNumber][originLetter] = "_"
//	board[targetNumber][targetLetter] = movedPieceId
//	return board, boardRep
//}

func main() {
	knight := PieceKind{"H"}
	queen := PieceKind{"Q"}
	king := PieceKind{"K"}
	bishop := PieceKind{"B"}
	tower := PieceKind{"T"}
	pawn := PieceKind{"P"}
	// Create the pawns
	g := &Game{pieces: make(map[int]Piece)}

	// Initialize the pawns
	i := 1
	for ; i <= 8; i ++ {
		k := i + 100
		g.pieces[i] = Piece{pawn, Black, i}
		g.board.Positions[0][i-1] = i
		g.pieces[k] = Piece{pawn, White, k}
		g.board.Positions[7][i-1] = k
	}

	// Initialize other pieces
	kinds := []PieceKind{tower, knight, bishop, queen, king, bishop, knight, tower}
	for j := 0; j <= 7; j ++ {
		i ++
		k := i + 100
		g.pieces[i] = Piece{kinds[j], Black, i}
		g.board.Positions[1][j] = i
		g.pieces[k] = Piece{kinds[j], White, k}
		g.board.Positions[6][j] = k
	}
	for i := 0; i < 400; i ++ {
		if a, ok := g.pieces[i]; ok {
			fmt.Println(a)
		}
	}
	fmt.Println(len(g.pieces))
	//
	////--------------------- GAME INITIALIZATION------------------------------
	//// We create the board that will hold the position of the pieces
	//board := [8][8]int{}
	//// We create its string representation
	//boardRep := [8][8]string{}
	//for i := 0; i < 8; i++ {
	//	for j := 0; j < 8; j++ {
	//		boardRep[i][j] = "_"
	//
	//	}
	//}
	////Create the mapping between moves and coordinates
	//lettersToInt := make(map[byte]int)
	//letters := [8]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	//for index, letter := range letters {
	//	lettersToInt[letter] = index + 1
	//}
	//
	////Each pieces appears a certain number of time
	//chessSet := make(map[int][]ChessPieceType)
	//chessSet[1] = []ChessPieceType{queen, king}
	//chessSet[2] = []ChessPieceType{bishop, knight, tower}
	//chessSet[8] = []ChessPieceType{pawn}
	//
	////We create the game with the previously defined variables
	//chessGame := make(map[int]ChessPiece)
	//for number, pieceSlice := range chessSet {
	//	for i := 0; i < len(pieceSlice); i++ {
	//		pieceType := pieceSlice[i]
	//		for j := 0; j < number; j++ {
	//			id_white := len(chessGame) + 2
	//			id_black := id_white + 1
	//			//We create each piece for one colour
	//			chessGame[id_white] = ChessPiece{pieceType, "white", id_white, j}
	//			chessGame[id_black] = ChessPiece{pieceType, "black", id_black, j}
	//		}
	//	}
	//}
	//// After building the set, we must place it on the board
	//for _, piece := range chessGame {
	//	id := piece.id
	//	number := piece.number
	//	pieceType := piece.chessPieceType
	//	colour := piece.colour
	//	// Depending on its number, we place it differently
	//	j := pieceType.originalX[number]
	//	// Depending  on the colour, it is either on the top or bottom row
	//	var colourId int
	//	if colour == "white" {
	//		colourId = 1
	//	} else {
	//		colourId = 0
	//	}
	//	i := pieceType.originalY[colourId]
	//	boardRep[i][j] = pieceType.asciiRep
	//	// if the piece is white, put it in lowercase
	//	if colour == "white" {
	//		boardRep[i][j] = strings.ToLower(pieceType.asciiRep)
	//	}
	//	board[i][j] = id
	//}
	//
	////--------------------------- BEGINNING THE ACTUAL GAME ---------------------------------------
	//
	//continueGame := true          // Will be false whenever there is a checkmate
	//turnColour := Colour{"white"} // Holds the color of the next player to play a move
	//for continueGame {
	//	PrintBoard(boardRep)
	//	isMoveFalse := false
	//	userMove, wrongFormat := GetUserInput(turnColour)
	//	for wrongFormat {
	//		fmt.Println("Please respect the move format !")
	//		userMove, wrongFormat = GetUserInput(turnColour)
	//	}
	//
	//	coordinateMove, isMoveFalse := TranslateMove(userMove, lettersToInt)
	//	if isMoveFalse {
	//		fmt.Println("Incorrect move entered !")
	//		continue
	//	}
	//	board, boardRep = playMove(coordinateMove, board, boardRep, chessGame)
	//
	//	//if ThereIsCheckmate(board) {
	//	//continueGame = false
	//	//fmt.Println("%s won the game !", turnColour)
	//	//}
	//	turnColour = Change(turnColour)
	//
	//}
	//
}
