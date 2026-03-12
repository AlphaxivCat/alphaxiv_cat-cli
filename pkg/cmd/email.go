// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
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

var emailsKickoffCommentUpdate = cli.Command{
	Name:    "kickoff-comment-update",
	Usage:   "Kicks off a background job to send comment digest emails to users",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "User role to filter by",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "window",
			Usage:    "Time window in hours",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "custom",
			Usage:    "Whether to use custom digest",
			Required: true,
		},
	},
	Action:          handleEmailsKickoffCommentUpdate,
	HideHelpCommand: true,
}

var emailsKickoffGeneralUpdate = cli.Command{
	Name:    "kickoff-general-update",
	Usage:   "Kicks off a background job to send general digest emails to users",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "User role to filter by",
			Required: true,
		},
	},
	Action:          handleEmailsKickoffGeneralUpdate,
	HideHelpCommand: true,
}

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

var emailsProcessCommentUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "process-comment-update",
	Usage:   "Process comment digest email for a user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
			BodyPath: "userId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-content",
			BodyPath: "customContent",
		},
	},
	Action:          handleEmailsProcessCommentUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"custom-content": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "custom-content.events",
			InnerField: "events",
		},
		&requestflag.InnerFlag[string]{
			Name:       "custom-content.intro-text",
			InnerField: "introText",
		},
		&requestflag.InnerFlag[string]{
			Name:       "custom-content.subject",
			InnerField: "subject",
		},
	},
})

var emailsProcessGeneralUpdate = cli.Command{
	Name:    "process-general-update",
	Usage:   "Process general digest email for a user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
			BodyPath: "userId",
		},
	},
	Action:          handleEmailsProcessGeneralUpdate,
	HideHelpCommand: true,
}

func handleEmailsCaptureBouncedEmails(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.EmailCaptureBouncedEmailsParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.CaptureBouncedEmails(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "emails capture-bounced-emails", obj, format, transform)
}

func handleEmailsCaptureResendBouncedEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.EmailCaptureResendBouncedEmailParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.CaptureResendBouncedEmail(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "emails capture-resend-bounced-email", obj, format, transform)
}

func handleEmailsKickoffCommentUpdate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("custom") && len(unusedArgs) > 0 {
		cmd.Set("custom", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.EmailKickoffCommentUpdateParams{
		Role:   cmd.Value("role").(string),
		Window: cmd.Value("window").(string),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.KickoffCommentUpdate(
		ctx,
		alphaxivcat.EmailKickoffCommentUpdateParamsCustom(cmd.Value("custom").(string)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "emails kickoff-comment-update", obj, format, transform)
}

func handleEmailsKickoffGeneralUpdate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("role") && len(unusedArgs) > 0 {
		cmd.Set("role", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.KickoffGeneralUpdate(ctx, cmd.Value("role").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "emails kickoff-general-update", obj, format, transform)
}

func handleEmailsProcessBouncedEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.EmailProcessBouncedEmailParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.ProcessBouncedEmail(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "emails process-bounced-email", obj, format, transform)
}

func handleEmailsProcessCommentUpdate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.EmailProcessCommentUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.ProcessCommentUpdate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "emails process-comment-update", obj, format, transform)
}

func handleEmailsProcessGeneralUpdate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.EmailProcessGeneralUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Emails.ProcessGeneralUpdate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "emails process-general-update", obj, format, transform)
}
