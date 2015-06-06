package uaaclientcredentials

import "net/url"

//UaaClientCredentials provides a token for a given clientId and clientSecret.
//The token is refreshed for you according to expires_in
type UaaClientCredentials struct {
	uaaURI             *url.URL
	clientID           string
	clientSecret       string
	authorizationToken string
	expiresAt          string
}

//GetBearerToken returns a currently valid bearer token to use against the
//CF API. You should not cache the token as the library will handle updating
//it if it's expired.
func (creds *UaaClientCredentials) GetBearerToken() string {
	return ""
}

//New UaaClientCredentials factory
func New(uaaURI *url.URL, clientID string, clientSecret string) (*UaaClientCredentials, error) {

	creds := &UaaClientCredentials{
		uaaURI:       uaaURI,
		clientID:     clientID,
		clientSecret: clientSecret,
	}

	return creds, nil
}
