package main

import (
	"Go-Todo/src/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./test.db") // flag.Args로 실행 인자로 db 파일 명을 설정할 수 있다.
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
