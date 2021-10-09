package main

import (
	"bytes"
	"compress/gzip"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
)

type submition struct {
	Id string
	Gif bool
	Content string
}

func main() {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {log.Fatal(err)}
	data, err := os.ReadFile("testfile.webp")
	if err != nil {log.Fatal(err)}
	if _, err := gz.Write(data); err != nil {
		log.Fatal(err)
	}	
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	content := b.String()

	app := fiber.New()

	app.Get("/raw.webp", func(c *fiber.Ctx) error {
		c.Type("webp")
		// c.Set("Content-Encoding", "gzip") // Comment & uncomment this line!
		return c.SendString(content)
	})

	app.Get("/i", func(c *fiber.Ctx) error {
		c.Type("html")
		return c.SendString(`<!DOCTYPE html>
<html lang="en">
<head>
<title> Sick ass epic image server </title>
<meta name="viewport" content="width=device-width,initial-scale=1">
<meta property="og:title" content="Shady's image server" />
<meta property="og:image" content="/raw.webp" />
<meta property="og:url" content="/i" />
<meta property="twitter:title" content="Shady's image server" />
<meta property="twitter:image" content="/raw.webp" />
<meta name="theme-color" content="#5655b0">
<meta name="twitter:card" content="summary_large_image">
<style>
:root {
	background-color: #202124 !important;
}
*, :after, :before {
	box-sizing: border-box;
	margin: 0 !important;
}
</style></head>
<body><img style="height: 100vh; margin: 0 auto !important; display: block;" src="/raw.webp" /></body>`)
	})

	app.Listen(":5999")
}