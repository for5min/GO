package main

import (
	"html/template"
	//"os"
	"reflect"
	"log"
	"net/http"
)

type Node struct {
	Host     string
	IP       string
	CPU      string
	MEM      string
	Weight   string
}

var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

// In the template, we use rangeStruct to turn our struct values
// into a slice we can iterate over
//var htmlTemplate = `
//{{range .}}<tr>
//{{range rangeStruct .}} <td>{{.}}</td>
//{{end}}
//<td><button type="button" class="btn btn-success">Move</button><button type="button" class="btn btn-danger">Halt</button></td>
//</tr>
//{{end}}
//`


func writedata(w http.ResponseWriter, r *http.Request) {
	container := []Node{
		{ "a", "192.168.1.3", "10", "10", "10"},
        { "b", "192.168.1.4", "20", "20", "20"},
	}

	t := template.New("t").Funcs(templateFuncs)
	t, err := t.ParseFiles("table.html")
	//t.Execute(w,htmlTemplate)
	//if err != nil {
	//	panic(err)
	//}

	//err = t.Execute(w, container)
	err = t.ExecuteTemplate(w, "table.html", container)
	if err != nil {
		panic(err)
	}


}

// RangeStructer takes the first argument, which must be a struct, and
// returns the value of each field in a slice. It will return nil
// if there are no arguments or first argument is not a struct
func RangeStructer(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}

	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return nil
	}

	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}

	return out
}

func main() {
	http.HandleFunc("/htmlfile", writedata)
	log.Println("Starting Fileserver")
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}