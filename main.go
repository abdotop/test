package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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
	fmt.Fprintf(w, "Votre IP publique est : %s ==> %s", strings.Split(ip, ":")[0],strings.Split(public_ip, ":")[0])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
