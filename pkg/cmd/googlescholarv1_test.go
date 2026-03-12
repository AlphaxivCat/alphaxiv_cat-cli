// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestGoogleScholarV1Connect(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "connect",
			"--api-key", "string",
			"--google-scholar-id", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("googleScholarId: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "google-scholar:v1", "connect",
			"--api-key", "string",
		)
	})
}

func TestGoogleScholarV1DeleteConnection(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "delete-connection",
			"--api-key", "string",
		)
	})
}

func TestGoogleScholarV1GetReport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "get-report",
			"--api-key", "string",
		)
	})
}

func TestGoogleScholarV1GetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "get-status",
			"--api-key", "string",
		)
	})
}

func TestGoogleScholarV1Resync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "resync",
			"--api-key", "string",
			"--mode", "all",
		)
	})
}

func TestGoogleScholarV1SetEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "set-email",
			"--api-key", "string",
			"--local-part", "localPart",
			"--verify-account-email=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"localPart: localPart\n" +
			"verifyAccountEmail: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "google-scholar:v1", "set-email",
			"--api-key", "string",
		)
	})
}

func TestGoogleScholarV1Sync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "sync",
			"--api-key", "string",
		)
	})
}

func TestGoogleScholarV1VerifyEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "google-scholar:v1", "verify-email",
			"--api-key", "string",
			"--code", "xxxxxx",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("code: xxxxxx")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "google-scholar:v1", "verify-email",
			"--api-key", "string",
		)
	})
}
