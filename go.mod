module github.com/cdvelop/sessionbackend

go 1.20

require github.com/cdvelop/model v0.0.77

require github.com/cdvelop/sessionhandler v0.0.4

require (
	github.com/cdvelop/object v0.0.41 // indirect
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/token v0.0.3 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/sessionhandler => ../sessionhandler
