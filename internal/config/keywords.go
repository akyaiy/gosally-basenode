package config

type _configConst struct {
	Keywords      []string
	Abbreviations map[string]string
}

var Defines = _configConst{
	Keywords: []string{
		"listen-address", "listen-port",
		"debug",
	},
	Abbreviations: map[string]string{
		"d": "debug",
		"a": "listen-address",
		"p": "listen-port",
	},
}
