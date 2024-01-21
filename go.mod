module github.com/cdvelop/sessionbackend

go 1.20

require (
	github.com/cdvelop/gookie v0.0.14
	github.com/cdvelop/model v0.0.119
	github.com/cdvelop/sessionhandler v0.0.24
	github.com/cdvelop/timetools v0.0.40
	github.com/cdvelop/token v0.0.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	golang.org/x/oauth2 v0.16.0
)

require (
	github.com/cdvelop/object v0.0.75 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/sessionhandler => ../sessionhandler
