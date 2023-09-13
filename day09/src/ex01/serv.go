package main

import (
  "fmt"
  "net/http"
//   "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
//   time.Sleep(10000 * time.Millisecond)
  fmt.Fprint(w, "hello")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}