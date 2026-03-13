// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestCommentsV2Delete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "comments:v2", "delete",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCommentsV2Downvote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "comments:v2", "downvote",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCommentsV2Flag(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "comments:v2", "flag",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--reason", "reason",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: reason")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "comments:v2", "flag",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCommentsV2ToggleEndorse(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "comments:v2", "toggle-endorse",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCommentsV2Upvote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "comments:v2", "upvote",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
