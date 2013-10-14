/* meeetup.com API */
package lag

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"

	"appengine"
	"appengine/urlfetch"
)

const (
	apiURL    = "http://api.meetup.com/2/events"
	key       = "3f27774e3316b736c4762233f53a6f"
	groupName = "Los-Angeles-Gophers"
)

var laLoc *time.Location

type Meetup struct {
	URL  string    `json:"event_url"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

func init() {
	laLoc, _ = time.LoadLocation("America/Los_Angeles")
}

func parseJSON(in io.Reader) (*Meetup, error) {
	dec := json.NewDecoder(in)
	var reply struct {
		Results []struct {
			URL  string `json:"event_url"`
			Name string `json:"name"`
			Time int64  `json:"time"`
		}
	}

	if err := dec.Decode(&reply); err != nil {
		return nil, err
	}

	if len(reply.Results) == 0 {
		return nil, nil
	}

	m := reply.Results[0]
	t := time.Unix(m.Time/1000, 0)
	if laLoc != nil {
		t = t.In(laLoc)
	}
	return &Meetup{m.URL, m.Name, t}, nil
}

func nextMeetup(ctx appengine.Context) (*Meetup, error) {
	query := url.Values{"key": {key}, "group_urlname": {groupName}}
	u := fmt.Sprintf("%s?%s", apiURL, query.Encode())

    client := urlfetch.Client(ctx)
	resp, err := client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return parseJSON(resp.Body)
}
