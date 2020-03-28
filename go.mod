module github.com/cuttle-ai/go-sdk

replace github.com/cuttle-ai/db-toolkit => ../db-toolkit/

replace github.com/cuttle-ai/octopus => ../octopus/

go 1.13

require (
	github.com/cuttle-ai/db-toolkit v0.0.0-00010101000000-000000000000
	github.com/cuttle-ai/octopus v0.0.0-00010101000000-000000000000
	github.com/gojektech/heimdall v5.0.2+incompatible
	github.com/gojektech/valkyrie v0.0.0-20190210220504-8f62c1e7ba45 // indirect
	github.com/hashicorp/consul/api v1.4.0
)
