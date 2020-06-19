package gil

import (
	"github.com/ArturMartini/gel"
	"os"
	"strings"
)

var (
	envsParams     = map[string]string{}
	envsConfigured = map[string]string{}
	envsRequired   = []string{}
)

func loadEnvs() {
	envsConfigured = gel.GetMap("gil_env.values")
	envsRequired = gel.GetList("gil_env.required")
	envsParams =  envsConfigured
	for _, arg := range os.Args {
		keyValue := strings.Split(arg, "=")
		if len(keyValue) > 1 {
			key := keyValue[0]
			value := keyValue[1]
			if _, ok := envsConfigured[key]; ok {
				envsParams[key] = value
			}
		}
	}
}
