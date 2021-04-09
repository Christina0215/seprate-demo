package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	query := request.URL.Query()
	a, _ := strconv.Atoi(query.Get("a"))
	b, _ := strconv.Atoi(query.Get("b"))

	ans := a+b
	fmt.Print(ans)
	fmt.Fprintln(writer,ans)

}

func main() {
	http.HandleFunc("/", handleIndex)

	fmt.Println("Running at port 8000 ...")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}