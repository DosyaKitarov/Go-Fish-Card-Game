package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var game bool
var playerPoints int
var oponentPoints int
var who bool

type card struct {
	suit  string
	value string
}

func main() {
	deck := createDeck()
	deck = shuffleDeck(deck)
	oponent := deck[0:7]
	deck = deck[7:]
	you := deck[0:7]
	deck = deck[7:]
	game = true
	startGame(you, oponent, deck)

}

func startGame(you []card, oponent []card, deck []card) {
	fmt.Println(" ____________________________")
	fmt.Println("|  ________________________  |")
	fmt.Println("| |                        | |")
	fmt.Println("| |   Welcome to Go Fish!  | |")
	fmt.Println("| |                        | |")
	time.Sleep(time.Second * 1)
	fmt.Println("| |   You know Rules? Y/N: | |")
	fmt.Println("| |________________________| |")
	fmt.Println("|____________________________|")
	var rules string
	fmt.Scan(&rules)
	if rules == "N" || rules == "n" {
		printRules()
	}
	fmt.Println(" ____________________________")
	fmt.Println("|  ________________________  |")
	fmt.Println("| |                        | |")
	fmt.Println("| |   	  Game Start!      | |")
	fmt.Println("| |________________________| |")
	fmt.Println("|____________________________|")

	time.Sleep(time.Second * 1)
	for game {
		yourTurn(&you, &oponent, &deck)
		oponentsTurn(&you, &oponent, &deck)
	}

}
func shuffleDeck(deck []card) []card {
	for i := 0; i < len(deck); i++ {
		j := rand.Intn(len(deck))
		temp := deck[i]
		deck[i] = deck[j]
		deck[j] = temp
	}
	return deck
}
func createDeck() []card {
	var cards []card
	cards = append(cards, createSuit("Spades")...)
	cards = append(cards, createSuit("Hearts")...)
	cards = append(cards, createSuit("Clubs")...)
	cards = append(cards, createSuit("Diamonds")...)
	return cards
}
func createSuit(suit string) []card {
	var cards []card
	cards = append(cards, card{suit, "A"})
	for i := 2; i <= 9; i++ {
		cards = append(cards, card{suit, fmt.Sprint(i)})
	}
	cards = append(cards, card{suit, "10"})
	cards = append(cards, card{suit, "J"})
	cards = append(cards, card{suit, "Q"})
	cards = append(cards, card{suit, "K"})

	return cards
}
func printRules() {
	fmt.Println(" ____________________________")
	fmt.Println("|  ________________________  |")
	fmt.Println("| |        Rules:          | |")
	fmt.Println("| |------------------------| |")
	fmt.Println("| | You will be playing    | |")
	fmt.Println("| | against the computer.  | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | You will be dealt 7    | |")
	fmt.Println("| | cards.                 | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)

	fmt.Println("| | The goal of the game   | |")
	fmt.Println("| | is to get 4 of a kind. | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)

	fmt.Println("| | You will ask the       | |")
	fmt.Println("| | computer for a card.   | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)

	fmt.Println("| | If the computer has    | |")
	fmt.Println("| | the card, you will     | |")
	fmt.Println("| | get it and get to go   | |")
	fmt.Println("| | again.                 | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)

	fmt.Println("| | If the computer does   | |")
	fmt.Println("| | not have the card,     | |")
	fmt.Println("| | you will draw a card   | |")
	fmt.Println("| | from the deck.         | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | If you get 4 of a kind,| |")
	fmt.Println("| | you will get a point   | |")
	fmt.Println("| | and get to go again.   | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | If you do not get 4    | |")
	fmt.Println("| | of a kind, the         | |")
	fmt.Println("| | computer will go.      | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)

	fmt.Println("| | The computer will ask  | |")
	fmt.Println("| | you for a card.        | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | If you have the card,  | |")
	fmt.Println("| | the computer will get  | |")
	fmt.Println("| | it and get to go       | |")
	fmt.Println("| | again.                 | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | If you do not have     | |")
	fmt.Println("| | the card, the          | |")
	fmt.Println("| | computer will draw a   | |")
	fmt.Println("| | card from the deck.    | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | If the computer gets   | |")
	fmt.Println("| | 4 of a kind, it will   | |")
	fmt.Println("| | get a point and get    | |")
	fmt.Println("| | to go again.           | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | If the computer does   | |")
	fmt.Println("| | not get 4 of a kind,   | |")
	fmt.Println("| | you will go.           | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | The game will continue | |")
	fmt.Println("| | until the deck is      | |")
	fmt.Println("| | empty.                 | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | The player with the    | |")
	fmt.Println("| | most points at the end | |")
	fmt.Println("| | of the game wins.      | |")
	fmt.Println("| |------------------------| |")
	time.Sleep(time.Second * 2)
	fmt.Println("| | Good luck!             | |")
	fmt.Println("| |________________________| |")
	fmt.Println("|____________________________|")
}

func countCards(you []card) {
	cardCounts := make(map[string]int)

	for _, c := range you {
		cardCounts[c.value]++
	}

	for value, count := range cardCounts {
		if value == "10" {
			fmt.Printf("| |           %s:%d         | |\n", value, count)
		} else {
			fmt.Printf("| |           %s:%d          | |\n", value, count)

		}
		fmt.Println("| |------------------------| |")

	}
}
func yourTurn(you *[]card, oponent *[]card, deck *[]card) {
	who = true

	fmt.Println(" ____________________________")
	fmt.Println("|  ________________________  |")
	fmt.Println("| |                        | |")
	fmt.Println("| |   	  Your turn!       | |")
	fmt.Println("| |________________________| |")
	fmt.Println("|____________________________|")
	time.Sleep(time.Second * 1)
	fmt.Println(" ____________________________")
	fmt.Println("|  ________________________  |")
	fmt.Println("| |                        | |")
	fmt.Println("| |	   Your deck:      | |")
	fmt.Println("| |------------------------| |")
	countCards(*you)
	fmt.Println("| |________________________| |")
	fmt.Println("|____________________________|")
	time.Sleep(time.Second * 1)

	var card string
	if len(*you) != 0 {
	LOOP:
		for {
			fmt.Println(" ______________________________________________________")
			fmt.Println("|  __________________________________________________  |")
			fmt.Println("| |                                                  | |")
			fmt.Println("| |   	  What card would you like to ask for?       | |")
			fmt.Println("| |__________________________________________________| |")
			fmt.Println("|______________________________________________________|")
			fmt.Scan(&card)
			for _, v := range *you {

				if v.value == strings.ToUpper(card) {
					break LOOP
				}
			}
			fmt.Println(" ______________________________________________________")
			fmt.Println("|  __________________________________________________  |")
			fmt.Println("| |                                                  | |")
			fmt.Println("| |   	  You do not have that card! Try again.      | |")
			fmt.Println("| |__________________________________________________| |")
			fmt.Println("|______________________________________________________|")
		}
	}
	found := false
	for _, v := range *oponent {
		if v.value == card {
			found = true
			break
		}
	}
	if found {
		takeCard(you, oponent, card)
		if checkPoints(you) {
			playerPoints++
		}
		yourTurn(you, oponent, deck)
	} else {
		fmt.Println(" ____________________________")
		fmt.Println("|  ________________________  |")
		fmt.Println("| |                        | |")
		fmt.Println("| |   	    Go fish!       | |")
		fmt.Println("| |________________________| |")
		fmt.Println("|____________________________|")
		if checkWin(deck) {
			return
		}
		checkWin(deck)
		time.Sleep(time.Second * 1)
		*you = append(*you, (*deck)[0])
		*deck = (*deck)[1:]
		fmt.Println(" ____________________________")
		fmt.Println("|  ________________________  |")
		fmt.Println("| |                        | |")
		fmt.Printf("| |   	  You took: %s      | |\n", (*you)[len(*you)-1].value)
		fmt.Println("| |________________________| |")
		fmt.Println("|____________________________|")
		time.Sleep(time.Second * 1)
		if checkPoints(you) {
			playerPoints++
		}
	}
}

func takeCard(you *[]card, oponent *[]card, vs string) {
	var indexses []int
	for i, v := range *oponent {
		if v.value == vs {
			indexses = append(indexses, i)
		}
	}

	for _, v := range indexses {
		*you = append(*you, (*oponent)[v])
	}
	resutl := []card{}
	for _, v := range *oponent {
		if v.value != vs {
			resutl = append(resutl, v)
		}
	}
	*oponent = resutl

	if who {
		fmt.Println(" ______________________________________________________")
		fmt.Println("|  __________________________________________________  |")
		fmt.Println("| |                                                  | |")
		fmt.Printf("| |   	     You get %v cards from oponent!           | |\n", len(indexses))
		fmt.Println("| |__________________________________________________| |")
		fmt.Println("|______________________________________________________|")
		time.Sleep(time.Second * 1)
	} else {
		fmt.Println(" ______________________________________________________")
		fmt.Println("|  __________________________________________________  |")
		fmt.Println("| |                                                  | |")
		fmt.Printf("| |          Oponent gets %v cards from you!          | |\n", len(indexses))
		fmt.Println("| |__________________________________________________| |")
		fmt.Println("|______________________________________________________|")
		time.Sleep(time.Second * 1)
	}
}
func oponentsTurn(you *[]card, oponent *[]card, deck *[]card) {
	who = false
	fmt.Println(" ____________________________")
	fmt.Println("|  ________________________  |")
	fmt.Println("| |                        | |")
	fmt.Println("| |     Oponents turn!     | |")
	fmt.Println("| |________________________| |")
	fmt.Println("|____________________________|")
	time.Sleep(time.Second * 1)
	card := (*oponent)[rand.Intn(len(*oponent))].value
	fmt.Println(" ____________________________")
	fmt.Println("|  ________________________  |")
	fmt.Println("| |                        | |")
	fmt.Printf("| |  Oponent asks for %v    | |\n", card)
	fmt.Println("| |________________________| |")
	fmt.Println("|____________________________|")
	time.Sleep(time.Second * 1)
	found := false
	for _, v := range *you {
		if v.value == card {
			found = true
			break
		}
	}
	if found {
		takeCard(oponent, you, card)
		if checkPoints(oponent) {
			oponentPoints++
		}

		oponentsTurn(you, oponent, deck)
	} else {

		fmt.Println(" ____________________________")
		fmt.Println("|  ________________________  |")
		fmt.Println("| |                        | |")
		fmt.Println("| |   	    Go fish!       | |")
		fmt.Println("| |________________________| |")
		fmt.Println("|____________________________|")
		if checkWin(deck) {
			return
		}
		time.Sleep(time.Second * 1)
		*oponent = append(*oponent, (*deck)[0])
		*deck = (*deck)[1:]
		if checkPoints(oponent) {
			oponentPoints++
		}

	}
}
func checkPoints(you *[]card) bool {
	cardCounts := make(map[string]int)

	for _, c := range *you {
		cardCounts[c.value]++
	}

	for value, count := range cardCounts {
		if count == 4 {
			fmt.Println(" ____________________________")
			fmt.Println("|  ________________________  |")
			fmt.Println("| |                        | |")
			fmt.Println("| |   	  Get a point!     | |")
			fmt.Println("| |________________________| |")
			fmt.Println("|____________________________|")
			time.Sleep(time.Second * 1)
			removeCard(you, value)
			return true
		}
	}
	return false
}
func removeCard(you *[]card, card string) {
	var indices []int
	for i := len(*you) - 1; i >= 0; i-- {
		if (*you)[i].value == card {
			indices = append(indices, i)
		}
	}
	for _, index := range indices {
		*you = append((*you)[:index], (*you)[index+1:]...)
	}
}
func checkWin(deck *[]card) bool {
	if len(*deck) == 0 {
		if playerPoints > oponentPoints {
			fmt.Println(" ____________________________")
			fmt.Println("|  ________________________  |")
			fmt.Println("| |                        | |")
			fmt.Println("| |        You win!        | |")
			fmt.Println("| |________________________| |")
			fmt.Println("|____________________________|")
			fmt.Printf("You: %v\n", playerPoints)
			fmt.Printf("Oponent: %v\n", oponentPoints)

		} else {
			fmt.Println(" ____________________________")
			fmt.Println("|  ________________________  |")
			fmt.Println("| |                        | |")
			fmt.Println("| |        You lose!       | |")
			fmt.Println("| |________________________| |")
			fmt.Println("|____________________________|")
			fmt.Printf("You: %v\n", playerPoints)
			fmt.Printf("Oponent: %v\n", oponentPoints)
			time.Sleep(time.Second * 2)
			fmt.Println(" ____________________________")
			fmt.Println("|  ________________________  |")
			fmt.Println("| |                        | |")
			fmt.Println("| | As punishment,         | |")
			fmt.Println("| | I delete your system32 | |")
			fmt.Println("| |________________________| |")
			fmt.Println("|____________________________|")
		}
		game = false
		return true
	}
	return false
}
