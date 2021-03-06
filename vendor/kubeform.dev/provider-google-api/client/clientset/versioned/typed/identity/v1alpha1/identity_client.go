/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-google-api/apis/identity/v1alpha1"
	"kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type IdentityV1alpha1Interface interface {
	RESTClient() rest.Interface
	PlatformDefaultSupportedIdpConfigsGetter
	PlatformInboundSamlConfigsGetter
	PlatformOauthIdpConfigsGetter
	PlatformTenantsGetter
	PlatformTenantDefaultSupportedIdpConfigsGetter
	PlatformTenantInboundSamlConfigsGetter
	PlatformTenantOauthIdpConfigsGetter
}

// IdentityV1alpha1Client is used to interact with features provided by the identity.google.kubeform.com group.
type IdentityV1alpha1Client struct {
	restClient rest.Interface
}

func (c *IdentityV1alpha1Client) PlatformDefaultSupportedIdpConfigs(namespace string) PlatformDefaultSupportedIdpConfigInterface {
	return newPlatformDefaultSupportedIdpConfigs(c, namespace)
}

func (c *IdentityV1alpha1Client) PlatformInboundSamlConfigs(namespace string) PlatformInboundSamlConfigInterface {
	return newPlatformInboundSamlConfigs(c, namespace)
}

func (c *IdentityV1alpha1Client) PlatformOauthIdpConfigs(namespace string) PlatformOauthIdpConfigInterface {
	return newPlatformOauthIdpConfigs(c, namespace)
}

func (c *IdentityV1alpha1Client) PlatformTenants(namespace string) PlatformTenantInterface {
	return newPlatformTenants(c, namespace)
}

func (c *IdentityV1alpha1Client) PlatformTenantDefaultSupportedIdpConfigs(namespace string) PlatformTenantDefaultSupportedIdpConfigInterface {
	return newPlatformTenantDefaultSupportedIdpConfigs(c, namespace)
}

func (c *IdentityV1alpha1Client) PlatformTenantInboundSamlConfigs(namespace string) PlatformTenantInboundSamlConfigInterface {
	return newPlatformTenantInboundSamlConfigs(c, namespace)
}

func (c *IdentityV1alpha1Client) PlatformTenantOauthIdpConfigs(namespace string) PlatformTenantOauthIdpConfigInterface {
	return newPlatformTenantOauthIdpConfigs(c, namespace)
}

// NewForConfig creates a new IdentityV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*IdentityV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IdentityV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new IdentityV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *IdentityV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new IdentityV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *IdentityV1alpha1Client {
	return &IdentityV1alpha1Client{c}
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
func (c *IdentityV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
