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

var papersKickoffDailyGitHubStarsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Kickoff background job to update daily GitHub stars with max limit",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "max",
			Required: true,
		},
	},
	Action:          handlePapersKickoffDailyGitHubStarsUpdate,
	HideHelpCommand: true,
}

var papersKickoffDailyGitHubStarsKickoffDailyGitHubStars = cli.Command{
	Name:            "kickoff-daily-github-stars",
	Usage:           "Kickoff background job to update daily GitHub stars",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersKickoffDailyGitHubStarsKickoffDailyGitHubStars,
	HideHelpCommand: true,
}

func handlePapersKickoffDailyGitHubStarsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("max") && len(unusedArgs) > 0 {
		cmd.Set("max", unusedArgs[0])
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
	_, err = client.Papers.KickoffDailyGitHubStars.Update(ctx, cmd.Value("max").(string), options...)
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
		Title:          "papers:kickoff-daily-github-stars update",
		Transform:      transform,
	})
}

func handlePapersKickoffDailyGitHubStarsKickoffDailyGitHubStars(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.KickoffDailyGitHubStars.KickoffDailyGitHubStars(ctx, options...)
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
		Title:          "papers:kickoff-daily-github-stars kickoff-daily-github-stars",
		Transform:      transform,
	})
}
