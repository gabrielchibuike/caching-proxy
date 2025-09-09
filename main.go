package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func main() {

	port := flag.Int("port", 8080, "Port to run the proxy server on")
	origin := flag.String("origin", "", "Origin server to forward requests to (required)")
	flag.Parse()

	if *origin == "" {
		log.Fatal("Error: --origin is required")
	}

	var respCache = cache.New(60*time.Second, 120*time.Second)

	app := fiber.New()

	app.Get("/*", func(c *fiber.Ctx) error {

		url := fmt.Sprintf("%s%s", *origin, c.OriginalURL())

		// url := c.Query("url")

		// if url == "" {
		// 	return c.Status(400).SendString("Missing URL")
		// }

		if cachedResp, found := respCache.Get(url); found {
			c.Set("X-Cache", "HIT")
			return c.Send(cachedResp.([]byte))
		}

		resp, err := http.Get(url)

		if err != nil || resp.StatusCode != 200 {
			return c.Status(502).SendString("Bad Gateway")
		}

		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		// Save to cache
		respCache.Set(url, body, cache.DefaultExpiration)

		c.Set("X-Cache", "MISS")

		return c.Send(body)

	})

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).SendString("Hello, World!")
	// })

	// app.Listen(":5000")
	// fmt.Println("Server running on port 5000")

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting caching proxy on %s, forwarding to %s\n", addr, *origin)
	log.Fatal(app.Listen(addr))
}
