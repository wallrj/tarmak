// Copyright Jetstack Ltd. See LICENSE for details.

// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/jetstack/tarmak/pkg/wing/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.SharedInformerFactory}
}
