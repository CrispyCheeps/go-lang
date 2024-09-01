/*
	Go ini punya package sendiri namanya net/http isinya berbagai fitur
	utk keperluan pembuatan aplikasi berbasis web. (web server, routing, templating, dll)

	ga kyk php yg perlu apache, dan .NET yg perlu IIS
*/

// A51.1 Membuat Aplikasi Web Sederhana.
// contoh sederhana utk munculin text di web ketika url tertentu diakses
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa kabar!")
}

func main() {




	
	// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
	// 	fmt.Fprintln(w, "halo!")
	// })

	//Digunakan utk routing aplikasi web
	/*
		HandleFunc ini punya 2 parameter:
		1. rute yg diinginkan
		2. callback / aksi ketika rute tsb diakses.
	*/
	// http.HandleFunc("/index", index)

	// fmt.Println("Starting web server at http://localhost:8080/")

	/*
		digunakan utk menghidupkan server sekaligus menjalankan aplikasi menggunakan
		server tsb. di GO 1 web aplikasi adalah 1 buah server beda
	*/
	// http.ListenAndServe(":8080", nil)





	//--------------------------------------------------
	//A51.2 Penggunaan template web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name" : "john wick ",
			"Message" : "have a nice day",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

