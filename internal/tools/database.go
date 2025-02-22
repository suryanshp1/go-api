package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database Collection

type LoginDetails struct {
	Username  string
	AuthToken string
}

type CoinDetails struct {
	Username string
	Coins    float64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error("Failed to setup database")
		return nil, err
	}

	return &database, nil
}