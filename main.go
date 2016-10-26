package main

import (
  "net/http"
  "log"
  "github.com/julienschmidt/httprouter"
  c "./controllers"
)

func main()  {
  r := httprouter.New()
  r.GET("/", c.Index)
  r.NotFound = http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/")))
  log.Fatal(http.ListenAndServe(":8080", r))
}
