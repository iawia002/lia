package object

import (
	"reflect"
	"sort"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestContainsAnnotation(t *testing.T) {
	tests := []struct {
		name   string
		obj    metav1.Object
		key    string
		wanted bool
	}{
		{
			name: "contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "aa",
			wanted: true,
		},
		{
			name: "not contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "bb",
			wanted: false,
		},
		{
			name:   "nil test",
			obj:    &corev1.Node{},
			key:    "aa",
			wanted: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAnnotation(tt.obj, tt.key); got != tt.wanted {
				t.Errorf("ContainsAnnotation() = %v, want %v", got, tt.wanted)
			}
		})
	}
}

func TestGetAnnotation(t *testing.T) {
	tests := []struct {
		name   string
		obj    metav1.Object
		key    string
		wanted string
	}{
		{
			name: "contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "aa",
			wanted: "bb",
		},
		{
			name: "not contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "bb",
			wanted: "",
		},
		{
			name:   "nil test",
			obj:    &corev1.Node{},
			key:    "aa",
			wanted: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAnnotation(tt.obj, tt.key); got != tt.wanted {
				t.Errorf("GetAnnotation() = %v, want %v", got, tt.wanted)
			}
		})
	}
}

func TestAddAnnotation(t *testing.T) {
	tests := []struct {
		name    string
		obj     metav1.Object
		k, v    string
		updated bool
		wanted  map[string]string
	}{
		{
			name: "normal test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"aa": "bb",
					},
				},
			},
			k:       "aa",
			v:       "cc",
			updated: true,
			wanted: map[string]string{
				"aa": "cc",
			},
		},
		{
			name: "not updated test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"aa": "bb",
					},
				},
			},
			k:       "aa",
			v:       "bb",
			updated: false,
			wanted: map[string]string{
				"aa": "bb",
			},
		},
		{
			name:    "nil test",
			obj:     &corev1.Node{},
			k:       "aa",
			v:       "bb",
			updated: true,
			wanted: map[string]string{
				"aa": "bb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAnnotation(tt.obj, tt.k, tt.v); !reflect.DeepEqual(got, tt.updated) {
				t.Errorf("AddAnnotation() = %v, want %v", got, tt.wanted)
			}
			if !reflect.DeepEqual(tt.obj.GetAnnotations(), tt.wanted) {
				t.Errorf("AddAnnotation() = %v, want %v", tt.obj.GetAnnotations(), tt.wanted)
			}
		})
	}
}

func TestContainsLabel(t *testing.T) {
	tests := []struct {
		name   string
		obj    metav1.Object
		key    string
		wanted bool
	}{
		{
			name: "contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "aa",
			wanted: true,
		},
		{
			name: "not contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "bb",
			wanted: false,
		},
		{
			name:   "nil test",
			obj:    &corev1.Node{},
			key:    "aa",
			wanted: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsLabel(tt.obj, tt.key); got != tt.wanted {
				t.Errorf("ContainsLabel() = %v, want %v", got, tt.wanted)
			}
		})
	}
}

func TestGetLabel(t *testing.T) {
	tests := []struct {
		name   string
		obj    metav1.Object
		key    string
		wanted string
	}{
		{
			name: "contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "aa",
			wanted: "bb",
		},
		{
			name: "not contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"aa": "bb",
					},
				},
			},
			key:    "bb",
			wanted: "",
		},
		{
			name:   "nil test",
			obj:    &corev1.Node{},
			key:    "aa",
			wanted: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLabel(tt.obj, tt.key); got != tt.wanted {
				t.Errorf("GetLabel() = %v, want %v", got, tt.wanted)
			}
		})
	}
}

func TestAddLabel(t *testing.T) {
	tests := []struct {
		name    string
		obj     metav1.Object
		k, v    string
		updated bool
		wanted  map[string]string
	}{
		{
			name: "normal test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"aa": "bb",
					},
				},
			},
			k:       "aa",
			v:       "cc",
			updated: true,
			wanted: map[string]string{
				"aa": "cc",
			},
		},
		{
			name: "not updated test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"aa": "bb",
					},
				},
			},
			k:       "aa",
			v:       "bb",
			updated: false,
			wanted: map[string]string{
				"aa": "bb",
			},
		},
		{
			name:    "nil test",
			obj:     &corev1.Node{},
			k:       "aa",
			v:       "bb",
			updated: true,
			wanted: map[string]string{
				"aa": "bb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddLabel(tt.obj, tt.k, tt.v); !reflect.DeepEqual(got, tt.updated) {
				t.Errorf("AddLabel() = %v, want %v", got, tt.wanted)
			}
			if !reflect.DeepEqual(tt.obj.GetLabels(), tt.wanted) {
				t.Errorf("AddLabel() = %v, want %v", tt.obj.GetLabels(), tt.wanted)
			}
		})
	}
}

func TestContainsFinalizer(t *testing.T) {
	tests := []struct {
		name      string
		obj       metav1.Object
		finalizer string
		wanted    bool
	}{
		{
			name: "contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{"aa", "bb"},
				},
			},
			finalizer: "aa",
			wanted:    true,
		},
		{
			name: "not contains test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{"aa", "bb"},
				},
			},
			finalizer: "cc",
			wanted:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsFinalizer(tt.obj, tt.finalizer); got != tt.wanted {
				t.Errorf("ContainsFinalizer() = %v, want %v", got, tt.wanted)
			}
		})
	}
}

func TestAddFinalizer(t *testing.T) {
	tests := []struct {
		name             string
		obj              metav1.Object
		finalizer        string
		updated          bool
		wantedFinalizers []string
	}{
		{
			name: "updated test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{"aa", "bb"},
				},
			},
			finalizer:        "cc",
			updated:          true,
			wantedFinalizers: []string{"aa", "bb", "cc"},
		},
		{
			name: "not updated test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{"aa", "bb"},
				},
			},
			finalizer:        "aa",
			updated:          false,
			wantedFinalizers: []string{"aa", "bb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddFinalizer(tt.obj, tt.finalizer); got != tt.updated {
				t.Errorf("AddFinalizer() = %v, want %v", got, tt.updated)
			}
			finalizers := tt.obj.GetFinalizers()
			sort.Strings(finalizers)
			if !reflect.DeepEqual(finalizers, tt.wantedFinalizers) {
				t.Errorf("AddFinalizer() = %v, want %v", finalizers, tt.wantedFinalizers)
			}
		})
	}
}

func TestRemoveFinalizer(t *testing.T) {
	tests := []struct {
		name             string
		obj              metav1.Object
		finalizer        string
		updated          bool
		wantedFinalizers []string
	}{
		{
			name: "updated test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{"aa", "bb"},
				},
			},
			finalizer:        "aa",
			updated:          true,
			wantedFinalizers: []string{"bb"},
		},
		{
			name: "not updated test",
			obj: &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{"aa", "bb"},
				},
			},
			finalizer:        "cc",
			updated:          false,
			wantedFinalizers: []string{"aa", "bb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFinalizer(tt.obj, tt.finalizer); got != tt.updated {
				t.Errorf("RemoveFinalizer() = %v, want %v", got, tt.updated)
			}
			finalizers := tt.obj.GetFinalizers()
			sort.Strings(finalizers)
			if !reflect.DeepEqual(finalizers, tt.wantedFinalizers) {
				t.Errorf("RemoveFinalizer() = %v, want %v", finalizers, tt.wantedFinalizers)
			}
		})
	}
}
