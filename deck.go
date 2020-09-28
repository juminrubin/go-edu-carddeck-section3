package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
type deck []Card

func newDeck() deck {
	cards := deck{}

	for i := 0; i < len(cardSuites); i++ {
		for j := 0; j < len(cardValues); j++ {
			cards = append(cards, Card{CardValue(j + 1), CardSuite(i + 1)})
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card.String())
	}
}

func (d deck) String() string {
	var deckString []string
	for _, card := range d {
		deckString = append(deckString, card.String())
	}
	return strings.Join(deckString, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.String()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Parsing deck string
	deckStringArray := strings.Split(string(bs), ",")
	cards := deck{}
	for _, cardString := range deckStringArray {
		cards = append(cards, cardFromString(cardString))
	}

	return cards
}

func (d deck) shuffle() {
	theSize := len(d)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		targetPostIndex := r.Intn(theSize)
		if i == targetPostIndex {
			continue
		}
		d[i], d[targetPostIndex] = d[targetPostIndex], d[i]
	}
}

// Card data structure
type Card struct {
	value CardValue
	suite CardSuite
}

func (c Card) String() string {
	return c.value.String() + " of " + c.suite.String()
}

func cardFromString(s string) Card {
	token := strings.Split(s, " ")
	cardValueIndex := indexOf(cardValues[:], token[0])
	cardSuiteIndex := indexOf(cardSuites[:], token[2])
	return Card{CardValue(cardValueIndex + 1), CardSuite(cardSuiteIndex + 1)}
}

func indexOf(slice []string, item string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// CardSuite Index
type CardSuite int

// CardSuite enumeration
const (
	Diamonds CardSuite = 1 + iota
	Hearts
	Spades
	Clubs
)

var cardSuites = [...]string{
	"Diamonds",
	"Hearts",
	"Spades",
	"Clubs",
}

func (suite CardSuite) String() string {
	return cardSuites[suite-1]
}

// CardValue Index
type CardValue int

// CardValue enumeration
const (
	Ace CardValue = 1 + iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var cardValues = [...]string{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

func (cardValue CardValue) String() string {
	return cardValues[cardValue-1]
}
