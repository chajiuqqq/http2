package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Img:{{ .filename }}</h1>
  <img src="www/{{ .filename }}"></img>
</body>
</html>
`))

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	crt:=flag.String("cert","./testdata/server.pem","cert")
	key:=flag.String("key","./testdata/server.key","key")
	flag.Parse()

	r := gin.Default()
	r.SetHTMLTemplate(html)
	r.Static("/www", "./www")

	r.GET("/html/:filename", func(c *gin.Context) {
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
			"filename": c.Param("filename"),
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", *crt,*key)
}
