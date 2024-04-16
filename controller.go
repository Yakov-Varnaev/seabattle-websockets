package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Yakov-Varnaev/seabattle-websockets/store"
	"github.com/gin-gonic/gin"
)

type CreateController struct{}

func (c *CreateController) Act() (gin.H, error) {
	game := store.NewGame()
	game, err := store.GetStore().Save(game)
	if err != nil {
		slog.Error("Game was not created.", "error", err.Error())
		return nil, err
	}
	slog.Info("Game was successfully created.")
	data := gin.H{
		"player1": fmt.Sprintf("/game/%s/player/%d", game.Id, 1),
		"player2": fmt.Sprintf("/game/%s/player/%d", game.Id, 2),
	}

	return data, nil
}

func createGame(c *gin.Context) {
	var controller CreateController
	data, err := controller.Act()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

type GameData struct {
	GameId   string
	PlayerId string
}

func GetGameDataFromContext(c *gin.Context) (*GameData, error) {
	gameId := c.Param("gameId")
	playerId := c.Param("playerId")

	if gameId == "" && playerId == "" {
		return nil, fmt.Errorf("gamedId and playerId are required")
	}

	gd := GameData{GameId: gameId, PlayerId: playerId}

	return &gd, nil
}

type RetrieveController struct {
	GameData
}

func (ctrl *RetrieveController) Act() (*store.Game, error) {
	slog.Info("Retrieving game", "id", ctrl.GameId, "player", ctrl.PlayerId)
	game, err := store.GetStore().Get(ctrl.GameId)
	if err != nil {
		slog.Error("Error while retrieving game instance", "id", ctrl.GameId)
		return nil, err
	}
	return game, nil
}

func NewRetrieveController(c *gin.Context) (*RetrieveController, error) {
	// get id from url params
	// get player id from url pararms

	gd, err := GetGameDataFromContext(c)
	if err != nil {
		slog.Error("Can get game data from gin.Context", "error", err.Error())
		return nil, err
	}
	controller := RetrieveController{*gd}

	return &controller, nil
}

func retrieveGame(c *gin.Context) {
	controller, err := NewRetrieveController(c)
	if err != nil {
		slog.Error("Error while creating controller", "error", err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	game, err := controller.Act()
	if err != nil {
		slog.Error("Error while retrieving game instance", "error", err.Error())
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusOK, game)
}

type UpdateController struct {
	GameData
}

func NewUpdateController(c *gin.Context) (*UpdateController, error) {
	gd, err := GetGameDataFromContext(c)
	if err != nil {
		slog.Error("Can get game data from gin.Context", "error", err.Error())
		return nil, err
	}
	controller := UpdateController{*gd}

	return &controller, nil
}

func updateGame(c *gin.Context) {}
