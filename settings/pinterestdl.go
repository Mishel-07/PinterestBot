package settings 

import (
        "encoding/json"
        "io/ioutil"
        "net/http"
)

func PinterestDownload(link string) (string, error) {
        url := "https://horrid-api.vercel.app/download_pin?url=" + link
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
                Link string `json:"link"`
        }
        err = json.Unmarshal(body, &data)
        if err != nil {
                return "", err
        }
        return data.Link, nil
}
