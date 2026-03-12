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

var commentsV2Delete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a comment by UUID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
	},
	Action:          handleCommentsV2Delete,
	HideHelpCommand: true,
}

var commentsV2Downvote = cli.Command{
	Name:    "downvote",
	Usage:   "Downvote a comment by UUID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
	},
	Action:          handleCommentsV2Downvote,
	HideHelpCommand: true,
}

var commentsV2Flag = cli.Command{
	Name:    "flag",
	Usage:   "Flag a comment for moderator review",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Required: true,
			BodyPath: "reason",
		},
	},
	Action:          handleCommentsV2Flag,
	HideHelpCommand: true,
}

var commentsV2ToggleEndorse = cli.Command{
	Name:    "toggle-endorse",
	Usage:   "Toggle author endorsement of a comment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
	},
	Action:          handleCommentsV2ToggleEndorse,
	HideHelpCommand: true,
}

var commentsV2Upvote = cli.Command{
	Name:    "upvote",
	Usage:   "Upvote a comment by UUID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
	},
	Action:          handleCommentsV2Upvote,
	HideHelpCommand: true,
}

func handleCommentsV2Delete(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
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

	return client.Comments.V2.Delete(ctx, cmd.Value("comment").(string), options...)
}

func handleCommentsV2Downvote(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
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

	return client.Comments.V2.Downvote(ctx, cmd.Value("comment").(string), options...)
}

func handleCommentsV2Flag(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.CommentV2FlagParams{}

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

	return client.Comments.V2.Flag(
		ctx,
		cmd.Value("comment").(string),
		params,
		options...,
	)
}

func handleCommentsV2ToggleEndorse(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
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

	return client.Comments.V2.ToggleEndorse(ctx, cmd.Value("comment").(string), options...)
}

func handleCommentsV2Upvote(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
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

	return client.Comments.V2.Upvote(ctx, cmd.Value("comment").(string), options...)
}
