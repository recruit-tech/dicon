package main

import "github.com/recruit-tech/dicon/examples/helloworld/di"

func main() {
	container := di.NewDIContainer()
	s, _ := container.SampleComponent()
	s.Exec()
}
