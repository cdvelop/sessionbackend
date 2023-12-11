module github.com/cdvelop/sessionbackend

go 1.20

require github.com/cdvelop/model v0.0.78

require github.com/cdvelop/sessionhandler v0.0.6

require (
	github.com/golang/protobuf v1.5.3 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

require (
	github.com/cdvelop/gookie v0.0.1
	github.com/cdvelop/object v0.0.42 // indirect
	github.com/cdvelop/strings v0.0.8 // indirect
	github.com/cdvelop/token v0.0.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	golang.org/x/oauth2 v0.15.0
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/sessionhandler => ../sessionhandler

replace github.com/cdvelop/gookie => ../gookie
