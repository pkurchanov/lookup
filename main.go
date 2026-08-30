package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// TODO:
// [~] Satellites (https://www.n2yo.com/api/)
// [ ] Astronomy (https://docs.astronomyapi.com/)
// [ ] Aviation (https://openskynetwork.github.io/opensky-api/rest.html)

func main() {
	godotenv.Load()
	fmt.Println("Hello, world!")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	lat := 56
	lng := 93
	alt := 162
	r := 30
	cat := 0
	n2yoKey := os.Getenv("N2YO_KEY")
	satEndpoint := fmt.Sprintf("https://api.n2yo.com/rest/v1/satellite/above/%d/%d/%d/%d/%d/&apiKey=%s", lat, lng, alt, r, cat, n2yoKey)

	resp, err := client.Get(satEndpoint)
	if err != nil {
		log.Fatalf("Request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Couldn't read body: %w", err)
	}

	fmt.Println("|", resp.Status, "|")
	fmt.Println(string(body))
}
