// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersV3XMentionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3:x-mentions", "update",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--dry-run=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("dryRun: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3:x-mentions", "update",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3XMentionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3:x-mentions", "delete",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
