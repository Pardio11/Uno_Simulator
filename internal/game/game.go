package game

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"unoSimulator/internal/deck"
	"unoSimulator/internal/discard"
	"unoSimulator/internal/player"
)

type Game struct{
	Id string `json:"id"`
	Players players `json:"players"`
	Deck deck.Deck `json:"deck"`
	DiscardDeck discard.Discard `json:"discard"`
	Turn int `json:"turn"`
	Direction bool `json:"direction"`
	CardCount int `json:"cardCount"`
}

type players []player.Player

func CreateGame() (Game,error) {
	id,err:= generateRandomID(8)
	if err != nil {
		var g Game
		return g, err 
	}

	d := deck.NewDeck()
	d.Shuffle()
	p := players{}

	g := Game{
		Id: id,
		Players: p,
		Deck: d,
		DiscardDeck: discard.CreateDiscard(),
	}
	return g,nil
}

func (g *Game) AddPlayer(name string) (player.Player,error) {
	dPlayer, dGame,err := deck.Deal(g.Deck, 7)
	var p player.Player
	if err != nil {
		return p,err
	}
	
	p, err = player.CreatePlayer(name, dPlayer)

	if err != nil {
		return p,err
	}

	g.Deck=dGame
	g.Players = append(g.Players, p)
	return p,nil
}

func (g *Game) StartGame() error {
	if len(g.DiscardDeck)>0{
		return fmt.Errorf("Game is already started")
	}
	validTop:=false
	for !validTop {
		g.DiscardDeck.Discard(&(g.Deck),len(g.Deck)-1)
		if g.DiscardDeck.TopCard().Color != "Any" {
			validTop=true
		}
	}
	
	return nil
}

func(g *Game) NextPlayer() *player.Player{
	return &g.Players[g.Turn]
}

func(g *Game) PlayTurn(index int, color string) error{
	if len(g.Deck)<g.CardCount+1 {
		g.joinDecks()
	}
	if index == -1{
		for i:=0;i<g.CardCount || i==0;i++{
			err:=g.NextPlayer().TakeCard(&g.Deck)
			fmt.Println("took card")
			if err !=nil{
				return err
			}
		}
		g.CardCount=0
		return nil
	}else{
		if index>=len(g.NextPlayer().Deck) && index<0{
			return errors.New("index card does not exist")
		}
		card:=g.NextPlayer().Deck[index]
		var pass bool
		if g.CardCount>0{
			pass=g.DiscardDeck.ValidPlusCardPlay(card)
		}else{
			pass=g.DiscardDeck.ValidPlay(card)
		}
		if pass {
			if card.Color =="Any"{
				if color!="Red"&&color!="Green"&&color!="Blue"&&color!="Yellow"{
					return errors.New("must select valid color")
				}
				g.NextPlayer().Deck[index].Color=color
			}
			
			g.playCard(index)
			return nil
		}
		fmt.Printf("Invalid Play\n")
		return errors.New("not valid card to play")
	}
}

func (g *Game) joinDecks(){
	g.Deck = append(g.Deck, g.DiscardDeck[:len(g.DiscardDeck)-1]...)
	g.DiscardDeck=g.DiscardDeck[len(g.DiscardDeck)-1:]
	g.Deck.Shuffle()
}

func generateRandomID(length int) (string, error) {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, length)

	// Get random numbers to fill the bytes slice
	for i := range bytes {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		bytes[i] = charset[num.Int64()]
	}

	return string(bytes), nil
}

func (g *Game) CountCards(){
	count:=len(g.Deck)+len(g.DiscardDeck)
	for _,player:= range g.Players{
		count+= len(player.Deck)
	}
	fmt.Printf("Total cards: %v\n",count)
}