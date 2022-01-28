/*
Boilerplate 2022 test.
*/

// Code generated by reverse-kube-resource. DO NOT EDIT.

package examples

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clusterRoleClusterRole = rbacv1.ClusterRole{
	ObjectMeta: metav1.ObjectMeta{
		Name: "cluster-role",
	},
	Rules: []rbacv1.PolicyRule{
		rbacv1.PolicyRule{
			Verbs: []string{
				"get",
				"list",
				"watch",
				"patch",
			},
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"persistentvolumes",
			},
		},
		rbacv1.PolicyRule{
			Verbs: []string{
				"get",
				"list",
				"watch",
			},
			APIGroups: []string{
				"storage.k8s.io",
			},
			Resources: []string{
				"csinodes",
			},
		},
		rbacv1.PolicyRule{
			Verbs: []string{
				"get",
				"list",
				"watch",
				"patch",
			},
			APIGroups: []string{
				"storage.k8s.io",
			},
			Resources: []string{
				"volumeattachments",
			},
		},
		rbacv1.PolicyRule{
			Verbs: []string{
				"patch",
			},
			APIGroups: []string{
				"storage.k8s.io",
			},
			Resources: []string{
				"volumeattachments/status",
			},
		},
	},
}
