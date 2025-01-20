package settings 

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type BingResponse struct {
    Result []struct {
        URL string `json:"images"`
    } `json:"result"` 
}

func SearchBing(query string) (BingResponse, error) {
    url := fmt.Sprintf("https://horrid-api.vercel.app/images?page=10&query=%s", query)
    resp, err := http.Get(url)
    if err != nil {
        return BingResponse{}, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return BingResponse{}, err
    }

    var result BingResponse
    err = json.Unmarshal(body, &result)
    if err != nil {
        return BingResponse{}, err
    }

    return result, nil
}

func SearchBingInline(query string) (BingResponse, error) {
    url := fmt.Sprintf("https://horrid-api.vercel.app/images?page=40&query=%s", query)
    resp, err := http.Get(url)
    if err != nil {
        return BingResponse{}, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return BingResponse{}, err
    }

    var result BingResponse
    err = json.Unmarshal(body, &result)
    if err != nil {
        return BingResponse{}, err
    }

    return result, nil
}
