package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	oandastuff "github.com/youthtrouble/congenial-goggles/oanda-stuff"
)

func OandaHandler(w http.ResponseWriter, r *http.Request) {

	oandaRates, time, err := oandastuff.FetchCurrentOandaRates()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseMessaage := OandaResponse{
		Message: fmt.Sprintf("Current GBP/NGN rates: ₦%s\n Time: %s\n", oandaRates.Response[0].AverageBid, *time),
	}
	json.NewEncoder(w).Encode(responseMessaage)
}

type OandaResponse struct {
	Message string `json:"message"`
}
