// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
)

func TestEmailsCaptureBouncedEmails(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "capture-bounced-emails",
			"--api-key", "string",
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
			t, pipeData, "emails", "capture-bounced-emails",
			"--api-key", "string",
		)
	})
}

func TestEmailsCaptureResendBouncedEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "capture-resend-bounced-email",
			"--api-key", "string",
			"--data", "{to: [string]}",
			"--type", "type",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailsCaptureResendBouncedEmail)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "capture-resend-bounced-email",
			"--api-key", "string",
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
			t, pipeData, "emails", "capture-resend-bounced-email",
			"--api-key", "string",
		)
	})
}

func TestEmailsKickoffCommentUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "kickoff-comment-update",
			"--api-key", "string",
			"--role", "role",
			"--window", "window",
			"--custom", "true",
		)
	})
}

func TestEmailsKickoffGeneralUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "kickoff-general-update",
			"--api-key", "string",
			"--role", "role",
		)
	})
}

func TestEmailsProcessBouncedEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "process-bounced-email",
			"--api-key", "string",
			"--email", "dev@stainless.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("email: dev@stainless.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "emails", "process-bounced-email",
			"--api-key", "string",
		)
	})
}

func TestEmailsProcessCommentUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "process-comment-update",
			"--api-key", "string",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--custom-content", "{events: [{date: date, description: description, link: link, title: title, endTimeRaw: endTimeRaw, startTimeRaw: startTimeRaw}], introText: introText, subject: subject}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailsProcessCommentUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "process-comment-update",
			"--api-key", "string",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--custom-content.events", "[{date: date, description: description, link: link, title: title, endTimeRaw: endTimeRaw, startTimeRaw: startTimeRaw}]",
			"--custom-content.intro-text", "introText",
			"--custom-content.subject", "subject",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"userId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"customContent:\n" +
			"  events:\n" +
			"    - date: date\n" +
			"      description: description\n" +
			"      link: link\n" +
			"      title: title\n" +
			"      endTimeRaw: endTimeRaw\n" +
			"      startTimeRaw: startTimeRaw\n" +
			"  introText: introText\n" +
			"  subject: subject\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "emails", "process-comment-update",
			"--api-key", "string",
		)
	})
}

func TestEmailsProcessGeneralUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "emails", "process-general-update",
			"--api-key", "string",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("userId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "emails", "process-general-update",
			"--api-key", "string",
		)
	})
}
