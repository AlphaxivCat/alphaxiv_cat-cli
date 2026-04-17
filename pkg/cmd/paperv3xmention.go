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

var papersV3XMentionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Search for X (Twitter) mentions of a paper, filter for quality, and save to\ndatabase",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Usage:    "Paper group ID (UUID)",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "dry-run",
			Usage:    "If true, only logs results without saving to database",
			Default:  false,
			BodyPath: "dryRun",
		},
	},
	Action:          handlePapersV3XMentionsUpdate,
	HideHelpCommand: true,
}

var papersV3XMentionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete all X (Twitter) mentions for a paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Usage:    "Paper group ID (UUID)",
			Required: true,
		},
	},
	Action:          handlePapersV3XMentionsDelete,
	HideHelpCommand: true,
}

func handlePapersV3XMentionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-group-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3XMentionUpdateParams{}

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
	_, err = client.Papers.V3.XMentions.Update(
		ctx,
		cmd.Value("paper-group-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3:x-mentions update",
		Transform:      transform,
	})
}

func handlePapersV3XMentionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-group-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-group-id", unusedArgs[0])
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

	return client.Papers.V3.XMentions.Delete(ctx, cmd.Value("paper-group-id").(string), options...)
}
