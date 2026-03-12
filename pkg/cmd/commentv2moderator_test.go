// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestCommentsV2ModeratorSendFeedback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "comments:v2:moderator", "send-feedback",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message", "message",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("message: message")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "comments:v2:moderator", "send-feedback",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCommentsV2ModeratorToggleFlags(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "comments:v2:moderator", "toggle-flags",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--addressed=true",
			"--closed=true",
			"--flag-addressed=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"addressed: true\n" +
			"closed: true\n" +
			"flagAddressed: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "comments:v2:moderator", "toggle-flags",
			"--api-key", "string",
			"--comment", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
