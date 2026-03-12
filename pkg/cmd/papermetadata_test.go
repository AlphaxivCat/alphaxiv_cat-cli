// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersMetadataRetrieveLatestMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:metadata", "retrieve-latest-metadata",
			"--api-key", "string",
			"--upid", "x",
			"--prevent-tracking", "preventTracking",
		)
	})
}

func TestPapersMetadataRetrieveVersionMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:metadata", "retrieve-version-metadata",
			"--api-key", "string",
			"--upid", "x",
			"--version-order", "x",
			"--prevent-tracking", "preventTracking",
		)
	})
}
