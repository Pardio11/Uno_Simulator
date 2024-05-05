package main

import (
	"fmt"
	"unoSimulator/internal/game"
)
func main (){
	fmt.Println("Running")
	g,_:=game.CreateGame("Carlos")
	g.AddPlayer("Pedro")
	fmt.Println(g)
}