package main

import (
	"fmt"
	// "net/http"
	"os"
	"github.com/abdotop/octopus"
	// "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	public_ip := r.Header.Get("X-Forwarded-For")
	ip := r.RemoteAddr
	// if ip == "" {
	// 	 // Fallback to RemoteAddr if no X-Forwarded-For header is present
	// } else {
	// 	// X-Forwarded-For peut contenir plusieurs adresses IP séparées par des virgules
	// 	// Nous prenons la première qui est normalement l'adresse IP du client
	// 	ip = strings.Split(ip, ",")[0]
	// }
	fmt.Fprintf(w, "Votre IP publique est : %s ==> %s", ip,public_ip)
}

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
// }

func main() {
	app := octopus.New()

	app.Get("/", func(c *octopus.Ctx) {
		fmt.Println(c.RemoteIP())
	})

	app.Run(":" + os.Getenv("PORT"))
}
