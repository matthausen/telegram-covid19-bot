package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot"
)

type Response struct {
	Data []Data
}

type Data struct {
	Country   string
	Confirmed int
	Recovered int
	Critical  int
	Deaths    int
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func fetchCovidData(c string) string {

	url := "https://covid-19-data.p.rapidapi.com/country?format=json&name=" + c

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "covid-19-data.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", goDotEnvVariable("RAPIDAPI_KEY"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var responseObject []Data
	json.Unmarshal([]byte(body), &responseObject)

	data := Data{
		responseObject[0].Country,
		responseObject[0].Confirmed,
		responseObject[0].Recovered,
		responseObject[0].Critical,
		responseObject[0].Deaths,
	}

	confirmed := strconv.Itoa(data.Confirmed)
	recovered := strconv.Itoa(data.Recovered)
	critical := strconv.Itoa(data.Critical)
	deaths := strconv.Itoa(data.Deaths)

	return ("🌎 Country: " + data.Country + "\n" +
		"🤒 Confirmed: " + confirmed + "\n" +
		"🙂 Recovered: " + recovered + "\n" +
		"🤕 Critical: " + critical + "\n" +
		"💀 Deaths: " + deaths)
}

func main() {
	bot := tbot.New(goDotEnvVariable("TELEGRAM_TOKEN"))
	c := bot.Client()
	bot.HandleMessage("/help", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID,
			"Check the covid-19 stats per country, simply typing 'covid' followed by the country name: e.g. - covid usa or covid china")
	})
	bot.HandleMessage("covid .+", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		text := strings.TrimPrefix(m.Text, "covid ")
		c.SendMessage(m.Chat.ID, fetchCovidData(text))
	})
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", nil)
}
