package main

import (
	"fmt"
	"strings"
)

var board [5][5]string
var rows = len(board)
var columns = len(board[0])
var teamOne [5]string
var teamTwo [5]string
var playerOne string = "P1"
var playerTwo string = "P2"

const (
	F = "F"
	B = "B"
	L = "L"
	R = "R"
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
	fmt.Println("player one starts")
	for {
		fmt.Println("player one (pawn:direction): ")
		m := input(playerOne)
		fmt.Println(m)
	}
}

func input(one string) int {
	var input string
	fmt.Scan(&input)
	var pawn, dir = getDirectionAndPawnToMove(input)
	prefix := "A-"
	pawnWithPrefix := prefix + pawn
	fmt.Println("lol", pawnWithPrefix)
	p := playerOnePoints.pointMap[pawnWithPrefix]
	fmt.Println("now", p)
	switch dir {
	case F:
		modified := Point{
			x: p.x,
			y: p.y,
		}
		fmt.Println("here", modified)
	case B:
	case L:
	case R:

	}
	return 1

}

func getDirectionAndPawnToMove(input string) (string, string) {
	var pawnAndDir = strings.Split(input, ":")
	println(pawnAndDir[0])
	println(pawnAndDir[1])
	return pawnAndDir[0], pawnAndDir[1]
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
			fmt.Println("logs", teamOne[i], playerOnePoints.pointMap[teamOne[i]])
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
