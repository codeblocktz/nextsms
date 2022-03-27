package nextsms

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	req, err := http.NewRequest("POST", BASE_URL+"/api/sms/v1/single", httpBody)

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

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	fmt.Println(string(body))

	return resp, true
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

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	fmt.Println(string(body))

	return resp, true
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
