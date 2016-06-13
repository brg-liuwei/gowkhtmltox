package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	img "github.com/brg-liuwei/gowkhtmltox/gowkhtmltoimage"
)

var port = flag.Int("port", 9999, "http listen port")

func urlRender(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusForbidden)
		return
	}
	url := r.FormValue("url")
	if len(url) == 0 {
		http.Error(w, "need get param: url=urlToRender", http.StatusOK)
		fmt.Println("no param")
		return
	}
	conv := img.NewConvertor()
	conv.SetProperty("web.loadImages", "true")
	conv.SetProperty("fmt", "jpg")
	conv.SetProperty("in", "url")
	if err := conv.Ready(); err != nil {
		http.Error(w, "render ready error", http.StatusOK)
		fmt.Println("url render ready error")
		return
	}
	if err := conv.Run(); err != nil {
		http.Error(w, "render error", http.StatusOK)
		fmt.Println("url render run error")
		return
	}
	w.Header().Set("Content-Type", "image/jpg")
	w.Write(conv.GetImage())
}

func htmlRender(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusForbidden)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body error: "+err.Error(), http.StatusOK)
		fmt.Println("html render read body error")
		return
	}
	conv := img.NewConvertor()
	conv.SetProperty("web.loadImages", "true")
	conv.SetProperty("fmt", "jpg")
	conv.SetHtml(string(body))
	if err := conv.Ready(); err != nil {
		http.Error(w, "render ready error", http.StatusOK)
		fmt.Println("html render ready error")
		return
	}
	if err := conv.Run(); err != nil {
		http.Error(w, "render error", http.StatusOK)
		fmt.Println("html render run error")
		return
	}
	w.Header().Set("Content-Type", "image/jpg")
	w.Write(conv.GetImage())
}

func init() {
	flag.Parse()
}

func main() {
	img.Init(false)

	http.HandleFunc("/render/url", urlRender)
	http.HandleFunc("/render/html", htmlRender)

	fmt.Println("port = ", *port)

	panic(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

	img.DeInit()
}
