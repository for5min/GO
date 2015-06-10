package main
import (
	"log"
	"encoding/json"
	"strings"
	"io"
	"fmt"
)



func main() {

	//var (
	//	config map[string]string
	//	t      *html.Template
	//)

	//host := "a"
	//ip := "192.168.1.3"
	//cpu := "10"
	//mem := "10"
	//weight := "10"


	//config = map[string]string{
	//	"HOST": host,
	//	"IP": ip,
	//	"CPU": cpu,
	//	"MEM": mem,
	//	"WEIGHT": weight,
	//}

	const jsondata = `
	{
  "host": "a",
    "ip": "192.168.1.3",
    "cpu": "10",
    "mem": "10",
    "weight": "10"
}
{
  "host": "b",
    "ip": "192.168.1.4",
    "cpu": "20",
    "mem": "20",
    "weight": "20"
}
	`
	type jsdata struct {
		HOST, IP, CPU, MEM, WEIGHT string
	}

	dec := json.NewDecoder(strings.NewReader(jsondata))

	for {
		var m jsdata
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s, %s, %s, %s, %s", m.HOST, m.IP, m.CPU, m.MEM, m.WEIGHT)

	}
}
