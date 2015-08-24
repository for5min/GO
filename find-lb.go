package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "net/http"
    simplejson "github.com/bitly/go-simplejson"
)


var Location = []byte(`
{
  "AMS1": "10.3.253.14",
  "CAL1": "10.4.224.245",
  "FRA4": "10.17.127.14",
  "RUH1": "10.113.254.10",
  "SIN1": "10.112.254.10",
  "SYD1": "10.7.63.16",
  "VA2": "172.18.120.20",
  "VA3": "172.27.2.20"
}
`)

type mytype map[string]string



func main() {


    ip := os.Args[1]


    var data mytype
    err := json.Unmarshal(Location, &data)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(data)
    for k, v := range data {
        //filter http://172.27.2.20/nitro/v1/config/lbvserver?filter=ipv46:69.196.224.253
        url := "http://" + v + "/nitro/v1/config/lbvserver?filter=ipv46:" + ip
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            log.Fatal(err)
        }
        req.SetBasicAuth("atu", "@Wsxcde3")
        client := http.Client{}
        res, err := client.Do(req)
        if err != nil {
            log.Fatal(err)
        }

        body, err := ioutil.ReadAll(res.Body)
        res.Body.Close()
        if err != nil {
            log.Fatal(err)
        }

        js, err := simplejson.NewJson(body)
        _, ok := js.CheckGet("lbvserver")
        if ok {
            fmt.Println(k,"(",v,")","has",ip)
            break
        }
    }
}