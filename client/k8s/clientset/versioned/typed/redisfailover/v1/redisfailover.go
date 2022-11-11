// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/spotahome/redis-operator/api/redisfailover/v1"
	scheme "github.com/spotahome/redis-operator/client/k8s/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedisFailoversGetter has a method to return a RedisFailoverInterface.
// A group's client should implement this interface.
type RedisFailoversGetter interface {
	RedisFailovers(namespace string) RedisFailoverInterface
}

// RedisFailoverInterface has methods to work with RedisFailover resources.
type RedisFailoverInterface interface {
	Create(ctx context.Context, redisFailover *v1.RedisFailover, opts metav1.CreateOptions) (*v1.RedisFailover, error)
	Update(ctx context.Context, redisFailover *v1.RedisFailover, opts metav1.UpdateOptions) (*v1.RedisFailover, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RedisFailover, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RedisFailoverList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RedisFailover, err error)
	RedisFailoverExpansion
}

// redisFailovers implements RedisFailoverInterface
type redisFailovers struct {
	client rest.Interface
	ns     string
}

// newRedisFailovers returns a RedisFailovers
func newRedisFailovers(c *DatabasesV1Client, namespace string) *redisFailovers {
	return &redisFailovers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisFailover, and returns the corresponding redisFailover object, and an error if there is any.
func (c *redisFailovers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RedisFailover, err error) {
	result = &v1.RedisFailover{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisFailovers that match those selectors.
func (c *redisFailovers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RedisFailoverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RedisFailoverList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisfailovers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisFailovers.
func (c *redisFailovers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisfailovers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a redisFailover and creates it.  Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *redisFailovers) Create(ctx context.Context, redisFailover *v1.RedisFailover, opts metav1.CreateOptions) (result *v1.RedisFailover, err error) {
	result = &v1.RedisFailover{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisfailovers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisFailover).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a redisFailover and updates it. Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *redisFailovers) Update(ctx context.Context, redisFailover *v1.RedisFailover, opts metav1.UpdateOptions) (result *v1.RedisFailover, err error) {
	result = &v1.RedisFailover{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(redisFailover.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisFailover).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the redisFailover and deletes it. Returns an error if one occurs.
func (c *redisFailovers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisFailovers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisfailovers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched redisFailover.
func (c *redisFailovers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RedisFailover, err error) {
	result = &v1.RedisFailover{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
