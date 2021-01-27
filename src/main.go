package main

import (
	"Go-Todo/src/app"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./test.db") // flag.Args로 실행 인자로 db 파일 명을 설정할 수 있다.
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
