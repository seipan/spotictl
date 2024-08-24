package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Auth(ctx context.Context) {

}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	token, err := authHandler.Token(r.Context(), "state-token", r)
	if err != nil {
		http.Error(w, "failed to get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != "state-token" {
		http.NotFound(w, r)
		log.Fatalf("no match state: %s != %s\n", st, "state-token")
	}

	fmt.Fprintf(w, "success! you can now use the client.")
	ch <- token
}
