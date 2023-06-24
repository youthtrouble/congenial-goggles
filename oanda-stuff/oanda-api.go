package oandastuff

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/youthtrouble/congenial-goggles/utils"
)

const oandaURL = "https://fxds-public-exchange-rates-api.oanda.com/cc-api/currencies?"

func FetchCurrentOandaRates() (*OandaRateResponse, *string, error) {

	lagos, _ := time.LoadLocation("Africa/Lagos")
	today := time.Now().In(lagos).Format("2006-01-02")
	todayNoFormat := time.Now().In(lagos)
	todayOneAM := time.Date(todayNoFormat.Year(), todayNoFormat.Month(), todayNoFormat.Day(), 1, 0, 0, 0, lagos)
	yesterday := time.Now().In(lagos).AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayNoFormat := time.Now().In(lagos).AddDate(0, 0, -1)
	dayBeforeYesterday := time.Now().In(lagos).AddDate(0, 0, -2).Format("2006-01-02")

	endDate := todayNoFormat.Format(time.RFC850)
	endpoint := fmt.Sprintf("base=GBP&quote=NGN&data_type=general_currency_pair&start_date=%s&end_date=%s", yesterday, today)
	if todayNoFormat.Before(todayOneAM) {
		endpoint = fmt.Sprintf("base=GBP&quote=NGN&data_type=general_currency_pair&start_date=%s&end_date=%s", dayBeforeYesterday, yesterday)
		endDate = yesterdayNoFormat.Format(time.RFC850)
	}

	var response OandaRateResponse
	err := executeOandaRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, nil, err
	}

	return &response, &endDate, nil
}

func executeOandaRequest(method, endpoint string, requestData, destination interface{}) error {

	url := fmt.Sprintf("%s%s", oandaURL, endpoint)
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	var req *http.Request

	if requestData == nil {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}
	}

	req.Header.Add("Host", "fxds-public-exchange-rates-api.oanda.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Origin", "https://www.oanda.com")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")

	var response *http.Response
	log.Print("request: ", req)
	response, err = utils.GetDebugClient().Do(req)
	if err != nil {
		return err
	}

	responseCode := response.StatusCode
	if responseCode != 200 && responseCode != 201 {
		log.Print("error processing request: ", response)
		return errors.New("error processing request")
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, destination)
}
