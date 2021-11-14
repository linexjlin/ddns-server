package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var service *Service

func main() {
	addr := flag.String("l", ":8010", "-l :8010")
	flag.Parse()
	cf, err := NewCloudFlare(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	if err != nil {
		panic(err)
	}
	service = &Service{cf: cf}
	http.HandleFunc("/UpdateDNS", UpdateDNS)
	log.Println("Server Listen on:", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
