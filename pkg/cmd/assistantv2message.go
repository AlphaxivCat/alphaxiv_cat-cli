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

var assistantV2MessagesList = cli.Command{
	Name:    "list",
	Usage:   "Get all messages for an llm chat",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "llm-chat",
			Required: true,
		},
	},
	Action:          handleAssistantV2MessagesList,
	HideHelpCommand: true,
}

var assistantV2MessagesSelect = cli.Command{
	Name:    "select",
	Usage:   "Select an llm chat message by id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "llm-chat",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "message",
			Required: true,
		},
	},
	Action:          handleAssistantV2MessagesSelect,
	HideHelpCommand: true,
}

var assistantV2MessagesSetFeedback = requestflag.WithInnerFlags(cli.Command{
	Name:    "set-feedback",
	Usage:   "Set or update feedback for a message",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "feedback",
			Required: true,
			BodyPath: "feedback",
		},
	},
	Action:          handleAssistantV2MessagesSetFeedback,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"feedback": {
		&requestflag.InnerFlag[string]{
			Name:       "feedback.type",
			InnerField: "type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "feedback.category",
			InnerField: "category",
		},
		&requestflag.InnerFlag[any]{
			Name:       "feedback.details",
			InnerField: "details",
		},
	},
})

func handleAssistantV2MessagesList(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("llm-chat") && len(unusedArgs) > 0 {
		cmd.Set("llm-chat", unusedArgs[0])
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
	_, err = client.Assistant.V2.Messages.List(ctx, cmd.Value("llm-chat").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "assistant:v2:messages list", obj, format, transform)
}

func handleAssistantV2MessagesSelect(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message") && len(unusedArgs) > 0 {
		cmd.Set("message", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AssistantV2MessageSelectParams{
		LlmChat: cmd.Value("llm-chat").(string),
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

	return client.Assistant.V2.Messages.Select(
		ctx,
		cmd.Value("message").(string),
		params,
		options...,
	)
}

func handleAssistantV2MessagesSetFeedback(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AssistantV2MessageSetFeedbackParams{}

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

	return client.Assistant.V2.Messages.SetFeedback(
		ctx,
		cmd.Value("message-id").(string),
		params,
		options...,
	)
}
