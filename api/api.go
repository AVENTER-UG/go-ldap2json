package api

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/mqu/openldap"
	//"io/ioutil"

	"net/http"
)

type userStruct struct {
	UUID     int
	USERNAME string
}

var LdapServer string   // LDAP SERVER
var LdapDN string       // LDAP DN
var LdapUser string     // LDAP Username
var LdapPassword string // LDAP Password
var LdapBase string     // LDAP Base
var ldap *openldap.Ldap

func init() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/versions", apiVersion).Methods("GET")
	rtr.HandleFunc("/api", apiVersion).Methods("GET")
	rtr.HandleFunc("/api/v0", apiV0Version).Methods("GET")
	rtr.HandleFunc("/api/v0/version", apiV0Version).Methods("GET")

	rtr.HandleFunc("/api/v0/getUser", apiV0GetUser).Methods("GET")

	//rtr.HandleFunc("/user/dieter/profile", api_user_profile).Methods("GET")
	http.Handle("/", rtr)
}

func apiVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")
	w.Write([]byte("/api/v0"))
}

func apiV0Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v0")
	w.Write([]byte("v0.1"))
}

func apiV0GetUser(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var username = params["username"]

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")

	d, _ := json.Marshal(searchLDAP(username))

	w.Write([]byte(string(d)))
}

func InitLDAP() bool {
	ldap, err := openldap.Initialize("ldap://" + LdapServer)

	if err != nil {
		fmt.Printf("LDAP::Initialize() : connection error\n")
		return false
	}

	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

	err = ldap.Bind(LdapUser, LdapPassword)
	if err != nil {
		fmt.Printf("LDAP::Bind() : bind error\n")
		fmt.Println(err)
		return false
	}
	defer ldap.Close()

	return true
}

func searchLDAP(username string) *openldap.LdapSearchResult {
	var scope = openldap.LDAP_SCOPE_SUBTREE // LDAP_SCOPE_BASE, LDAP_SCOPE_ONELEVEL, LDAP_SCOPE_SUBTREE
	var filter = "cn=" + username
	var attributes = []string{"cn", "uuid", "givenname", "mail"}
	var result *openldap.LdapSearchResult
	var err error

	result, err = ldap.SearchAll(LdapBase, scope, filter, attributes)

	if err != nil {
		fmt.Println("LDAP Search error: ", err)
	}

	return result
}
