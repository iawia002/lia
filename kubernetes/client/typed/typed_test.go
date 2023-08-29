package typed

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var (
	pod1Name = "pod1"
	pod1     = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: pod1Name,
		},
	}
)

func TestTypedGet(t *testing.T) {
	tests := []struct {
		name   string
		objs   []runtime.Object
		gvk    schema.GroupVersionKind
		key    types.NamespacedName
		isErr  bool
		wanted string
	}{
		{
			name: "normal test",
			objs: []runtime.Object{pod1},
			gvk:  corev1.SchemeGroupVersion.WithKind("Pod"),
			key: types.NamespacedName{
				Name: pod1Name,
			},
			wanted: pod1Name,
		},
		{
			name: "not exists test",
			objs: []runtime.Object{pod1},
			gvk:  corev1.SchemeGroupVersion.WithKind("Pod"),
			key: types.NamespacedName{
				Name: "aaa",
			},
			isErr: true,
		},
		{
			name: "wrong type test",
			objs: []runtime.Object{pod1},
			gvk:  corev1.SchemeGroupVersion.WithKind("Pod1"),
			key: types.NamespacedName{
				Name: pod1Name,
			},
			isErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := fake.NewFakeClient(tt.objs...)
			typedClient, _ := NewTypedClient(tt.gvk, WithClientReader(c))

			got, err := typedClient.Get(context.TODO(), tt.key)
			if tt.isErr != (err != nil) {
				t.Errorf("%s Get() unexpected error: %v", tt.name, err)
			}
			if tt.isErr {
				return
			}
			pod := got.(*corev1.Pod)
			if pod.Name != tt.wanted {
				t.Errorf("Get() = %v, want %v", pod.Name, tt.wanted)
			}
		})
	}
}

func TestTypedList(t *testing.T) {
	tests := []struct {
		name   string
		objs   []runtime.Object
		gvk    schema.GroupVersionKind
		isErr  bool
		wanted int
	}{
		{
			name:   "normal test",
			objs:   []runtime.Object{pod1},
			gvk:    corev1.SchemeGroupVersion.WithKind("Pod"),
			wanted: 1,
		},
		{
			name:  "wrong type test",
			objs:  []runtime.Object{pod1},
			gvk:   corev1.SchemeGroupVersion.WithKind("Pod1"),
			isErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := fake.NewFakeClient(tt.objs...)
			typedClient, _ := NewTypedClient(tt.gvk, WithClientReader(c))

			got, err := typedClient.List(context.TODO(), metav1.NamespaceAll)
			if tt.isErr != (err != nil) {
				t.Errorf("%s List() unexpected error: %v", tt.name, err)
			}
			if tt.isErr {
				return
			}
			pods := got.(*corev1.PodList)
			if len(pods.Items) != tt.wanted {
				t.Errorf("List() = %v, want %v", len(pods.Items), tt.wanted)
			}
		})
	}
}
