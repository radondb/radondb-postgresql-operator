package kubeapi

/*
 Copyright 2020 - 2021 Crunchy Data Solutions, Inc.
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

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	radondb "github.com/radondb/radondb-postgresql-operator/pkg/generated/clientset/versioned"
	radondbscheme "github.com/radondb/radondb-postgresql-operator/pkg/generated/clientset/versioned/scheme"
	radondbv1 "github.com/radondb/radondb-postgresql-operator/pkg/generated/clientset/versioned/typed/radondb.com/v1"
)

func init() {
	// Register all types of our clientset into the standard scheme.
	_ = radondbscheme.AddToScheme(scheme.Scheme)
}

type Interface interface {
	kubernetes.Interface
	RadondbV1() radondbv1.RadondbV1Interface
}

// Interface should satisfy both our typed Interface and the standard one.
var (
	_ radondb.Interface    = Interface(nil)
	_ kubernetes.Interface = Interface(nil)
)

// Client provides methods for interacting with Kubernetes resources.
// It implements both kubernetes and radondb clientset Interfaces.
type Client struct {
	*rest.Config
	*kubernetes.Clientset

	radondbV1 *radondbv1.RadondbV1Client
}

// Client should satisfy Interface.
var _ Interface = &Client{}

// RadondbV1 retrieves the RadondbV1Client
func (c *Client) RadondbV1() radondbv1.RadondbV1Interface { return c.radondbV1 }

// LoadClientConfig prepares a configuration from the environment or home directory,
// falling back to in-cluster when applicable.
func LoadClientConfig() (*rest.Config, error) {
	// The default loading rules try to read from the files specified in the
	// environment or from the home directory.
	loader := clientcmd.NewDefaultClientConfigLoadingRules()

	// The deferred loader tries an in-cluster config if the default loading
	// rules produce no results.
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loader, &clientcmd.ConfigOverrides{},
	).ClientConfig()
}

// NewClient returns a kubernetes.Clientset and its underlying configuration.
func NewClient() (*Client, error) {
	config, err := LoadClientConfig()
	if err != nil {
		return nil, err
	}

	return NewClientForConfig(config)
}

// NewClientForConfig returns a kubernetes.Clientset using config.
func NewClientForConfig(config *rest.Config) (*Client, error) {
	var err error

	// Match the settings applied by sigs.k8s.io/controller-runtime@v0.6.0;
	// see https://github.com/kubernetes-sigs/controller-runtime/issues/365.
	if config.QPS == 0.0 {
		config.QPS = 20.0
		config.Burst = 30.0
	}

	client := &Client{Config: config}
	client.Clientset, err = kubernetes.NewForConfig(client.Config)

	if err == nil {
		client.radondbV1, err = radondbv1.NewForConfig(client.Config)
	}

	return client, err
}
