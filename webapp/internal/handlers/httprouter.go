package handlers

import (
	"fmt"
	"goclub/webapp/internal/files"
	"goclub/webapp/internal/logs"
	"net/http"
)

const pkgLogName = "handlers"

const (
	HandlePattern = "%s %s"
	PathToRoot    = "/"
	PathToPing    = "/ping"
	PathToUI      = "/ui/"
)

type HTTPRouter struct {
	http.ServeMux
}

/*
	mux.HandleFunc("/", defaultHandler)

	fileServer := http.FileServer(http.Dir("/temp/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

*/

func NewHTTPRouter() *HTTPRouter {
	router := &HTTPRouter{}
	router.HandleFunc(fmt.Sprintf(HandlePattern, http.MethodGet, PathToRoot), router.HandleRoot)
	router.HandleFunc(fmt.Sprintf(HandlePattern, http.MethodGet, PathToPing), router.HandlePing)
	router.HandleFunc(fmt.Sprintf(HandlePattern, http.MethodGet, PathToUI), router.HandleUI)

	return router
}

func (router *HTTPRouter) HandleRoot(_response http.ResponseWriter, _reqest *http.Request) {
	http.ServeFileFS(_response, _reqest, files.DataFS, "data/hello.html")
}

func (router *HTTPRouter) HandlePing(_response http.ResponseWriter, _reqest *http.Request) {
	_response.WriteHeader(http.StatusOK)
	fmt.Fprint(_response, "Pong!")
}

func (router *HTTPRouter) HandleUI(_response http.ResponseWriter, _reqest *http.Request) {
	const fnLogName = "HandleUI"
	logs.Debug(pkgLogName+"."+fnLogName, "path", _reqest.URL.Path)
	http.ServeFileFS(_response, _reqest, files.UIFS, _reqest.URL.Path)
}
