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

var adminV1GetModeratorFeed = cli.Command{
	Name:    "get-moderator-feed",
	Usage:   "Get page of comments for moderator feed",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "feed-type",
			Usage:     `Allowed values: "all", "flagged", "recent".`,
			Default:   "all",
			QueryPath: "feedType",
		},
		&requestflag.Flag[string]{
			Name:      "page",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "page-size",
			QueryPath: "pageSize",
		},
	},
	Action:          handleAdminV1GetModeratorFeed,
	HideHelpCommand: true,
}

var adminV1LookupUserByEmail = cli.Command{
	Name:    "lookup-user-by-email",
	Usage:   "Look up a user by email address (admin only)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email",
			Required:  true,
			QueryPath: "email",
		},
	},
	Action:          handleAdminV1LookupUserByEmail,
	HideHelpCommand: true,
}

func handleAdminV1GetModeratorFeed(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AdminV1GetModeratorFeedParams{}

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
	_, err = client.Admin.V1.GetModeratorFeed(ctx, params, options...)
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
		Title:          "admin:v1 get-moderator-feed",
		Transform:      transform,
	})
}

func handleAdminV1LookupUserByEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AdminV1LookupUserByEmailParams{}

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
	_, err = client.Admin.V1.LookupUserByEmail(ctx, params, options...)
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
		Title:          "admin:v1 lookup-user-by-email",
		Transform:      transform,
	})
}
