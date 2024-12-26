package auth

import "testing"

func TestToken(t *testing.T) {
	token, err := NewToken("123", "user")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	auth, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzUyMDYzMzcsIklkZW50aXR5IjoiMTIzIiwiUm9sZSI6InVzZXIifQ._7plkoQRMMoGhFEztFYdNpLxC_SThlALcGLTTn90A3U")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(auth)
}
