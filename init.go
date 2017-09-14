package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	api "./api"
)

// APIServerPort is the port of the server
var APIServerPort string

func init() {
	flag.StringVar(&APIServerPort, "server", "127.0.0.1:10777", "api server port")
	flag.StringVar(&api.LDAPServer, "ldap_server", "127.0.0.1:1389", "ldap server")
	flag.StringVar(&api.LDAPUser, "ldap_username", "", "ldap username")
	flag.StringVar(&api.LDAPPassword, "ldap_password", "", "ldap password")
	flag.StringVar(&api.LDAPBase, "ldap_base", "", "ldap base")
	flag.Parse()
}

func main() {
	fmt.Println("API_SERVERPORT=", APIServerPort)
	fmt.Println("LDAP_SERVER=", api.LDAPServer)
	fmt.Println("LDAP_USER=", api.LDAPUser)
	fmt.Println("LDAP_PASSWORD=", api.LDAPPassword)
	fmt.Println("LDAP_BASE=", api.LDAPBase)

	log.Println("Listening..." + APIServerPort)
	http.ListenAndServe(APIServerPort, nil)
}
