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
	id string
	players players
	deck deck.Deck
	DiscardDeck discard.Discard
	Turn int 
	direction bool
	cardCount int
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
		DiscardDeck: discard.CreateDiscard(),
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
	if len(g.DiscardDeck)>0{
		return fmt.Errorf("Game is already started")
	}
	validTop:=false
	for !validTop {
		g.DiscardDeck.Discard(&(g.deck),len(g.deck)-1)
		if g.DiscardDeck.TopCard().Color != "Any" {
			validTop=true
		}
	}
	
	return nil
}
func(g *Game) NextPlayer() *player.Player{
	return &g.players[g.Turn]
}

func(g *Game) PlayTurn(index int, color string) error{
	if index == -1{
		err:=g.NextPlayer().TakeCard(&g.deck)
		fmt.Println("took card")
		if err !=nil{
			return err
		}
		return nil
	}else{
		if index>=len(g.NextPlayer().Deck) && index<0{
			return errors.New("index card does not exist")
		}
		card:=g.NextPlayer().Deck[index]
		
		if g.DiscardDeck.ValidPlay(card){
			if card.Color =="Any"&&color!="Red"&&color!="Green"&&color!="Blue"&&color!="Yellow"{
				return errors.New("must select valid color")
			}
			card.Color=color
			g.playCard(index)
			return nil
		}
		fmt.Printf("Invalid Play\n")
		return errors.New("not valid card to play")
	}
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
	count:=len(g.deck)+len(g.DiscardDeck)
	for _,player:= range g.players{
		count+= len(player.Deck)
	}
	fmt.Printf("Total cards: %v\n",count)
}