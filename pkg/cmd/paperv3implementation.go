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

var papersV3ImplementationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Add an implementation (AlphaXiv, Marimo, Author, or Other)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "implementation-type",
			Required: true,
			BodyPath: "implementationType",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handlePapersV3ImplementationsCreate,
	HideHelpCommand: true,
}

var papersV3ImplementationsList = cli.Command{
	Name:    "list",
	Usage:   "Get all implementations for a paper (AlphaXiv, Marimo, Author, and Other)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Required: true,
		},
	},
	Action:          handlePapersV3ImplementationsList,
	HideHelpCommand: true,
}

var papersV3ImplementationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an implementation (AlphaXiv, Marimo, Author, or Other)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "implementation-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Required:  true,
			QueryPath: "type",
		},
	},
	Action:          handlePapersV3ImplementationsDelete,
	HideHelpCommand: true,
}

func handlePapersV3ImplementationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-group-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3ImplementationNewParams{}

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
	_, err = client.Papers.V3.Implementations.New(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3:implementations create", obj, format, transform)
}

func handlePapersV3ImplementationsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.Implementations.List(ctx, cmd.Value("paper-group-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3:implementations list", obj, format, transform)
}

func handlePapersV3ImplementationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("implementation-id") && len(unusedArgs) > 0 {
		cmd.Set("implementation-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3ImplementationDeleteParams{
		PaperGroupID: cmd.Value("paper-group-id").(string),
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

	return client.Papers.V3.Implementations.Delete(
		ctx,
		cmd.Value("implementation-id").(string),
		params,
		options...,
	)
}
