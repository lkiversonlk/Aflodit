package main

import (
	"github.com/lkiversonlk/aflodit/context"
	"fmt"
	"os"
	"github.com/lkiversonlk/aflodit/bidder"
	"net/http"
	"log"
	"encoding/json"
)

func bid(w http.ResponseWriter, r *http.Request)  {
	
}

func main() {
	context := context.NewBidderContext()
	err := context.ConnectMongo("localhost", "aflodit")
	if err != nil {
		fmt.Println("error connecting mongodb")
		os.Exit(-1)
	} else {
		fmt.Println("connected to mongodb")
	}

	randomBidder := bidder.NewRandomBidder(context)

	http.HandleFunc("/bid", func(w http.ResponseWriter, r *http.Request) {
		result := randomBidder.Bid(nil)
		if bytes, err := json.Marshal(result); err != nil {
			fmt.Fprintf(w, "no result")
		}else {
			fmt.Fprintf(w, string(bytes))
		}

	})

	err = http.ListenAndServe(":5041", nil)

	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
