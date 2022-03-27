package nextsms

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BASE_URL     = "https://messaging-service.co.tz"
	CONTENT_TYPE = "application/json"
)

type NextSms interface {
	Send(c *Credentials) (interface{}, bool)
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SingleSms struct {
	Text string `json:"text"`
	To   string `json:"to"`
	From string `json:"from"`
}

type statusResponse struct {
	groupId     int
	groupName   string
	id          int
	name        string
	description string
}

type messageResponse struct {
	to       string
	status   statusResponse
	smsCount int
}

type NextSmsResponse struct {
	messages messageResponse
}

type MultipleSms struct {
	Messages []SingleSms `json:"messages"`
}

func (n SingleSms) Send(c *Credentials) (interface{}, bool) {

	body, err := json.Marshal(n)

	if err != nil {
		fmt.Println(err)
		return err, false
	}

	httpBody := bytes.NewBuffer(body)

	client := &http.Client{}
	req, err := http.NewRequest("POST", BASE_URL+"/api/sms/v1/text/single", httpBody)

	if err != nil {
		fmt.Println(err)
		return err, false
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	defer res.Body.Close()

	var result *NextSmsResponse
	error := json.NewDecoder(res.Body).Decode(&result)

	if res.StatusCode != 200 || error != nil {
		return result, false
	}
	return result, true
}

func (n MultipleSms) Send(c *Credentials) (interface{}, bool) {

	body, err := json.Marshal(n)

	if err != nil {
		fmt.Println(err)
		return err, false
	}

	httpBody := bytes.NewBuffer(body)

	client := &http.Client{}
	req, err := http.NewRequest("POST", BASE_URL+"/api/sms/v1/text/multi", httpBody)

	if err != nil {
		fmt.Println(err)
		return err, false
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	defer res.Body.Close()

	var result *NextSmsResponse
	error := json.NewDecoder(res.Body).Decode(&result)

	if res.StatusCode != 200 || error != nil {
		return result, false
	}
	return result, true
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
