package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Token contains typical values returned from REST API.
type Token struct {
	Raw       string                 `json:"raw"`       // The raw token.  Populated when you Parse a token
	Header    map[string]interface{} `json:"headers"`   // The first segment of the token
	Claims    StandardClaims         `json:"claims"`    // The second segment of the token
	Signature string                 `json:"signature"` // The third segment of the token.  Populated when you Parse a token
	Valid     bool                   `json:"valid"`     // Is the token valid?  Populated when you Parse/Verify a token
}

type StandardClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
	Name      string `json:"name,omitempty"`
}

type MapClaims map[string]interface{}

func decodeTokenSegment(segment string) ([]byte, error) {
	if l := len(segment) % 4; l > 0 {
		segment += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(segment)
}

func decodeAccessToken(accessToken string, claims StandardClaims) (*Token, error) {
	var err error
	token := &Token{Raw: accessToken}
	parts := strings.Split(accessToken, ".")
	var headerBytes []byte
	headerBytes, err = decodeTokenSegment(parts[0])
	if err = json.Unmarshal(headerBytes, &token.Header); err != nil {
		return token, fmt.Errorf("An error occurred unmarshalling token")
	}
	var claimBytes []byte
	if claimBytes, err = decodeTokenSegment(parts[1]); err != nil {
		return token, fmt.Errorf("An error occurred unmarshalling token")
	}
	dec := json.NewDecoder(bytes.NewBuffer(claimBytes))
	err = dec.Decode(&claims)
	token.Claims = claims
	return token, err
}

var (
	jwt = flag.String("jwt", "", "Please provide a jwt to be decoded")
)

func main() {
	flag.Parse()

	if len(*jwt) == 0 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	jwt, err := decodeAccessToken(*jwt, StandardClaims{})
	if err != nil {
		log.Fatal("Could not parse the JSON Web Token")
	}
	prettyJWT, err := json.MarshalIndent(jwt, "", " ")
	if err != nil {
		log.Fatal("Could not Marshal this type")
	}
	addNewLine := append(prettyJWT, '\n')
	os.Stdout.Write(addNewLine)
}
