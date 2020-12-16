package controller

import (
	"github.com/tkestack/yarn-opterator/pkg/controller/nodemanagerset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nodemanagerset.Add)
}
