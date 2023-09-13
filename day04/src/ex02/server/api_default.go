package main

import "C"
import (
	"encoding/json"
	"errors"
	// "fmt"
	"log"
	"net/http"
	"strconv"
)

func getPrice(candyType string) (int, error) {
	switch candyType {
	case "CE":
		return 10, nil
	case "AA":
		return 15, nil
	case "NT":
		return 17, nil
	case "DE":
		return 21, nil
	case "YR":
		return 23, nil
	default:
		return 0, errors.New("wrong candy type")
	}
}

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data Order

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Fatalln(err)
		}
		var response interface{}
		candyPrice, err := getPrice(data.CandyType)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response = struct {
				Error string
			}{"Wrong candy type!"}
		} else if data.CandyCount < 0 {
			w.WriteHeader(http.StatusBadRequest)
			response = struct {
				Error string
			}{"Negative candy count!"}
		} else if candyPrice * data.CandyCount > data.Money {
			amount := candyPrice * data.CandyCount - data.Money
			response = struct {
				Error string
			}{"You need " + strconv.Itoa(amount) + " more money!"}
			w.WriteHeader(http.StatusPaymentRequired)
		} else {
			change := data.Money - candyPrice*data.CandyCount
			thanks := ask_cow()
			response = struct {
				Change int
				Thanks string
			}{change, thanks}
			w.WriteHeader(http.StatusCreated)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
