package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yakov-Varnaev/seabattle-websockets/store"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"
)

type CreateApiSuite struct {
	suite.Suite
}

func (suite *CreateApiSuite) SetupSuite() {
	store.Init("local")
}

func (suite *CreateApiSuite) TestGameCreate() {
	router := CreateRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/game", nil)
	router.ServeHTTP(w, req)
	games, _ := store.GetStore().List()
	game := games[0]

	expectedData := game.GetPlayerLinks()

	var responseData map[string]string

	_ = json.Unmarshal([]byte(w.Body.String()), &responseData)

	suite.Equal(http.StatusOK, w.Code)
	suite.Equal(expectedData, responseData)
}

func TestCreateApiSuite(t *testing.T) {
	suite.Run(t, new(CreateApiSuite))
}

type ReadApiSuite struct {
	suite.Suite
	game *store.Game
}

func (suite *ReadApiSuite) SetupSuite() {
	store.Init("local")
	game := store.NewGame()
	game, err := store.GetStore().Save(game)
	if err != nil {
		panic(err.Error())
	}
	suite.game = game
}

func (suite *ReadApiSuite) TestRetriveExistingGame() {
	var ok bool
	router := CreateRouter()

	url := fmt.Sprintf("/game/%s/player/1", suite.game.Id)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	ok = suite.Equal(http.StatusOK, w.Code)
	if !ok {
		return
	}

	var resultGame *store.Game
	err := json.Unmarshal([]byte(w.Body.String()), &resultGame)
	ok = suite.NoError(err)
	if !ok {
		return
	}

	suite.True(cmp.Equal(suite.game, resultGame))
}

func TestReadApiSuite(t *testing.T) {
	suite.Run(t, new(ReadApiSuite))
}
