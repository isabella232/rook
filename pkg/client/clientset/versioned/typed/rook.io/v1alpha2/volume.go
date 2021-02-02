/*
Copyright The Kubernetes Authors.

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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/rook/rook/pkg/apis/rook.io/v1alpha2"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumesGetter has a method to return a VolumeInterface.
// A group's client should implement this interface.
type VolumesGetter interface {
	Volumes(namespace string) VolumeInterface
}

// VolumeInterface has methods to work with Volume resources.
type VolumeInterface interface {
	Create(ctx context.Context, volume *v1alpha2.Volume, opts v1.CreateOptions) (*v1alpha2.Volume, error)
	Update(ctx context.Context, volume *v1alpha2.Volume, opts v1.UpdateOptions) (*v1alpha2.Volume, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.Volume, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.VolumeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Volume, err error)
	VolumeExpansion
}

// volumes implements VolumeInterface
type volumes struct {
	client rest.Interface
	ns     string
}

// newVolumes returns a Volumes
func newVolumes(c *RookV1alpha2Client, namespace string) *volumes {
	return &volumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volume, and returns the corresponding volume object, and an error if there is any.
func (c *volumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Volume, err error) {
	result = &v1alpha2.Volume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Volumes that match those selectors.
func (c *volumes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.VolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.VolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumes.
func (c *volumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volume and creates it.  Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Create(ctx context.Context, volume *v1alpha2.Volume, opts v1.CreateOptions) (result *v1alpha2.Volume, err error) {
	result = &v1alpha2.Volume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volume).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volume and updates it. Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Update(ctx context.Context, volume *v1alpha2.Volume, opts v1.UpdateOptions) (result *v1alpha2.Volume, err error) {
	result = &v1alpha2.Volume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumes").
		Name(volume.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volume).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volume and deletes it. Returns an error if one occurs.
func (c *volumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volume.
func (c *volumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Volume, err error) {
	result = &v1alpha2.Volume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
