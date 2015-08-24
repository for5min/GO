package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os/user"
	"log"
	"os"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	config := &ssh.ClientConfig{
		User: usr.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password("@Wsxcde3"),
		},
	}
	client, err := ssh.Dial("tcp", "fgprd-cs-ldap-va3-app001.mhint:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	name := os.Args[1]
	command := "ldapsearch -xLLLZ -h localhost cn='" + name + "' uid | grep ^uid | awk '{print $2}'"
	if err := session.Run(command ); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println("uid is:", b.String())
}
