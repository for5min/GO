package main

import (
	"net/http"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"fmt"
	"os"
	"log"
	"bytes"
)

func main() {

	longurl := os.Args[1]

	con, err := simplejson.NewJson([]byte(`{}`))
	if err != nil {
		fmt.Println(err)
	}

	con.Set("longUrl", longurl)
	//fmt.Println(con)

	jsonStr, err := con.MarshalJSON()
	//jsonStr, _ := json.Marshal(con)
	//fmt.Println(string(jsonStr))


	//con := map[string]string{"longUrl": longurl}


	os.Setenv("HTTP_PROXY", "http://10.75.106.10:3128")

	API := "AIzaSyBxfayswzAvRVCKc3C6HybF60GDk-JdY7k"
	url := "https://www.googleapis.com/urlshortener/v1/url?key=" + API



	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}
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
	shorten := js.Get("id").MustString()


	fmt.Println(shorten)
}