// Copyright 2020 The Khulnasoft Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommit_Submodule(t *testing.T) {
	c, err := testrepo.CatFileCommit("4e59b72440188e7c2578299fc28ea425fbe9aece")
	if err != nil {
		t.Fatal(err)
	}

	mod, err := c.Submodule("khulnasoft/docs-api")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "khulnasoft/docs-api", mod.Name)
	assert.Equal(t, "https://github.com/khulnasoft/docs-api.git", mod.URL)
	assert.Equal(t, "6b08f76a5313fa3d26859515b30aa17a5faa2807", mod.Commit)

	_, err = c.Submodule("404")
	assert.Equal(t, ErrSubmoduleNotExist, err)
}
