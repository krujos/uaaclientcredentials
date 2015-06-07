package uaaclientcredentials

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//UaaClientCredentials provides a token for a given clientId and clientSecret.
//The token is refreshed for you according to expires_in
type UaaClientCredentials struct {
	uaaURI            *url.URL
	clientID          string
	clientSecret      string
	accessToken       string
	expiresAt         time.Time
	scope             string
	skipSSLValidation bool
}

//UAATokenResponse is the struct version of the json /oauth/token gives us
//when we ask for client credentials.
type UAATokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Jti         string `json:"jti"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

//GetBearerToken returns a currently valid bearer token to use against the
//CF API. You should not cache the token as the library will handle updating
//it if it's expired.
func (creds *UaaClientCredentials) GetBearerToken() string {
	return "bearer " + creds.accessToken
}

//New UaaClientCredentials factory
func New(uaaURI *url.URL, skipSSLValidation bool, clientID string,
	clientSecret string) (*UaaClientCredentials, error) {

	if len(clientID) < 1 {
		return nil, errors.New("clientID cannot be empty")
	}

	if len(clientSecret) < 1 {
		return nil, errors.New("clientSecret cannot be empty")
	}

	creds := &UaaClientCredentials{
		uaaURI:            uaaURI,
		clientID:          clientID,
		clientSecret:      clientSecret,
		skipSSLValidation: skipSSLValidation,
	}

	return creds, nil
}

func (creds *UaaClientCredentials) getTLSConfig() *tls.Config {
	if creds.skipSSLValidation {
		return &tls.Config{InsecureSkipVerify: true}
	}
	return &tls.Config{}
}

func (creds *UaaClientCredentials) getClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: creds.getTLSConfig(),
		},
	}
}

func (creds *UaaClientCredentials) getJSON() ([]byte, error) {
	client := creds.getClient()
	url := creds.uaaURI.String() + "/oauth/token?grant_type=client_credentials"

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(creds.clientID, creds.clientSecret)

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}

	log.Println("Server responded with status code", resp.StatusCode)
	if resp.StatusCode != 200 {
		return nil, errors.New("UAA responded with bad status (" +
			strconv.Itoa(resp.StatusCode) + ")")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (creds *UaaClientCredentials) getToken() error {

	body, err := creds.getJSON()

	var token UAATokenResponse
	json.Unmarshal(body, &token)

	if nil != err {
		return err
	}

	creds.accessToken = token.AccessToken
	duration, _ := time.ParseDuration(strconv.Itoa(token.ExpiresIn) + "s")
	creds.expiresAt = time.Now().Add(duration)
	return nil

}
