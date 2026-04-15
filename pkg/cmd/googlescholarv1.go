// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var googleScholarV1Connect = cli.Command{
	Name:    "connect",
	Usage:   "Start connecting Google Scholar profile to user account, or replace a pending\nconnection with a different profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "google-scholar-id",
			Usage:    "Google Scholar ID",
			Required: true,
			BodyPath: "googleScholarId",
		},
	},
	Action:          handleGoogleScholarV1Connect,
	HideHelpCommand: true,
}

var googleScholarV1DeleteConnection = cli.Command{
	Name:            "delete-connection",
	Usage:           "Remove Google Scholar ID and queued papers from a user",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleGoogleScholarV1DeleteConnection,
	HideHelpCommand: true,
}

var googleScholarV1GetReport = cli.Command{
	Name:            "get-report",
	Usage:           "Get a full report of a user's Google Scholar sync (including lists of papers)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleGoogleScholarV1GetReport,
	HideHelpCommand: true,
}

var googleScholarV1GetStatus = cli.Command{
	Name:            "get-status",
	Usage:           "Get status of user's Google Scholar sync",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleGoogleScholarV1GetStatus,
	HideHelpCommand: true,
}

var googleScholarV1Resync = cli.Command{
	Name:    "resync",
	Usage:   "Start a new Google Scholar sync for this user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    `Allowed values: "all", "new".`,
			Required: true,
		},
	},
	Action:          handleGoogleScholarV1Resync,
	HideHelpCommand: true,
}

var googleScholarV1SetEmail = cli.Command{
	Name:    "set-email",
	Usage:   "Set the local-part of a user's Google Scholar institutional email and send a\nverification email to that address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "local-part",
			Usage:    "Institutional email local-part",
			Required: true,
			BodyPath: "localPart",
		},
		&requestflag.Flag[bool]{
			Name:     "verify-account-email",
			Usage:    "Send verification code to the user's primary email instead of the institutional email. Admin only.",
			Required: true,
			BodyPath: "verifyAccountEmail",
		},
	},
	Action:          handleGoogleScholarV1SetEmail,
	HideHelpCommand: true,
}

var googleScholarV1Sync = cli.Command{
	Name:            "sync",
	Usage:           "Make some progress syncing a user's Google Scholar papers",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleGoogleScholarV1Sync,
	HideHelpCommand: true,
}

var googleScholarV1VerifyEmail = cli.Command{
	Name:    "verify-email",
	Usage:   "Verify a user's Google Scholar email by entering the code sent to that email",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "code",
			Required: true,
			BodyPath: "code",
		},
	},
	Action:          handleGoogleScholarV1VerifyEmail,
	HideHelpCommand: true,
}

func handleGoogleScholarV1Connect(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.GoogleScholarV1ConnectParams{}

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
	_, err = client.GoogleScholar.V1.Connect(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "google-scholar:v1 connect", obj, format, explicitFormat, transform)
}

func handleGoogleScholarV1DeleteConnection(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	return client.GoogleScholar.V1.DeleteConnection(ctx, options...)
}

func handleGoogleScholarV1GetReport(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.GoogleScholar.V1.GetReport(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "google-scholar:v1 get-report", obj, format, explicitFormat, transform)
}

func handleGoogleScholarV1GetStatus(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.GoogleScholar.V1.GetStatus(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "google-scholar:v1 get-status", obj, format, explicitFormat, transform)
}

func handleGoogleScholarV1Resync(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("mode") && len(unusedArgs) > 0 {
		cmd.Set("mode", unusedArgs[0])
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
	_, err = client.GoogleScholar.V1.Resync(ctx, alphaxivcat.GoogleScholarV1ResyncParamsMode(cmd.Value("mode").(string)), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "google-scholar:v1 resync", obj, format, explicitFormat, transform)
}

func handleGoogleScholarV1SetEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.GoogleScholarV1SetEmailParams{}

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

	return client.GoogleScholar.V1.SetEmail(ctx, params, options...)
}

func handleGoogleScholarV1Sync(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.GoogleScholar.V1.Sync(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "google-scholar:v1 sync", obj, format, explicitFormat, transform)
}

func handleGoogleScholarV1VerifyEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.GoogleScholarV1VerifyEmailParams{}

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
	_, err = client.GoogleScholar.V1.VerifyEmail(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "google-scholar:v1 verify-email", obj, format, explicitFormat, transform)
}
