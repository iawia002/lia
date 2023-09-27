package client

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestApply(t *testing.T) {
	pod1 := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod1",
		},
	}

	tests := []struct {
		name  string
		obj   client.Object
		objs  []runtime.Object
		isErr bool
	}{
		{
			name:  "create test",
			obj:   pod1,
			isErr: false,
		},
		{
			name:  "update test",
			obj:   pod1,
			objs:  []runtime.Object{pod1},
			isErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := fake.NewFakeClient(tt.objs...)
			if err := Apply(
				context.TODO(), c, tt.obj,
				WithDryRun(true), WithForce(true), WithFieldManager("test"),
			); tt.isErr != (err != nil) {
				t.Errorf("%s Apply() unexpected error: %v", tt.name, err)
			}
		})
	}
}
