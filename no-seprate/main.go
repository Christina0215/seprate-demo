package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
	"strconv"
)

// 返回静态页面
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(writer, nil)
	query := request.URL.Query()


	a, _ := strconv.Atoi(query.Get("a"))
	b, _ := strconv.Atoi(query.Get("b"))
	ans := a+b
	fmt.Fprintf(writer, `ans:%d`,ans)

}

func main() {
	http.HandleFunc("/", handleIndex)

	fmt.Println("Running at port 3000 ...")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}