// Copyright (C) Couchbase, Inc 2020

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/couchbase/service-broker/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ServiceBrokerConfigs returns a ServiceBrokerConfigInformer.
	ServiceBrokerConfigs() ServiceBrokerConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ServiceBrokerConfigs returns a ServiceBrokerConfigInformer.
func (v *version) ServiceBrokerConfigs() ServiceBrokerConfigInformer {
	return &serviceBrokerConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
