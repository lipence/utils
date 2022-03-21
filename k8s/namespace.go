package k8sUtils

import (
	"bytes"
	"os"
)

var nameSpace string

const EnvNameSpace = "POD_NAMESPACE"
const PathNameSpaceFile = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

func init() {
	if data, err := os.ReadFile(PathNameSpaceFile); err == nil {
		nameSpace = string(bytes.TrimSpace(data))
	}
}

func NameSpace(defaultName string) string {
	if ns, ok := os.LookupEnv(EnvNameSpace); ok {
		return ns
	}
	if nameSpace == "" {
		return defaultName
	}
	return nameSpace
}
