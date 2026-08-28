// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var emailsCaptureBouncedEmails = cli.Command{
	Name:    "capture-bounced-emails",
	Usage:   "Receives bounce notifications from AWS SES via SNS",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message",
			Usage:    "Stringified JSON message containing bounce/complaint data",
			Required: true,
			BodyPath: "Message",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "SNS notification type",
			Required: true,
			BodyPath: "Type",
		},
	},
	Action:          handleEmailsCaptureBouncedEmails,
	HideHelpCommand: true,
}

var emailsCaptureResendBouncedEmail = requestflag.WithInnerFlags(cli.Command{
	Name:    "capture-resend-bounced-email",
	Usage:   "Receives bounce notifications from Resend",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Usage:    "Event data containing bounced emails",
			Required: true,
			BodyPath: "data",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Event type from Resend",
			Required: true,
			BodyPath: "type",
		},
	},
	Action:          handleEmailsCaptureResendBouncedEmail,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data": {
		&requestflag.InnerFlag[[]string]{
			Name:       "data.to",
			Usage:      "Bounced email addresses",
			InnerField: "to",
		},
	},
})

var emailsProcessBouncedEmail = cli.Command{
	Name:    "process-bounced-email",
	Usage:   "Process a bounced email and update user preferences",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "email",
			Required: true,
			BodyPath: "email",
		},
	},
	Action:          handleEmailsProcessBouncedEmail,
	HideHelpCommand: true,
}

func handleEmailsCaptureBouncedEmails(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := alphaxivcat.EmailCaptureBouncedEmailsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.CaptureBouncedEmails(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "emails capture-bounced-emails",
		Transform:      transform,
	})
}

func handleEmailsCaptureResendBouncedEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := alphaxivcat.EmailCaptureResendBouncedEmailParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.CaptureResendBouncedEmail(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "emails capture-resend-bounced-email",
		Transform:      transform,
	})
}

func handleEmailsProcessBouncedEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := alphaxivcat.EmailProcessBouncedEmailParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.ProcessBouncedEmail(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "emails process-bounced-email",
		Transform:      transform,
	})
}
