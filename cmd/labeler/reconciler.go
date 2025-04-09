package main

import (
	"context"
	"encoding/json"
	alpha1 "github.com/pramodbindal/my-custom-labeler/pkg/apis/pramodbindal/v1alpha1"
	"github.com/pramodbindal/my-custom-labeler/pkg/client/informers/externalversions/pramodbindal/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	v1 "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"

	pipelinerunInformer "github.com/tektoncd/pipeline/pkg/client/injection/informers/pipeline/v1/pipelinerun"
)

type Reconciler struct {
	labelerInformer    v1alpha1.LabelerInformer
	deploymentInformer v1.DeploymentInformer
	kubeclient         kubernetes.Interface
}

func (r Reconciler) ReconcileKind(ctx context.Context, labeler *alpha1.Labeler) reconciler.Event {
	logger := logging.FromContext(ctx)
	logger.Infof("Reconciling Labeler : %s", labeler.Name)

	// Get List of Deployments
	deployments, err := r.deploymentInformer.Lister().Deployments("mynamespace").List(labels.Everything())
	if err != nil {
		logger.Errorf("Error listing deployments: %v", err)
	}

	pipelinerunInformer.Get(ctx).Lister().List(labels.Everything())

	patchData := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": labeler.Spec.CustomLabels,
		},
	}

	patchJson, err := json.Marshal(patchData)
	logger.Infof("Patch Json : %s", patchJson)
	if err != nil {
		logger.Errorf("Failed to marshal patch data: %v", err)
	}

	// Apply labels
	for _, deployment := range deployments {
		_, err := r.kubeclient.AppsV1().Deployments(deployment.Namespace).Patch(ctx, deployment.Name, types.MergePatchType, patchJson, metav1.PatchOptions{})
		if err != nil {
			logger.Errorf("Failed to patch deployment: %v", err)
		}
	}
	return nil
}
