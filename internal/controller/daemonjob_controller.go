package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "github.com/philomathesinc/daemonset-cronjob-hybrid/api/v1"
	corev1 "k8s.io/api/core/v1"
)

// DaemonjobReconciler reconciles a Daemonjob object
type DaemonjobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.example.com,resources=daemonjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.example.com,resources=daemonjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.example.com,resources=daemonjobs/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=pods,verbs=create;list;get;watch
//+kubebuilder:rbac:groups=core,resources=node,verbs=list

func (r *DaemonjobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Call kubernetes API to get nodes
	var nodes corev1.NodeList
	if err := r.List(ctx, &nodes); err != nil {
		log.Error(err, "unable to fetch nodes")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	for _, n := range nodes.Items {
		log.Info(n.Name)
	}

	// Daemonset -> Pod
	// nodes = GetNodes()
	// for _, node := range nodes {
	// 		Create(pod)
	// }
	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DaemonjobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Daemonjob{}).
		Complete(r)
}
