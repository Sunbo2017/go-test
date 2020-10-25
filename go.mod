module go-test

go 1.12

require (
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.1.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/unidoc/unipdf/v3 v3.13.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
	rsc.io/pdf v0.1.1
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190422165155-953cdadca894
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190723021737-8bb11ff117ca
)
