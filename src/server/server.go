/*--------------------------------------------------------------*/
/*							Webserver Module					*/
/*							Jacob Mussler						*/
/*							Staff Engineer						*/
/*							26 AUG 2021							*/
/*							REV1_00								*/
/*--------------------------------------------------------------*/

package server

import (
	"fmt"
	"log"
	"net/http"
)

/* Server Structure */
type WebData struct {
	ServerName string
}

/* Initizalization Of The Webserver */
func Init(webStuct WebData) {
	fmt.Println("Webserver: ", webStuct.ServerName, " has started")
	http.HandleFunc("/", webStuct.handler)
	log.Fatal(http.ListenAndServeTLS(":8043", "server.crt", "server.key", nil))
}

func (web *WebData) handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		/* Call The Get Function */
		web.get(res, req)
	}
}
