// Code generated by reverse-kube-resource. DO NOT EDIT.

package examples

import (
	corev1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// Namespace "test2"
	test2Namespace = corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test2",
		},
		Spec:   corev1.NamespaceSpec{},
		Status: corev1.NamespaceStatus{},
	}

	// PersistentVolume "pv0004"
	pv0004VolumeModeFilesystem corev1.PersistentVolumeMode = "Filesystem"

	pv0004PersistentVolume = corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pv0004",
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: map[corev1.ResourceName]apiresource.Quantity{
				"storage": apiresource.MustParse("5Gi"),
			},
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				NFS: &corev1.NFSVolumeSource{
					Server: "172.17.0.2",
					Path:   "/tmp",
				},
			},
			AccessModes: []corev1.PersistentVolumeAccessMode{
				"ReadWriteOnce",
			},
			PersistentVolumeReclaimPolicy: "Recycle",
			StorageClassName:              "slow",
			MountOptions: []string{
				"hard",
				"nfsvers=4.1",
			},
			VolumeMode: &pv0004VolumeModeFilesystem,
		},
	}
)
