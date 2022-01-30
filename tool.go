//go:build tools
// +build tools

package buttonsapi

import (
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/kisielk/godepgraph"
)
