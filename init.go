package main

import (
	"flag"
)

func init() {
	flag.StringVar(&API_SERVERPORT, "server", "127.0.0.1:10777", "api server port")
	flag.StringVar(&LDAP_SERVER, "ldap_server", "127.0.0.1:1389", "ldap server")
	flag.StringVar(&LDAP_DN, "ldap_dn", "user", "ldap dn")
	flag.Parse()
}
