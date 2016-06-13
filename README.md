# gowkhtmltox
===

gowkhtmltox is a golang wrapper for [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)

### Installation

    go get -u github.com/brg-liuwei/gowkhtmltox/gowkhtmltoimage

### Demo

    go get -u github.com/brg-liuwei/gowkhtmltox/gowkhtmltox-demo

    # run service
    gowkhtmltox-demo -port 9999

    # test (http GET)
    curl "http://127.0.0.1:9999/render/url?url=https://www.baidu.com"

    # test (http POST)
    curl http://127.0.0.1:9999/render/html -d '<html><body><h1>Hello World</h1></body></html>'
