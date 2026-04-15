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

var foldersV3Create = cli.Command{
	Name:    "create",
	Usage:   "Create a new folder for the current user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-name",
			Required: true,
			BodyPath: "folderName",
		},
	},
	Action:          handleFoldersV3Create,
	HideHelpCommand: true,
}

var foldersV3List = cli.Command{
	Name:            "list",
	Usage:           "Get all folders for the current user",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleFoldersV3List,
	HideHelpCommand: true,
}

var foldersV3Delete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a folder",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-id",
			Required: true,
		},
	},
	Action:          handleFoldersV3Delete,
	HideHelpCommand: true,
}

var foldersV3AddPapers = cli.Command{
	Name:    "add-papers",
	Usage:   "Add papers to a folder (without removing from other folders)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "paper-group-id",
			BodyPath: "paperGroupIds",
		},
		&requestflag.Flag[[]string]{
			Name:     "universal-id",
			BodyPath: "universalIds",
		},
	},
	Action:          handleFoldersV3AddPapers,
	HideHelpCommand: true,
}

var foldersV3MovePapers = cli.Command{
	Name:    "move-papers",
	Usage:   "Move papers from source folder to destination folder",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "paper-group-id",
			Required: true,
			BodyPath: "paperGroupIds",
		},
		&requestflag.Flag[string]{
			Name:     "from-folder-id",
			BodyPath: "fromFolderId",
		},
	},
	Action:          handleFoldersV3MovePapers,
	HideHelpCommand: true,
}

var foldersV3RemovePapers = cli.Command{
	Name:    "remove-papers",
	Usage:   "Remove papers from a specific folder",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "paper-group-id",
			Required: true,
			BodyPath: "paperGroupIds",
		},
	},
	Action:          handleFoldersV3RemovePapers,
	HideHelpCommand: true,
}

var foldersV3ToggleSharing = cli.Command{
	Name:    "toggle-sharing",
	Usage:   "Toggle whether a folder and all its descendant folders are shared or not",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "sharing-status",
			Usage:    `Allowed values: "not_shared", "shared".`,
			Required: true,
			BodyPath: "sharingStatus",
		},
	},
	Action:          handleFoldersV3ToggleSharing,
	HideHelpCommand: true,
}

var foldersV3UpdateName = cli.Command{
	Name:    "update-name",
	Usage:   "Rename a folder",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleFoldersV3UpdateName,
	HideHelpCommand: true,
}

var foldersV3UpdateParent = cli.Command{
	Name:    "update-parent",
	Usage:   "Update a folder's parent (for nesting)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "folder-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "parent-id",
			Required: true,
			BodyPath: "parentId",
		},
	},
	Action:          handleFoldersV3UpdateParent,
	HideHelpCommand: true,
}

func handleFoldersV3Create(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.FolderV3NewParams{}

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
	_, err = client.Folders.V3.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "folders:v3 create", obj, format, explicitFormat, transform)
}

func handleFoldersV3List(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Folders.V3.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "folders:v3 list", obj, format, explicitFormat, transform)
}

func handleFoldersV3Delete(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("folder-id") && len(unusedArgs) > 0 {
		cmd.Set("folder-id", unusedArgs[0])
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

	return client.Folders.V3.Delete(ctx, cmd.Value("folder-id").(string), options...)
}

func handleFoldersV3AddPapers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("folder-id") && len(unusedArgs) > 0 {
		cmd.Set("folder-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.FolderV3AddPapersParams{}

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
	_, err = client.Folders.V3.AddPapers(
		ctx,
		cmd.Value("folder-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "folders:v3 add-papers", obj, format, explicitFormat, transform)
}

func handleFoldersV3MovePapers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("folder-id") && len(unusedArgs) > 0 {
		cmd.Set("folder-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.FolderV3MovePapersParams{}

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
	_, err = client.Folders.V3.MovePapers(
		ctx,
		cmd.Value("folder-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "folders:v3 move-papers", obj, format, explicitFormat, transform)
}

func handleFoldersV3RemovePapers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("folder-id") && len(unusedArgs) > 0 {
		cmd.Set("folder-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.FolderV3RemovePapersParams{}

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

	return client.Folders.V3.RemovePapers(
		ctx,
		cmd.Value("folder-id").(string),
		params,
		options...,
	)
}

func handleFoldersV3ToggleSharing(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("folder-id") && len(unusedArgs) > 0 {
		cmd.Set("folder-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.FolderV3ToggleSharingParams{}

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
	_, err = client.Folders.V3.ToggleSharing(
		ctx,
		cmd.Value("folder-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "folders:v3 toggle-sharing", obj, format, explicitFormat, transform)
}

func handleFoldersV3UpdateName(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("folder-id") && len(unusedArgs) > 0 {
		cmd.Set("folder-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.FolderV3UpdateNameParams{}

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

	return client.Folders.V3.UpdateName(
		ctx,
		cmd.Value("folder-id").(string),
		params,
		options...,
	)
}

func handleFoldersV3UpdateParent(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("folder-id") && len(unusedArgs) > 0 {
		cmd.Set("folder-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.FolderV3UpdateParentParams{}

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

	return client.Folders.V3.UpdateParent(
		ctx,
		cmd.Value("folder-id").(string),
		params,
		options...,
	)
}
