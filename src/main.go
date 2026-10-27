package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Папка зі сторінкою: за замовчуванням ./html, або перший аргумент запуску
	dir := "./html"
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))

	log.Printf("Сервер працює на порту 8080, віддає папку %s", dir)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
