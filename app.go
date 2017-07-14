package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	api "./api"
)

var API_SERVERPORT string
var LDAP_SERVER string
var LDAP_DN string
var LDAP_USER string
var LDAP_PASSWORD string

func main() {
	fmt.Println("API_SERVERPORT=", API_SERVERPORT)
	fmt.Println("LDAP_SERVER=", LDAP_SERVER)
	fmt.Println("LDAP_DN=", LDAP_DN)
	fmt.Println("LDAP_USER=", LDAP_USER)
	fmt.Println("LDAP_PASSWORD=", LDAP_PASSWORD)

	api.ldapServer = LDAP_SERVER
	api.ldapDN = LDAP_DN
	api.ldapUser = LDAP_USERNAME
	api.ldapPassword = LDAP_PASSWORD

	log.Println("Listening..." + API_SERVERPORT)
	http.ListenAndServe(API_SERVERPORT, nil)
}
