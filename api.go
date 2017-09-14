package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"ghe.iparadigms.com/seu/arena/resources"

	"github.com/unrolled/render"
)

// APISomething ...
type APISomething struct {
	renderer *render.Render

	staticFileHander http.Handler
}

// NewAPISomething
func NewAPISomething() APISomething {
	// if c.DocumentRoot == "" {
	// c.DocumentRoot = "./resources"
	// }
	documentRoot := "./resources"
	isDev := true
	ro := render.Options{
		Asset:                     nil,
		AssetNames:                nil,
		Layout:                    "layout",
		Extensions:                []string{".tmpl"},
		Funcs:                     []template.FuncMap{},
		Delims:                    render.Delims{"{{", "}}"},
		Charset:                   "UTF-8",
		IndentJSON:                false,
		IndentXML:                 false,
		PrefixJSON:                []byte(""),
		PrefixXML:                 []byte(""),
		HTMLContentType:           "text/html",
		IsDevelopment:             isDev,
		UnEscapeHTML:              false,
		StreamingJSON:             false,
		RequirePartials:           false,
		DisableHTTPErrorRendering: false,
	}
	var fileServer http.Handler
	if isDev {
		ro.Directory = filepath.Join(documentRoot, "templates")
		fileServer = http.FileServer(http.Dir(documentRoot))
	} else {
		ro.Directory = "templates"
		ro.Asset = resources.Asset
		ro.AssetNames = resources.AssetNames
		// fileServer = http.FileServer(&assetfs.AssetFS{Asset: resources.Asset, AssetDir: resources.AssetDir, AssetInfo: resources.AssetInfo, Prefix: ""})
	}

	return APISomething{
		renderer:         render.New(ro),
		staticFileHander: fileServer,
	}
}

func (api APISomething) IndexHandlerGET(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templateParams := api.NewTemplateParams(w, r)
	api.renderer.HTML(w, http.StatusOK, "index", templateParams)
}

func (api APISomething) AboutHandlerGET(w http.ResponseWriter, r *http.Request) {
	templateParams := api.NewTemplateParams(w, r)
	api.renderer.HTML(w, http.StatusOK, "about", templateParams)
}
