package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

//go:embed ui/build/*
var ui embed.FS

func main() {
	// Prefix removal from embedded file system
	// https://stackoverflow.com/questions/66274816/go-1-16-how-to-use-strip-prefix-in-goembed
	fsys, err := fs.Sub(ui, "ui/build")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))
	go func() {
		err = http.ListenAndServe(":8101", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// https://github.com/bitcoin/bips/blob/master/bip-0021.mediawiki
	router := gin.Default()
	router.GET("/qr", func(c *gin.Context) {
		var png []byte
		scheme := "bitcoin:%s?amount=%f&label=%s&message=%s"
		amount, err := strconv.ParseFloat(c.Query("amount"), 32)
		if err != nil {
			c.AbortWithError(500, err)
		}
		uri := fmt.Sprintf(
			scheme,
			c.Query("address"),
			amount,
			url.QueryEscape(c.Query("label")),
			url.QueryEscape(c.Query("message")),
		)
		log.Printf(uri)
		png, _ = qrcode.Encode(uri, qrcode.Medium, 256)
		c.Data(200, "image/png", png)
	})

	err = router.Run(":8100")
	if err != nil {
		log.Print(err)
	}
}