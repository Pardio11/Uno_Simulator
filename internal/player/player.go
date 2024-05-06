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
 func (p *Player) TakeCard(d *deck.Deck)error{
	c,dTable,err:=deck.Deal(*d,1)
	*d=dTable
	if err != nil{
		return err
	}
	p.Deck = append(p.Deck, c[0])
	return nil
 }