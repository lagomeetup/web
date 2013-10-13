/* meeetup.com API */
package lag

import (
	"fmt"
	/*
		"time"
		"net/http"
		"net/url"

		"log"
	*/
	"encoding/json"
	"io"
	"os"
)

const (
	apiURL    = "http://api.meetup.com/2/events"
	key       = "3f27774e3316b736c4762233f53a6f"
	groupName = "Los-Angeles-Gophers"
)

type Meetup struct {
	URL  string `json:"event_url"`
	Name string `json:"name"`
	Time int64  `json:"time"`
}

func parseJSON(in io.Reader) (*Meetup, error) {
	dec := json.NewDecoder(in)
	var reply struct {
		Results []Meetup `json:"results"`
	}

	if err := dec.Decode(&reply); err != nil {
		return nil, err
	}

	if len(reply.Results) == 0 {
		return nil, nil
	}

	return &reply.Results[0], nil
}

func main() {
	/*
		query := url.Values{"key": {key}, "group_urlname": {groupName}}
		u := fmt.Sprintf("%s?%s", apiURL, query.Encode())

		resp, err := http.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		//t := time.Unix(1381971600, 0)
		fmt.Println(resp.Status)
	*/

	file, err := os.Open("res.json")
	if err != nil {
		fmt.Printf("error open: %s\n", err)
		os.Exit(1)
	}

	meetup, err := parseJSON(file)
	if err != nil {
		fmt.Printf("error decode: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", *meetup)
}
