package deck

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Card struct {
	Color  string `json:"Color"`
	Number int `json:"Number"`
	Power  string `json:"Power"`
}

type Deck []Card

func NewDeck() Deck {
	var cards Deck
	cardColors := []string{"Red", "Blue", "Green", "Yellow"}
	cardPowers := []string{"+2","Reverse", "Skip"}
	for _, color := range cardColors {
		for i := 0; i <= 9; i++ {
			if i!=0 {
				card := Card{Color: color, Number: i}
				cards = append(cards, card)
			}
			card := Card{Color: color, Number: i}
			cards = append(cards, card)
		}
	}
	for _, color := range cardColors {
		for _,power := range cardPowers {
			card := Card{Color: color, Number: -1, Power:power}
			cards = append(cards, card)
			cards = append(cards, card)
		}
	}
	for i := 0; i < 4; i++ {
		card := Card{Color: "Any", Number:-1,Power: "wild" }
		cards = append(cards, card)
	}
	for i := 0; i < 4; i++ {
		card := Card{Color: "Any", Number:-1,Power: "+4"}
		cards = append(cards, card)
	}
	return cards
}

func (d Deck) ShowDeck() {
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
func (d Deck) Shuffle(){
	for i := range d{
		newPos := rand.Intn(len(d)-1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func Deal(d Deck,  handSize int) (Deck, Deck,error) {
	if len(d)<handSize {
		return nil, nil, fmt.Errorf("Deck size is smaller than hand requested\nDeck Size: %v, Hand Size:%v",len(d),handSize)
	}
	return d[:handSize],d[handSize:], nil
}

func (d Deck) CardExists(card Card) bool {
	for _, c := range d {
		if c.Color == card.Color && c.Number == card.Number && c.Power == card.Power {
			return true
		}
	}
	return false
}

func (d Deck) SaveToFile(fileName string) error {
	jsonCards, err := json.Marshal(d)
    if err != nil {
        panic (err)
    }
	return os.WriteFile(fileName, []byte(jsonCards), 0666) 
}

func NewDeckFromFile(filename string) Deck{
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var cards Deck
	json.Unmarshal(bs, &cards)
	return cards
}
