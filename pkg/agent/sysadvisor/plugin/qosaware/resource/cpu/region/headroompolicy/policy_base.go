/*
Copyright 2022 The Katalyst Authors.

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

package headroompolicy

import (
	"github.com/kubewharf/katalyst-core/pkg/agent/sysadvisor/metacache"
	"github.com/kubewharf/katalyst-core/pkg/agent/sysadvisor/types"
	"github.com/kubewharf/katalyst-core/pkg/metaserver"
)

type PolicyBase struct {
	RegionName string
	PodSet     types.PodSet
	Total      int

	MetaCache  *metacache.MetaCacheImp
	MetaServer *metaserver.MetaServer
}

func NewPolicyBase(regionName string, metaCache *metacache.MetaCacheImp, metaServer *metaserver.MetaServer) *PolicyBase {
	cp := &PolicyBase{
		RegionName: regionName,
		PodSet:     make(types.PodSet),

		MetaCache:  metaCache,
		MetaServer: metaServer,
	}
	return cp
}

func (p *PolicyBase) SetPodSet(podSet types.PodSet) {
	p.PodSet = podSet.Clone()
}

func (p *PolicyBase) SetEssentials(total int) {
	p.Total = total
}
