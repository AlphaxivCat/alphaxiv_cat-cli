// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestEmailsCaptureBouncedEmails(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"emails", "capture-bounced-emails",
			"--message", "Message",
			"--type", "Type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"Message: Message\n" +
			"Type: Type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"emails", "capture-bounced-emails",
		)
	})
}

func TestEmailsCaptureResendBouncedEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"emails", "capture-resend-bounced-email",
			"--data", "{to: [string]}",
			"--type", "type",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailsCaptureResendBouncedEmail)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"emails", "capture-resend-bounced-email",
			"--data.to", "[string]",
			"--type", "type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  to:\n" +
			"    - string\n" +
			"type: type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"emails", "capture-resend-bounced-email",
		)
	})
}

func TestEmailsProcessBouncedEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"emails", "process-bounced-email",
			"--email", "dev@stainless.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("email: dev@stainless.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"emails", "process-bounced-email",
		)
	})
}
