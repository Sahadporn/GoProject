package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		City             string `json:"city"`
		Country          string `json:"country"`
		PopulationCounts []struct {
			Year       string `json:"year"`
			Value      string `json:"value"`
			Sex        string `json:"sex"`
			Reliabilty string `json:"reliabilty"`
		} `json:"populationCounts"`
	} `json:"data"`
}

var cities = os.Getenv("CITIES")

func callAPI(city string) {
	var requestBody = []byte(fmt.Sprintf(`
		{
			"city": "%s"
		}
	`, city))

	url := os.Getenv("API_URL")

	fmt.Println("\nCalling API for " + strings.Title(city) + "...")

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)

	if !responseObject.Error {
		fmt.Printf("Country: %s ✦ City: %s ✦ Population: %s ✦ Year: %s \n", responseObject.Data.Country, responseObject.Data.City, responseObject.Data.PopulationCounts[0].Value, responseObject.Data.PopulationCounts[0].Year)
	} else {
		fmt.Println(responseObject.Msg)
	}

}

func main() {
	fmt.Println("Hello")
	var citiesArray = strings.Split(cities, ",")
	for _, city := range citiesArray {
		callAPI(string(city))
	}

	fmt.Println("\nGoodbye")
}
