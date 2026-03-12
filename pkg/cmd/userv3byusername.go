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

var usersV3ByUsernameGetProfilePage = cli.Command{
	Name:    "get-profile-page",
	Usage:   "This route is specifically for the Client's user page",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "username",
			Required: true,
		},
	},
	Action:          handleUsersV3ByUsernameGetProfilePage,
	HideHelpCommand: true,
}

var usersV3ByUsernameGetUser = cli.Command{
	Name:    "get-user",
	Usage:   "Retrieve a user's basic information given its username",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "username",
			Required: true,
		},
	},
	Action:          handleUsersV3ByUsernameGetUser,
	HideHelpCommand: true,
}

func handleUsersV3ByUsernameGetProfilePage(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("username") && len(unusedArgs) > 0 {
		cmd.Set("username", unusedArgs[0])
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
	_, err = client.Users.V3.ByUsername.GetProfilePage(ctx, cmd.Value("username").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3:by-username get-profile-page", obj, format, transform)
}

func handleUsersV3ByUsernameGetUser(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("username") && len(unusedArgs) > 0 {
		cmd.Set("username", unusedArgs[0])
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
	_, err = client.Users.V3.ByUsername.GetUser(ctx, cmd.Value("username").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3:by-username get-user", obj, format, transform)
}
