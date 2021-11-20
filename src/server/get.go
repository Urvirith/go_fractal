/*--------------------------------------------------------------*/
/*							Get Module							*/
/*							Jake Mussler						*/
/*							20 Nov 2021							*/
/*							REV1_00								*/
/*--------------------------------------------------------------*/

package server

import (
	"net/http"
	"net/url"
)

func (web *WebData) get(res http.ResponseWriter, req *http.Request) {
	path, err := url.Parse(req.URL.Path)

	if err != nil {
		println(err)
	} else {
		qry := req.URL.Query()
		if len(qry) == 0 {
			web.getPage(res, req, path)
		} else {
			/* Unknown Query */
			res.WriteHeader(http.StatusNotImplemented)
			res.Write([]byte("Unknown Query String"))
			return
		}
	}
}

func (web *WebData) getPage(res http.ResponseWriter, req *http.Request, path *url.URL) {
	var prefix = "./src/website"
	var fileloc = prefix + path.String()

	http.ServeFile(res, req, fileloc)
}
