// a minimalist "echo" webserver
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Client represents the OAuth client information from the book
type Client struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURIs []string `json:"redirect_uris"`
	Scope        string   `json:"scope"`
}

// AuthServer holds the endpoint configuration
type AuthServer struct {
	AuthorizationEndpoint string
	TokenEndpoint         string
}

// Global variables to mimic the stateful nature of the JS example
// In a production app, these would be stored in a session or database
var (
	// Authorization server information
	authServer = AuthServer{
		AuthorizationEndpoint: "http://localhost:9001/authorize",
		TokenEndpoint:         "http://localhost:9001/token",
	}

	// Client information array
	clients = []Client{
		{
			ClientID:     "oauth-client-1",
			ClientSecret: "oauth-client-secret-1",
			RedirectURIs: []string{"http://localhost:9000/callback"},
			Scope:        "foo bar",
		},
	}

	// Global maps to store codes and requests (mimicking the JS logic)
	codes    = make(map[string]map[string]interface{})
	requests = make(map[string]map[string]interface{})

	// Templates
	templates = template.Must(template.ParseFiles("exercises/ap-A-ex-0/authorizationServer/files/index.html"))
)

func main() {
	// Note we're relying on the DefaultServerMux here. No explicit declaration
	// of a servemux. Which may be contrary to nbest practices?
	// 1. Static file serving (Equivalent to express.static)
	// This serves images, CSS, or JS files from the directory
	fs := http.FileServer(http.Dir("files"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.HandleFunc registers a handler function for the given pattern
	// 2. Root Route (Equivalent to app.get('/'))
	http.HandleFunc("/", indexHandler) // each request calls handler

	// Start Server (Equivalent to app.listen)
	host := "localhost"
	port := "9001"
	log.Printf("OAuth Authorization Server is listening at http://%s:%s", host, port)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

// handler echos the path component of the requested URL
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Only handle the exact root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Data to pass to the template
	data := map[string]interface{}{
		"Clients":    clients,
		"AuthServer": authServer,
	}

	// Render the template (Equivalent to res.render)
	err := templates.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
