// a minimalist "echo" webserver
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Global variables to mimic the stateful nature of the JS example
// In a production app, these would be stored in a session or database
var (
	accessToken  string
	refreshToken string
	scope        string
	// Parse templates once at startup for efficiency
	templates = template.Must(template.ParseFiles("exercises/ap-A-ex-0/client/files/index.html"))
)

func main() {
	// 1. Static file serving (Equivalent to express.static)
	// This serves images, CSS, or JS files from the directory
	fs := http.FileServer(http.Dir("files"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.HandleFunc registers a handler function for the given pattern
	// 2. Root Route (Equivalent to app.get('/'))
	http.HandleFunc("/", indexHandler) // each request calls handler

	// 3. Start Server (Equivalent to app.listen)
	// Start the server on port 9000
	host := "localhost"
	port := "9000"
	log.Printf("OAuth Client is listening at http://%s:%s", host, port)

	//fmt.Println("Server is starting on port 9000...")
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

	// Data object to pass to the template
	data := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"scope":         scope,
	}

	// Render the template (Equivalent to res.render)
	err := templates.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
