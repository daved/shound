module github.com/daved/shound

go 1.21.9

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/daved/clic v0.0.0
	github.com/daved/flagset v0.0.0
)

replace (
	github.com/daved/clic => ../clic
	github.com/daved/flagset => ../flagset
)
