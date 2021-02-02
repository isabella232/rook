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

package versioned

import (
	"fmt"

	cassandrav1alpha1 "github.com/rook/rook/pkg/client/clientset/versioned/typed/cassandra.rook.io/v1alpha1"
	cephv1 "github.com/rook/rook/pkg/client/clientset/versioned/typed/ceph.rook.io/v1"
	nfsv1alpha1 "github.com/rook/rook/pkg/client/clientset/versioned/typed/nfs.rook.io/v1alpha1"
	rookv1 "github.com/rook/rook/pkg/client/clientset/versioned/typed/rook.io/v1"
	rookv1alpha2 "github.com/rook/rook/pkg/client/clientset/versioned/typed/rook.io/v1alpha2"
	yugabytedbv1alpha1 "github.com/rook/rook/pkg/client/clientset/versioned/typed/yugabytedb.rook.io/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CassandraV1alpha1() cassandrav1alpha1.CassandraV1alpha1Interface
	CephV1() cephv1.CephV1Interface
	NfsV1alpha1() nfsv1alpha1.NfsV1alpha1Interface
	RookV1() rookv1.RookV1Interface
	RookV1alpha2() rookv1alpha2.RookV1alpha2Interface
	YugabytedbV1alpha1() yugabytedbv1alpha1.YugabytedbV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	cassandraV1alpha1  *cassandrav1alpha1.CassandraV1alpha1Client
	cephV1             *cephv1.CephV1Client
	nfsV1alpha1        *nfsv1alpha1.NfsV1alpha1Client
	rookV1             *rookv1.RookV1Client
	rookV1alpha2       *rookv1alpha2.RookV1alpha2Client
	yugabytedbV1alpha1 *yugabytedbv1alpha1.YugabytedbV1alpha1Client
}

// CassandraV1alpha1 retrieves the CassandraV1alpha1Client
func (c *Clientset) CassandraV1alpha1() cassandrav1alpha1.CassandraV1alpha1Interface {
	return c.cassandraV1alpha1
}

// CephV1 retrieves the CephV1Client
func (c *Clientset) CephV1() cephv1.CephV1Interface {
	return c.cephV1
}

// NfsV1alpha1 retrieves the NfsV1alpha1Client
func (c *Clientset) NfsV1alpha1() nfsv1alpha1.NfsV1alpha1Interface {
	return c.nfsV1alpha1
}

// RookV1 retrieves the RookV1Client
func (c *Clientset) RookV1() rookv1.RookV1Interface {
	return c.rookV1
}

// RookV1alpha2 retrieves the RookV1alpha2Client
func (c *Clientset) RookV1alpha2() rookv1alpha2.RookV1alpha2Interface {
	return c.rookV1alpha2
}

// YugabytedbV1alpha1 retrieves the YugabytedbV1alpha1Client
func (c *Clientset) YugabytedbV1alpha1() yugabytedbv1alpha1.YugabytedbV1alpha1Interface {
	return c.yugabytedbV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.cassandraV1alpha1, err = cassandrav1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cephV1, err = cephv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.nfsV1alpha1, err = nfsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rookV1, err = rookv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rookV1alpha2, err = rookv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.yugabytedbV1alpha1, err = yugabytedbv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.cassandraV1alpha1 = cassandrav1alpha1.NewForConfigOrDie(c)
	cs.cephV1 = cephv1.NewForConfigOrDie(c)
	cs.nfsV1alpha1 = nfsv1alpha1.NewForConfigOrDie(c)
	cs.rookV1 = rookv1.NewForConfigOrDie(c)
	cs.rookV1alpha2 = rookv1alpha2.NewForConfigOrDie(c)
	cs.yugabytedbV1alpha1 = yugabytedbv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.cassandraV1alpha1 = cassandrav1alpha1.New(c)
	cs.cephV1 = cephv1.New(c)
	cs.nfsV1alpha1 = nfsv1alpha1.New(c)
	cs.rookV1 = rookv1.New(c)
	cs.rookV1alpha2 = rookv1alpha2.New(c)
	cs.yugabytedbV1alpha1 = yugabytedbv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
