/*
Copyright 2021 KubeCube Authors

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

package multicluster

import (
	"github.com/kubecube-io/kubecube/pkg/multicluster/fake"
	"github.com/kubecube-io/kubecube/pkg/multicluster/manager"
)

// Interface the way to be used outside for multi cluster manager
func Interface() manager.MultiClustersManager {
	if fake.IsFake {
		return fake.FakeMultiClusterMgr
	}
	return manager.MultiClusterMgr
}
