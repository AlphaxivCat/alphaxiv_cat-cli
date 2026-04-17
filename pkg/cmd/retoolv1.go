// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var retoolV1GetCumulativeUsers = cli.Command{
	Name:            "get-cumulative-users",
	Usage:           "Get cumulative user count over time (all time)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetCumulativeUsers,
	HideHelpCommand: true,
}

var retoolV1GetDailyConversations = cli.Command{
	Name:            "get-daily-conversations",
	Usage:           "Get daily conversation counts by variant (all time)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetDailyConversations,
	HideHelpCommand: true,
}

var retoolV1GetDailyNewAccounts = cli.Command{
	Name:            "get-daily-new-accounts",
	Usage:           "Get daily new account counts (all time)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetDailyNewAccounts,
	HideHelpCommand: true,
}

var retoolV1GetDailyUserChatMessages = cli.Command{
	Name:            "get-daily-user-chat-messages",
	Usage:           "Get daily user chat message counts by variant (all time)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetDailyUserChatMessages,
	HideHelpCommand: true,
}

var retoolV1GetTotalCommentCount = cli.Command{
	Name:            "get-total-comment-count",
	Usage:           "Get total count of all comments",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetTotalCommentCount,
	HideHelpCommand: true,
}

var retoolV1GetTotalPaperCount = cli.Command{
	Name:            "get-total-paper-count",
	Usage:           "Get total count of public, non-blocked papers",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetTotalPaperCount,
	HideHelpCommand: true,
}

var retoolV1GetTotalPrivateNotesCount = cli.Command{
	Name:            "get-total-private-notes-count",
	Usage:           "Get total count of private notes (comments with tag=personal)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetTotalPrivateNotesCount,
	HideHelpCommand: true,
}

var retoolV1GetTotalUserCount = cli.Command{
	Name:            "get-total-user-count",
	Usage:           "Get total count of non-deleted users",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetTotalUserCount,
	HideHelpCommand: true,
}

var retoolV1GetWeeklyMessageCountsByUser = cli.Command{
	Name:            "get-weekly-message-counts-by-user",
	Usage:           "Get assistant input message counts per user for the last week, sorted by count\ndescending",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetWeeklyMessageCountsByUser,
	HideHelpCommand: true,
}

var retoolV1GetWeeklyPrivateNotes = cli.Command{
	Name:            "get-weekly-private-notes",
	Usage:           "Get weekly private note counts (all time)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetWeeklyPrivateNotes,
	HideHelpCommand: true,
}

var retoolV1GetWeeklyPublicComments = cli.Command{
	Name:            "get-weekly-public-comments",
	Usage:           "Get weekly public comment counts (all time)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRetoolV1GetWeeklyPublicComments,
	HideHelpCommand: true,
}

func handleRetoolV1GetCumulativeUsers(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetCumulativeUsers(ctx, options...)
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
		Title:          "retool:v1 get-cumulative-users",
		Transform:      transform,
	})
}

func handleRetoolV1GetDailyConversations(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetDailyConversations(ctx, options...)
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
		Title:          "retool:v1 get-daily-conversations",
		Transform:      transform,
	})
}

func handleRetoolV1GetDailyNewAccounts(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetDailyNewAccounts(ctx, options...)
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
		Title:          "retool:v1 get-daily-new-accounts",
		Transform:      transform,
	})
}

func handleRetoolV1GetDailyUserChatMessages(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetDailyUserChatMessages(ctx, options...)
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
		Title:          "retool:v1 get-daily-user-chat-messages",
		Transform:      transform,
	})
}

func handleRetoolV1GetTotalCommentCount(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetTotalCommentCount(ctx, options...)
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
		Title:          "retool:v1 get-total-comment-count",
		Transform:      transform,
	})
}

func handleRetoolV1GetTotalPaperCount(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetTotalPaperCount(ctx, options...)
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
		Title:          "retool:v1 get-total-paper-count",
		Transform:      transform,
	})
}

func handleRetoolV1GetTotalPrivateNotesCount(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetTotalPrivateNotesCount(ctx, options...)
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
		Title:          "retool:v1 get-total-private-notes-count",
		Transform:      transform,
	})
}

func handleRetoolV1GetTotalUserCount(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetTotalUserCount(ctx, options...)
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
		Title:          "retool:v1 get-total-user-count",
		Transform:      transform,
	})
}

func handleRetoolV1GetWeeklyMessageCountsByUser(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetWeeklyMessageCountsByUser(ctx, options...)
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
		Title:          "retool:v1 get-weekly-message-counts-by-user",
		Transform:      transform,
	})
}

func handleRetoolV1GetWeeklyPrivateNotes(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetWeeklyPrivateNotes(ctx, options...)
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
		Title:          "retool:v1 get-weekly-private-notes",
		Transform:      transform,
	})
}

func handleRetoolV1GetWeeklyPublicComments(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Retool.V1.GetWeeklyPublicComments(ctx, options...)
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
		Title:          "retool:v1 get-weekly-public-comments",
		Transform:      transform,
	})
}
