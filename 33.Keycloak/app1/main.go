package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var(
	clientID="app"
	clientSecret="..."
)

func main() {
	ctx := context.Background()
	provider, err :=  oidc.NewProvider(ctx, "http://10.39.52.113:8080/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID = clientID,
		ClientSecret = clientSecret,
		Endpoint: provider.Endpoint(),
		RedirectURL: "http://10.39.52.113:8081/auth/callback",
		Scopes: []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "exemplo"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.Redirect (w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request){
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State doesnt match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange (ctx, r.URL.Query().Get("code"))

		if err != nil{
			http.Error(w, "problema ao trocar token", http. StatusInternalError )
			return
		}

		rawIDToken, err := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "problema ao pegar id token", http. StatusInternalError )
			return
		}

		res := struct {
			OAuth2Token *oauth.Token
			IDToken string
		}{
			OAuth2Token, rawIDToken
		}

		data, _ := json.MarshalIndent(res, "", "    " )
		w.Write(data)

	})

	log.Fatal(http.listenAndServe(":8081", nil))
}