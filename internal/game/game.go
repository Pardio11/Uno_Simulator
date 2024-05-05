package game

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"unoSimulator/internal/deck"
	"unoSimulator/internal/discard"
	"unoSimulator/internal/player"
)

type Game struct{
	id string
	players players
	deck deck.Deck
	discardDeck discard.Discard
	Turn int 
}

type players []player.Player

func CreateGame(name string) (Game,error) {
	id,err:= generateRandomID(8)
	if err != nil {
		var g Game
		return g, err 
	}

	d := deck.NewDeck()
	d.Shuffle()
	dPlayer, dGame,err := deck.Deal(d,7)
	if err != nil {
		var g Game
		return g, err 
	}

	player,err := player.CreatePlayer(name,dPlayer)
	if err != nil {
		var g Game
		return g, err 
	}
	p := players{player}

	g := Game{
		id: id,
		players: p,
		deck: dGame,
		discardDeck: discard.CreateDiscard(),
	}
	return g,nil
}

func (g *Game) AddPlayer(name string) error {
	dPlayer, dGame,err := deck.Deal(g.deck, 7)

	if err != nil {
		return err
	}
	
	p, err := player.CreatePlayer(name, dPlayer)

	if err != nil {
		return err
	}

	g.deck=dGame
	g.players = append(g.players, p)
	return nil
}

func (g *Game) StartGame() error {
	if len(g.discardDeck)>0{
		return fmt.Errorf("Game is already started")
	}
	dDiscard,dGame,err:=deck.Deal(g.deck,1)
	if err != nil {
		return err
	}
	g.discardDeck = discard.Discard(dDiscard)
	g.deck = dGame
	return nil
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