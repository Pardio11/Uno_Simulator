package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unoSimulator/internal/deck"
	"unoSimulator/internal/game"
)
type nameJSON struct{
	Game string `json:"game"`
	Name string `json:"name"`
}

type playerPass struct{
	Password string `json:"password"`
	Deck deck.Deck `json:"deck"`
}
type playerPlay struct{
	GameId string `json:"gameId"`
	Password string `json:"password"`
	Index int `json:"index"`
	Color string `json:"color"`
}
type gamePass struct{
	Id string `json:"id"`
}
type gameStart struct{
	Turn string `json:"turn"`
	TopCard deck.Card `json:"topCard"`
}

type gameStore interface {
    CreateEditGame(game.Game) error
	GameExist(string) bool
	GetGame(string) (game.Game,error)
    /* EndGame(name string)  error
    
	StartGame(name string, recipe game.Game) error
	NextTurn(name string, recipe game.Game) error
    PlayTurn(name string, recipe game.Game) error */
}

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Uno Simulator"))
}

type GameHandler struct {
    store gameStore
}

func NewGameHandler(s gameStore) *GameHandler {
    return &GameHandler{
        store: s,
    }
}

func (h GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Game")

	searchID := true
	var g game.Game
	var err error
	for searchID{
		g,err = game.CreateGame()
		if !h.store.GameExist(g.Id){
			searchID = false
		}
	}

	if err != nil {
		fmt.Println(err)
		InternalServerErrorHandler(w,r)
	}

	h.store.CreateEditGame(g)
	id:= gamePass{
		Id: g.Id,
	}

	json, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err)
		InternalServerErrorHandler(w,r)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func (h GameHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	var player nameJSON
	if err := json.NewDecoder(r.Body).Decode(&player); err!=nil{
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte("Payload Is Not Formatted as Expected"))
		return
	}
	if player.Name == "" {
        BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte(`Name can't be empty`))
		return
    }
	if !h.store.GameExist(player.Game) {
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte("Game doesn't exist"))
		return
	}
	g,err:= h.store.GetGame(player.Game)
	if err != nil{
		InternalServerErrorHandler(w,r)
		return
	}
	newPlayer,err:=g.AddPlayer(player.Name)
	if err != nil{
		InternalServerErrorHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte(err.Error()))
		return
	}
	p := playerPass{
		Password: newPlayer.Id,
		Deck: newPlayer.Deck,
	}
	h.store.CreateEditGame(g)
	json, _ := json.Marshal(p)
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func (h GameHandler) StartGame(w http.ResponseWriter, r *http.Request) {
	var game gamePass
	if err := json.NewDecoder(r.Body).Decode(&game); err!=nil{
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte("Payload Is Not Formatted as Expected"))
		return
	}

	if !h.store.GameExist(game.Id) {
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte("Game doesn't exist"))
		return
	}

	g,err:= h.store.GetGame(game.Id)
	if err != nil{
		InternalServerErrorHandler(w,r)
		return
	}

	if err = g.StartGame();err != nil{
		InternalServerErrorHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s:=gameStart{
		Turn: g.NextPlayer().Name,
		TopCard: g.DiscardDeck.TopCard(),
	}
	json, _ := json.Marshal(s)
	h.store.CreateEditGame(g)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func (h GameHandler) PlayTurn(w http.ResponseWriter, r *http.Request) {
	var play playerPlay
	if err := json.NewDecoder(r.Body).Decode(&play); err!=nil{
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte("Payload Is Not Formatted as Expected"))
		return
	}

	if !h.store.GameExist(play.GameId) {
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte("Game doesn't exist"))
		return
	}

	g,err:= h.store.GetGame(play.GameId)
	if err != nil{
		InternalServerErrorHandler(w,r)
		return
	}
	if play.Password != g.NextPlayer().Id{
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte("Bad Password"))
		return
	}
	if err = g.PlayTurn(play.Index, play.Color);err != nil{
		BadRequestHandler(w,r)
		w.Write([]byte("\n"))
		w.Write([]byte(err.Error()))
		return
	}
	h.store.CreateEditGame(g)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(g.NextPlayer().Name))
}


func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 Not Found"))
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 Bad Request"))
}