/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_AddonSpec = map[string]string{
	"": "AddonSpec describes the attributes on a Addon.",
}

func (AddonSpec) SwaggerDoc() map[string]string {
	return map_AddonSpec
}

var map_CSIOperator = map[string]string{
	"":     "CSIOperator is a operator to manages CSI external components.",
	"spec": "Spec defines the desired identities of storage operator.",
}

func (CSIOperator) SwaggerDoc() map[string]string {
	return map_CSIOperator
}

var map_CSIOperatorList = map[string]string{
	"":      "CSIOperatorList is the whole list of all storage operators which owned by a tenant.",
	"items": "List of storage operators.",
}

func (CSIOperatorList) SwaggerDoc() map[string]string {
	return map_CSIOperatorList
}

var map_CSIOperatorSpec = map[string]string{
	"":        "CSIOperatorSpec describes the attributes of a storage operator.",
	"version": "Version of the CSI operator.",
}

func (CSIOperatorSpec) SwaggerDoc() map[string]string {
	return map_CSIOperatorSpec
}

var map_CSIOperatorStatus = map[string]string{
	"":                            "CSIOperatorStatus is information about the current status of a storage operator.",
	"storageVendorVersion":        "StorageVendorVersion will be set to the config version of the storage vendor.",
	"phase":                       "Phase is the current lifecycle phase of the tapp controller of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (CSIOperatorStatus) SwaggerDoc() map[string]string {
	return map_CSIOperatorStatus
}

var map_CSIProxyOptions = map[string]string{
	"": "CSIProxyOptions is the query options to a kube-apiserver proxy call for CSI crd object.",
}

func (CSIProxyOptions) SwaggerDoc() map[string]string {
	return map_CSIProxyOptions
}

