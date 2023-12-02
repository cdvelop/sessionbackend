module github.com/cdvelop/sessionbackend

go 1.20

require github.com/cdvelop/model v0.0.74

require github.com/cdvelop/sessionhandler v0.0.1

require (
	github.com/cdvelop/input v0.0.57 // indirect
	github.com/cdvelop/object v0.0.38 // indirect
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timetools v0.0.23 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/sessionhandler => ../sessionhandler
