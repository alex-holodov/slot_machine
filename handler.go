package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type handler struct {
	machine slotMachine
	key     []byte
}

type Response struct {
	Total int64
	Spins []playResult
	Jwt   string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := h.doPlay(body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h handler) doPlay(token []byte) (*Response, error) {
	balance, err := h.getClaim(token)
	if err != nil {
		return nil, err
	}

	if balance.Bet > balance.Chips {
		return nil, fmt.Errorf("not enougth money")
	}

	result, err := h.machine.play(balance.Bet)
	if err != nil {
		return nil, err
	}

	var total int64
	for _, r := range result {
		total += r.Total
	}

	newBalance := balance.Chips - balance.Bet + total

	t, err := h.newToken(balance.Uid, newBalance, balance.Bet)
	if err != nil {
		return nil, err
	}

	return &Response{
		Total: total,
		Spins: result,
		Jwt:   t,
	}, nil
}

func (h handler) newToken(id string, chips int64, bet int64) (string, error) {
	data := balanceClaim{
		Uid:   id,
		Chips: chips,
		Bet:   bet,
		StandardClaims: jwt.StandardClaims{
			Issuer: "atkins_machine",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(h.key)
}

func (h handler) getClaim(token []byte) (*balanceClaim, error) {
	t, err := jwt.ParseWithClaims(string(token), &balanceClaim{}, func(token *jwt.Token) (interface{}, error) {
		return h.key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*balanceClaim)
	if !ok {
		return nil, fmt.Errorf("bad claim value")
	}

	return claims, nil
}

type balanceClaim struct {
	Uid   string
	Chips int64
	Bet   int64
	jwt.StandardClaims
}
