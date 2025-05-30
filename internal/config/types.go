package config

type Source int

type ConfValue struct {
	Value  string
	Source Source
}

type _config map[string]ConfValue

type TypeConfig _config

const (
	SrcNONE Source = iota
	SrcDEF
	SrcCMD
	SrcCFG
	SrcENV
	SrcCST
)

const (
	True = "true"
)

type Config struct {
	Vars TypeConfig
}
