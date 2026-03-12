// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestAdminV1GetModeratorFeed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "admin:v1", "get-moderator-feed",
			"--api-key", "string",
			"--feed-type", "all",
			"--page", "page",
			"--page-size", "pageSize",
		)
	})
}

func TestAdminV1LookupUserByEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "admin:v1", "lookup-user-by-email",
			"--api-key", "string",
			"--email", "dev@stainless.com",
		)
	})
}
