package services

import "strings"

const MaxChunkSize = 3000

// PrepareTextForSummarization cleans and chunks text
func PrepareTextForSummarization(text string) []string {
	text = normalizeText(text)

	if len(text) <= MaxChunkSize {
		return []string{text}
	}

	return chunkText(text, MaxChunkSize)
}

func normalizeText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.Join(strings.Fields(text), " ")
	return text
}

func chunkText(text string, chunkSize int) []string {
	var chunks []string

	for start := 0; start < len(text); start += chunkSize {
		end := start + chunkSize
		if end > len(text) {
			end = len(text)
		}
		chunks = append(chunks, text[start:end])
	}

	return chunks
}
