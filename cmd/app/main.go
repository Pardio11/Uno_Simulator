package main

import (
	"fmt"
	"unoSimulator/internal/game"
)
func main (){
	fmt.Println("Running")
	g,_:=game.CreateGame("Carlos")
	g.AddPlayer("Pedro")
	g.AddPlayer("Romero")
	
	g.StartGame()
	
for{
	fmt.Printf("Top Card: %v\n",g.DiscardDeck.TopCard())
	fmt.Println(g.NextPlayer())
	var input int
    fmt.Println("\nEnter something:")
    _, err := fmt.Scan(&input)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	g.PlayTurn(input)
	g.CountCards()
}
}