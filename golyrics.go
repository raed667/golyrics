// Package golyrics contains utility functions for getting lyrics.
package golyrics

	import (
	"log"
	"net/http"
	"io/ioutil"
	"net/url"
)

// Fetch used with 2 arguments, a : Artist and t : Tititle, return 1 string with lyrics or error message
func Fetch(a string, t string) string {

	artist:=a
	title:=t

	var Url *url.URL
	Url, err := url.Parse("http://makeitpersonal.co")
	if err != nil {
		panic("boom")
	}

	Url.Path += "/lyrics"
	parameters := url.Values{}
	parameters.Add("artist", artist)
	parameters.Add("title", title)
	Url.RawQuery = parameters.Encode()
	resp, err := http.Get(Url.String())

	if err != nil {
		log.Fatalf("http.Get => %v", err.Error())
	}
	body, _ := ioutil.ReadAll(resp.Body)
	
	return string(body)
}
