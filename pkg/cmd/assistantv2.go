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

var assistantV2Chat = requestflag.WithInnerFlags(cli.Command{
	Name:    "chat",
	Usage:   "Send a message to the AI assistant and receive streaming responses",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "file",
			Required: true,
			BodyPath: "files",
		},
		&requestflag.Flag[any]{
			Name:     "llm-chat-id",
			Required: true,
			BodyPath: "llmChatId",
		},
		&requestflag.Flag[string]{
			Name:     "message",
			Required: true,
			BodyPath: "message",
		},
		&requestflag.Flag[any]{
			Name:     "paper-version-id",
			Required: true,
			BodyPath: "paperVersionId",
		},
		&requestflag.Flag[any]{
			Name:     "parent-message-id",
			Required: true,
			BodyPath: "parentMessageId",
		},
		&requestflag.Flag[any]{
			Name:     "selection-page-range",
			Required: true,
			BodyPath: "selectionPageRange",
		},
		&requestflag.Flag[any]{
			Name:     "thinking",
			Required: true,
			BodyPath: "thinking",
		},
		&requestflag.Flag[string]{
			Name:     "web-search",
			Usage:    `Allowed values: "off", "full".`,
			Required: true,
			BodyPath: "webSearch",
		},
		&requestflag.Flag[string]{
			Name:     "assistant-variant",
			Usage:    `Allowed values: "homepage", "paper", "folder", "landing", "folder-add-papers".`,
			BodyPath: "assistantVariant",
		},
		&requestflag.Flag[bool]{
			Name:     "folder-add-papers",
			BodyPath: "folderAddPapers",
		},
		&requestflag.Flag[string]{
			Name:     "folder-id",
			BodyPath: "folderId",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    `Allowed values: "claude-opus-4.5", "claude-opus-4.6", "claude-opus-4.7", "claude-sonnet-4.5", "claude-sonnet-4.6", "gemini-2.5-flash", "gemini-2.5-pro", "gemini-3-flash", "gemini-3.1-pro", "glm-5-turbo", "glm-5.1", "gpt-5", "gpt-5.1", "gpt-5.2", "gpt-5.4", "gpt-5.4-mini", "gpt-5.4-nano", "kimi-k2.5", "kimi-k2.6", "mercury-2", "minimax-m2.5", "minimax-m2.7", "qwen-3.5", "fast", "smart", "pro", "gemini-3-pro", "claude-4.5-sonnet", "claude-4.6-sonnet".`,
			Default:  "fast",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "plan",
			Usage:    `Allowed values: "free", "pro".`,
			Default:  "free",
			BodyPath: "plan",
		},
		&requestflag.Flag[string]{
			Name:     "signature",
			BodyPath: "signature",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAssistantV2Chat,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"file": {
		&requestflag.InnerFlag[string]{
			Name:       "file.content-type",
			InnerField: "contentType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "file.url",
			InnerField: "url",
		},
	},
})

var assistantV2DeleteChat = cli.Command{
	Name:    "delete-chat",
	Usage:   "Delete an llm chat by id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "llm-chat",
			Required: true,
		},
	},
	Action:          handleAssistantV2DeleteChat,
	HideHelpCommand: true,
}

var assistantV2EditChat = cli.Command{
	Name:    "edit-chat",
	Usage:   "Updates properties on an LlmChat. Currently only supports title",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "llm-chat",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
	},
	Action:          handleAssistantV2EditChat,
	HideHelpCommand: true,
}

var assistantV2GetChats = cli.Command{
	Name:    "get-chats",
	Usage:   "Get llm chats for this user, filtered by variant, and optionally by paper\nversion",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "folder",
			QueryPath: "folder",
		},
		&requestflag.Flag[string]{
			Name:      "paper-version",
			QueryPath: "paperVersion",
		},
		&requestflag.Flag[string]{
			Name:      "variant",
			Usage:     `Allowed values: "homepage", "paper", "folder".`,
			QueryPath: "variant",
		},
	},
	Action:          handleAssistantV2GetChats,
	HideHelpCommand: true,
}

var assistantV2GetURLMetadata = cli.Command{
	Name:    "get-url-metadata",
	Usage:   "Fetch metadata (title and favicon) from a given URL",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "The URL to fetch metadata from",
			Required:  true,
			QueryPath: "url",
		},
	},
	Action:          handleAssistantV2GetURLMetadata,
	HideHelpCommand: true,
}

func handleAssistantV2Chat(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AssistantV2ChatParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	stream := client.Assistant.V2.ChatStreaming(ctx, params, options...)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(stream, maxItems, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "assistant:v2 chat",
		Transform:      transform,
	})
}

func handleAssistantV2DeleteChat(ctx context.Context, cmd *cli.Command) error {
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

	return client.Assistant.V2.DeleteChat(ctx, cmd.Value("llm-chat").(string), options...)
}

func handleAssistantV2EditChat(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("llm-chat") && len(unusedArgs) > 0 {
		cmd.Set("llm-chat", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AssistantV2EditChatParams{}

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

	return client.Assistant.V2.EditChat(
		ctx,
		cmd.Value("llm-chat").(string),
		params,
		options...,
	)
}

func handleAssistantV2GetChats(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AssistantV2GetChatsParams{}

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
	_, err = client.Assistant.V2.GetChats(ctx, params, options...)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "assistant:v2 get-chats",
		Transform:      transform,
	})
}

func handleAssistantV2GetURLMetadata(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.AssistantV2GetURLMetadataParams{}

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
	_, err = client.Assistant.V2.GetURLMetadata(ctx, params, options...)
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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "assistant:v2 get-url-metadata",
		Transform:      transform,
	})
}
