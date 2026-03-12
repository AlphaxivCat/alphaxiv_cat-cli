// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestAssistantV2MessagesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2:messages", "list",
			"--api-key", "string",
			"--llm-chat", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAssistantV2MessagesSelect(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2:messages", "select",
			"--api-key", "string",
			"--llm-chat", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAssistantV2MessagesSetFeedback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2:messages", "set-feedback",
			"--api-key", "string",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--feedback", "{type: upvote, category: category, details: details}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(assistantV2MessagesSetFeedback)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2:messages", "set-feedback",
			"--api-key", "string",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--feedback.type", "upvote",
			"--feedback.category", "category",
			"--feedback.details", "details",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"feedback:\n" +
			"  type: upvote\n" +
			"  category: category\n" +
			"  details: details\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "assistant:v2:messages", "set-feedback",
			"--api-key", "string",
			"--message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
