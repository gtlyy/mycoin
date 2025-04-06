module mycoin

go 1.19

require (
	github.com/btcsuite/btcd v0.22.2
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/stretchr/testify v1.8.2
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	// mymariadb v0.0.0-00010101000000-000000000000
)

require (
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd // indirect
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace mytime => ../mytime
// replace myfun => ../myfun
// replace mymariadb => ../mymariadb
