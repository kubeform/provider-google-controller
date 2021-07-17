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
	v1alpha1 "kubeform.dev/provider-google-api/apis/binaryauthorization/v1alpha1"
	"kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type BinaryauthorizationV1alpha1Interface interface {
	RESTClient() rest.Interface
	AttestorsGetter
	AttestorIamBindingsGetter
	AttestorIamMembersGetter
	AttestorIamPoliciesGetter
	PoliciesGetter
}

// BinaryauthorizationV1alpha1Client is used to interact with features provided by the binaryauthorization.google.kubeform.com group.
type BinaryauthorizationV1alpha1Client struct {
	restClient rest.Interface
}

func (c *BinaryauthorizationV1alpha1Client) Attestors(namespace string) AttestorInterface {
	return newAttestors(c, namespace)
}

func (c *BinaryauthorizationV1alpha1Client) AttestorIamBindings(namespace string) AttestorIamBindingInterface {
	return newAttestorIamBindings(c, namespace)
}

func (c *BinaryauthorizationV1alpha1Client) AttestorIamMembers(namespace string) AttestorIamMemberInterface {
	return newAttestorIamMembers(c, namespace)
}

func (c *BinaryauthorizationV1alpha1Client) AttestorIamPolicies(namespace string) AttestorIamPolicyInterface {
	return newAttestorIamPolicies(c, namespace)
}

func (c *BinaryauthorizationV1alpha1Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

// NewForConfig creates a new BinaryauthorizationV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*BinaryauthorizationV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BinaryauthorizationV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new BinaryauthorizationV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BinaryauthorizationV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BinaryauthorizationV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *BinaryauthorizationV1alpha1Client {
	return &BinaryauthorizationV1alpha1Client{c}
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
func (c *BinaryauthorizationV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}