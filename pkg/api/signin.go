package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

var secretKey = []byte("super_secret_signature_key")

type SigninRequest struct {
	Password string `json:"password"`
}

type SigninResponse struct {
	Token string `json:"token"`
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var req SigninRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJson(w, ErrorResponse{Error: "wrong request format"})
		return
	}

	envPass := os.Getenv("TODO_PASSWORD")
	if envPass == "" {
		sendJson(w, ErrorResponse{Error: "auth is failed"})
		return
	}
	if req.Password != envPass {
		sendJson(w, ErrorResponse{Error: "wrong password"})
		return
	}
	passHash := hash(envPass)
	claims := jwt.MapClaims{
		"password_hash": passHash,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, e := token.SignedString(secretKey)
	if e != nil {
		sendJson(w, ErrorResponse{Error: e.Error()})
		return
	}
	sendJson(w, SigninResponse{Token: signedToken})
}
func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
