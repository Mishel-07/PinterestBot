package api

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
)

func PinterestDownload(link string) (string, error) {
  url := "https://horridapi.onrender.com/download_pin?link=" + link
  resp, err := http.Get(url)
  if err != nil {
   return "", err
  }
  defer resp.Body.Close()
 
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
   return "", err
  }

  var data struct {
   URL string `json:"url"`
  }
  err = json.Unmarshal(body, &data)
  if err != nil {
   return "", err
  }

  return data.URL, nil
 }