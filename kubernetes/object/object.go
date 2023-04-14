package object

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

// ContainsAnnotation determines whether the object contains an annotation.
func ContainsAnnotation(obj metav1.Object, key string) bool {
	if _, ok := obj.GetAnnotations()[key]; ok {
		return true
	}
	return false
}

// GetAnnotation returns the annotation value of the object.
func GetAnnotation(obj metav1.Object, key string) string {
	return obj.GetAnnotations()[key]
}

// ContainsLabel determines whether the object contains a label.
func ContainsLabel(obj metav1.Object, key string) bool {
	if _, ok := obj.GetLabels()[key]; ok {
		return true
	}
	return false
}

// GetLabel returns the label value of the object.
func GetLabel(obj metav1.Object, key string) string {
	return obj.GetLabels()[key]
}

// ContainsFinalizer determines whether the object contains a finalizer.
func ContainsFinalizer(obj metav1.Object, finalizer string) bool {
	finalizers := obj.GetFinalizers()
	for _, item := range finalizers {
		if item == finalizer {
			return true
		}
	}
	return false
}

// AddFinalizer adds a finalizer to the object, returns true if the object's finalizers are updated.
func AddFinalizer(obj metav1.Object, finalizer string) bool {
	finalizers := sets.New(obj.GetFinalizers()...)
	if finalizers.Has(finalizer) {
		return false
	}
	obj.SetFinalizers(finalizers.Insert(finalizer).UnsortedList())
	return true
}

// RemoveFinalizer removes the finalizer from the object, returns true if the object's finalizers are updated.
func RemoveFinalizer(obj metav1.Object, finalizer string) bool {
	finalizers := sets.New(obj.GetFinalizers()...)
	if !finalizers.Has(finalizer) {
		return false
	}
	obj.SetFinalizers(finalizers.Delete(finalizer).UnsortedList())
	return true
}
