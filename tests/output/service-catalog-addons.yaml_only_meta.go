// Code generated by reverse-kube-resource. DO NOT EDIT.

package examples

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	// Job "service-catalog-addons-service-binding-usage-controller-cleanup"
	serviceCatalogAddonsServiceBindingUsageControllerCleanupJob = batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-binding-usage-controller-cleanup",
			Labels: map[string]string{
				"app":      "service-catalog-addons-service-binding-usage-controller",
				"chart":    "service-binding-usage-controller-0.1.0",
				"heritage": "Helm",
				"release":  "service-catalog-addons",
			},
			Annotations: map[string]string{
				"helm.sh/hook":               "pre-delete",
				"helm.sh/hook-delete-policy": "before-hook-creation, hook-succeeded",
				"helm.sh/hook-weight":        "1",
			},
		},
	}

	// PodSecurityPolicy "service-catalog-addons-service-catalog-ui"
	serviceCatalogAddonsServiceCatalogUiPodSecurityPolicy = policyv1beta1.PodSecurityPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-catalog-ui",
			Labels: map[string]string{
				"app":                       "service-catalog-ui",
				"chart":                     "service-catalog-ui-0.1.0",
				"heritage":                  "Helm",
				"kyma-project.io/component": "frontend",
				"release":                   "service-catalog-addons",
			},
		},
	}

	// ServiceAccount "service-catalog-addons-service-binding-usage-controller"
	serviceCatalogAddonsServiceBindingUsageControllerServiceAccount = corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service-catalog-addons-service-binding-usage-controller",
			Namespace: "kyma-system",
			Labels: map[string]string{
				"app":      "service-catalog-addons-service-binding-usage-controller",
				"chart":    "service-binding-usage-controller-0.1.0",
				"heritage": "Helm",
				"release":  "service-catalog-addons",
			},
		},
	}

	// ServiceAccount "service-catalog-addons-service-catalog-ui"
	serviceCatalogAddonsServiceCatalogUiServiceAccount = corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service-catalog-addons-service-catalog-ui",
			Namespace: "kyma-system",
			Labels: map[string]string{
				"app": "service-catalog-addons-service-catalog-ui",
			},
		},
	}

	// ConfigMap "service-binding-usage-controller-process-sbu-spec"
	serviceBindingUsageControllerProcessSbuSpecConfigMap = corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-binding-usage-controller-process-sbu-spec",
			Labels: map[string]string{
				"app":   "service-catalog-addons-service-binding-usage-controller",
				"chart": "service-binding-usage-controller-0.1.0",
			},
		},
	}

	// ConfigMap "service-binding-usage-controller-dashboard"
	serviceBindingUsageControllerDashboardConfigMap = corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-binding-usage-controller-dashboard",
			Labels: map[string]string{
				"app":               "monitoring-grafana",
				"grafana_dashboard": "1",
			},
		},
	}

	// ConfigMap "service-catalog-ui"
	serviceCatalogUiConfigMap = corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-ui",
			Labels: map[string]string{
				"app":   "service-catalog-ui",
				"chart": "service-catalog-ui-0.1.0",
			},
		},
	}

	// ClusterRole "service-catalog-addons-service-binding-usage-controller"
	serviceCatalogAddonsServiceBindingUsageControllerClusterRole = rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-binding-usage-controller",
			Labels: map[string]string{
				"app":      "service-catalog-addons-service-binding-usage-controller",
				"chart":    "service-binding-usage-controller-0.1.0",
				"heritage": "Helm",
				"release":  "service-catalog-addons",
			},
		},
	}

	// ClusterRoleBinding "service-catalog-addons-service-binding-usage-controller"
	serviceCatalogAddonsServiceBindingUsageControllerClusterRoleBinding = rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-binding-usage-controller",
			Labels: map[string]string{
				"app":      "service-catalog-addons-service-binding-usage-controller",
				"chart":    "service-binding-usage-controller-0.1.0",
				"heritage": "Helm",
				"release":  "service-catalog-addons",
			},
		},
	}

	// Role "service-catalog-addons-service-catalog-ui"
	serviceCatalogAddonsServiceCatalogUiRole = rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service-catalog-addons-service-catalog-ui",
			Namespace: "kyma-system",
		},
	}

	// RoleBinding "service-catalog-addons-service-catalog-ui"
	serviceCatalogAddonsServiceCatalogUiRoleBinding = rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service-catalog-addons-service-catalog-ui",
			Namespace: "kyma-system",
		},
	}

	// Service "service-catalog-addons-service-binding-usage-controller"
	serviceCatalogAddonsServiceBindingUsageControllerService = corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-binding-usage-controller",
			Labels: map[string]string{
				"app":      "service-catalog-addons-service-binding-usage-controller",
				"chart":    "service-binding-usage-controller-0.1.0",
				"heritage": "Helm",
				"release":  "service-catalog-addons",
			},
		},
	}

	// Service "service-catalog-addons-service-catalog-ui"
	serviceCatalogAddonsServiceCatalogUiService = corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-catalog-ui",
			Labels: map[string]string{
				"app":      "service-catalog-ui",
				"chart":    "service-catalog-ui-0.1.0",
				"heritage": "Helm",
				"release":  "service-catalog-addons",
			},
		},
	}

	// Deployment "service-catalog-addons-service-binding-usage-controller"
	serviceCatalogAddonsServiceBindingUsageControllerDeployment = appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-binding-usage-controller",
			Labels: map[string]string{
				"app":                       "service-catalog-addons-service-binding-usage-controller",
				"chart":                     "service-binding-usage-controller-0.1.0",
				"heritage":                  "Helm",
				"kyma-project.io/component": "controller",
				"release":                   "service-catalog-addons",
			},
		},
	}

	// Deployment "service-catalog-addons-service-catalog-ui"
	serviceCatalogAddonsServiceCatalogUiDeployment = appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-catalog-addons-service-catalog-ui",
			Labels: map[string]string{
				"app":                       "service-catalog-ui",
				"chart":                     "service-catalog-ui-0.1.0",
				"heritage":                  "Helm",
				"kyma-project.io/component": "frontend",
				"release":                   "service-catalog-addons",
			},
		},
	}

	// Unstructured "service-catalog-addons-service-catalog-ui"
	serviceCatalogAddonsServiceCatalogUiUnstructuredDestinationRule = v1unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "networking.istio.io/v1alpha3",
			"kind":       "DestinationRule",
			"metadata": map[string]any{
				"name": "service-catalog-addons-service-catalog-ui",
			},
		},
	}

	// Unstructured "service-catalog-addons-service-binding-usage-controller"
	serviceCatalogAddonsServiceBindingUsageControllerUnstructuredPeerAuthentication = v1unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "security.istio.io/v1beta1",
			"kind":       "PeerAuthentication",
			"metadata": map[string]any{
				"name": "service-catalog-addons-service-binding-usage-controller",
			},
		},
	}

	// Unstructured "service-catalog-addons-service-binding-usage-controller"
	serviceCatalogAddonsServiceBindingUsageControllerUnstructuredServiceMonitor = v1unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "monitoring.coreos.com/v1",
			"kind":       "ServiceMonitor",
			"metadata": map[string]any{
				"labels": map[string]any{
					"app":        "service-catalog-addons-service-binding-usage-controller",
					"chart":      "service-binding-usage-controller-0.1.0",
					"heritage":   "Helm",
					"prometheus": "monitoring",
					"release":    "service-catalog-addons",
				},
				"name": "service-catalog-addons-service-binding-usage-controller",
			},
		},
	}

	// Unstructured "deployment"
	deploymentUnstructuredUsageKind = v1unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "servicecatalog.kyma-project.io/v1alpha1",
			"kind":       "UsageKind",
			"metadata": map[string]any{
				"name": "deployment",
			},
		},
	}

	// Unstructured "service-catalog-addons-service-catalog-ui-catalog"
	serviceCatalogAddonsServiceCatalogUiCatalogUnstructuredVirtualService = v1unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "networking.istio.io/v1alpha3",
			"kind":       "VirtualService",
			"metadata": map[string]any{
				"labels": map[string]any{
					"app":      "service-catalog-ui",
					"chart":    "service-catalog-ui-0.1.0",
					"heritage": "Helm",
					"release":  "service-catalog-addons",
				},
				"name": "service-catalog-addons-service-catalog-ui-catalog",
			},
		},
	}
)
