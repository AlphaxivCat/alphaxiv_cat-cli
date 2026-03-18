// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersV3OverviewRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3:overview", "retrieve",
			"--paper-version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--language", "am",
		)
	})
}

func TestPapersV3OverviewRetrieveStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3:overview", "retrieve-status",
			"--paper-version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
