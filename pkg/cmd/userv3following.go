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

var usersV3FollowingList = cli.Command{
	Name:    "list",
	Usage:   "List the users that the given user is following",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleUsersV3FollowingList,
	HideHelpCommand: true,
}

func handleUsersV3FollowingList(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Users.V3.Following.List(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3:following list", obj, format, transform)
}
