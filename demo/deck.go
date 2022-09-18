package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func createDeck() deck {
	cards := deck{}
	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(deck deck, handsize int) (deck, deck) {
	return deck[:handsize], deck[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	data := []byte(d.toString())
	return ioutil.WriteFile(filename, data, 0666)
}

func createDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), ","))

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		new_pos := r.Intn(len(d) - 1)
		d[index], d[new_pos] = d[new_pos], d[index]
	}
}

// func main() {
// 	cards := createDeck()
// 	cards.print()
// 	fmt.Println("========================")
// 	hand, remaingDeck := deal(cards, 5)
// 	hand.print()
// 	fmt.Println("========================")
// 	remaingDeck.print()
// 	fmt.Println("========================")
// 	cards.print()
// 	fmt.Println("========================")
// 	fmt.Println(cards.saveToFile("deck"))
// 	fmt.Println("========================")
// 	cards = createDeckFromFile("deck")
// 	cards.shuffle()
// 	cards.print()
// }
