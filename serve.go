package main

import (
	"log"
	"net/http"
)

func main() {
	go func() {
		server := http.FileServer(http.Dir("/home/dev/go/src/github.com/TeamFairmont/hugo/bookshelf/public"))
		http.Handle("/", server)
		http.Handle("/test", http.FileServer(http.Dir("/home/dev/go/src/github.com/TeamFairmont/hugo/bookshelf/public/test/")))
		http.Handle("/dist", http.FileServer(http.Dir("/home/dev/go/src/github.com/TeamFairmont/hugo/bookshelf/public/test/dist/")))

		log.Println("Listening...")
		http.ListenAndServe(":8181", nil)
	}()
	go func() {
		//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./public/test/"))))
		//log.Fatal(http.ListenAndServe("/dist/", http.FileServer(http.Dir("./test/dist/"))))
	}()

	forever := make(chan bool)
	<-forever
}
