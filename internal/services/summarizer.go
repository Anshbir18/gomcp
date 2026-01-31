package services

import (
	"context"
	"errors"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

const SummaryPrompt = `
You are a helpful assistant.
Summarize the following text in a clear and concise manner.
Focus only on the main content and ignore navigation or ads.
Make the Summary Usefull for a person and make it a bit detailed also.
`

// SummarizeChunks summarizes multiple text chunks and merges results
func SummarizeChunks(chunks []string) (string, error) {
	if len(chunks) == 0 {
		return "", errors.New("no text chunks provided")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OPENAI_API_KEY not set")
	}

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	var summaries []string

	for _, chunk := range chunks {
		resp, err := client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleSystem,
						Content: SummaryPrompt,
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: chunk,
					},
				},
				Temperature: 0.3,
			},
		)

		if err != nil {
			return "", err
		}

		summaries = append(summaries, resp.Choices[0].Message.Content)
	}

	// Merge all summaries into one
	return strings.Join(summaries, "\n\n"), nil
}
