package main

import (
	"fmt"
	"net/http"
	"log"
)

func httpSayHello(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)

	fmt.Fprintf(w,"hello tianhang !")
}

func main() {
	http.HandleFunc("/",httpSayHello)
	err := http.ListenAndServe(":9000",nil)

	if err != nil{
		log.Fatal("listen and server:",err)
	}else{
		log.Println("server listen on : 9000")
	}

}
