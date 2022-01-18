/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

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

// Code generated by reverse-kube-resource. DO NOT EDIT.

package csicinder

import (
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
)

// ServiceAccount "csi-cinder-controller-sa"
func csiCinderControllerSaServiceAccountCreator() reconciling.NamedServiceAccountCreatorGetter {
	return func() (string, reconciling.ServiceAccountCreator) {
		t := csiCinderControllerSaServiceAccount.DeepCopy()
		return t.Name, func(o *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
			return t, nil
		}
	}
}

// ServiceAccount "csi-cinder-node-sa"
func csiCinderNodeSaServiceAccountCreator() reconciling.NamedServiceAccountCreatorGetter {
	return func() (string, reconciling.ServiceAccountCreator) {
		t := csiCinderNodeSaServiceAccount.DeepCopy()
		return t.Name, func(o *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
			return t, nil
		}
	}
}

// StorageClass "csi-cinder-sc-delete"
func csiCinderScDeleteStorageClassCreator() reconciling.NamedStorageClassCreatorGetter {
	return func() (string, reconciling.StorageClassCreator) {
		t := csiCinderScDeleteStorageClass.DeepCopy()
		return t.Name, func(o *storagev1.StorageClass) (*storagev1.StorageClass, error) {
			return t, nil
		}
	}
}

// StorageClass "csi-cinder-sc-retain"
func csiCinderScRetainStorageClassCreator() reconciling.NamedStorageClassCreatorGetter {
	return func() (string, reconciling.StorageClassCreator) {
		t := csiCinderScRetainStorageClass.DeepCopy()
		return t.Name, func(o *storagev1.StorageClass) (*storagev1.StorageClass, error) {
			return t, nil
		}
	}
}

// ClusterRole "csi-attacher-role"
func csiAttacherRoleClusterRoleCreator() reconciling.NamedClusterRoleCreatorGetter {
	return func() (string, reconciling.ClusterRoleCreator) {
		t := csiAttacherRoleClusterRole.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {
			return t, nil
		}
	}
}

// ClusterRole "csi-provisioner-role"
func csiProvisionerRoleClusterRoleCreator() reconciling.NamedClusterRoleCreatorGetter {
	return func() (string, reconciling.ClusterRoleCreator) {
		t := csiProvisionerRoleClusterRole.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {
			return t, nil
		}
	}
}

// ClusterRole "csi-snapshotter-role"
func csiSnapshotterRoleClusterRoleCreator() reconciling.NamedClusterRoleCreatorGetter {
	return func() (string, reconciling.ClusterRoleCreator) {
		t := csiSnapshotterRoleClusterRole.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {
			return t, nil
		}
	}
}

// ClusterRole "csi-resizer-role"
func csiResizerRoleClusterRoleCreator() reconciling.NamedClusterRoleCreatorGetter {
	return func() (string, reconciling.ClusterRoleCreator) {
		t := csiResizerRoleClusterRole.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {
			return t, nil
		}
	}
}

// ClusterRole "csi-nodeplugin-role"
func csiNodepluginRoleClusterRoleCreator() reconciling.NamedClusterRoleCreatorGetter {
	return func() (string, reconciling.ClusterRoleCreator) {
		t := csiNodepluginRoleClusterRole.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {
			return t, nil
		}
	}
}

// ClusterRoleBinding "csi-attacher-binding"
func csiAttacherBindingClusterRoleBindingCreator() reconciling.NamedClusterRoleBindingCreatorGetter {
	return func() (string, reconciling.ClusterRoleBindingCreator) {
		t := csiAttacherBindingClusterRoleBinding.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			return t, nil
		}
	}
}

// ClusterRoleBinding "csi-provisioner-binding"
func csiProvisionerBindingClusterRoleBindingCreator() reconciling.NamedClusterRoleBindingCreatorGetter {
	return func() (string, reconciling.ClusterRoleBindingCreator) {
		t := csiProvisionerBindingClusterRoleBinding.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			return t, nil
		}
	}
}

// ClusterRoleBinding "csi-snapshotter-binding"
func csiSnapshotterBindingClusterRoleBindingCreator() reconciling.NamedClusterRoleBindingCreatorGetter {
	return func() (string, reconciling.ClusterRoleBindingCreator) {
		t := csiSnapshotterBindingClusterRoleBinding.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			return t, nil
		}
	}
}

// ClusterRoleBinding "csi-resizer-binding"
func csiResizerBindingClusterRoleBindingCreator() reconciling.NamedClusterRoleBindingCreatorGetter {
	return func() (string, reconciling.ClusterRoleBindingCreator) {
		t := csiResizerBindingClusterRoleBinding.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			return t, nil
		}
	}
}

// ClusterRoleBinding "csi-nodeplugin-binding"
func csiNodepluginBindingClusterRoleBindingCreator() reconciling.NamedClusterRoleBindingCreatorGetter {
	return func() (string, reconciling.ClusterRoleBindingCreator) {
		t := csiNodepluginBindingClusterRoleBinding.DeepCopy()
		return t.Name, func(o *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			return t, nil
		}
	}
}

// Role "external-resizer-cfg"
func externalResizerCfgRoleCreator() reconciling.NamedRoleCreatorGetter {
	return func() (string, reconciling.RoleCreator) {
		t := externalResizerCfgRole.DeepCopy()
		return t.Name, func(o *rbacv1.Role) (*rbacv1.Role, error) {
			return t, nil
		}
	}
}

// RoleBinding "csi-resizer-role-cfg"
func csiResizerRoleCfgRoleBindingCreator() reconciling.NamedRoleBindingCreatorGetter {
	return func() (string, reconciling.RoleBindingCreator) {
		t := csiResizerRoleCfgRoleBinding.DeepCopy()
		return t.Name, func(o *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error) {
			return t, nil
		}
	}
}

// Service "openstack-cinder-csi-controllerplugin"
func openstackCinderCsiControllerpluginServiceCreator() reconciling.NamedServiceCreatorGetter {
	return func() (string, reconciling.ServiceCreator) {
		t := openstackCinderCsiControllerpluginService.DeepCopy()
		return t.Name, func(o *corev1.Service) (*corev1.Service, error) {
			return t, nil
		}
	}
}

// DaemonSet "openstack-cinder-csi-nodeplugin"
func openstackCinderCsiNodepluginDaemonSetCreator() reconciling.NamedDaemonSetCreatorGetter {
	return func() (string, reconciling.DaemonSetCreator) {
		t := openstackCinderCsiNodepluginDaemonSet.DeepCopy()
		return t.Name, func(o *appsv1.DaemonSet) (*appsv1.DaemonSet, error) {
			return t, nil
		}
	}
}

// StatefulSet "openstack-cinder-csi-controllerplugin"
func openstackCinderCsiControllerpluginStatefulSetCreator() reconciling.NamedStatefulSetCreatorGetter {
	return func() (string, reconciling.StatefulSetCreator) {
		t := openstackCinderCsiControllerpluginStatefulSet.DeepCopy()
		return t.Name, func(o *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
			return t, nil
		}
	}
}

// CSIDriver "cinder.csi.openstack.org"
func cinderCsiOpenstackOrgCSIDriverCreator() reconciling.NamedCSIDriverCreatorGetter {
	return func() (string, reconciling.CSIDriverCreator) {
		t := cinderCsiOpenstackOrgCSIDriver.DeepCopy()
		return t.Name, func(o *storagev1.CSIDriver) (*storagev1.CSIDriver, error) {
			return t, nil
		}
	}
}