package main

import (
  "net/http"
  "notes"
)

func main() {
  http.HandleFunc("/", notes.ServeNotes)

  http.ListenAndServe(":3000", nil)
}
