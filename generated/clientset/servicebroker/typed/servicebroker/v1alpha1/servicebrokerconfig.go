// Copyright (C) Couchbase, Inc 2020

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	scheme "github.com/couchbase/service-broker/generated/clientset/servicebroker/scheme"
	v1alpha1 "github.com/couchbase/service-broker/pkg/apis/servicebroker/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceBrokerConfigsGetter has a method to return a ServiceBrokerConfigInterface.
// A group's client should implement this interface.
type ServiceBrokerConfigsGetter interface {
	ServiceBrokerConfigs(namespace string) ServiceBrokerConfigInterface
}

// ServiceBrokerConfigInterface has methods to work with ServiceBrokerConfig resources.
type ServiceBrokerConfigInterface interface {
	Create(*v1alpha1.ServiceBrokerConfig) (*v1alpha1.ServiceBrokerConfig, error)
	Update(*v1alpha1.ServiceBrokerConfig) (*v1alpha1.ServiceBrokerConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceBrokerConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceBrokerConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceBrokerConfig, err error)
	ServiceBrokerConfigExpansion
}

// serviceBrokerConfigs implements ServiceBrokerConfigInterface
type serviceBrokerConfigs struct {
	client rest.Interface
	ns     string
}

// newServiceBrokerConfigs returns a ServiceBrokerConfigs
func newServiceBrokerConfigs(c *ServicebrokerV1alpha1Client, namespace string) *serviceBrokerConfigs {
	return &serviceBrokerConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceBrokerConfig, and returns the corresponding serviceBrokerConfig object, and an error if there is any.
func (c *serviceBrokerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceBrokerConfig, err error) {
	result = &v1alpha1.ServiceBrokerConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceBrokerConfigs that match those selectors.
func (c *serviceBrokerConfigs) List(opts v1.ListOptions) (result *v1alpha1.ServiceBrokerConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceBrokerConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceBrokerConfigs.
func (c *serviceBrokerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a serviceBrokerConfig and creates it.  Returns the server's representation of the serviceBrokerConfig, and an error, if there is any.
func (c *serviceBrokerConfigs) Create(serviceBrokerConfig *v1alpha1.ServiceBrokerConfig) (result *v1alpha1.ServiceBrokerConfig, err error) {
	result = &v1alpha1.ServiceBrokerConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		Body(serviceBrokerConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceBrokerConfig and updates it. Returns the server's representation of the serviceBrokerConfig, and an error, if there is any.
func (c *serviceBrokerConfigs) Update(serviceBrokerConfig *v1alpha1.ServiceBrokerConfig) (result *v1alpha1.ServiceBrokerConfig, err error) {
	result = &v1alpha1.ServiceBrokerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		Name(serviceBrokerConfig.Name).
		Body(serviceBrokerConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceBrokerConfig and deletes it. Returns an error if one occurs.
func (c *serviceBrokerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceBrokerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceBrokerConfig.
func (c *serviceBrokerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceBrokerConfig, err error) {
	result = &v1alpha1.ServiceBrokerConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicebrokerconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
