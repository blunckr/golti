package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

type OidAuth struct {
	ResponseType   string `url:"response_type"`
	RedirectURI    string `url:"redirect_uri"`
	ResponseMode   string `url:"response_mode"`
	ClientID       string `url:"client_id"`
	Scope          string `url:"scope"`
	State          string `url:"state"`
	LoginHint      string `url:"login_hint"`
	Prompt         string `url:"prompt"`
	LtiMessageHint string `url:"lti_message_hint"`
	Nonce          string `url:"nonce"`
}

func initHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if err := r.ParseForm(); err != nil {
		// handle error
	}

	for key, values := range r.PostForm {
		fmt.Println("=")
		fmt.Println(key)
		fmt.Println(values)
	}
	auth := OidAuth{
		ResponseType:   "id_token",
		ResponseMode:   "form_post",
		Scope:          "openid",
		Prompt:         "none",
		RedirectURI:    r.FormValue("target_link_uri"),
		ClientID:       r.FormValue("client_id"),
		LoginHint:      r.FormValue("login_hint"),
		LtiMessageHint: r.FormValue("lti_message_hint"),
	}

	values, _ := query.Values(auth)

	http.Redirect(w, r, "https://canvas.instructure.com/api/lti/authorize_redirect?"+values.Encode(), http.StatusFound)
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if err := r.ParseForm(); err != nil {
		// handle error
	}

	for key, values := range r.PostForm {
		fmt.Println("=")
		fmt.Println(key)
		fmt.Println(values)
	}
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// TODO getting the id_token, need to validate and unpack
}

func main() {

	fmt.Println("vim-go")
	pubKey := MakeKeys()
	PrinkJwk(pubKey)

	http.HandleFunc("/init", initHandler)
	http.HandleFunc("/", allHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
