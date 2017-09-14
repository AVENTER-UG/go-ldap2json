package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	api "./api"
)

var API_SERVERPORT string
var LDAP_SERVER string
var LDAP_USER string
var LDAP_PASSWORD string
var LDAP_BASE string

func init() {
	flag.StringVar(&API_SERVERPORT, "server", "127.0.0.1:10777", "api server port")
	flag.StringVar(&LDAP_SERVER, "ldap_server", "127.0.0.1:1389", "ldap server")
	flag.StringVar(&LDAP_USER, "ldap_username", "", "ldap username")
	flag.StringVar(&LDAP_PASSWORD, "ldap_password", "", "ldap password")
	flag.StringVar(&LDAP_BASE, "ldap_base", "", "ldap base")
	flag.Parse()
}

func main() {
	fmt.Println("API_SERVERPORT=", API_SERVERPORT)
	fmt.Println("LDAP_SERVER=", LDAP_SERVER)
	fmt.Println("LDAP_USER=", LDAP_USER)
	fmt.Println("LDAP_PASSWORD=", LDAP_PASSWORD)
	fmt.Println("LDAP_BASE=", LDAP_BASE)

	log.Println("Listening..." + API_SERVERPORT)
	http.ListenAndServe(API_SERVERPORT, nil)

	api.LdapBase = LDAP_BASE
	api.LdapServer = LDAP_SERVER
	api.LdapUser = LDAP_USER
	api.LdapPassword = LDAP_PASSWORD
}
