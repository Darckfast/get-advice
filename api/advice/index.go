package getadvice

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Advice struct {
	Slip struct {
		ID     int    `json:"id"`
		Advice string `json:"advice"`
	} `json:"slip"`
}

func Handler(writer http.ResponseWriter, _ *http.Request) {
	log.Println("Incoming request")

	var res *http.Response

	res, _ = http.Get("https://api.adviceslip.com/advice")

	if res.StatusCode != 200 {
		writer.WriteHeader(res.StatusCode)

		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		log.Println("Error calling adviceslip api", res.StatusCode, string(body))
		log.Println("Sending response", res.StatusCode)
		return
	}

	var advice Advice

	json.NewDecoder(res.Body).Decode(&advice)

	writer.WriteHeader(200)

	writer.Header().Set("Cache-Control", "public, max-age=600")
	writer.Header().Set("CDN-Cache-Control", "public, max-age=600")
	writer.Header().Set("Vercel-CDN-Cache-Control", "public, max-age=600")

	writer.Write([]byte(advice.Slip.Advice))

	log.Println("Sending response", 200)
}
