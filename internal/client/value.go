package client

import (
	"net/http"

	"github.com/zmb3/spotify"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

var (
	redirectURI = "http://localhost:8080/callback"
	authHandler = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(
			spotify.ScopeUserReadPrivate,
			spotify.ScopeUserLibraryRead,
		),
	)
	ch  = make(chan *oauth2.Token)
	srv = &http.Server{Addr: ":8080"}
)
