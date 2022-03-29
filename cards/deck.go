package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

	// build the deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function
// range iterates over every element in the slice/array
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// turn deck into a string
func (d deck) toString() string {
	// returns comma seperated string of cards in the deck
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// if error returned
	if err != nil {
		// option 1: log the error and return a call to newDeck()
		// option 2: log the error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// split string byteSlice at the commas
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffle
func (d deck) shuffle() {
	// random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// for each index, card in cards
	for i := range d {
		// generate a random number between 0 and len(cards) - 1
		newPosition := r.Intn(len(d) - 1)
		// swap the current card and the card at cards[randomNumber]
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
