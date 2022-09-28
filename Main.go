package main

import "fmt"

var board [5][5]string
var rows = len(board)
var columns = len(board[0])
var teamOne [5]string
var teamTwo [5]string
var playerOne string = "P1"
var playerTwo string = "P2"

const (
	F = iota
	B
	L
	R
)

var playerOneScore = 0
var playerTwoScore = 0

type Point struct {
	x int
	y int
}

func (p Point) PointX() int {
	return p.x
}
func (p Point) PointY() int {
	return p.y
}

type pointsAndPos struct {
	pointMap map[string]Point
}

func (p pointsAndPos) pointsPos() {

}

var playerOnePoints pointsAndPos
var playerTwoPoints pointsAndPos

func main() {
	initialize()
	playerOnePoints.pointMap = map[string]Point{}
	playerTwoPoints.pointMap = map[string]Point{}
	println("*****************************")
	fmt.Println("enter your players ,player one : ")
	println("*****************************")
	inputPlayerAndPlace(playerOne)
	draw()
	println("*****************************")
	fmt.Println("enter your players ,player two : ")
	println("*****************************")
	inputPlayerAndPlace(playerTwo)

	for {
		draw()
		var input, err = Input()
		if err == nil {
			updatePos()
		}
		fmt.Println(input)
	}

}

func inputPlayerAndPlace(one string) {
	for i := 0; i < rows; i++ {
		if one == playerOne {
			var input string
			fmt.Scan(&input)
			teamOne[i] = "A-" + input
			board[columns-1][i] = teamOne[i]
			var p Point = Point{columns - 1, i}
			playerOnePoints.pointMap[teamOne[i]] = p
		} else {
			var input string
			fmt.Scan(&input)
			teamTwo[i] = "B-" + input
			board[0][i] = teamTwo[i]
			var p Point = Point{0, i}
			playerTwoPoints.pointMap[teamTwo[i]] = p
		}
	}
}

func updatePos() {

}

func Input() (string, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("error is " + err.Error())
	}
	return input, nil
}

func initialize() {

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			board[i][j] = "*   "
		}
		println()
	}
}

func draw() {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print(board[i][j] + "  ")
		}
		fmt.Println()
	}
}
