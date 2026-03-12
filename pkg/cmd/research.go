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

var researchCreateProject = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-project",
	Usage:   "Create a research project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "folder",
			Usage:    "Prefill the project with a set of research papers.",
			BodyPath: "folder",
		},
		&requestflag.Flag[[]string]{
			Name:     "initial-paper",
			Usage:    "Add these papers to the group on init",
			BodyPath: "initialPapers",
		},
		&requestflag.Flag[string]{
			Name:     "parent-id",
			BodyPath: "parentId",
		},
	},
	Action:          handleResearchCreateProject,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"folder": {
		&requestflag.InnerFlag[string]{
			Name:       "folder.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "folder.delete",
			Usage:      "If set true, the folder is deleted after copying it",
			InnerField: "delete",
		},
	},
})

func handleResearchCreateProject(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.ResearchNewProjectParams{}

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
	_, err = client.Research.NewProject(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "research create-project", obj, format, transform)
}
