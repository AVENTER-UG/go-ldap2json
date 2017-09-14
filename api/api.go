package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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
	var user userStruct
	getHTTPRequest(&user, r)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")

	d, _ := json.Marshal(searchLDAP(user.USERNAME))

	fmt.Println(d)
}

func InitLDAP() {
	var err error
	ldap, err = openldap.Initialize("ldap://" + LdapServer)

	fmt.Println("Init LDAP: " + LdapServer)

	logError(err)

	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

	err = ldap.Bind(LdapUser, LdapPassword)
	logError(err)

	defer ldap.Close()
}

func searchLDAP(username string) *openldap.LdapSearchResult {

	var scope = openldap.LDAP_SCOPE_SUBTREE // LDAP_SCOPE_BASE, LDAP_SCOPE_ONELEVEL, LDAP_SCOPE_SUBTREE
	var filter = "cn=" + username
	var attributes = []string{}
	var result *openldap.LdapSearchResult
	var err error

	fmt.Println("Search LDAP: " + username)
	fmt.Println("Scope: ", scope)
	fmt.Println("Filter: " + filter)

	result, err = ldap.SearchAll("", scope, filter, attributes)

	logError(err)

	return result
}

func getHTTPRequest(str interface{}, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	logError(err)

	err = json.Unmarshal(body, &str)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Println("ERROR: ", err.Error())
	}
}

func sendJSON(js []byte, err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
}
