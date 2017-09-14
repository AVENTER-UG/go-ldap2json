# go-ldapjson

go-ldapjson is a API gateway that will give a easy and JSON based access to authenticate user against LDAP.

To support our development, please take a small donation.
[![Donate](https://liberapay.com/assets/widgets/donate.svg)](https://liberapay.com/AVENTER/donate)

### How to run

**LDAP_USER** and **LDAP_PASSWORD** is the user to __use__ LDAP not to authenticate against. 

```
go run init.go API_SERVERPORT=8888 LDAP_SERVER=localhost:1368 LDAP_DN= LDAP_USER= LDAP_PASSWORD= LDAP_BASE=
```

### LDAP Test Server

To use the LDAP testserver, to the following steps. It important that you already have installed npm, nodejs and git. :-)

```
cd test_server
git clone https://github.com/mcavage/node-ldapjs.git  ldapjs
cd ldapjs
npm install
cd ..
node ldapserver.js
```



