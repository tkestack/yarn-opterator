package nodemanagerset

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	appv1alpha1 "github.com/tkestack/yarn-opterator/pkg/apis/app/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_nodemanagerset")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new NodeManagerSet Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNodeManagerSet{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("nodemanagerset-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource NodeManagerSet
	err = c.Watch(&source.Kind{Type: &appv1alpha1.NodeManagerSet{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner NodeManagerSet
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appv1alpha1.NodeManagerSet{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileNodeManagerSet implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileNodeManagerSet{}

// ReconcileNodeManagerSet reconciles a NodeManagerSet object
type ReconcileNodeManagerSet struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a NodeManagerSet object and makes changes based on the state read
// and what is in the NodeManagerSet.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileNodeManagerSet) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling NodeManagerSet")

	// Fetch the NodeManagerSet instance
	instance := &appv1alpha1.NodeManagerSet{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	reqLogger.Info(fmt.Sprintf("%+v", instance))

	// Define a new Pod object
	pod := newPodForCR(instance)

	// Set NodeManagerSet instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *appv1alpha1.NodeManagerSet) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}

	podSpec := cr.Spec.Template.Spec.DeepCopy()
	envs := getConfigEnv(cr)

	for idx := range podSpec.Containers {
		podSpec.Containers[idx].Env = append(podSpec.Containers[idx].Env, envs...)
	}

	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: *podSpec,
	}
}

func getConfigEnv(cr *appv1alpha1.NodeManagerSet) []corev1.EnvVar {
	var envs = []corev1.EnvVar{}

	if cr.Spec.ClusterSource.MapReduceCluster != nil {
		envs = append(envs, corev1.EnvVar{Name: "ClusterId", Value: cr.Spec.ClusterSource.MapReduceCluster.ClusterId})
		envs = append(envs, corev1.EnvVar{Name: "Identifier", Value: strconv.Itoa(int(cr.Spec.ClusterSource.MapReduceCluster.Identifier))})

		valueMaps := GetFieldNameAndValue(cr.Spec.ClusterSource.MapReduceCluster.Config)
		for key := range valueMaps {
			envs = append(envs, corev1.EnvVar{Name: key, Value: valueMaps[key]})
		}

	} else {
		log.Info("bbbbbbbb")
	}

	return envs
}

//find struct field name and value
func GetFieldNameAndValue(s interface{}) map[string]string {

	var valueMaps = make(map[string]string)
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return valueMaps
	}

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		switch v.FieldByName(t.Field(i).Name).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueMaps[t.Field(i).Name] = fmt.Sprintf("%d", v.FieldByName(t.Field(i).Name).Int())
		case reflect.String:
			valueMaps[t.Field(i).Name] = v.FieldByName(t.Field(i).Name).String()
		default:
			log.Info("Unsupported type, name %s", t.Field(i).Name)
			continue
		}
	}

	return valueMaps
}
