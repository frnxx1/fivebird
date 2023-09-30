package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func buildHandler(fn func(c *gin.Context, key, key2, key3, key4, key5 string), key, key2, key3, key4, key5 string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c, key, key2, key3, key4, key5)
	}
}

func indexView(ctx *gin.Context, key, key2, key3, key4, key5 string) {
	ctx.HTML(
		http.StatusOK,
		"index.html", gin.H{
			"key":  fmt.Sprintf("http://%s/%s", ctx.Request.Host, key),
			"key2": fmt.Sprintf("http://%s/%s", ctx.Request.Host, key2),
			"key3": fmt.Sprintf("http://%s/%s", ctx.Request.Host, key3),
			"key4": fmt.Sprintf("http://%s/%s", ctx.Request.Host, key4),
			"key5": fmt.Sprintf("http://%s/%s", ctx.Request.Host, key5),
		})

}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles(
		"templates/index.html",
	)
	router.Static("/bird", "./")
	router.Static("/css", "../templates/buttom.css")

	/* router.GET("/",indexView) */
	router.GET("/", buildHandler(
		indexView,
		"bird/senichka.html",
		"bird/vorobey.html",
		"bird/voron.html",
		"bird/orel.html",
		"bird/golub.html",
	))

	return router
}

func main() {
	router := getRouter()
	router.Run("localhost:8080")

}
