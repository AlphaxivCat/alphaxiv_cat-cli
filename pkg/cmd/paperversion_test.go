// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersVersionsRequestAIOverview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:versions", "request-ai-overview",
			"--api-key", "string",
			"--upid", "x",
			"--version-order", "x",
			"--preferred-language", "am",
		)
	})
}

func TestPapersVersionsRequestAITranslation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:versions", "request-ai-translation",
			"--api-key", "string",
			"--upid", "x",
			"--version-order", "x",
			"--language", "am",
		)
	})
}
