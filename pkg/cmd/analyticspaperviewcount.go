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

var analyticsPaperViewCountIngestEvent = cli.Command{
	Name:    "ingest-event",
	Usage:   "Track a paper view event for analytics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Usage:    "Paper ID to track view for",
			Required: true,
			BodyPath: "paperId",
		},
		&requestflag.Flag[string]{
			Name:     "created-at",
			Usage:    "Optional timestamp for the view event",
			BodyPath: "createdAt",
		},
		&requestflag.Flag[any]{
			Name:     "user-id",
			Usage:    "Optional user ID who viewed the paper",
			BodyPath: "userId",
		},
	},
	Action:          handleAnalyticsPaperViewCountIngestEvent,
	HideHelpCommand: true,
}

var analyticsPaperViewCountKickoffJob = cli.Command{
	Name:    "kickoff-job",
	Usage:   "Kicks off a background job to aggregate paper view counts",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Time period filter: 'all' or number of days",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "like",
			Usage:    "Optional filter for likes",
			BodyPath: "like",
		},
	},
	Action:          handleAnalyticsPaperViewCountKickoffJob,
	HideHelpCommand: true,
}

var analyticsPaperViewCountProcessJob = cli.Command{
	Name:    "process-job",
	Usage:   "Process view count aggregation for a specific paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Usage:    "Paper ID to process view counts for",
			Required: true,
			BodyPath: "paperId",
		},
		&requestflag.Flag[string]{
			Name:     "publication-date",
			Usage:    "Publication date for age decay calculation",
			Required: true,
			BodyPath: "publicationDate",
		},
		&requestflag.Flag[bool]{
			Name:     "like",
			Usage:    "Whether to add noise to votes",
			BodyPath: "like",
		},
	},
	Action:          handleAnalyticsPaperViewCountProcessJob,
	HideHelpCommand: true,
}

func handleAnalyticsPaperViewCountIngestEvent(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AnalyticsPaperViewCountIngestEventParams{}

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
	_, err = client.Analytics.PaperViewCount.IngestEvent(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "analytics:paper-view-count ingest-event", obj, format, explicitFormat, transform)
}

func handleAnalyticsPaperViewCountKickoffJob(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AnalyticsPaperViewCountKickoffJobParams{}

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
	_, err = client.Analytics.PaperViewCount.KickoffJob(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "analytics:paper-view-count kickoff-job", obj, format, explicitFormat, transform)
}

func handleAnalyticsPaperViewCountProcessJob(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AnalyticsPaperViewCountProcessJobParams{}

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
	_, err = client.Analytics.PaperViewCount.ProcessJob(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "analytics:paper-view-count process-job", obj, format, explicitFormat, transform)
}
