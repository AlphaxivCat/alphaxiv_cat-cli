// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersIngestIngestLatest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:ingest", "ingest-latest",
			"--upid", "x",
			"--prevent-tracking", "preventTracking",
		)
	})
}

func TestPapersIngestIngestVersion(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:ingest", "ingest-version",
			"--upid", "x",
			"--version-label", "x",
			"--prevent-tracking", "preventTracking",
		)
	})
}
