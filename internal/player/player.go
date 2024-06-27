package player

import (
	"crypto/rand"
	"math/big"
	"unoSimulator/internal/deck"
)

type Player struct {
	Id string
	Name string
	Deck deck.Deck
}

func CreatePlayer(name string,d deck.Deck) (Player,error) {
	dPlayer,_,err :=deck.Deal(d,7)
	if err !=nil{
		var p Player
		return p,err
	}
	id,err := generateRandomID(10)
	if err != nil {
		var p Player
		return p, err 
	}
	p:=Player{
		Id: id,
		Name:name,
		Deck:dPlayer,
	}
	return p, nil
}
 func (p *Player) TakeCard(d *deck.Deck)error{
	c,dTable,err:=deck.Deal(*d,1)
	*d=dTable
	if err != nil{
		return err
	}
	p.Deck = append(p.Deck, c[0])
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