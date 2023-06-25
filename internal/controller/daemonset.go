package controller

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	demov1 "github.com/philomathesinc/daemonset-cronjob-hybrid/api/v1"
)

// Daemonset -> Pod
// nodes = GetNodes()
//
//	for _, node := range nodes {
//			Create(pod)
//	}
func (r *DaemonjobReconciler) CreateDaemonset(ctx context.Context, req ctrl.Request, log logr.Logger, daemonjob *demov1.Daemonjob) (ctrl.Result, error) {
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

		if err := ctrl.SetControllerReference(daemonjob, pod, r.Scheme); err != nil {
			log.Error(err, "unable to set controller reference")
		}

		// call k8s api to create the pod
		if err := r.Create(ctx, pod); err != nil {
			log.Error(err, "unable to create pod")
		}
		log.Info("Pod created", "name", req.Name+"-"+n.Name, "namespace", req.Namespace)
	}

	return ctrl.Result{}, nil
}
