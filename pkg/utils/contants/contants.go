// Copyright © 2021 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contants

import (
	"github.com/mitchellh/go-homedir"
)

const (
	LvsCareStaticPodName = "kube-sealyun-lvscare"
	StaticPodFileSuffix  = "yaml"
	DefaultLvsCareImage  = "sealyun.hub:5000/sealyun/lvscare:latest"
	DefaultCNI           = "calico"
)

//CRD kind
const (
	Config   = "Config"
	Cluster  = "Cluster"
	Resource = "Resource"
	Kubeadm  = "Kubeadm"
)

func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		return "/root"
	}
	return home
}
