package base64

import (
	"testing"
)

func TestDecodeAccessToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	decodeToken, err := decodeAccessToken(token, StandardClaims{})
	if err != nil {
		t.Fatal("Error occured while base 64 decoding jwt")
	}
	if name := decodeToken.Claims.Name; name != "John Doe" {
		t.Errorf("Should equal the value of %s", name)
	}
	if subject := decodeToken.Claims.Subject; subject != "1234567890" {
		t.Errorf("Should equal the value of %s", subject)
	}
}
