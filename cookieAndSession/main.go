package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("content-type", "text/html")
		fmt.Print(t)
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
func indexPage(w http.ResponseWriter, r *http.Request) {
	t,_:=template.ParseFiles("index.gtpl")
	t.Execute(w,nil)
}
func init() {
	globalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		Maxlifetime: 3600,
		CookieName:  "gosessionid",
	})
}
func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
