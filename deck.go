package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "10", "9", "8", "7", "6"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // d - reciever of type 'deck'
	for i, card := range d {
		fmt.Println(i, card) // itirate over elements in a slice and print them out
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // generates a new int64 that is used as seed for random number generator
	r := rand.New(source)                           // generates a new seed based of int64 created before for random generator

	for i := range d { // i is index of the card in deck
		newPosition := r.Intn(len(d) - 1)           // generates random numbers that are asigned as indexes for cards
		d[i], d[newPosition] = d[newPosition], d[i] // shuffling cards indexes
	}
}
