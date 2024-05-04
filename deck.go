package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type card struct {
	Color  string `json:"color"`
	Number int `json:"number"`
	Power  string `json:"Power"`
}

type deck []card

func newDeck() deck {
	var cards deck
	cardColors := []string{"Red", "Blue", "Green", "Yellow"}
	cardPowers := []string{"+2","Reverse", "Skip"}
	for _, color := range cardColors {
		for i := 0; i <= 9; i++ {
			if i!=0 {
				card := card{Color: color, Number: i}
				cards = append(cards, card)
			}
			card := card{Color: color, Number: i}
			cards = append(cards, card)
		}
	}
	for _, color := range cardColors {
		for _,power := range cardPowers {
			card := card{Color: color, Number: -1, Power:power}
			cards = append(cards, card)
			cards = append(cards, card)
		}
	}
	for i := 0; i < 4; i++ {
		card := card{Color: "Any", Number:-1, }
		cards = append(cards, card)
	}
	for i := 0; i < 4; i++ {
		card := card{Color: "Any", Number:-1,Power: "+4"}
		cards = append(cards, card)
	}
	return cards
}

func (d deck) showDeck() {
	for _, card := range d {
		switch{
			case card.Number==-1 && card.Power == "":
				fmt.Printf("Color:%v\n",card.Color)
			case card.Number==-1:
				fmt.Printf("Color:%v Power:%v\n",card.Color,card.Power)
			default:
				fmt.Printf("Color:%v Number:%v\n",card.Color,card.Number)
		}
	
	}
}
func (d deck) shuffle(){
	for i := range d{
		newPos := rand.Intn(len(d)-1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func deal(d deck,  handSize int) (deck, deck) {
	return d[:handSize],d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	jsonCards, err := json.Marshal(d)
    if err != nil {
        panic (err)
    }
	return os.WriteFile(fileName, []byte(jsonCards), 0666) 
}

func newDeckFromFile(filename string) deck{
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var cards deck
	json.Unmarshal(bs, &cards)
	return cards
}
