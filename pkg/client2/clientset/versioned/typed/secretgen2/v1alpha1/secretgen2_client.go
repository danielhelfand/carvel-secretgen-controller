// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/carvel-secretgen-controller/pkg/apis/secretgen2/v1alpha1"
	"github.com/vmware-tanzu/carvel-secretgen-controller/pkg/client2/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SecretgenV1alpha1Interface interface {
	RESTClient() rest.Interface
	SecretExportsGetter
	SecretImportsGetter
}

// SecretgenV1alpha1Client is used to interact with features provided by the secretgen.carvel.dev group.
type SecretgenV1alpha1Client struct {
	restClient rest.Interface
}

func (c *SecretgenV1alpha1Client) SecretExports(namespace string) SecretExportInterface {
	return newSecretExports(c, namespace)
}

func (c *SecretgenV1alpha1Client) SecretImports(namespace string) SecretImportInterface {
	return newSecretImports(c, namespace)
}

// NewForConfig creates a new SecretgenV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*SecretgenV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SecretgenV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new SecretgenV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SecretgenV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SecretgenV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SecretgenV1alpha1Client {
	return &SecretgenV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SecretgenV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
