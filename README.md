Reverse Kube Resources
===

This tool is intended to generate golang code tightly coupled with [Kubernetes client-go](https://github.com/kubernetes/client-go).

The input are raw yaml manifests and the tool leverages Kubernetes scheme to generate valid golang structures compatible with client-go.

### Example:
The directory [`examples`](./examples) contains few sample yaml files with kubernetes resources.

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: test
spec: {}
status: {}
```

After executing the generator with following arguments:
```
$ go run github.com/wozniakjan/reverse-kube-resource -package test -src ./examples/ns.yaml > output/ns.go
```

The output is a buildable go file:
```go
// Code generated by reverse-kube-resource. DO NOT EDIT.

package test

import (
        corev1 "k8s.io/api/core/v1"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var testNamespace = corev1.Namespace{
        ObjectMeta: metav1.ObjectMeta{
                Name: "test",
        },
        Spec:   corev1.NamespaceSpec{},
        Status: corev1.NamespaceStatus{},
}
```
