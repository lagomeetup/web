package lag

import (
    "net/http"
)

func redirect(from, to string) {
	http.Handle(from, http.RedirectHandler(to, http.StatusTemporaryRedirect))
}

// Just a bunch of redirects, / is handled in app.yaml
func init() {
	redirect("/meetup", "http://www.meetup.com/Los-Angeles-Gophers/")
	redirect("/ideas", "https://trello.com/b/KQSMtS7P/l-a-gophers-meetup-ideas")
	redirect("/videos", "http://www.youtube.com/watch?v=l_8z4fc-qNw&list=PLxGjbtXdLtAMtjF5C6dAvO9IZGANBuYS0")
	redirect("/talks", "https://github.com/lagomeetup/talks")
}
