package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "github.com/philomathesinc/daemonset-cronjob-hybrid/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// Call kubernetes API to get nodes
	var nodes corev1.NodeList
	if err := r.List(ctx, &nodes); err != nil {
		log.Error(err, "unable to fetch nodes")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	for _, n := range nodes.Items {
		log.Info(n.Name)
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.Name + "-" + n.Name,
				Namespace: req.Namespace,
			},
			Spec: corev1.PodSpec{
				NodeName: n.Name,
				Containers: []corev1.Container{
					{
						Name:    "default",
						Image:   "busybox",
						Command: []string{"sh", "-c"},
						Args: []string{
							"sleep 3600",
						},
					},
				},
			},
		}

		if err := ctrl.SetControllerReference(&daemonjob, pod, r.Scheme); err != nil {
			log.Error(err, "unable to set controller reference")
		}

		// call k8s api to create the pod
		if err := r.Create(ctx, pod); err != nil {
			log.Error(err, "unable to create pod")
		}
		log.Info("Pod created", "name", req.Name+"-"+n.Name, "namespace", req.Namespace)
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
