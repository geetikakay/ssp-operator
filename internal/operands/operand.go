package operands

import (
	"sigs.k8s.io/controller-runtime/pkg/client"

	"kubevirt.io/ssp-operator/internal/common"
)

type Operand interface {
	// WatchTypes returns a slice of namespaced resources, that the operator should watch.
	WatchTypes() []client.Object

	// WatchClusterTypes returns a slice of cluster resources, that the operator should watch.
	WatchClusterTypes() []client.Object

	// Reconcile creates and updates resources.
	Reconcile(*common.Request) ([]common.ReconcileResult, error)

	// Cleanup removes any created cluster resources.
	// They don't use owner references, so the garbage collector will not remove them.
	Cleanup(*common.Request) error

	// Name returns the name of the operand
	Name() string
}
