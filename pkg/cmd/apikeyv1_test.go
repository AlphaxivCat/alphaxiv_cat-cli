// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestAPIKeysV1Create(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "api-keys:v1", "create",
			"--api-key", "string",
			"--label", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("label: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "api-keys:v1", "create",
			"--api-key", "string",
		)
	})
}

func TestAPIKeysV1List(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "api-keys:v1", "list",
			"--api-key", "string",
		)
	})
}

func TestAPIKeysV1CreateImpersonation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "api-keys:v1", "create-impersonation",
			"--api-key", "string",
			"--user", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("user: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "api-keys:v1", "create-impersonation",
			"--api-key", "string",
		)
	})
}

func TestAPIKeysV1Revoke(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "api-keys:v1", "revoke",
			"--api-key", "string",
			"--api-key-id", "own",
		)
	})
}
