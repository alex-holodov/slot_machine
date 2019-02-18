package main

import (
	"net/http"
)

const (
	secretKey = "secret_key"
)

func main() {
	h := handler{
		machine: newAtkinsGame(),
		key:     []byte(secretKey),
	}

	//t, _ := h.newToken("test", 10000, 10)
	//fmt.Printf("%s\n", t)

	http.Handle("/api/machines/atkins-diet/spins", h)
	http.ListenAndServe(":8081", nil)
}
