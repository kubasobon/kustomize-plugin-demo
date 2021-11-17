package main

import (
	"os"

	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"
	kyaml "sigs.k8s.io/kustomize/kyaml/yaml"
)

type Konfigure struct {
	Data map[string]string `yaml:"data" json:"data"`
}

func main() {
	config := new(Konfigure)
	fn := func(items []*kyaml.RNode) ([]*kyaml.RNode, error) {
		cm := corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "ConfigMap",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "from-plugin",
			},
			Data: config.Data,
		}

		bytes, err := yaml.Marshal(&cm)
		if err != nil {
			return nil, err
		}

		newItem, err := kyaml.Parse(string(bytes))
		if err != nil {
			return nil, err
		}

		items = append(items, newItem)
		return items, nil
	}
	p := framework.SimpleProcessor{Config: config, Filter: kio.FilterFunc(fn)}
	cmd := command.Build(p, command.StandaloneDisabled, false)
	command.AddGenerateDockerfile(cmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
