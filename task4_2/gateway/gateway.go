package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func createProxy(target string) *httputil.ReverseProxy {
	destination, err := url.Parse(target)
	if err != nil {
		log.Fatalf("invalid proxy target %s: %v", target, err)
	}
	proxy := httputil.NewSingleHostReverseProxy(destination)
	return proxy
}

func main() {
	pythonService := createProxy("http://localhost:8001")
	goService := createProxy("http://localhost:8002")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Printf("Маршрутизация запроса: %s %s", r.Method, path)

		if strings.HasPrefix(path, "/analytics") {
			r.URL.Path = strings.TrimPrefix(path, "/analytics")
			pythonService.ServeHTTP(w, r)
			return
		}

		if strings.HasPrefix(path, "/api/v1") {
			r.URL.Path = strings.TrimPrefix(path, "/api/v1")
			goService.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Маршрут не найден", http.StatusNotFound)
	})

	log.Println("Шлюз запущен на :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}