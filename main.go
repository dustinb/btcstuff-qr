package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	var requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "btcstuff",
			Name:      "qr_requests",
			Help:      "Request counter for /qr endpoint",
		}, []string{})
	prometheus.MustRegister(requestCounter)
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/qr", func(w http.ResponseWriter, r *http.Request) {
		requestCounter.WithLabelValues().Inc()

		values := url.Values{}
		amount := r.URL.Query().Get("amount")
		if amount != "" {
			_, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 32)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			values.Add("amount", r.URL.Query().Get("amount"))
		}
		values.Add("label", r.URL.Query().Get("label"))
		values.Add("message", r.URL.Query().Get("message"))

		// TODO: Verify bitcoin address format
		scheme := "bitcoin:%s?%s"
		uri := fmt.Sprintf(scheme, r.URL.Query().Get("address"),values.Encode())
		log.Print(values.Encode())
		log.Printf(uri)

		var png []byte
		png, _ = qrcode.Encode(uri, qrcode.Medium, 256)
		w.Header().Set("Content-Type", "image/png")
		w.Write(png)
	})

	err = http.ListenAndServe(":8101", nil)
	if err != nil {
		log.Fatal(err)
	}

}