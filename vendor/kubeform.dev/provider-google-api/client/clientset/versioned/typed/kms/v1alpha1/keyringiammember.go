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

	v1alpha1 "kubeform.dev/provider-google-api/apis/kms/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KeyRingIamMembersGetter has a method to return a KeyRingIamMemberInterface.
// A group's client should implement this interface.
type KeyRingIamMembersGetter interface {
	KeyRingIamMembers(namespace string) KeyRingIamMemberInterface
}

// KeyRingIamMemberInterface has methods to work with KeyRingIamMember resources.
type KeyRingIamMemberInterface interface {
	Create(ctx context.Context, keyRingIamMember *v1alpha1.KeyRingIamMember, opts v1.CreateOptions) (*v1alpha1.KeyRingIamMember, error)
	Update(ctx context.Context, keyRingIamMember *v1alpha1.KeyRingIamMember, opts v1.UpdateOptions) (*v1alpha1.KeyRingIamMember, error)
	UpdateStatus(ctx context.Context, keyRingIamMember *v1alpha1.KeyRingIamMember, opts v1.UpdateOptions) (*v1alpha1.KeyRingIamMember, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KeyRingIamMember, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KeyRingIamMemberList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeyRingIamMember, err error)
	KeyRingIamMemberExpansion
}

// keyRingIamMembers implements KeyRingIamMemberInterface
type keyRingIamMembers struct {
	client rest.Interface
	ns     string
}

// newKeyRingIamMembers returns a KeyRingIamMembers
func newKeyRingIamMembers(c *KmsV1alpha1Client, namespace string) *keyRingIamMembers {
	return &keyRingIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the keyRingIamMember, and returns the corresponding keyRingIamMember object, and an error if there is any.
func (c *keyRingIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KeyRingIamMember, err error) {
	result = &v1alpha1.KeyRingIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keyringiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KeyRingIamMembers that match those selectors.
func (c *keyRingIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeyRingIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KeyRingIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keyringiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keyRingIamMembers.
func (c *keyRingIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keyringiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a keyRingIamMember and creates it.  Returns the server's representation of the keyRingIamMember, and an error, if there is any.
func (c *keyRingIamMembers) Create(ctx context.Context, keyRingIamMember *v1alpha1.KeyRingIamMember, opts v1.CreateOptions) (result *v1alpha1.KeyRingIamMember, err error) {
	result = &v1alpha1.KeyRingIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keyringiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keyRingIamMember).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a keyRingIamMember and updates it. Returns the server's representation of the keyRingIamMember, and an error, if there is any.
func (c *keyRingIamMembers) Update(ctx context.Context, keyRingIamMember *v1alpha1.KeyRingIamMember, opts v1.UpdateOptions) (result *v1alpha1.KeyRingIamMember, err error) {
	result = &v1alpha1.KeyRingIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keyringiammembers").
		Name(keyRingIamMember.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keyRingIamMember).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *keyRingIamMembers) UpdateStatus(ctx context.Context, keyRingIamMember *v1alpha1.KeyRingIamMember, opts v1.UpdateOptions) (result *v1alpha1.KeyRingIamMember, err error) {
	result = &v1alpha1.KeyRingIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keyringiammembers").
		Name(keyRingIamMember.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keyRingIamMember).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the keyRingIamMember and deletes it. Returns an error if one occurs.
func (c *keyRingIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keyringiammembers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keyRingIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keyringiammembers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched keyRingIamMember.
func (c *keyRingIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeyRingIamMember, err error) {
	result = &v1alpha1.KeyRingIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keyringiammembers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
