package player

import "unoSimulator/internal/deck"

type Player struct {
	Name string
	Deck deck.Deck
}

func CreatePlayer(name string,d deck.Deck) (Player,error) {
	dPlayer,_,err :=deck.Deal(d,7)
	if err !=nil{
		var p Player
		return p,nil
	}
	p:=Player{
		Name:name,
		Deck:dPlayer,
	}
	return p, nil
}