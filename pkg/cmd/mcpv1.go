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

var mcpV1EstablishConnection = cli.Command{
	Name:    "establish-connection",
	Usage:   "Establish SSE connection for server-to-client MCP messages",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleMcpV1EstablishConnection,
	HideHelpCommand: true,
}

var mcpV1SendMessage = cli.Command{
	Name:    "send-message",
	Usage:   "Send a message to the MCP server",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleMcpV1SendMessage,
	HideHelpCommand: true,
}

var mcpV1TerminateSession = cli.Command{
	Name:            "terminate-session",
	Usage:           "Terminate an MCP session",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMcpV1TerminateSession,
	HideHelpCommand: true,
}

func handleMcpV1EstablishConnection(ctx context.Context, cmd *cli.Command) error {
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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.Mcp.V1.EstablishConnectionStreaming(ctx, options...)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(os.Stdout, "mcp:v1 establish-connection", stream, format, transform, maxItems)
}

func handleMcpV1SendMessage(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.McpV1SendMessageParams{}

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
	_, err = client.Mcp.V1.SendMessage(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mcp:v1 send-message", obj, format, transform)
}

func handleMcpV1TerminateSession(ctx context.Context, cmd *cli.Command) error {
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

	return client.Mcp.V1.TerminateSession(ctx, options...)
}
