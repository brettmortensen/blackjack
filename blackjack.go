package main

import (
	"fmt"
	"math/rand"
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

func main() {
	fmt.Println("Welcome to Blackjack!\n")
	var deck = newDeck()
	fmt.Println(deck)
	fmt.Printf("Deck size: %v\n", len(deck.cards))
}
