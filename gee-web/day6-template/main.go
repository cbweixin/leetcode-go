package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// ❯ curl -i http://localhost:9999/students
//HTTP/1.1 200 OK
//Content-Type: text/html
//Date: Tue, 20 Apr 2021 14:29:03 GMT
//Content-Length: 141
//
//
//<html>
//<body>
//    <p>hello, gee</p>
//
//    <p>0: Geektutu is 20 years old</p>
//
//    <p>1: Jack is 22 years old</p>
//
//</body>
//</html>%

// ❯ curl -i http://localhost:9999/date
//HTTP/1.1 200 OK
//Content-Type: text/html
//Date: Tue, 20 Apr 2021 14:27:24 GMT
//Content-Length: 80
//
//
//<html>
//<body>
//    <p>hello, gee</p>
//    <p>Date: 2019-08-17</p>
//</body>

// ❯ curl -i http://localhost:9999/
//HTTP/1.1 200 OK
//Content-Type: text/html
//Date: Tue, 20 Apr 2021 14:26:07 GMT
//Content-Length: 108
//
//<html>
//    <link rel="stylesheet" href="/assets/css/geektutu.css">
//    <p>geektutu.css is loaded</p>
//</html>%

// ❯ curl -i http://localhost:9999/assets/css/geektutu.css
//HTTP/1.1 200 OK
//Accept-Ranges: bytes
//Content-Length: 67
//Content-Type: text/css; charset=utf-8
//Last-Modified: Mon, 19 Apr 2021 14:56:29 GMT
//Date: Tue, 20 Apr 2021 14:50:35 GMT
//
//p {
//    color: orange;
//    font-weight: 700;
//    font-size: 20px;
//}%

// ❯ curl -i http://localhost:9999/assets/file1.txt
//HTTP/1.1 200 OK
//Accept-Ranges: bytes
//Content-Length: 29
//Content-Type: text/plain; charset=utf-8
//Last-Modified: Mon, 19 Apr 2021 14:57:13 GMT
//Date: Tue, 20 Apr 2021 14:52:19 GMT
//
//I'm file1
//I'm file1
//I'm file1%
func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")

}
