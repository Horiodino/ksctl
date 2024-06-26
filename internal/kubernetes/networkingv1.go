package kubernetes

import (
	"context"

	ksctlErrors "github.com/ksctl/ksctl/pkg/helpers/errors"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Kubernetes) netPolicyApply(o *networkingv1.NetworkPolicy) error {

	ns := o.Namespace
	_, err := k.clientset.
		NetworkingV1().
		NetworkPolicies(ns).
		Create(context.Background(), o, v1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			_, err = k.clientset.
				NetworkingV1().
				NetworkPolicies(ns).
				Update(context.Background(), o, v1.UpdateOptions{})
			if err != nil {
				return ksctlErrors.ErrFailedKubernetesClient.Wrap(
					log.NewError(kubernetesCtx, "netpol apply failed", "Reason", err),
				)
			}
		} else {
			return ksctlErrors.ErrFailedKubernetesClient.Wrap(
				log.NewError(kubernetesCtx, "netpol apply failed", "Reason", err),
			)
		}
	}
	return nil
}

func (k *Kubernetes) netPolicyDelete(o *networkingv1.NetworkPolicy) error {

	ns := o.Namespace
	err := k.clientset.
		NetworkingV1().
		NetworkPolicies(ns).
		Delete(context.Background(), o.Name, v1.DeleteOptions{})
	if err != nil {
		return ksctlErrors.ErrFailedKubernetesClient.Wrap(
			log.NewError(kubernetesCtx, "netpol delete failed", "Reason", err),
		)
	}
	return nil
}
