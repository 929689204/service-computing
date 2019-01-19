package service

import (
	"os"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
)


func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		var loginPage string = os.Getenv("GOPATH") + "/src/cloudgo/resource/login.gtpl"
		t, _ := template.ParseFiles(loginPage)
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Fprintf(w, "username: %s\n", r.Form["username"])
		fmt.Fprintf(w, "password: %s\n", r.Form["password"])
	}
}

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!\n")
	})
	mux.HandleFunc("/login", login);

	n := negroni.Classic() //创建一个negroni
	n.UseHandler(mux)
	return n;
}
