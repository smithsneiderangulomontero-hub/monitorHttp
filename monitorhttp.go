package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := []string{
		"https://www.google.com",
		"https://github.com",
		"https://www.twitter.com",
	}
	for _, url := range url {
		revisarEstado(url)
	}
}

func revisarEstado(url string) {
	// Creamos un cliente con un tiempo límite de 5 segundos
	cliente := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := cliente.Get(url)
	if err != nil {
		fmt.Printf("❌ %s está caído o la URL es inválida.\n", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("✅ %s está funcionando correctamente.\n", url)
	} else {
		fmt.Printf("⚠️ %s devolvió el código: %d\n", url, resp.StatusCode)
	}
}