// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openshift/api/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsolesGetter has a method to return a ConsoleInterface.
// A group's client should implement this interface.
type ConsolesGetter interface {
	Consoles() ConsoleInterface
}

// ConsoleInterface has methods to work with Console resources.
type ConsoleInterface interface {
	Create(ctx context.Context, console *v1.Console, opts metav1.CreateOptions) (*v1.Console, error)
	Update(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (*v1.Console, error)
	UpdateStatus(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (*v1.Console, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Console, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ConsoleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Console, err error)
	ConsoleExpansion
}

// consoles implements ConsoleInterface
type consoles struct {
	client rest.Interface
}

// newConsoles returns a Consoles
func newConsoles(c *OperatorV1Client) *consoles {
	return &consoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the console, and returns the corresponding console object, and an error if there is any.
func (c *consoles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Get().
		Resource("consoles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Consoles that match those selectors.
func (c *consoles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConsoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ConsoleList{}
	err = c.client.Get().
		Resource("consoles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consoles.
func (c *consoles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("consoles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a console and creates it.  Returns the server's representation of the console, and an error, if there is any.
func (c *consoles) Create(ctx context.Context, console *v1.Console, opts metav1.CreateOptions) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Post().
		Resource("consoles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(console).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a console and updates it. Returns the server's representation of the console, and an error, if there is any.
func (c *consoles) Update(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Put().
		Resource("consoles").
		Name(console.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(console).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *consoles) UpdateStatus(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Put().
		Resource("consoles").
		Name(console.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(console).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the console and deletes it. Returns an error if one occurs.
func (c *consoles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("consoles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consoles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("consoles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched console.
func (c *consoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Console, err error) {
	result = &v1.Console{}
	err = c.client.Patch(pt).
		Resource("consoles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
