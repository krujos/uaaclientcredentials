package uaaclientcredentials

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//UaaClientCredentials provides a token for a given clientId and clientSecret.
//The token is refreshed for you according to expires_in
type UaaClientCredentials struct {
	uaaURI             *url.URL
	clientID           string
	clientSecret       string
	authorizationToken string
	expiresAt          time.Time
	scope              string
	skipSSLValidation  bool
}

//GetBearerToken returns a currently valid bearer token to use against the
//CF API. You should not cache the token as the library will handle updating
//it if it's expired.
func (creds *UaaClientCredentials) GetBearerToken() string {
	return "bearer " + creds.authorizationToken
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

func (creds *UaaClientCredentials) getToken() error {

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: creds.getTLSConfig(),
		},
	}

	resp, err := client.Get(creds.uaaURI.String() + "/oauth/token?grant_type=client_credentials")
	if nil != err {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	type UaaTokenRepsone struct {
		expireIn           int
		authorizationToken string
		scope              string
	}

	token := &UaaTokenRepsone{}
	json.Unmarshal(body, token)

	if nil != err {
		return err
	}

	creds.authorizationToken = token.authorizationToken
	duration, _ := time.ParseDuration(strconv.Itoa(token.expireIn) + "s")
	creds.expiresAt = time.Now().Add(duration)
	return nil

}
