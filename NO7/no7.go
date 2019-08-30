package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config/env.json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error config file!, &s", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, you've requested: %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":"+viper.GetString("server.port"), nil)
}
