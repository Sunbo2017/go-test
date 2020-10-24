module go-test

go 1.12

require (
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.1.0
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/tools/gopls v0.5.1 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190422165155-953cdadca894
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190723021737-8bb11ff117ca
)
