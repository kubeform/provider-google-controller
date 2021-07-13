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
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-google-api/apis/cloud/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RunServiceIamPoliciesGetter has a method to return a RunServiceIamPolicyInterface.
// A group's client should implement this interface.
type RunServiceIamPoliciesGetter interface {
	RunServiceIamPolicies(namespace string) RunServiceIamPolicyInterface
}

// RunServiceIamPolicyInterface has methods to work with RunServiceIamPolicy resources.
type RunServiceIamPolicyInterface interface {
	Create(ctx context.Context, runServiceIamPolicy *v1alpha1.RunServiceIamPolicy, opts v1.CreateOptions) (*v1alpha1.RunServiceIamPolicy, error)
	Update(ctx context.Context, runServiceIamPolicy *v1alpha1.RunServiceIamPolicy, opts v1.UpdateOptions) (*v1alpha1.RunServiceIamPolicy, error)
	UpdateStatus(ctx context.Context, runServiceIamPolicy *v1alpha1.RunServiceIamPolicy, opts v1.UpdateOptions) (*v1alpha1.RunServiceIamPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RunServiceIamPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RunServiceIamPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RunServiceIamPolicy, err error)
	RunServiceIamPolicyExpansion
}

// runServiceIamPolicies implements RunServiceIamPolicyInterface
type runServiceIamPolicies struct {
	client rest.Interface
	ns     string
}

// newRunServiceIamPolicies returns a RunServiceIamPolicies
func newRunServiceIamPolicies(c *CloudV1alpha1Client, namespace string) *runServiceIamPolicies {
	return &runServiceIamPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the runServiceIamPolicy, and returns the corresponding runServiceIamPolicy object, and an error if there is any.
func (c *runServiceIamPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RunServiceIamPolicy, err error) {
	result = &v1alpha1.RunServiceIamPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RunServiceIamPolicies that match those selectors.
func (c *runServiceIamPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RunServiceIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RunServiceIamPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested runServiceIamPolicies.
func (c *runServiceIamPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a runServiceIamPolicy and creates it.  Returns the server's representation of the runServiceIamPolicy, and an error, if there is any.
func (c *runServiceIamPolicies) Create(ctx context.Context, runServiceIamPolicy *v1alpha1.RunServiceIamPolicy, opts v1.CreateOptions) (result *v1alpha1.RunServiceIamPolicy, err error) {
	result = &v1alpha1.RunServiceIamPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runServiceIamPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a runServiceIamPolicy and updates it. Returns the server's representation of the runServiceIamPolicy, and an error, if there is any.
func (c *runServiceIamPolicies) Update(ctx context.Context, runServiceIamPolicy *v1alpha1.RunServiceIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.RunServiceIamPolicy, err error) {
	result = &v1alpha1.RunServiceIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		Name(runServiceIamPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runServiceIamPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *runServiceIamPolicies) UpdateStatus(ctx context.Context, runServiceIamPolicy *v1alpha1.RunServiceIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.RunServiceIamPolicy, err error) {
	result = &v1alpha1.RunServiceIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		Name(runServiceIamPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runServiceIamPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the runServiceIamPolicy and deletes it. Returns an error if one occurs.
func (c *runServiceIamPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *runServiceIamPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched runServiceIamPolicy.
func (c *runServiceIamPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RunServiceIamPolicy, err error) {
	result = &v1alpha1.RunServiceIamPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("runserviceiampolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
