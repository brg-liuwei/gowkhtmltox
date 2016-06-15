package gowkhtmltoimage_test

import (
	"testing"

	img "github.com/brg-liuwei/gowkhtmltox/gowkhtmltoimage"
)

func init() {
	img.Init(false)
}

func BenchmarkHtmlJpg(b *testing.B) {
	htmlSnippet := `
        <html>
            <body>
                <h1>Hello World</h1>
            </body>
        </html>
    `

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conv := img.NewConvertor()
		conv.SetProperty("fmt", "jpg")
		conv.SetHtml(htmlSnippet)
		if err := conv.Ready(); err != nil {
			b.Error(err)
		}
		if err := conv.Run(); err != nil {
			b.Error("i = ", i, "; err: ", err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkHtmlJpeg(b *testing.B) {
	htmlSnippet := `
        <html>
            <body>
                <h1>Hello World</h1>
            </body>
        </html>
    `

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conv := img.NewConvertor()
		conv.SetProperty("fmt", "jpeg")
		conv.SetHtml(htmlSnippet)
		if err := conv.Ready(); err != nil {
			b.Error(err)
		}
		if err := conv.Run(); err != nil {
			b.Error("i = ", i, "; err: ", err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkHtmlPng(b *testing.B) {
	htmlSnippet := `
        <html>
            <body>
                <h1>Hello World</h1>
            </body>
        </html>
    `

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conv := img.NewConvertor()
		conv.SetProperty("fmt", "png")
		conv.SetHtml(htmlSnippet)
		if err := conv.Ready(); err != nil {
			b.Error(err)
		}
		if err := conv.Run(); err != nil {
			b.Error("i = ", i, "; err: ", err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkHtmlSvg(b *testing.B) {
	htmlSnippet := `
        <html>
            <body>
                <h1>Hello World</h1>
            </body>
        </html>
    `

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conv := img.NewConvertor()
		conv.SetProperty("fmt", "svg")
		conv.SetHtml(htmlSnippet)
		if err := conv.Ready(); err != nil {
			b.Error(err)
		}
		if err := conv.Run(); err != nil {
			b.Error("i = ", i, "; err: ", err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkHtmlBmp(b *testing.B) {
	htmlSnippet := `
        <html>
            <body>
                <h1>Hello World</h1>
            </body>
        </html>
    `

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conv := img.NewConvertor()
		conv.SetProperty("fmt", "bmp")
		conv.SetHtml(htmlSnippet)
		if err := conv.Ready(); err != nil {
			b.Error(err)
		}
		if err := conv.Run(); err != nil {
			b.Error("i = ", i, "; err: ", err)
		}
	}

	b.ReportAllocs()
}
