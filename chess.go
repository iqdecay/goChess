package main

import (
	"bufio"
	"fmt"
	"os"
)

type Color = int

const (
	White = iota
	Black = iota
)

type Game struct {
	board  [8][8]int
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

func (g Game) represent() string {
	positions := g.board
	pieces := g.pieces
	var display string
	var str string
	for _, line := range positions {
		str = ""
		for _, id := range line {
			if id == 0 {
				str += "_"
			} else {
				str += pieces[id].kind.symbol
			}
		}
		display += "\n" + str
	}
	return display
}

func GetUserInput() (userMove string, wrongFormat error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the next %s move, under lnln form, where l is a letter and n a number, then press Enter:\n", c.colour)
	userMove, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("the input couldn't be read")
	} else if len(userMove) != 5 {
		return "", fmt.Errorf("the input wasn't properly formatted")
	} else {
		return userMove, nil
	}
}

func main() {
	// Initialize Piece types
	knight := PieceKind{"H"}
	queen := PieceKind{"Q"}
	king := PieceKind{"K"}
	bishop := PieceKind{"B"}
	tower := PieceKind{"T"}
	pawn := PieceKind{"P"}
	g := &Game{pieces: make(map[int]Piece)}

	// Initialize the pawns, 1xx pieces are White, 0xx pieces are Black
	i := 1
	for ; i <= 8; i ++ {
		k := i + 100
		g.pieces[i] = Piece{pawn, Black, i}
		g.board[1][i-1] = i
		g.pieces[k] = Piece{pawn, White, k}
		g.board[6][i-1] = k
	}
	// Initialize other pieces according to their place on the board
	kinds := []PieceKind{tower, knight, bishop, queen, king, bishop, knight, tower}
	for j := 0; j <= 7; j ++ {
		i ++
		k := i + 100
		g.pieces[i] = Piece{kinds[j], Black, i}
		g.board[0][j] = i
		g.pieces[k] = Piece{kinds[j], White, k}
		g.board[7][j] = k
	}

	fmt.Println(len(g.pieces))
	fmt.Println(g.represent())
	////--------------------- GAME INITIALIZATION------------------------------
	//Create the mapping between moves and coordinates
	lettersToInt := make(map[byte]int)
	letters := [8]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	for index, letter := range letters {
		lettersToInt[letter] = index + 1
	}

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

//func TranslateMove(userMove string, lettersToInt map[byte]int) ([4]int, bool) {
//	a := lettersToInt[userMove[0]]
//	a--
//	b, _ := strconv.Atoi(string(userMove[1]))
//	b--
//	c := lettersToInt[userMove[2]]
//	c--
//	d, _ := strconv.Atoi(string(userMove[3]))
//	d--
//	translatedMove := [4]int{a, b, c, d}
//	for _, i := range translatedMove {
//		if i < 0 || i > 7 {
//			return translatedMove, true
//		}
//	}
//	return translatedMove, false
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
