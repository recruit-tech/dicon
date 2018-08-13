// Code generated by "dicon"; DO NOT EDIT.

package di

import (
	"fmt"

	"github.com/akito0107/dicon/examples/helloworld/component"
	"github.com/pkg/errors"
)

type dicontainer struct {
	store map[string]interface{}
}

func NewDIContainer() DIContainer {
	return &dicontainer{
		store: map[string]interface{}{},
	}
}

func (d *dicontainer) DependencyComponent() (component.DependencyComponent, error) {
	if i, ok := d.store["DependencyComponent"]; ok {
		instance, ok := i.(component.DependencyComponent)
		if !ok {
			return nil, fmt.Errorf("invalid instance is cached %v", instance)
		}
		return instance, nil
	}
	instance, err := component.NewDependencyComponent()
	if err != nil {
		return nil, errors.Wrap(err, "creation DependencyComponent failed at DICON")
	}
	d.store["DependencyComponent"] = instance
	return instance, nil
}
func (d *dicontainer) SampleComponent() (component.SampleComponent, error) {
	if i, ok := d.store["SampleComponent"]; ok {
		instance, ok := i.(component.SampleComponent)
		if !ok {
			return nil, fmt.Errorf("invalid instance is cached %v", instance)
		}
		return instance, nil
	}
	dep0, err := d.DependencyComponent()
	if err != nil {
		return nil, errors.Wrap(err, "resolve DependencyComponent failed at DICON")
	}
	instance, err := component.NewSampleComponent(dep0)
	if err != nil {
		return nil, errors.Wrap(err, "creation SampleComponent failed at DICON")
	}
	d.store["SampleComponent"] = instance
	return instance, nil
}