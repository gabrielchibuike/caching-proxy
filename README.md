# 🚀 Caching Proxy Server (Go + Fiber)

A simple **CLI-based caching proxy server** built with [Fiber](https://gofiber.io/) in Go.  
It forwards requests to an origin server and caches the responses for a configurable time (default: 60s).  
If the same request is made again within the cache TTL, the proxy serves the cached response instead of hitting the origin server again.

---

## ✨ Features

- ✅ CLI tool with `--port` and `--origin` flags
- ✅ Caches responses in memory with a TTL
- ✅ Adds custom header `X-Cache: HIT/MISS` for debugging
- ✅ Works with any endpoint (catch-all routing)
- ✅ Reduces load on origin servers and speeds up repeated requests

---

## 🛠 Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/caching-proxy.git
   cd caching-proxy
   ```
2. go mod tidy
3. go build caching-proxy

##🚀 Usage

Run the proxy with a port and origin server:

./caching-proxy --port 3000 --origin http://dummyjson.com

Now, you can make requests through the proxy:

curl -i http://localhost:3000/products/1

##Example response headers:

First request (cache MISS):

HTTP/1.1 200 OK
Content-Type: application/json
X-Cache: MISS

Second request (cache HIT):

HTTP/1.1 200 OK
Content-Type: application/json
X-Cache: HIT

##📂 Project Structure

.
├── main.go # Main entry point (CLI + Fiber app)
├── go.mod # Go module file
├── go.sum
└── README.md # Project documentation

##📌 Why a Caching Proxy?

Improves performance by serving repeated requests faster

Reduces server load on the origin server

Increases reliability if the origin server is slow or temporarily unavailable

Mimics how real-world CDNs (e.g., Cloudflare, Akamai) work
