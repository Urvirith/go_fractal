/*--------------------------------------------------------------*/
/*							Main Module							*/
/*							Jake Mussler						*/
/*							20 Nov 2021							*/
/*							REV1_00								*/
/*--------------------------------------------------------------*/

package main

import (
	"fmt"
	"os"

	browser "github.com/Urvirith/go_fractal/src/browser"
	config "github.com/Urvirith/go_fractal/src/config"
	server "github.com/Urvirith/go_fractal/src/server"
)

func main() {
	argswithoutprog := os.Args[1:]

	if len(argswithoutprog) == 0 {
		conf, err := config.Confget()

		if err != nil {
			fmt.Println(err)
			return
		}

		go browser.OpenBrowser("https://localhost:8043/")
		server.Init(conf.Webserv)
	} else {
		switch argswithoutprog[0] {
		case "-c":
			config.Setup()
		case "--config":
			config.Setup()
		default:
			fmt.Println("Unknown Pass")
		}
	}
}
