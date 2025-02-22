package handlers

import (
	"encoding/json"
	"net/http"

	"goapi/api"
	"goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)


func GetCoinsBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}

	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error("Failed to decode request")
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error("Failed to setup database")
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	var response = api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: tokenDetails.Coins,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)


	if err != nil {
		log.Error("Failed to encode response")
		api.InternalErrorHandler(w)
		return
	}

}