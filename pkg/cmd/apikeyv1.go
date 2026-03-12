// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
	"github.com/stainless-sdks/alphaxiv_cat-go"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var apiKeysV1Create = cli.Command{
	Name:    "create",
	Usage:   "Create a new API key for the current user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "label",
			Required: true,
			BodyPath: "label",
		},
	},
	Action:          handleAPIKeysV1Create,
	HideHelpCommand: true,
}

var apiKeysV1List = cli.Command{
	Name:            "list",
	Usage:           "List API keys for the current user.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAPIKeysV1List,
	HideHelpCommand: true,
}

var apiKeysV1CreateImpersonation = cli.Command{
	Name:    "create-impersonation",
	Usage:   "Create a new API key for the current user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user",
			Required: true,
			BodyPath: "user",
		},
	},
	Action:          handleAPIKeysV1CreateImpersonation,
	HideHelpCommand: true,
}

var apiKeysV1Revoke = cli.Command{
	Name:    "revoke",
	Usage:   "Revoke an API key for the authenticated user. No-op if already revoked.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "api-key-id",
			Required: true,
		},
	},
	Action:          handleAPIKeysV1Revoke,
	HideHelpCommand: true,
}

func handleAPIKeysV1Create(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.APIKeyV1NewParams{}

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
	_, err = client.APIKeys.V1.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys:v1 create", obj, format, transform)
}

func handleAPIKeysV1List(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.APIKeys.V1.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys:v1 list", obj, format, transform)
}

func handleAPIKeysV1CreateImpersonation(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.APIKeyV1NewImpersonationParams{}

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
	_, err = client.APIKeys.V1.NewImpersonation(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys:v1 create-impersonation", obj, format, transform)
}

func handleAPIKeysV1Revoke(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("api-key-id") && len(unusedArgs) > 0 {
		cmd.Set("api-key-id", unusedArgs[0])
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
	_, err = client.APIKeys.V1.Revoke(ctx, alphaxivcat.APIKeyV1RevokeParamsAPIKeyID(cmd.Value("api-key-id").(string)), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys:v1 revoke", obj, format, transform)
}
