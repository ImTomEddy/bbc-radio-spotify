package spotify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var token = ""

var client = &http.Client{}

var userID = os.Getenv("USER_ID")

const (
	tokenEndpoint  = "https://accounts.spotify.com/api/token"
	searchEndpoint = "https://api.spotify.com/v1/search"
)

//Authenticate authenticates the app using the refresh token
func Authenticate() (err error) {
	q := url.Values{}

	q.Add("grant_type", "refresh_token")
	q.Add("refresh_token", os.Getenv("REFRESH_TOKEN"))

	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(q.Encode()))

	if err != nil {
		return err
	}

	b64 := base64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID") + ":" + os.Getenv("CLIENT_SECRET")))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+b64)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return errors.New(fmt.Sprint("spotify auth status ", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var b struct {
		Token string `json:"access_token"`
	}

	err = json.Unmarshal(body, &b)

	if err != nil {
		return err
	}

	token = b.Token

	return err
}
