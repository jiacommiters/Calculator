package main

import (

	"fmt"
	"html/template"
	"net/http"
	"strconv"

)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		
		data := map[string]string{
			"Result": "",
			"Error":  "",
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/kalkulator",  func(w http.ResponseWriter, r *http.Request) {




	


	})

}