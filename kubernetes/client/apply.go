package client

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ApplyOptions defines options needed to create or patch an object.
type ApplyOptions struct {
	dryRun       bool
	fieldManager string
	force        bool
}

// WithDryRun sets the DryRun option when creating or patching the object.
// The default value is `false`.
func WithDryRun(dryRun bool) func(*ApplyOptions) {
	return func(o *ApplyOptions) {
		o.dryRun = dryRun
	}
}

// WithFieldManager sets the FieldManager option when creating or patching the object.
// The default value is `controller-runtime`.
func WithFieldManager(fieldManager string) func(*ApplyOptions) {
	return func(o *ApplyOptions) {
		o.fieldManager = fieldManager
	}
}

// WithForce sets the Force option when patching the object.
// The default value is `true`.
func WithForce(force bool) func(*ApplyOptions) {
	return func(o *ApplyOptions) {
		o.force = force
	}
}

// Apply creates the given object or updates the existing object to the given one using apply patch.
func Apply(ctx context.Context, c client.Client, obj client.Object, options ...func(*ApplyOptions)) error {
	opts := &ApplyOptions{
		dryRun:       false,
		fieldManager: "controller-runtime",
		force:        true,
	}
	for _, f := range options {
		f(opts)
	}

	key := client.ObjectKeyFromObject(obj)
	newObj := obj.DeepCopyObject().(client.Object)
	if err := c.Get(ctx, key, newObj); err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		createOptions := []client.CreateOption{
			client.FieldOwner(opts.fieldManager),
		}
		if opts.dryRun {
			createOptions = append(createOptions, client.DryRunAll)
		}
		return c.Create(ctx, obj, createOptions...)
	}

	patchOptions := []client.PatchOption{
		client.FieldOwner(opts.fieldManager),
	}
	if opts.dryRun {
		patchOptions = append(patchOptions, client.DryRunAll)
	}
	if opts.force {
		patchOptions = append(patchOptions, client.ForceOwnership)
	}
	return c.Patch(ctx, obj, client.Apply, patchOptions...)
}
