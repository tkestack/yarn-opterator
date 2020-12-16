package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Represents the source of a yarn cluster.
// Only one of its members may be specified.
type ClusterSource struct {
	// Open Source MapReduce struct.
	// +optional
	MapReduceCluster *MapReduceClusterSource `json:"mapReduceCluster,omitempty" protobuf:"bytes,1,opt,name=mapReduceCluster"`

	// Tencent cloud Elastic MapReduce struct.
	// +optional
	TEMapReduceCluster *TEMapReduceClusterSource `json:"tEMapReduceCluster,omitempty" protobuf:"bytes,2,opt,name=tEMapReduceCluster"`
}

// Open Source MapReduce source.
type MapReduceClusterSource struct {
	// cluster unique id for cluster
	ClusterId string `json:"clusterId" protobuf:"bytes,1,opt,name=clusterId"`

	// identifier of emr flowId, be used to unique the yarn node
	Identifier int64 `json:"identifier" protobuf:"bytes,2,opt,name=identifier"`

	// map reduce config
	Config MapReduceConfig `json:"config" protobuf:"bytes,3,opt,name=config"`
}

// Open Source MapReduce config.
type MapReduceConfig struct {
	// resource manager active address
	RMActiveAddress string `json:"rmActiveAddress" protobuf:"bytes,1,opt,name=rmActiveAddress"`

	// resource manager standby address
	RMStandbyAddress string `json:"rmStandbyAddress" protobuf:"bytes,2,opt,name=rmStandbyAddress"`

	// resource manager active admin address
	RMActiveAdmin string `json:"rmActiveAdmin" protobuf:"bytes,3,opt,name=rmActiveAdmin"`

	// resource manager standby admin address
	RMStandbyAdmin string `json:"rmStandbyAdmin" protobuf:"bytes,4,opt,name=rmStandbyAdmin"`

	// resource manager active hostname
	RMActiveHostname string `json:"rmActiveHostname" protobuf:"bytes,5,opt,name=rmActiveHostname"`

	// resource manager standby hostname
	RMStandbyHostname string `json:"rmStandbyHostname" protobuf:"bytes,6,opt,name=rmStandbyHostname"`

	// resource manager active tracker
	RMActiveTracker string `json:"rmActiveTracker" protobuf:"bytes,7,opt,name=rmActiveTracker"`

	// resource manager standby tracker
	RMStandbyTracker string `json:"rmStandbyTracker" protobuf:"bytes,8,opt,name=rmStandbyTracker"`

	// resource manager active scheduler
	RMActiveScheduler string `json:"rmActiveScheduler" protobuf:"bytes,9,opt,name=rmActiveScheduler"`

	// resource manager standby scheduler
	RMStandbyScheduler string `json:"rmStandbyScheduler" protobuf:"bytes,10,opt,name=rmStandbyScheduler"`

	// resource manager active webapp
	RMActiveWebapp string `json:"rmActiveWebapp" protobuf:"bytes,11,opt,name=rmActiveWebapp"`

	// resource manager standby webapp
	RMStandbyWebapp string `json:"rmStandbyWebapp" protobuf:"bytes,12,opt,name=rmStandbyWebapp"`

	// resource manager zookeeper address, comma separated
	RMZookeeperAddress string `json:"rmZookeeperAddress" protobuf:"bytes,13,opt,name=rmZookeeperAddress"`

	// resource manager zookeeper path
	RMZookeeperPath string `json:"rmZookeeperPath" protobuf:"bytes,14,opt,name=rmZookeeperPath"`
}

// Tencent cloud Elastic MapReduce source.
type TEMapReduceClusterSource struct {
	// cluster unique id for emr cluster
	ClusterId string `json:"clusterId" protobuf:"bytes,1,opt,name=clusterId"`

	// identifier of emr flowId, be used to unique the yarn node
	Identifier int64 `json:"identifier" protobuf:"bytes,2,opt,name=identifier"`

	// generate submount path enable
	GenerateSubMountPath bool `json:"generatesubmountpath" protobuf:"bytes,3,opt,name=generatesubmountpath"`
}

// NodeManagerSetStatus defines the observed state of NodeManagerSet
type NodeManagerSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// NodeManagerSetSpec defines the desired state of NodeManagerSet
type NodeManagerSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Count         int                    `json:"count,omitempty" protobuf:"bytes,1,opt,name=count"`
	Excludes      []string               `json:"excludes,omitempty" protobuf:"bytes,2,opt,name=excludes"`
	Template      corev1.PodTemplateSpec `json:"template" protobuf:"bytes,3,opt,name=template"`
	ClusterSource ClusterSource          `json:"clusterSource" protobuf:"bytes,4,opt,name=clusterSource"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeManagerSet is the Schema for the nodemanagersets API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodemanagersets,scope=Namespaced
type NodeManagerSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeManagerSetSpec   `json:"spec,omitempty"`
	Status NodeManagerSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeManagerSetList contains a list of NodeManagerSet
type NodeManagerSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeManagerSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeManagerSet{}, &NodeManagerSetList{})
}
