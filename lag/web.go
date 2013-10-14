package lag

import (
	"fmt"
	"net/http"

	"appengine"
)

func redirect(from, to string) {
	http.Handle(from, http.RedirectHandler(to, http.StatusTemporaryRedirect))
}

func next(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	mp, err := nextMeetup(ctx)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	if mp == nil {
		fmt.Fprintf(w, "No meetup scheduled")
		return
	}

	time := mp.Time.Format("02/01/2006 15:04")
	fmt.Fprintf(w, "<a href=\"%s\">%s</a> [%s]", mp.URL, mp.Name, time)
}

// Just a bunch of redirects, / is handled in app.yaml
func init() {
	redirect("/meetup", "http://www.meetup.com/Los-Angeles-Gophers/")
	redirect("/ideas", "https://trello.com/b/KQSMtS7P/l-a-gophers-meetup-ideas")
	redirect("/videos", "http://www.youtube.com/watch?v=l_8z4fc-qNw&list=PLxGjbtXdLtAMtjF5C6dAvO9IZGANBuYS0")
	redirect("/talks", "https://github.com/lagomeetup/talks")

	http.HandleFunc("/next", next)
}
