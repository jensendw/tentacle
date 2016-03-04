package main


import (
  "gopkg.in/resty.v0"
  "fmt"
  "github.com/antonholmquist/jason"
  "log"
  "github.com/dchest/uniuri"
)

func main() {
  // GET request
  resp, err := resty.R().Get("http://httpbin.org/get")
  v, err := jason.NewObjectFromBytes([]byte(resp.Body()))
  url, err := v.GetString("url")
  fmt.Printf("The URL: %v\n", url)
  s := uniuri.New()

  fmt.Printf(s)
  if err != nil {
    log.Fatal(err)
  }
}
