package base64

import (
	"fmt"
	"testing"
)

func TestDecodeAccessToken(t *testing.T) {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlJEUTNSa1U0UXpJMU1rRTBSakV6T1RKR09UQTNPVEUwTjBNME1EZ3pRMFJDT0VFelF6UTBOdyJ9.eyJpc3MiOiJodHRwczovL2piZWxtb250LmF1dGgwLmNvbS8iLCJzdWIiOiJHemZnNzk3bGZMbUY5WWJLQlRmendiN1ZHaWhqRXpZVEBjbGllbnRzIiwiYXVkIjoiaHR0cHM6Ly9qYmVsbW9udC5hdXRoMC5jb20vYXBpL3YyLyIsImlhdCI6MTUxNTg2MTYxMiwiZXhwIjoxNTE1OTQ4MDEyLCJhenAiOiJHemZnNzk3bGZMbUY5WWJLQlRmendiN1ZHaWhqRXpZVCIsImd0eSI6ImNsaWVudC1jcmVkZW50aWFscyJ9.hAVg-T7BJXLZgTDfPEafWAuTp3J18iwNaHsHvRBdlLl75xGzulrBXWnEKGmEmoOwKyyFfu7RV_OJ_HFNxxy-ibcymFChHNjeNAQfnxbCe9lzCyFwulDEX84CrnBYdTuPDAzehKVhM1nIyVBt0ZC4tWHQElqYwD1bxD143HVPCfPKLPtsEZh1ErW29PH-NWc7S02BqVNKZuT5emXZh7ITWud87kcG2BAmaIds6a3b-iGmqxcLKMTaEAl4Xd0LrHyTBrZHKXSin43gX8GXmxg54IojGDmf2yGFVUIkXBTKHdy1M7ZHpwMFKpYEeZEFCqF8GHv-ngbgF46AWNezZD6sOQ"
	decodeToken, err := decodeAccessToken(token, StandardClaims{})
	if err != nil {
		t.Fatal("Error occured while base 64 decoding jwt")
	}
	fmt.Println(decodeToken.Claims)
}
