package main

import (
	"Go-Todo/src/app"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeNewHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
