package foo

import (
	"fmt"
     "github.com/rubensseva/module-test/internal/shared"
)

func Foo() {
	fmt.Println("foo")
	shared.PrintTheThing()
}
