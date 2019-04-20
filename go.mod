module petmate

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20190130090550-b01c7a725664 => github.com/golang/crypto v0.0.0-20190130090550-b01c7a725664
	golang.org/x/sync v0.0.0 => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys v0.0.0-20190129075346-302c3dd5f1cc => github.com/golang/sys v0.0.0-20190129075346-302c3dd5f1cc
	golang.org/x/text v0.0.0 => github.com/golang/text v0.3.0
)

require (
	github.com/go-stack/stack v1.8.0 // indirect; indirect.0
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.2.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/panjf2000/ants v4.0.2+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	golang.org/x/crypto v0.0.0-20190130090550-b01c7a725664 // indirect
	golang.org/x/sync v0.0.0 // indirect
	golang.org/x/text v0.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
