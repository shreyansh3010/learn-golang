package main

import "fmt"

type deck []string

func newDeck() deck {

	cards := deck{}

	cardHouses := []string{"Spade", "Diamond", "Heart", "Club"}

	cardNos := []string{"Ace", "Two", "Three", "Four"}

	for _, house := range cardHouses {
		for _, no := range cardNos {
			cards = append(cards, no+" of "+house)
		}
	}

	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
