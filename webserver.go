package main

import (
	//"bytes"
	//"io/ioutil"
	html "html/template"
	"log"
	"net/http"
	"text/template"
)

func textwrite(w http.ResponseWriter, r *http.Request) {

	const letter = `{{.A}}
{{.B}}
`

	var (
		config map[string]string
		t      *template.Template
	)

	// END 0 OMIT

	// START 1 OMIT
	//baseUrl := os.Getenv("IPXE_BASEURL")
	A := "test1"
	B := "test2"

	config = map[string]string{
		"A": A,
		"B": B,
	}
	t = template.Must(template.New("ipxebootscript").Parse(letter))
	log.Printf("creating boot script for %s", r.RemoteAddr)
	if err := t.Execute(w, config); err != nil {
		http.Error(w, "Error generating text file", 500)
		return
	}
}

func textwritefile(w http.ResponseWriter, r *http.Request) {

	var (
		config map[string]string
		t      *template.Template
	)

	// END 0 OMIT

	// START 1 OMIT
	//baseUrl := os.Getenv("IPXE_BASEURL")
	A := "This is content from template"
	B := "This one is the sam"
	test := "I'm IP(10.75.145.37)"

	config = map[string]string{
		"A": A,
		"B": B,
		"C": test,
	}
	t = template.Must(template.New("ipxebootscript").ParseFiles("test.gtpl"))
	log.Printf("creating boot script for %s", r.RemoteAddr)
	if err := t.ExecuteTemplate(w, "test.gtpl", config); err != nil {
		http.Error(w, "Error generating textfile", 500)
		return
	}
}

func htmlwrite(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	t, err := html.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>"); err != nil {
		http.Error(w, "Error generating textfile", 500)
		return
	}
}

func htmlwritefile(w http.ResponseWriter, r *http.Request) {

	var (
		config map[string]string
		t      *html.Template
	)

	name := "Alex"
	message := "Hello world"

	config = map[string]string{
		"FirstName": name,
		"Message": message,
	}

	t, err := html.New("test").ParseFiles("tmpl.html")
	if err = t.ExecuteTemplate(w, "tmpl.html", config); err != nil {
		http.Error(w, "Error", 500)
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It Works!"))
}

func main() {
	http.Handle("/file", http.StripPrefix("/file", http.FileServer(http.Dir("/Users/atu/Downloads"))))
	http.HandleFunc("/", hello)
	http.HandleFunc("/text", textwrite)
	http.HandleFunc("/textfile", textwritefile)
	http.HandleFunc("/html", htmlwrite)
	http.HandleFunc("/htmlfile", htmlwritefile)
	log.Println("Starting Fileserver")
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}