var map_Cluster = map[string]string{
	"":     "Cluster is a Kubernetes cluster in",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (Cluster) SwaggerDoc() map[string]string {
	return map_Cluster
}

var map_ClusterAddon = map[string]string{
	"":     "ClusterAddon contains the Addon component for the current kubernetes cluster",
	"spec": "Spec defines the desired identities of addons in this set.",
}

func (ClusterAddon) SwaggerDoc() map[string]string {
	return map_ClusterAddon
}

var map_ClusterAddonList = map[string]string{
	"":      "ClusterAddonList is the whole list of all ClusterAddon.",
	"items": "List of ClusterAddon",
}

func (ClusterAddonList) SwaggerDoc() map[string]string {
	return map_ClusterAddonList
}

var map_ClusterAddonSpec = map[string]string{
	"":        "ClusterAddonSpec indicates the specifications of the ClusterAddon.",
	"type":    "Addon type, one of PersistentEvent or LogCollector etc.",
	"level":   "AddonLevel is level of cluster addon.",
	"version": "Version",
}

func (ClusterAddonSpec) SwaggerDoc() map[string]string {
	return map_ClusterAddonSpec
}

var map_ClusterAddonStatus = map[string]string{
	"":       "ClusterAddonStatus is information about the current status of a ClusterAddon.",
	"phase":  "Phase is the current lifecycle phase of the addon of cluster.",
	"reason": "Reason is a brief CamelCase string that describes any failure.",
}

func (ClusterAddonStatus) SwaggerDoc() map[string]string {
	return map_ClusterAddonStatus
}

var map_ClusterAddonType = map[string]string{
	"":              "ClusterAddonType records the all addons of cluster available.",
	"type":          "Addon type, one of Helm, PersistentEvent or LogCollector etc.",
	"level":         "AddonLevel is level of cluster addon.",
	"latestVersion": "LatestVersion is latest version of the addon.",
	"description":   "Description is desc of the addon.",
}

func (ClusterAddonType) SwaggerDoc() map[string]string {
	return map_ClusterAddonType
}

var map_ClusterAddonTypeList = map[string]string{
	"": "ClusterAddonTypeList is a resource containing a list of ClusterAddonType objects.",
}

func (ClusterAddonTypeList) SwaggerDoc() map[string]string {
	return map_ClusterAddonTypeList
}

var map_ClusterAddress = map[string]string{
	"":     "ClusterAddress contains information for the cluster's address.",
	"type": "Cluster address type, one of Public, ExternalIP or InternalIP.",
	"host": "The cluster address.",
}

func (ClusterAddress) SwaggerDoc() map[string]string {
	return map_ClusterAddress
}

var map_ClusterApplyOptions = map[string]string{
	"": "ClusterApplyOptions is the query options to a kube-apiserver proxy call for cluster object.",
}

func (ClusterApplyOptions) SwaggerDoc() map[string]string {
	return map_ClusterApplyOptions
}

var map_ClusterComponent = map[string]string{
	"": "ClusterComponent records the number of copies of each component of the cluster master.",
}

func (ClusterComponent) SwaggerDoc() map[string]string {
	return map_ClusterComponent
}

var map_ClusterComponentReplicas = map[string]string{
	"": "ClusterComponentReplicas records the number of copies of each state of each component of the cluster master.",
}

func (ClusterComponentReplicas) SwaggerDoc() map[string]string {
	return map_ClusterComponentReplicas
}

var map_ClusterCondition = map[string]string{
	"":                   "ClusterCondition contains details for the current condition of this cluster.",
	"type":               "Type is the type of the condition.",
	"status":             "Status is the status of the condition. Can be True, False, Unknown.",
	"lastProbeTime":      "Last time we probed the condition.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "Unique, one-word, CamelCase reason for the condition's last transition.",
	"message":            "Human-readable message indicating details about last transition.",
}

func (ClusterCondition) SwaggerDoc() map[string]string {
	return map_ClusterCondition
}

var map_ClusterCredential = map[string]string{
	"":                  "ClusterCredential records the credential information needed to access the cluster.",
	"etcdCACert":        "For TKE in global reuse",
	"caCert":            "For connect the cluster",
	"clientCert":        "For kube-apiserver X509 auth",
	"clientKey":         "For kube-apiserver X509 auth",
	"token":             "For kube-apiserver token auth",
	"bootstrapToken":    "For kubeadm init or join",
	"certificateKey":    "For kubeadm init or join",
	"username":          "Username is the username for basic authentication to the kubernetes cluster.",
	"as":                "Impersonate is the username to act-as.",
	"as-groups":         "ImpersonateGroups is the groups to imperonate.",
	"as-user-extra":     "ImpersonateUserExtra contains additional information for impersonated user.",
	"serverCrt":         "For kube-apiserver server crt",
	"serverKey":         "For kube-apiserver server key",
	"serviceAccountKey": "For kube-apiserver issue ServiceAccount",
	"kubeletPasswd":     "For kubelet token auth",
	"kubeProxyPasswd":   "For kubeProxy token auth",
}

func (ClusterCredential) SwaggerDoc() map[string]string {
	return map_ClusterCredential
}

var map_ClusterCredentialList = map[string]string{
	"":      "ClusterCredentialList is the whole list of all ClusterCredential which owned by a tenant.",
	"items": "List of clusters",
}

func (ClusterCredentialList) SwaggerDoc() map[string]string {
	return map_ClusterCredentialList
}

var map_ClusterFeature = map[string]string{
	"":                 "ClusterFeature records the features that are enabled by the cluster.",
	"authzWebhookAddr": "For kube-apiserver authorization webhook",
	"upgrade":          "Upgrade control upgrade process.",
}

func (ClusterFeature) SwaggerDoc() map[string]string {
	return map_ClusterFeature
}

var map_ClusterGroupAPIResourceItem = map[string]string{
	"":             "ClusterGroupAPIResourceItem specifies the name of a resource and whether it is namespaced.",
	"name":         "name is the plural name of the resource.",
	"singularName": "singularName is the singular name of the resource.  This allows clients to handle plural and singular opaquely. The singularName is more correct for reporting status on a single item and both singular and plural are allowed from the kubectl CLI interface.",
	"namespaced":   "namespaced indicates if a resource is namespaced or not.",
	"group":        "group is the preferred group of the resource.  Empty implies the group of the containing resource list. For subresources, this may have a different value, for example: Scale\".",
	"version":      "version is the preferred version of the resource.  Empty implies the version of the containing resource list For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)\".",
	"kind":         "kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')",
	"verbs":        "verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)",
	"shortNames":   "shortNames is a list of suggested short names of the resource.",
	"categories":   "categories is a list of the grouped resources this resource belongs to (e.g. 'all')",
}

func (ClusterGroupAPIResourceItem) SwaggerDoc() map[string]string {
	return map_ClusterGroupAPIResourceItem
}

var map_ClusterGroupAPIResourceItems = map[string]string{
	"":             "ClusterGroupAPIResourceItems contains the GKV for the current kubernetes cluster",
	"groupVersion": "groupVersion is the group and version this APIResourceList is for.",
	"resources":    "resources contains the name of the resources and if they are namespaced.",
}

func (ClusterGroupAPIResourceItems) SwaggerDoc() map[string]string {
	return map_ClusterGroupAPIResourceItems
}

var map_ClusterGroupAPIResourceItemsList = map[string]string{
	"":                 "ClusterGroupAPIResourceItemsList is the whole list of all ClusterAPIResource.",
	"items":            "List of ClusterGroupAPIResourceItems",
	"failedGroupError": "Failed Group Error",
}

func (ClusterGroupAPIResourceItemsList) SwaggerDoc() map[string]string {
	return map_ClusterGroupAPIResourceItemsList
}

var map_ClusterGroupAPIResourceOptions = map[string]string{
	"": "ClusterGroupAPIResourceOptions is the query options.",
}

func (ClusterGroupAPIResourceOptions) SwaggerDoc() map[string]string {
	return map_ClusterGroupAPIResourceOptions
}

var map_ClusterList = map[string]string{
	"":      "ClusterList is the whole list of all clusters which owned by a tenant.",
	"items": "List of clusters",
}

func (ClusterList) SwaggerDoc() map[string]string {
	return map_ClusterList
}

var map_ClusterMachine = map[string]string{
	"":       "ClusterMachine is the master machine definition of cluster.",
	"taints": "If specified, the node's taints.",
}

func (ClusterMachine) SwaggerDoc() map[string]string {
	return map_ClusterMachine
}

var map_ClusterMachineProxy = map[string]string{
	"": "ClusterMachine is the proxy definition of ClusterMachine.",
}

func (ClusterMachineProxy) SwaggerDoc() map[string]string {
	return map_ClusterMachineProxy
}

var map_ClusterProperty = map[string]string{
	"": "ClusterProperty records the attribute information of the cluster.",
}

func (ClusterProperty) SwaggerDoc() map[string]string {
	return map_ClusterProperty
}

var map_ClusterResource = map[string]string{
	"":            "ClusterResource records the current available and maximum resource quota information for the cluster.",
	"capacity":    "Capacity represents the total resources of a cluster.",
	"allocatable": "Allocatable represents the resources of a cluster that are available for scheduling. Defaults to Capacity.",
}

func (ClusterResource) SwaggerDoc() map[string]string {
	return map_ClusterResource
}

var map_ClusterSpec = map[string]string{
	"":                     "ClusterSpec is a description of a cluster.",
	"finalizers":           "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
	"serviceCIDR":          "ServiceCIDR is used to set a separated CIDR for k8s service, it's exclusive with MaxClusterServiceNum.",
	"dnsDomain":            "DNSDomain is the dns domain used by k8s services. Defaults to \"cluster.local\".",
	"clusterCredentialRef": "ClusterCredentialRef for isolate sensitive information. If not specified, cluster controller will create one; If specified, provider must make sure is valid.",
	"etcd":                 "Etcd holds configuration for etcd.",
	"hostnameAsNodename":   "If true will use hostname as nodename, if false will use machine IP as nodename.",
	"bootstrapApps":        "BootstrapApps will install apps during creating cluster",
	"appVersion":           "AppVersion is the overall version of system components",
	"clusterLevel":         "ClusterLevel is the expect level of cluster",
}

func (ClusterSpec) SwaggerDoc() map[string]string {
	return map_ClusterSpec
}

var map_ClusterStatus = map[string]string{
	"":               "ClusterStatus represents information about the status of a cluster.",
	"message":        "A human readable message indicating details about why the cluster is in this condition.",
	"reason":         "A brief CamelCase message indicating details about why the cluster is in this state.",
	"addresses":      "List of addresses reachable to the cluster.",
	"appVersion":     "AppVersion is the overall version of system components",
	"componentPhase": "ComponentPhase is the status of components, contains \"deployed\", \"pending-upgrade\", \"failed\" status",
	"clusterLevel":   "ClusterLevel is the real level of cluster",
}

func (ClusterStatus) SwaggerDoc() map[string]string {
	return map_ClusterStatus
}

var map_ConfigMap = map[string]string{
	"":           "ConfigMap holds configuration data for tke to consume.",
	"data":       "Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.",
	"binaryData": "BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process.",
}

func (ConfigMap) SwaggerDoc() map[string]string {
	return map_ConfigMap
}

var map_ConfigMapList = map[string]string{
	"":      "ConfigMapList is a resource containing a list of ConfigMap objects.",
	"items": "Items is the list of ConfigMaps.",
}

func (ConfigMapList) SwaggerDoc() map[string]string {
	return map_ConfigMapList
}

var map_CronHPA = map[string]string{
	"":     "CronHPA is a new kubernetes workload.",
	"spec": "Spec defines the desired identities of CronHPA.",
}

func (CronHPA) SwaggerDoc() map[string]string {
	return map_CronHPA
}

var map_CronHPAList = map[string]string{
	"":      "CronHPAList is the whole list of all CronHPAs which owned by a tenant.",
	"items": "List of CronHPAs",
}

func (CronHPAList) SwaggerDoc() map[string]string {
	return map_CronHPAList
}

var map_CronHPAProxyOptions = map[string]string{
	"": "CronHPAProxyOptions is the query options to a kube-apiserver proxy call.",
}

func (CronHPAProxyOptions) SwaggerDoc() map[string]string {
	return map_CronHPAProxyOptions
}

var map_CronHPASpec = map[string]string{
	"": "CronHPASpec describes the attributes on a CronHPA.",
}

func (CronHPASpec) SwaggerDoc() map[string]string {
	return map_CronHPASpec
}

var map_CronHPAStatus = map[string]string{
	"":                            "CronHPAStatus is information about the current status of a CronHPA.",
	"phase":                       "Phase is the current lifecycle phase of the CronHPA of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (CronHPAStatus) SwaggerDoc() map[string]string {
	return map_CronHPAStatus
}

var map_Etcd = map[string]string{
	"":         "Etcd contains elements describing Etcd configuration.",
	"local":    "Local provides configuration knobs for configuring the local etcd instance Local and External are mutually exclusive",
	"external": "External describes how to connect to an external etcd cluster Local and External are mutually exclusive",
}

func (Etcd) SwaggerDoc() map[string]string {
	return map_Etcd
}

var map_ExternalEtcd = map[string]string{
	"":          "ExternalEtcd describes an external etcd cluster. Kubeadm has no knowledge of where certificate files live and they must be supplied.",
	"endpoints": "Endpoints of etcd members. Required for ExternalEtcd.",
	"caFile":    "CAFile is an SSL Certificate Authority file used to secure etcd communication. Required if using a TLS connection.",
	"certFile":  "CertFile is an SSL certification file used to secure etcd communication. Required if using a TLS connection.",
	"keyFile":   "KeyFile is an SSL key file used to secure etcd communication. Required if using a TLS connection.",
}

func (ExternalEtcd) SwaggerDoc() map[string]string {
	return map_ExternalEtcd
}

var map_LocalEtcd = map[string]string{
	"":               "LocalEtcd describes that kubeadm should run an etcd cluster locally",
	"dataDir":        "DataDir is the directory etcd will place its data. Defaults to \"/var/lib/etcd\".",
	"extraArgs":      "ExtraArgs are extra arguments provided to the etcd binary when run inside a static pod.",
	"serverCertSANs": "ServerCertSANs sets extra Subject Alternative Names for the etcd server signing cert.",
	"peerCertSANs":   "PeerCertSANs sets extra Subject Alternative Names for the etcd peer signing cert.",
}

func (LocalEtcd) SwaggerDoc() map[string]string {
	return map_LocalEtcd
}

var map_Machine = map[string]string{
	"":     "Machine instance in Kubernetes cluster",
	"spec": "Spec defines the desired identities of the Machine.",
}

func (Machine) SwaggerDoc() map[string]string {
	return map_Machine
}

var map_MachineAddress = map[string]string{
	"":        "MachineAddress contains information for the machine's address.",
	"type":    "Machine address type, one of Public, ExternalIP or InternalIP.",
	"address": "The machine address.",
}

func (MachineAddress) SwaggerDoc() map[string]string {
	return map_MachineAddress
}

var map_MachineCondition = map[string]string{
	"":                   "MachineCondition contains details for the current condition of this Machine.",
	"type":               "Type is the type of the condition.",
	"status":             "Status is the status of the condition. Can be True, False, Unknown.",
	"lastProbeTime":      "Last time we probed the condition.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "Unique, one-word, CamelCase reason for the condition's last transition.",
	"message":            "Human-readable message indicating details about last transition.",
}

func (MachineCondition) SwaggerDoc() map[string]string {
	return map_MachineCondition
}

var map_MachineList = map[string]string{
	"":      "MachineList is the whole list of all machine in an cluster.",
	"items": "List of clusters",
}

func (MachineList) SwaggerDoc() map[string]string {
	return map_MachineList
}

var map_MachineSpec = map[string]string{
	"":           "MachineSpec is a description of machine.",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
	"taints":     "If specified, the node's taints.",
}

func (MachineSpec) SwaggerDoc() map[string]string {
	return map_MachineSpec
}

var map_MachineStatus = map[string]string{
	"":            "MachineStatus represents information about the status of an machine.",
	"message":     "A human readable message indicating details about why the machine is in this condition.",
	"reason":      "A brief CamelCase message indicating details about why the machine is in this state.",
	"addresses":   "List of addresses reachable to the machine.",
	"machineInfo": "Set of ids/uuids to uniquely identify the node.",
}

func (MachineStatus) SwaggerDoc() map[string]string {
	return map_MachineStatus
}

var map_MachineSystemInfo = map[string]string{
	"":                        "MachineSystemInfo is a set of ids/uuids to uniquely identify the node.",
	"machineID":               "MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html",
	"systemUUID":              "SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html",
	"bootID":                  "Boot ID reported by the node.",
	"kernelVersion":           "Kernel Version reported by the node.",
	"osImage":                 "OS Image reported by the node.",
	"containerRuntimeVersion": "ContainerRuntime Version reported by the node.",
	"kubeletVersion":          "Kubelet Version reported by the node.",
	"kubeProxyVersion":        "KubeProxy Version reported by the node.",
	"operatingSystem":         "The Operating System reported by the node",
	"architecture":            "The Architecture reported by the node",
}

func (MachineSystemInfo) SwaggerDoc() map[string]string {
	return map_MachineSystemInfo
}

var map_PersistentBackEnd = map[string]string{
	"": "PersistentBackEnd indicates the backend type and attributes of the persistent log store.",
}

func (PersistentBackEnd) SwaggerDoc() map[string]string {
	return map_PersistentBackEnd
}

var map_PersistentEvent = map[string]string{
	"":     "PersistentEvent is a recorder of kubernetes event.",
	"spec": "Spec defines the desired identities of clusters in this set.",
}

func (PersistentEvent) SwaggerDoc() map[string]string {
	return map_PersistentEvent
}

var map_PersistentEventList = map[string]string{
	"":      "PersistentEventList is the whole list of all clusters which owned by a tenant.",
	"items": "List of PersistentEvents",
}

func (PersistentEventList) SwaggerDoc() map[string]string {
	return map_PersistentEventList
}

var map_PersistentEventSpec = map[string]string{
	"": "PersistentEventSpec describes the attributes on a PersistentEvent.",
}

func (PersistentEventSpec) SwaggerDoc() map[string]string {
	return map_PersistentEventSpec
}

var map_PersistentEventStatus = map[string]string{
	"":                            "PersistentEventStatus is information about the current status of a PersistentEvent.",
	"phase":                       "Phase is the current lifecycle phase of the persistent event of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (PersistentEventStatus) SwaggerDoc() map[string]string {
	return map_PersistentEventStatus
}

var map_ProxyOptions = map[string]string{
	"":     "ProxyOptions is the query options to a proxy call.",
	"path": "Path is the URL path to use for the current proxy request.",
}

func (ProxyOptions) SwaggerDoc() map[string]string {
	return map_ProxyOptions
}

var map_Registry = map[string]string{
	"": "Registry records the third-party image repository information stored by the user.",
}

func (Registry) SwaggerDoc() map[string]string {
	return map_Registry
}

var map_RegistryList = map[string]string{
	"": "RegistryList is a resource containing a list of Registry objects.",
}

func (RegistryList) SwaggerDoc() map[string]string {
	return map_RegistryList
}

var map_RegistrySpec = map[string]string{
	"": "RegistrySpec indicates the specifications of the third-party image repository.",
}

func (RegistrySpec) SwaggerDoc() map[string]string {
	return map_RegistrySpec
}

var map_ResourceRequirements = map[string]string{
	"": "ResourceRequirements describes the compute resource requirements.",
}

func (ResourceRequirements) SwaggerDoc() map[string]string {
	return map_ResourceRequirements
}

var map_StorageBackEndCLS = map[string]string{
	"": "StorageBackEndCLS records the attributes required when the backend storage type is CLS.",
}

func (StorageBackEndCLS) SwaggerDoc() map[string]string {
	return map_StorageBackEndCLS
}

var map_StorageBackEndES = map[string]string{
	"": "StorageBackEndES records the attributes required when the backend storage type is ElasticSearch.",
}

func (StorageBackEndES) SwaggerDoc() map[string]string {
	return map_StorageBackEndES
}

var map_TappController = map[string]string{
	"":     "TappController is a new kubernetes workload.",
	"spec": "Spec defines the desired identities of tapp controller.",
}

func (TappController) SwaggerDoc() map[string]string {
	return map_TappController
}

var map_TappControllerList = map[string]string{
	"":      "TappControllerList is the whole list of all tapp controllers which owned by a tenant.",
	"items": "List of tapp controllers",
}

func (TappControllerList) SwaggerDoc() map[string]string {
	return map_TappControllerList
}

var map_TappControllerProxyOptions = map[string]string{
	"": "TappControllerProxyOptions is the query options to a kube-apiserver proxy call.",
}

func (TappControllerProxyOptions) SwaggerDoc() map[string]string {
	return map_TappControllerProxyOptions
}

var map_TappControllerSpec = map[string]string{
	"": "TappControllerSpec describes the attributes on a tapp controller.",
}

func (TappControllerSpec) SwaggerDoc() map[string]string {
	return map_TappControllerSpec
}

var map_TappControllerStatus = map[string]string{
	"":                            "TappControllerStatus is information about the current status of a tapp controller.",
	"phase":                       "Phase is the current lifecycle phase of the tapp controller of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (TappControllerStatus) SwaggerDoc() map[string]string {
	return map_TappControllerStatus
}

var map_Upgrade = map[string]string{
	"mode":     "Upgrade mode, default value is Auto.",
	"strategy": "Upgrade strategy config.",
}

func (Upgrade) SwaggerDoc() map[string]string {
	return map_Upgrade
}

var map_UpgradeStrategy = map[string]string{
	"":                       "UpgradeStrategy used to control the upgrade process.",
	"maxUnready":             "The maximum number of pods that can be unready during the upgrade. 0% means all pods need to be ready after evition. 100% means ignore any pods unready which may be used in one worker node, use this carefully! default value is 0%.",
	"drainNodeBeforeUpgrade": "Whether drain node before upgrade. Draining node before upgrade is recommended. But not all pod running as cows, a few running as pets. If your pod can not accept be expelled from current node, this value should be false.",
}

func (UpgradeStrategy) SwaggerDoc() map[string]string {
	return map_UpgradeStrategy
}

// AUTO-GENERATED FUNCTIONS END HERE
