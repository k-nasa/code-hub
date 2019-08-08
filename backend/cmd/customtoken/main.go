package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/voyagegroup/treasure-app/firebase"
)

type FirebaseCustomToken struct {
	Kind         string `json:"kind"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	IsNewUser    bool   `json:"isNewUser"`
}

// https://firebase.google.com/docs/auth/admin/create-custom-tokens#sign_in_using_custom_tokens_on_clients
func main() {

	if len(os.Args) != 3 {
		log.Fatal("Need 2 argument but got ", len(os.Args)-1)
	}

	uid := os.Args[1]
	if len(uid) == 0 {
		log.Fatal("uid flag is missing.")
	}

	if len(uid) == 0 {
		log.Fatal("uid flag is missing.")
	}

	tokenFileName := os.Args[2]
	if len(tokenFileName) == 0 {
		log.Fatal("token file name flag is missing.")
	}

	if len(tokenFileName) == 0 {
		log.Fatal("token file name flag is missing.")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	client, err := firebase.InitAuthClient()
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomToken(context.Background(), uid)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	webapikey := os.Getenv("FIREBASE_WEB_API_KEY")
	endpoint := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s", webapikey)

	if webapikey == "" {
		log.Fatal("firebase Web API key is missing")
	}

	body := []byte(fmt.Sprintf(`
	{
		"token":"%s",
		"returnSecureToken":true
	}
	`, token))
	values := url.Values{}
	values.Set("returnSecureToken", "true")
	values.Set("token", token)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{}
	resp, err := c.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	firebaseCustomToken := &FirebaseCustomToken{}
	if err := json.Unmarshal(respBytes, firebaseCustomToken); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(tokenFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, firebaseCustomToken.IDToken)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBytes))
	fmt.Printf("%s created", tokenFileName)
}
