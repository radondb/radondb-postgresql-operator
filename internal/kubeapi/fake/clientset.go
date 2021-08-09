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

	"github.com/qingcloud/postgres-operator/internal/kubeapi"
	fakeqingcloud "github.com/qingcloud/postgres-operator/pkg/generated/clientset/versioned/fake"
	qingcloudv1 "github.com/qingcloud/postgres-operator/pkg/generated/clientset/versioned/typed/qingcloud.com/v1"
)

type Clientset struct {
	*fakekubernetes.Clientset
	PGOClientset *fakeqingcloud.Clientset
}

var _ kubeapi.Interface = &Clientset{}

// QingcloudV1 retrieves the QingcloudV1Client
func (c *Clientset) QingcloudV1() qingcloudv1.QingcloudV1Interface {
	return c.PGOClientset.QingcloudV1()
}
