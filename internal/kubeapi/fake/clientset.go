package fake

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
	fakekubernetes "k8s.io/client-go/kubernetes/fake"

	"github.com/randondb/postgres-operator/internal/kubeapi"
	fakerandondb "github.com/randondb/postgres-operator/pkg/generated/clientset/versioned/fake"
	randondbv1 "github.com/randondb/postgres-operator/pkg/generated/clientset/versioned/typed/randondb.com/v1"
)

type Clientset struct {
	*fakekubernetes.Clientset
	PGOClientset *fakerandondb.Clientset
}

var _ kubeapi.Interface = &Clientset{}

// RadondbV1 retrieves the RadondbV1Client
func (c *Clientset) RadondbV1() randondbv1.RadondbV1Interface {
	return c.PGOClientset.RadondbV1()
}
