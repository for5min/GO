package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	//	"os"
	"text/template"
)

var (
	config map[string]string
	t      *template.Template
)

const ipxeBootScript = `#!ipxe
set base-url http://{{.BaseUrl}}
kernel ${base-url}/coreos_production_pxe.vmlinuz sshkey="{{.SSHKey}}"
initrd ${base-url}/coreos_production_pxe_image.cpio.gz
boot
`

// END 0 OMIT

// START 1 OMIT
func init() {
	//baseUrl := os.Getenv("IPXE_BASEURL")
	baseUrl := "0.0.0.0:9000"
	sshKey, err := ioutil.ReadFile("/Users/atu/.ssh/id_rsa.pub")
	if err != nil {
		log.Fatal(err)
	}
	config = map[string]string{
		"BaseUrl": baseUrl,
		"SSHKey":  string(bytes.TrimSpace(sshKey)),
	}
	t = template.Must(template.New("ipxebootscript").Parse(ipxeBootScript))
}

// END 1 OMIT

// START 2 OMIT
func ipxeBootScriptServer(w http.ResponseWriter, r *http.Request) {
	log.Printf("creating boot script for %s", r.RemoteAddr)
	if err := t.Execute(w, config); err != nil {
		http.Error(w, "Error generating the iPXE boot script", 500)
		return
	}
}

func main() {
	http.HandleFunc("/boot", ipxeBootScriptServer)
	http.Handle("/", http.FileServer(http.Dir("/Users/atu/Downloads")))
	log.Println("Starting CoreOS iPXE Server...")
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}
