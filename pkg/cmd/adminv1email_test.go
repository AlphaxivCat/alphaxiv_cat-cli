// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestAdminV1EmailsSendWeeklyDigest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"admin:v1:emails", "send-weekly-digest",
			"--a", "{introText: introText, subject: subject}",
			"--b", "{introText: introText, subject: subject}",
			"--event", "{date: date, description: description, link: link, title: title, ctaText: ctaText, endTimeRaw: endTimeRaw, startTimeRaw: startTimeRaw}",
			"--role", "admin",
			"--test-batch-size", "1",
			"--test-email", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(adminV1EmailsSendWeeklyDigest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"admin:v1:emails", "send-weekly-digest",
			"--a.intro-text", "introText",
			"--a.subject", "subject",
			"--b.intro-text", "introText",
			"--b.subject", "subject",
			"--event.date", "date",
			"--event.description", "description",
			"--event.link", "link",
			"--event.title", "title",
			"--event.cta-text", "ctaText",
			"--event.end-time-raw", "endTimeRaw",
			"--event.start-time-raw", "startTimeRaw",
			"--role", "admin",
			"--test-batch-size", "1",
			"--test-email", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"a:\n" +
			"  introText: introText\n" +
			"  subject: subject\n" +
			"b:\n" +
			"  introText: introText\n" +
			"  subject: subject\n" +
			"events:\n" +
			"  - date: date\n" +
			"    description: description\n" +
			"    link: link\n" +
			"    title: title\n" +
			"    ctaText: ctaText\n" +
			"    endTimeRaw: endTimeRaw\n" +
			"    startTimeRaw: startTimeRaw\n" +
			"role: admin\n" +
			"testBatchSize: 1\n" +
			"testEmails:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"admin:v1:emails", "send-weekly-digest",
		)
	})
}
