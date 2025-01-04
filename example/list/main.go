package main

import (
	"net/http"

	"github.com/troygilman0/gong"
)

func main() {
	db := newUserDatabase()

	g := gong.New(http.NewServeMux())

	g.Route("/", homeView{}, func(r gong.Route) {
		r.Route("users", listView{
			db: db,
			UserView: userView{
				db: db,
			},
		}, nil)
		r.Route("user/{name}", testView{
			db: db,
			UserView: userView{
				db: db,
			},
		}, nil)
	})

	if err := http.ListenAndServe(":8080", g); err != nil {
		panic(err)
	}
}
