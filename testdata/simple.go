//nolint // this is test
package options

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Simple struct {
	//options:ignore
	StringVal string
	IntVal    int
}
