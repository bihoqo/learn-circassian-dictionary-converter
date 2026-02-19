package code

import (
	"encoding/json"
	"fmt"
	"learn-circassian-helper/modals"
	"learn-circassian-helper/utils"
	"os"
	"path/filepath"
	"strings"
)

// CallConvertPhase03ToPhase04 reads all Phase 03 HTML JSON files and merges
// them into a single key-value database where each word maps to an array of
// dictionary entries from different sources.
func CallConvertPhase03ToPhase04() {
	srcDir := "content/phase-03-html-data"
	distDir := "content/phase-04-merged-database"
	distPath := filepath.Join(distDir, "merged-database.json")

	if err := os.MkdirAll(distDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create output directory: %v", err))
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to read source directory: %v", err))
	}

	merged := make(map[string][]modals.MergedDictEntry)
	dictionaries := make([]modals.DictionaryInfo, 0)
	seenDictIDs := make(map[int]bool)

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(srcDir, entry.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading %s: %s\n", filePath, err)
			continue
		}

		var dictObj modals.DictObjectHTML
		if err := json.Unmarshal(data, &dictObj); err != nil {
			fmt.Printf("Error parsing %s: %s\n", filePath, err)
			continue
		}

		if !seenDictIDs[dictObj.Id] {
			seenDictIDs[dictObj.Id] = true
			dictionaries = append(dictionaries, modals.DictionaryInfo{
				Id:       dictObj.Id,
				Title:    dictObj.Title,
				FromLang: dictObj.FromLang,
				ToLang:   dictObj.ToLang,
			})
		}

		fmt.Printf("Merging into Phase 04: %s (%d words)\n", entry.Name(), len(dictObj.WordsToHtmlMap))

		for word, htmlValues := range dictObj.WordsToHtmlMap {
			if len(word) > 50 {
				fmt.Printf("  Skipping long key (%d chars): %s\n", len(word), word)
				continue
			}
			merged[word] = append(merged[word], modals.MergedDictEntry{
				Id:   dictObj.Id,
				Html: strings.Join(htmlValues, ""),
			})
		}
	}

	if err := utils.SaveDictToJSON(distPath, merged); err != nil {
		panic(err)
	}

	dictsPath := filepath.Join(distDir, "dictionaries.json")
	if err := utils.SaveDictToJSON(dictsPath, dictionaries); err != nil {
		panic(err)
	}

	fmt.Printf("Phase 03 → Phase 04 merge complete. Total words: %d, dictionaries: %d\n", len(merged), len(dictionaries))
}
