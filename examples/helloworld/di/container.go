package di

import (
	"github.com/recruit-tech/dicon/examples/helloworld/component"
)

// +DICON
type DIContainer interface {
	SampleComponent() (component.SampleComponent, error)
	DependencyComponent() (component.DependencyComponent, error)
}
