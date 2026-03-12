// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
)

func TestResearchCreateProject(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "research", "create-project",
			"--api-key", "string",
			"--name", "name",
			"--folder", "{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, delete: true}",
			"--initial-paper", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--parent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(researchCreateProject)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "research", "create-project",
			"--api-key", "string",
			"--name", "name",
			"--folder.id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--folder.delete=true",
			"--initial-paper", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--parent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"folder:\n" +
			"  id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  delete: true\n" +
			"initialPapers:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"parentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "research", "create-project",
			"--api-key", "string",
		)
	})
}
