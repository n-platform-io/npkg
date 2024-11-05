package actions

import (
	"fmt"
	"github.com/n-platform-io/npkg/types"
)

func Build(pkg *types.PackageBuild) {
	fmt.Println("Building...", pkg)
}
