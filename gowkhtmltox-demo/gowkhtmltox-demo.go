package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	img "github.com/brg-liuwei/gowkhtmltox/gowkhtmltoimage"
)

var port = flag.Int("port", 9999, "http listen port")
var renderService *RenderService

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
	conv.SetProperty("in", url)
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

	now := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body error: "+err.Error(), http.StatusOK)
		fmt.Println("html render read body error")
		return
	}

	fmt.Println("read body: ", time.Since(now))

	now = time.Now()

	ch := renderService.AddHtml(string(body))

	fmt.Println("ready: ", time.Since(now))

	now = time.Now()

	m := <-ch

	if m["errmsg"] != "ok" {
		http.Error(w, "render error: "+m["errmsg"], http.StatusOK)
	}

	fmt.Println("run: ", time.Since(now))

	now = time.Now()

	w.Header().Set("Content-Type", "image/jpg")
	w.Write([]byte(m["jpg"]))

	fmt.Println("output: ", time.Since(now))
}

func init() {
	flag.Parse()
	img.Init(false)
}

func main_http() {
	runtime.GOMAXPROCS(8)
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	renderService = NewRenderService()
	go renderService.Run()

	http.HandleFunc("/render/url", urlRender)
	http.HandleFunc("/render/html", htmlRender)

	panic(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func main() {
	renderService = NewRenderService()
	go renderService.Run()

	data := "<html><body><h1>Hello HtmlToImage</h1></body><html>"

	for i := 0; i != 10; i++ {
		now := time.Now()
		ch := renderService.AddHtml(data)
		for m := range ch {
			fmt.Println(time.Since(now), " >>> ", m["errmsg"], len(m["jpg"]))
		}
	}
}
