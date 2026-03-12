// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestFoldersV3Create(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "create",
			"--api-key", "string",
			"--folder-name", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("folderName: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "folders:v3", "create",
			"--api-key", "string",
		)
	})
}

func TestFoldersV3List(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "list",
			"--api-key", "string",
		)
	})
}

func TestFoldersV3Delete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "delete",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersV3AddPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "add-papers",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--universal-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"paperGroupIds:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"universalIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "folders:v3", "add-papers",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersV3MovePapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "move-papers",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--from-folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"paperGroupIds:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"fromFolderId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "folders:v3", "move-papers",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersV3RemovePapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "remove-papers",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"paperGroupIds:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "folders:v3", "remove-papers",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersV3ToggleSharing(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "toggle-sharing",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sharing-status", "not_shared",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("sharingStatus: not_shared")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "folders:v3", "toggle-sharing",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersV3UpdateName(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "update-name",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "folders:v3", "update-name",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFoldersV3UpdateParent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "folders:v3", "update-parent",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--parent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("parentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "folders:v3", "update-parent",
			"--api-key", "string",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
