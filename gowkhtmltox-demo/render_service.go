package main

import (
	img "github.com/brg-liuwei/gowkhtmltox/gowkhtmltoimage"
)

type renderHolder struct {
	ch   chan map[string]string
	html string
	url  string
}

type RenderService struct {
	ch chan *renderHolder
}

func NewRenderService() *RenderService {
	return &RenderService{
		ch: make(chan *renderHolder, 1024),
	}
}

func (s *RenderService) add(html, url string) <-chan map[string]string {
	resultChan := make(chan map[string]string, 1)
	if len(html) != 0 {
		s.ch <- &renderHolder{
			ch:   resultChan,
			html: html,
		}
	} else {
		s.ch <- &renderHolder{
			ch:  resultChan,
			url: url,
		}
	}
	return resultChan
}

func (s *RenderService) AddHtml(html string) <-chan map[string]string {
	return s.add(html, "")
}

func (s *RenderService) AddUrl(url string) <-chan map[string]string {
	return s.add("", url)
}

func (s *RenderService) Run() {
	for data := range s.ch {

		outChan := data.ch
		conv := img.NewConvertor()
		conv.SetProperty("web.loadImages", "true")
		conv.SetProperty("fmt", "jpg")

		if len(data.html) != 0 {
			conv.SetHtml(data.html)
		} else if len(data.url) != 0 {
			conv.SetProperty("url", data.url)
		} else {
			outChan <- map[string]string{
				"errmsg": "no html or url",
			}
			close(outChan)
			continue
		}

		if err := conv.Ready(); err != nil {
			outChan <- map[string]string{
				"errmsg": err.Error(),
			}
			close(outChan)
			continue
		}

		if err := conv.Run(); err != nil {
			outChan <- map[string]string{
				"errmsg": err.Error(),
			}
			close(outChan)
			continue
		}

		outChan <- map[string]string{
			"errmsg": "ok",
			"jpg":    string(conv.GetImage()),
		}
		close(outChan)
	}
}
