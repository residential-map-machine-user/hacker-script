package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Set initial variables
	accountSid := strings.TrimSpace(os.Getenv("TWILIO_ACCOUNT_SID"))
	authToken := strings.TrimSpace(os.Getenv("TWILIO_AUTH_TOKEN"))
	myNumber := strings.TrimSpace(os.Getenv("MY_NUMBER"))
	twilioNumber := strings.TrimSpace(os.Getenv("TWILIO_NUMBER"))
	fmt.Println(accountSid)
	fmt.Println(authToken)
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	// Build out the data for our message
	v := url.Values{}
	v.Set("To", "+8109052330123")
	v.Set("From", "+14356592852")
	v.Set("Body", "Brooklyn's in the house!")
	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(byteArray))

}
