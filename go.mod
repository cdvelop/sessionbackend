module github.com/cdvelop/sessionbackend

go 1.20

require (
	github.com/cdvelop/gookie v0.0.14
	github.com/cdvelop/model v0.0.113
	github.com/cdvelop/sessionhandler v0.0.22
	github.com/cdvelop/token v0.0.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	golang.org/x/oauth2 v0.15.0
)

require (
	github.com/cdvelop/object v0.0.71 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/sessionhandler => ../sessionhandler
