package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "github.com/philomathesinc/daemonset-cronjob-hybrid/api/v1"
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
//+kubebuilder:rbac:groups=core,resources=nodes,verbs=list;watch

func (r *DaemonjobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// A variable to hold the CR data
	var daemonjob demov1.Daemonjob

	// Call kubernetes API to get the CR data
	if err := r.Get(ctx, req.NamespacedName, &daemonjob); err != nil {
		log.Error(err, "unable to fetch daemonjob")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// r.CreateDaemonset(ctx, req, log, &daemonjob)
	r.CreateCronJob(ctx, req, log, &daemonjob)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DaemonjobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Daemonjob{}).
		Complete(r)
}
