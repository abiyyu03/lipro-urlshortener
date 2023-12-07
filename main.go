package main

import (
	"url-shortener/config"
	"url-shortener/handler"

	"github.com/gofiber/fiber/v2"
)

// mekanisme penyimpanan url
// - ambil data dari (url, custom-url) dari request (post)
// - store data request ke db

// verifikasi link
// - get custom-url segment 2
// - cocokan dengan data dari database
// - ambil url aslinya
// - redirect kesitu
func main() {
	app := fiber.New()
	config.Init()

	storeUrlHandler := new(handler.StoreUrl)
	resolveUrlHandler := new(handler.ResolveUrl)
	app.Post("/", storeUrlHandler.CreateCustomUrl)
	app.Get("/:short_url", resolveUrlHandler.Resolve)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Write([]byte("hello again"))
	// })
	// http.HandleFunc("/:shortUrl", func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Write([]byte("hello again"))
	// })
	app.Listen(":8080")
	// fmt.Println(http.ListenAndServe(":8080", nil))
}
