package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Color = int

const (
	White = iota
	Black = iota
)

var lettersToInt = map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4, "f": 5, "g": 6, "h": 7}

var validInput = regexp.MustCompile(`([a-h][1-8]){2}`)

type Game struct {
	colour Color
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
	fmt.Printf("Enter the next move, under lnln form, where l is a letter and n a number, then press Enter:\n")
	userMove, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("the input couldn't be read")
	} else if len(userMove) > 5 {
		return "", fmt.Errorf("the input was too long")
	} else if !validInput.MatchString(userMove) {
		return "", fmt.Errorf("the input wasn't properly formatted")
	} else {
		return userMove[:4], nil // Remove the delimiter
	}
}

func translateInput(i string) [4]int {
	l1 := lettersToInt[i[0:1]] - 1
	l2 := lettersToInt[i[2:3]] - 1
	n1, _ := strconv.Atoi(i[1:2])
	n2, _ := strconv.Atoi(i[3:4])
	n1 --
	n2 --
	return [4]int{l1, n1, l2, n2}

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

	continueGame := true
	turnColour := White
	for continueGame {
		fmt.Println(g.represent())
		fmt.Println(turnColour)
		input, err := GetUserInput()
		for err != nil {
			fmt.Println(err)
			input, err = GetUserInput()
		}
		move := translateInput(input)

	}

}


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
