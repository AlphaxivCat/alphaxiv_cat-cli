// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/urfave/cli/v3"
)

var adminV1EmailsSendMonthlyDigest = cli.Command{
	Name:    "send-monthly-digest",
	Usage:   "Queue monthly digest emails to users via Upstash",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "Filter by user role",
			BodyPath: "role",
		},
	},
	Action:          handleAdminV1EmailsSendMonthlyDigest,
	HideHelpCommand: true,
}

var adminV1EmailsSendWeeklyDigest = requestflag.WithInnerFlags(cli.Command{
	Name:    "send-weekly-digest",
	Usage:   "Queue weekly digest emails to users via Upstash",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "event",
			Usage:    "Custom events to include",
			BodyPath: "events",
		},
		&requestflag.Flag[string]{
			Name:     "intro-text",
			Usage:    "Custom intro message",
			BodyPath: "introText",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "Filter by user role",
			BodyPath: "role",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Custom email subject",
			BodyPath: "subject",
		},
	},
	Action:          handleAdminV1EmailsSendWeeklyDigest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"event": {
		&requestflag.InnerFlag[string]{
			Name:       "event.date",
			InnerField: "date",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.link",
			InnerField: "link",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.title",
			InnerField: "title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.end-time-raw",
			InnerField: "endTimeRaw",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.start-time-raw",
			InnerField: "startTimeRaw",
		},
	},
})

func handleAdminV1EmailsSendMonthlyDigest(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AdminV1EmailSendMonthlyDigestParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Admin.V1.Emails.SendMonthlyDigest(ctx, params, options...)
}

func handleAdminV1EmailsSendWeeklyDigest(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AdminV1EmailSendWeeklyDigestParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Admin.V1.Emails.SendWeeklyDigest(ctx, params, options...)
}
