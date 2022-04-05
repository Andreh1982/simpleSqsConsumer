package sqshandlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"simpleSQSconsumer/models"
)

func SendToAlertsApi(messageNew models.AlertSQS) {

	method := "POST"
	apiUrl := "http://localhost:8081/v1/alerts"
	alertSQS, err := json.Marshal(messageNew)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(alertSQS)
	client := &http.Client{}
	req, err := http.NewRequest(method, apiUrl, reader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
