package tools

import (
	"time"
)

type mockDB struct {}


var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234",
		Username:  "alex",
	},
	"bob": {
		AuthToken: "5678",
		Username:  "bob",
	},
	"charlie": {
		AuthToken: "9012",
		Username:  "charlie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"bob": {
		Coins: 200,
		Username: "bob",
	},
	"charlie": {
		Coins: 300,
		Username: "charlie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB calls

	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB calls

	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}