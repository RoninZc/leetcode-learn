package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		if input == "" {
			break
		}
		compare(input)
	}
}

func compare(input string) {
	left, right := InitInput(input)
	leftPokes, rightPokes := score(left), score(right)
	if leftPokes == nil || rightPokes == nil {
		fmt.Println("ERROR")
		return
	}
	if leftPokes.Type == BigBoom {
		OutPut(left)
		return
	}
	if rightPokes.Type == BigBoom {
		OutPut(right)
		return
	}
	if leftPokes.Type == BOOM || rightPokes.Type == BOOM {
		if leftPokes.Type != BOOM {
			OutPut(right)
			return
		}
		if rightPokes.Type != BOOM {
			OutPut(left)
			return
		}
		if leftPokes.Score > rightPokes.Score {
			OutPut(left)
			return
		} else {
			OutPut(right)
			return
		}
	}
	if leftPokes.Type != rightPokes.Type {
		fmt.Println("ERROR")
		return
	}
	if leftPokes.Score > rightPokes.Score {
		OutPut(left)
	} else {
		OutPut(right)
	}
}

func OutPut(pokes []string) {
	for _, poke := range pokes {
		fmt.Print(poke, " ")
	}
	fmt.Println()
}

func InitInput(input string) (left []string, right []string) {
	left, right, leftOk, point := make([]string, 0), make([]string, 0), false, 0
	for i, char := range input {
		switch char {
		case ' ':
			if !leftOk {
				left = append(left, input[point:i])
			} else {
				right = append(right, input[point:i])
			}
			point = i + 1
		case '-':
			leftOk = true
			left = append(left, input[point:i])
			point = i + 1
		}
	}
	right = append(right, input[point:len(input)-1])
	return
}

const (
	ONE = iota + 1
	DOUBLE
	THREE
	BOOM
	SHUNZI
	BigBoom
)

type Pokes struct {
	Type  int
	Score int
}

func score(pokes []string) *Pokes {
	switch len(pokes) {
	case 1:
		return &Pokes{
			Type:  ONE,
			Score: GetPokeScore(pokes[0]),
		}
	case 2:
		for i, poke := range pokes {
			if poke == "joker" || poke == "JOKER" {
				if i == 1 {
					return &Pokes{
						Type:  BigBoom,
						Score: 0,
					}
				}
				continue
			}
			break
		}
		return &Pokes{
			Type:  DOUBLE,
			Score: GetPokeScore(pokes[0]),
		}
	case 3:
		return &Pokes{
			Type:  THREE,
			Score: GetPokeScore(pokes[0]),
		}
	case 4:
		return &Pokes{
			Type:  BOOM,
			Score: GetPokeScore(pokes[0]),
		}
	case 5:
		return &Pokes{
			Type:  SHUNZI,
			Score: GetPokeScore(pokes[0]),
		}
	}
	return nil
}

func GetPokeScore(poke string) int {
	switch poke {
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "10":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	case "2":
		return 15
	case "joker":
		return 16
	case "JOKER":
		return 17
	default:
		return 0
	}
}
