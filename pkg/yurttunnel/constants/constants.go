/*
Copyright 2020 The OpenYurt Authors.

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

package constants

const (
	YurttunnelServerAgentPort          = 10262
	YurttunnelServerMasterPort         = 10263
	YurttunnelServerMasterInsecurePort = 10264
	YurttunnelServerServiceNs          = "kube-system"
	YurttunnelServerServiceName        = "x-tunnel-server-svc"
	YurttunnelServerAgentPortName      = "tcp"
	YurttunnelServerExternalAddrKey    = "x-tunnel-server-external-addr"
	YurttunnelEndpointsNs              = "kube-system"
	YurttunnelEndpointsName            = "x-tunnel-server-svc"

	// yurttunnel PKI related constants
	YurttunnelCSROrg                 = "openyurt:yurttunnel"
	YurttunnelAgentCSRCN             = "yurttunnel-agent"
	YurttunneServerCSROrg            = "system:masters"
	YurttunneServerCSRCN             = "kube-apiserver-kubelet-client"
	YurttunnelCAFile                 = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	YurttunnelTokenFile              = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	YurttunnelServerCertDir          = "/var/lib/yurttunnel-server/pki"
	YurttunnelAgentCertDir           = "/var/lib/yurttunnel-agent/pki"
	YurttunnelCSRApproverThreadiness = 2

	YurtEdgeNodeLabel          = "alibabacloud.com/is-edge-worker"
	YurttunnelEnableAgentLabel = "alibabacloud.com/enable-yurttunnel-agent"
	YurttunnelAgentPodIPEnv    = "POD_IP"
)
