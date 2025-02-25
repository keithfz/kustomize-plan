package internal

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name string `yaml:"name"`
}

type Manifest struct {
	ObjMeta    Metadata `yaml:"metadata"`
	Kind       string   `yaml:"kind"`
	ApiVersion string   `yaml:"apiVersion"`
}

func ParseFile(filepath string) (map[string]string, error) {
	kustomization, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	manifests := strings.Split(string(kustomization), "\n---\n")
	manifestMap := make(map[string]string)
	var parsedManifest Manifest
	var key string
	for _, manifest := range manifests {
		err := yaml.Unmarshal([]byte(manifest), &parsedManifest)
		if err != nil {
			return nil, err
		}

		key = fmt.Sprintf("%s %s %s", parsedManifest.ObjMeta.Name, parsedManifest.Kind, parsedManifest.ApiVersion)
		manifestMap[key] = manifest
	}
	return manifestMap, nil

}
