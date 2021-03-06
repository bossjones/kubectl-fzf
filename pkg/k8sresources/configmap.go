package k8sresources

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

const ConfigMapHeader = "Namespace Name Age Labels\n"

// ConfigMap is the summary of a kubernetes configMap
type ConfigMap struct {
	ResourceMeta
}

// NewConfigMapFromRuntime builds a pod from informer result
func NewConfigMapFromRuntime(obj interface{}, config CtorConfig) K8sResource {
	c := &ConfigMap{}
	c.FromRuntime(obj, config)
	return c
}

// FromRuntime builds object from the informer's result
func (c *ConfigMap) FromRuntime(obj interface{}, config CtorConfig) {
	configMap := obj.(*corev1.ConfigMap)
	c.FromObjectMeta(configMap.ObjectMeta)
}

// HasChanged returns true if the resource's dump needs to be updated
func (c *ConfigMap) HasChanged(k K8sResource) bool {
	return true
}

// ToString serializes the object to strings
func (c *ConfigMap) ToString() string {
	line := strings.Join([]string{c.namespace,
		c.name,
		c.resourceAge(),
		c.labelsString(),
	}, " ")
	return fmt.Sprintf("%s\n", line)
}
