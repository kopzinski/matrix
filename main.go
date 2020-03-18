package main

import (
	"echo"
	"flatten"
	"invert"
	"multiply"
	"net/http"
	"sum"
	"fmt"
)

func main() {
	http.HandleFunc("/echo", echo.EchoHandler)
	http.HandleFunc("/invert", invert.InvertHandler)
	http.HandleFunc("/flatten", flatten.FlattenHandler)
	http.HandleFunc("/sum", sum.SumHandler)
	http.HandleFunc("/multiply", multiply.MultiplyHandler)

	fmt.Println("Matrix listening at port 8080")
	http.ListenAndServe(":8080", nil)
}
