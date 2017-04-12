package notes

import (
  "net/http"
  "encoding/json"
)

func ServeNotes(w http.ResponseWriter, r *http.Request) {
  notes, err := GetNotes()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  response, err := json.Marshal(notes)
  if err != nil {
    panic(err)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}
