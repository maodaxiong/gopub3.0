package main

import (
	"net/http"

	"github.com/kataras/iris"
	_ "gopub3.0/model"
	"gopub3.0/mssh"
	"gopub3.0/route"
)

func main() {
	go mssh.Init()
	// mssh.Socks5ProxyStart()
	// mssh.Dial2local(client)

	app := iris.New()

	route.Init(app)

	app.Build()
	srv := &http.Server{Handler: app, Addr: ":8088"}
	println("Start a server listening on http://localhost:8088")

	srv.ListenAndServe()

}