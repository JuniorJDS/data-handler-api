package route

import (
	"fmt"
	"net/http"
)

type HelloWorld struct {
}

func NewHelloWorld() *HelloWorld {
	return &HelloWorld{}
}

func (h *HelloWorld) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world!\n")
}
