package main

import (
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc("/notfound", http.NotFound)
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		// 返回数组
		id := query["id"]
		log.Println(id)


		// 返回第一个参数
		name := query.Get("name")
		log.Println(name)
	})
	
	
	
	
	
	
	

}





