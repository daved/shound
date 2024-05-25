module github.com/daved/shound

go 1.21.9

require (
	github.com/daved/clic v0.0.0
	github.com/daved/flagset v0.0.0
	gopkg.in/yaml.v3 v3.0.1
)

replace (
	github.com/daved/clic => ../clic
	github.com/daved/flagset => ../flagset
)
