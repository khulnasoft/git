// Copyright 2022 The Khulnasoft Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_CatFileBlob(t *testing.T) {
	t.Run("not a blob", func(t *testing.T) {
		_, err := testrepo.CatFileBlob("007cb92318c7bd3b56908ea8c2e54370245562f8")
		assert.Equal(t, ErrNotBlob, err)
	})

	t.Run("get a blob, no full rev hash", func(t *testing.T) {
		b, err := testrepo.CatFileBlob("021a")
		require.NoError(t, err)
		assert.True(t, b.IsBlob())
	})

	t.Run("get a blob", func(t *testing.T) {
		b, err := testrepo.CatFileBlob("021a721a61a1de65865542c405796d1eb985f784")
		require.NoError(t, err)
		assert.True(t, b.IsBlob())
	})
}
