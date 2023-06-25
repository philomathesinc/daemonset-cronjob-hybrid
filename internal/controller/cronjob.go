package controller

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	demov1 "github.com/philomathesinc/daemonset-cronjob-hybrid/api/v1"
)

// Cronjob -> Job -> Pod
//
//	if correctTime {
//		Create(pod)
//	}
//
// 1. We are not reading anything from the spec.
// 2. Run a busybox container exiting with 0.
// 3. Running the pod every minute.
func (r *DaemonjobReconciler) CreateCronJob(ctx context.Context, req ctrl.Request, log logr.Logger, daemonjob *demov1.Daemonjob) (ctrl.Result, error) {
	if daemonjob.Status.LastRun != nil {
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
		},
		Spec: corev1.PodSpec{
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

	ctime := metav1.Now()
	daemonjob.Status.LastRun = &ctime

	if err := r.Status().Update(ctx, daemonjob); err != nil {
		log.Error(err, "unable to update status")
	}

	return ctrl.Result{}, nil
}
