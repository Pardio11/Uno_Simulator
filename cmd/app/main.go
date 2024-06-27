package main

import (
	"net/http"
	"os"

	"unoSimulator/internal/handlers"
	"unoSimulator/internal/store"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)
func main (){
	
    router := mux.NewRouter()
	store := store.NewRedisHandler()
    gameHandler := handlers.NewGameHandler(store)

    home := handlers.HomeHandler{}
	router.HandleFunc("/join_uno", gameHandler.AddPlayer).Methods("POST")
    router.HandleFunc("/create_uno", gameHandler.CreateGame).Methods("POST")
    router.HandleFunc("/start_game", gameHandler.StartGame).Methods("POST")
	router.HandleFunc("/play", gameHandler.PlayTurn).Methods("POST")
    router.HandleFunc("/", home.ServeHTTP)
	godotenv.Load()
	go_port := os.Getenv("GO_PORT")
    http.ListenAndServe(go_port, router)

	/* fmt.Println("Running")
	g,_:=game.CreateGame("Carlos")
	g.CountCards()
	fmt.Printf("Pedro: %v\n",g.AddPlayer("Pedro"))
	fmt.Printf("Romero: %v\n",g.AddPlayer("Romero"))
	
	g.StartGame()
	
	for{
		fmt.Printf("Top Card: %v\n",g.DiscardDeck.TopCard())
		fmt.Println(g.NextPlayer())
		var input int
		input2:=""
		fmt.Println("\nEnter index:")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if input>=0 && input<len(g.NextPlayer().Deck) && g.NextPlayer().Deck[input].Color=="Any"{
			fmt.Println("\nEnter color:")
			_, err = fmt.Scan(&input2)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		g.PlayTurn(input,input2)
		fmt.Println(g.DiscardDeck)
		g.CountCards()
	} */
}