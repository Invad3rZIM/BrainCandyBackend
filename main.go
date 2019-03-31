package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/hey" {
		//http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "hi")
}

func main() {

	h := NewHandler()

	http.HandleFunc("/users/new", h.NewUserHandler) // done

	http.HandleFunc("/secrets/new", h.NewSecretHandler)  //done
	http.HandleFunc("/secrets/get", h.GetSecretsHandler) //working on

	http.HandleFunc("/messages/send", h.SendMessageHandler) //done

	http.HandleFunc("/messages/getnew", h.GetNewMessagesHandler)
	http.HandleFunc("/messages/getold", h.GetOldMessagesHandler)
	http.HandleFunc("/messages/hasnew", h.HasNewMessagesHandler)
	http.HandleFunc("/messages/hasold", h.HasOldMessagesHandler)

	appengine.Main() // Starts the server to receive requests
}
