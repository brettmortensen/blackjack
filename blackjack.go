package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Card struct {
	suit  string
	value string
}

func newCard(suit, value string) *Card {
	return &Card{suit: suit, value: value}
}

type Deck struct {
	cards []Card
}

func newDeck() *Deck {
	var cards []Card
	var suits = []string{"♣️", "♦️", "♥️", "♠️"}
	var values = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			cards = append(cards, *newCard(suits[i], values[j]))
		}
	}

	// https://pkg.go.dev/math/rand#Seed
	rand.Seed(time.Now().UnixNano())

	// https://pkg.go.dev/math/rand#Shuffle
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return &Deck{cards: cards}
}

type Hand struct {
	cards []Card
}

func newHand(deck *Deck) *Hand {
	hand := Hand{}
	hand.drawCard(deck)
	hand.drawCard(deck)
	return &hand
}

func (hand *Hand) drawCard(deck *Deck) Card {
	var card = deck.cards[0]
	deck.cards = deck.cards[1:]
	hand.cards = append(hand.cards, card)
	return card
}

func (hand *Hand) value() int {
	var aceCount int
	var value int

	for _, card := range hand.cards {
		switch card.value {
		case "A":
			aceCount = aceCount + 1
		case "2", "3", "4", "5", "6", "7", "8", "9", "10":
			cardValue, _ := strconv.Atoi(card.value)
			value = value + cardValue
		case "J", "Q", "K":
			value = value + 10
		}
	}

	switch aceCount {
	case 1:
		if value+11 <= 21 {
			value = value + 11
		} else {
			value = value + 1
		}
	case 2:
		if value+12 <= 21 {
			value = value + 12
		} else {
			value = value + 2
		}
	case 3:
		if value+13 <= 21 {
			value = value + 13
		} else {
			value = value + 3
		}
	case 4:
		if value+14 <= 21 {
			value = value + 14
		} else {
			value = value + 4
		}
	}

	return value
}

func main() {
	fmt.Println("Welcome to Go Blackjack!\n")

	fmt.Println("Shuffling deck...\n")
	time.Sleep(time.Second)
	var deck = newDeck()

	fmt.Println("Dealing cards...\n")
	time.Sleep(time.Second)
	var usersHand = newHand(deck)
	var dealersHand = newHand(deck)

	fmt.Println("Dealer peeks at their second card...\n")
	time.Sleep(time.Second)

	var input string

	fmt.Printf("Dealer's hand: %v\n", dealersHand.cards[0])
	fmt.Printf("Your hand: %v\n", usersHand.cards)

	dealersHandValue := dealersHand.value()
	usersHandValue := usersHand.value()

	if usersHandValue == 21 {
		fmt.Println("You are dealt blackjack, you win!")
		os.Exit(0)
	} else if dealersHandValue == 21 {
		fmt.Printf("Dealer's hand: %v\n", dealersHand.cards)
		fmt.Println("Dealer is dealt blackjack, dealer wins!")
		os.Exit(0)
	}

	for {
		fmt.Print("(h)it or (s)tand?: ")
		fmt.Scan(&input)
		fmt.Println()

		switch input {
		case "h":
			fmt.Printf("You receive the card: %v\n\n", usersHand.drawCard(deck))
			time.Sleep(time.Second)
			fmt.Printf("Your hand: %v\n", usersHand.cards)

			if usersHand.value() > 21 {
				fmt.Printf("You bust at %d, dealer wins!\n", usersHand.value())
				os.Exit(0)
			} else if usersHand.value() == 21 {
				fmt.Println("You get blackjack, you win!")
				os.Exit(0)
			}
		case "s":
			fmt.Printf("Dealer flips over their second card, dealer's hand: %v\n", dealersHand.cards)

			for {
				fmt.Printf("Dealer receives the card: %v\n\n", dealersHand.drawCard(deck))
				time.Sleep(time.Second)
				fmt.Printf("Dealer's hand: %v\n", dealersHand.cards)

				dealersHandValue := dealersHand.value()
				usersHandValue := usersHand.value()

				if dealersHandValue > 21 {
					fmt.Printf("Dealer busts at %d, you win!\n", dealersHandValue)
					os.Exit(0)
				} else if dealersHandValue == 21 {
					fmt.Println("Dealer gets blackjack, dealer wins!")
					os.Exit(0)
				} else if dealersHandValue >= 17 && dealersHandValue <= 21 {
					if dealersHandValue > usersHandValue {
						fmt.Printf("Dealer's %d beats your %d, dealer wins!\n", dealersHandValue, usersHandValue)
						os.Exit(0)
					} else if dealersHandValue == usersHandValue {
						fmt.Printf("Dealer's %d ties your %d, push!\n", dealersHandValue, usersHandValue)
						os.Exit(0)
					} else {
						fmt.Printf("Your %d beats dealer's %d, you win!\n", usersHandValue, dealersHandValue)
						os.Exit(0)
					}
				}
			}
		default:
			fmt.Println("Not a valid choice. Please enter 'h' or 's'.")
		}
	}
}
