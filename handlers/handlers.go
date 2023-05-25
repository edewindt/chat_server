package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode())

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "index.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, tpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tpl)
	if err != nil {
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		return err
	}
	return nil
}
