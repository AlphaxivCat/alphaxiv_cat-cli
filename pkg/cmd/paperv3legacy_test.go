// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersV3LegacyRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3:legacy", "retrieve",
			"--api-key", "string",
			"--unresolved", "unresolved",
		)
	})
}

func TestPapersV3LegacyRetrieveComments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3:legacy", "retrieve-comments",
			"--api-key", "string",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
