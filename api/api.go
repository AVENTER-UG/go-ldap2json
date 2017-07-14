package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mqu/openldap"
	//"io/ioutil"

	"net/http"
)

type userStruct struct {
	UUID     int
	USERNAME string
}

var ldapServer string   // LDAP SERVER
var ldapDN string       // LDAP DN
var ldapUser string     // LDAP Username
var ldapPassword string // LDAP Password

func init() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/versions", apiVersions).Methods("GET")
	rtr.HandleFunc("/api", apiVersions).Methods("GET")
	rtr.HandleFunc("/api/v0", apiV0Version).Methods("GET")
	rtr.HandleFunc("/api/v0/version", apiV0Version).Methods("GET")

	rtr.HandleFunc("/api/v0/getUser", apiV0GetUser).Methods("GET")

	//rtr.HandleFunc("/user/dieter/profile", api_user_profile).Methods("GET")
	http.Handle("/", rtr)

	initLDAP()
}

func apiVersions(w http.ResponseWriter, r *http.Request) {
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

	var user userStruct

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")

	d, _ := json.Marshal(searchLDAP(username))

	w.Write([]byte(string(d)))
}

func initLDAP() {
	ldap, err := openldap.Initialize(ldapServer)

	if err != nil {
		fmt.Printf("LDAP::Initialize() : connection error\n")
		return
	}

	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

	err = ldap.Bind(ldapUser, ldapPassword)
	if err != nil {
		fmt.Printf("LDAP::Bind() : bind error\n")
		fmt.Println(err)
		return
	}
	defer ldap.Close()
}

func searchLDAP(username string) string {
	var scope = openldap.LDAP_SCOPE_SUBTREE // LDAP_SCOPE_BASE, LDAP_SCOPE_ONELEVEL, LDAP_SCOPE_SUBTREE
	var filter = "cn=" + username
	var attributes = []string{"cn", "uuid", "givenname", "mail"}

	result, err := ldap.SearchAll(base, scope, filter, attributes)

	if err != nil {
		fmt.Println("LDAP Search Error: %s", err)
		return
	}

	// Only Debug
	fmt.Printf("# num results : %d\n", result.Count())
	fmt.Printf("# search : %s\n", result.Filter())
	fmt.Printf("# base : %s\n", result.Base())
	fmt.Printf("# attributes : [%s]\n", strings.Join(result.Attributes(), ", "))

	var res = result.Attributes()

	return res
}
