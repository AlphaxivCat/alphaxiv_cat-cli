// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
)

func TestAssistantV2Chat(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2", "chat",
			"--api-key", "string",
			"--max-items", "10",
			"--file", "{contentType: contentType, url: https://example.com}",
			"--llm-chat-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message", "message",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--parent-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--selection-page-range", "[0, 0]",
			"--thinking=true",
			"--web-search", "off",
			"--assistant-variant", "homepage",
			"--folder-add-papers=true",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--model", "gemini-2.5-flash",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(assistantV2Chat)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2", "chat",
			"--api-key", "string",
			"--max-items", "10",
			"--file.content-type", "contentType",
			"--file.url", "https://example.com",
			"--llm-chat-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message", "message",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--parent-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--selection-page-range", "[0, 0]",
			"--thinking=true",
			"--web-search", "off",
			"--assistant-variant", "homepage",
			"--folder-add-papers=true",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--model", "gemini-2.5-flash",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"files:\n" +
			"  - contentType: contentType\n" +
			"    url: https://example.com\n" +
			"llmChatId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"message: message\n" +
			"paperVersionId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"selectionPageRange:\n" +
			"  - 0\n" +
			"  - 0\n" +
			"thinking: true\n" +
			"webSearch: 'off'\n" +
			"assistantVariant: homepage\n" +
			"folderAddPapers: true\n" +
			"folderId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"model: gemini-2.5-flash\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "assistant:v2", "chat",
			"--api-key", "string",
			"--max-items", "10",
		)
	})
}

func TestAssistantV2DeleteChat(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2", "delete-chat",
			"--api-key", "string",
			"--llm-chat", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAssistantV2EditChat(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2", "edit-chat",
			"--api-key", "string",
			"--llm-chat", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("title: title")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "assistant:v2", "edit-chat",
			"--api-key", "string",
			"--llm-chat", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAssistantV2GetChats(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2", "get-chats",
			"--api-key", "string",
			"--folder", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--paper-version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--variant", "homepage",
		)
	})
}

func TestAssistantV2GetURLMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "assistant:v2", "get-url-metadata",
			"--api-key", "string",
			"--url", "https://example.com",
		)
	})
}
