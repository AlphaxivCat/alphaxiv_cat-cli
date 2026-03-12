// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersV3ImplementationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3:implementations", "create",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--implementation-type", "alphaxiv",
			"--url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"implementationType: alphaxiv\n" +
			"url: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3:implementations", "create",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3ImplementationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3:implementations", "list",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3ImplementationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3:implementations", "delete",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--implementation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--type", "alphaxiv",
		)
	})
}
