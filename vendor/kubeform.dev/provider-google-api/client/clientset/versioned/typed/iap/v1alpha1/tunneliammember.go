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

	v1alpha1 "kubeform.dev/provider-google-api/apis/iap/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TunnelIamMembersGetter has a method to return a TunnelIamMemberInterface.
// A group's client should implement this interface.
type TunnelIamMembersGetter interface {
	TunnelIamMembers(namespace string) TunnelIamMemberInterface
}

// TunnelIamMemberInterface has methods to work with TunnelIamMember resources.
type TunnelIamMemberInterface interface {
	Create(ctx context.Context, tunnelIamMember *v1alpha1.TunnelIamMember, opts v1.CreateOptions) (*v1alpha1.TunnelIamMember, error)
	Update(ctx context.Context, tunnelIamMember *v1alpha1.TunnelIamMember, opts v1.UpdateOptions) (*v1alpha1.TunnelIamMember, error)
	UpdateStatus(ctx context.Context, tunnelIamMember *v1alpha1.TunnelIamMember, opts v1.UpdateOptions) (*v1alpha1.TunnelIamMember, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TunnelIamMember, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TunnelIamMemberList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TunnelIamMember, err error)
	TunnelIamMemberExpansion
}

// tunnelIamMembers implements TunnelIamMemberInterface
type tunnelIamMembers struct {
	client rest.Interface
	ns     string
}

// newTunnelIamMembers returns a TunnelIamMembers
func newTunnelIamMembers(c *IapV1alpha1Client, namespace string) *tunnelIamMembers {
	return &tunnelIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tunnelIamMember, and returns the corresponding tunnelIamMember object, and an error if there is any.
func (c *tunnelIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TunnelIamMember, err error) {
	result = &v1alpha1.TunnelIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tunneliammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TunnelIamMembers that match those selectors.
func (c *tunnelIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TunnelIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TunnelIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tunneliammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tunnelIamMembers.
func (c *tunnelIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tunneliammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tunnelIamMember and creates it.  Returns the server's representation of the tunnelIamMember, and an error, if there is any.
func (c *tunnelIamMembers) Create(ctx context.Context, tunnelIamMember *v1alpha1.TunnelIamMember, opts v1.CreateOptions) (result *v1alpha1.TunnelIamMember, err error) {
	result = &v1alpha1.TunnelIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tunneliammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tunnelIamMember).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tunnelIamMember and updates it. Returns the server's representation of the tunnelIamMember, and an error, if there is any.
func (c *tunnelIamMembers) Update(ctx context.Context, tunnelIamMember *v1alpha1.TunnelIamMember, opts v1.UpdateOptions) (result *v1alpha1.TunnelIamMember, err error) {
	result = &v1alpha1.TunnelIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tunneliammembers").
		Name(tunnelIamMember.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tunnelIamMember).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tunnelIamMembers) UpdateStatus(ctx context.Context, tunnelIamMember *v1alpha1.TunnelIamMember, opts v1.UpdateOptions) (result *v1alpha1.TunnelIamMember, err error) {
	result = &v1alpha1.TunnelIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tunneliammembers").
		Name(tunnelIamMember.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tunnelIamMember).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tunnelIamMember and deletes it. Returns an error if one occurs.
func (c *tunnelIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tunneliammembers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tunnelIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tunneliammembers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tunnelIamMember.
func (c *tunnelIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TunnelIamMember, err error) {
	result = &v1alpha1.TunnelIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tunneliammembers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}