// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestSitemapsListOverviews(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sitemaps", "list-overviews",
			"--api-key", "string",
			"--limit", "limit",
			"--page", "page",
		)
	})
}

func TestSitemapsListPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sitemaps", "list-papers",
			"--api-key", "string",
			"--limit", "limit",
			"--page", "page",
		)
	})
}

func TestSitemapsListUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sitemaps", "list-users",
			"--api-key", "string",
			"--limit", "limit",
			"--page", "page",
		)
	})
}
