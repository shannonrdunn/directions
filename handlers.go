package main


import (
  "time"
  "net/http"
  )

type Fastest struct {
    Path     string  `json:"path"`
    Eta      string  `json:"eta"`
    Duration string  `json:"duration"`
}

func Route(w http.ResponseWriter, r *http.Request) {
  route := Fastest{
    Path: "South"
    Eta: "5"
    Duration: "10"
  }
}
