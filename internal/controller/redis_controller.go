/*
Copyright 2023.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	qwoptcontrollerv1beta1 "k8s.io/redis/api/v1beta1"
)

// RedisReconciler reconciles a Redis object
type RedisReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=qwoptcontroller.k8s.io,resources=redis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=qwoptcontroller.k8s.io,resources=redis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=qwoptcontroller.k8s.io,resources=redis/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Redis object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *RedisReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	log := log.FromContext(ctx).WithValues("redis", req.NamespacedName)
	log.Info("reconciling redis")

	// 不重新入队列
	forget := ctrl.Result{
		Requeue:      false,
		RequeueAfter: 0,
	}

	// redis 实例
	rds := &qwoptcontrollerv1beta1.Redis{}
	if err := r.Client.Get(ctx, req.NamespacedName, rds); err != nil {
		if errors.IsNotFound(err) {
			log.Error(err, "unable to get redis")
			return forget, nil
		}
		return forget, err
	}

	log.Info("Add/Update/Delete for redis", "name", rds.GetName())

	return forget, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RedisReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&qwoptcontrollerv1beta1.Redis{}).
		Complete(r)
}
