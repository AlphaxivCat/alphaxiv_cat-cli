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

var commentsV2ModeratorSendFeedback = cli.Command{
	Name:    "send-feedback",
	Usage:   "Send moderator feedback to comment author",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "message",
			Required: true,
			BodyPath: "message",
		},
	},
	Action:          handleCommentsV2ModeratorSendFeedback,
	HideHelpCommand: true,
}

var commentsV2ModeratorToggleFlags = cli.Command{
	Name:    "toggle-flags",
	Usage:   "Toggle one (or more) of a set of comment moderation flags",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "addressed",
			BodyPath: "addressed",
		},
		&requestflag.Flag[bool]{
			Name:     "closed",
			BodyPath: "closed",
		},
		&requestflag.Flag[bool]{
			Name:     "flag-addressed",
			BodyPath: "flagAddressed",
		},
	},
	Action:          handleCommentsV2ModeratorToggleFlags,
	HideHelpCommand: true,
}

func handleCommentsV2ModeratorSendFeedback(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.CommentV2ModeratorSendFeedbackParams{}

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

	return client.Comments.V2.Moderator.SendFeedback(
		ctx,
		cmd.Value("comment").(string),
		params,
		options...,
	)
}

func handleCommentsV2ModeratorToggleFlags(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.CommentV2ModeratorToggleFlagsParams{}

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
	_, err = client.Comments.V2.Moderator.ToggleFlags(
		ctx,
		cmd.Value("comment").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "comments:v2:moderator toggle-flags", obj, format, explicitFormat, transform)
}
