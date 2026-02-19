package code

import (
	"encoding/json"
	"fmt"
	"html"
	"learn-circassian-helper/modals"
	"learn-circassian-helper/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var pipeMarkerRegex = regexp.MustCompile(`\|([^|]+)\|`)

// formatTextWithBoldMarkers converts |text| markers into bold HTML spans.
// Text outside markers is expected to already be HTML-escaped.
func formatTextWithBoldMarkers(escapedText string) string {
	return pipeMarkerRegex.ReplaceAllString(escapedText, "<span style='font-weight:bold'>$1</span>")
}

// meaningToHTML converts a plain text definition (with \n and \t formatting)
// into HTML divs with indentation. Supports |text| bold markers.
func meaningToHTML(meaning string) string {
	lines := strings.Split(meaning, "\n")
	var sb strings.Builder
	for _, line := range lines {
		tabs := strings.Count(line, "\t")
		line = strings.TrimLeft(line, "\t")
		if line == "" {
			continue
		}
		escaped := html.EscapeString(line)
		processed := formatTextWithBoldMarkers(escaped)
		sb.WriteString(fmt.Sprintf("<div style='margin-left:%dem'>%s</div>", tabs, processed))
	}
	return sb.String()
}

// wordObjectToHTML renders a DictObjectJsonObj WordObject into a full HTML string.
func wordObjectToHTML(key string, w *modals.WordObject) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("<div><h2>%s</h2>", html.EscapeString(key)))

	if w.Type != "" {
		sb.WriteString(fmt.Sprintf("<p>Type: %s</p>", html.EscapeString(w.Type)))
	}

	if len(w.Cognates) > 0 {
		sb.WriteString("<h3>Cognates:</h3>")
		for _, cognate := range w.Cognates {
			sb.WriteString(fmt.Sprintf("<div style='margin-left:1em'>%s: %s</div>",
				html.EscapeString(cognate.Dialect), html.EscapeString(cognate.Word)))
		}
	}

	if w.Redirect != "" {
		sb.WriteString(fmt.Sprintf("<p>Redirect: %s</p>", html.EscapeString(w.Redirect)))
	}

	if len(w.Definitions) > 0 {
		sb.WriteString("<h3>Definitions:</h3>")
		for i, def := range w.Definitions {
			sb.WriteString(fmt.Sprintf("<div style='margin-left:1em; margin-bottom:1em'><font color='darkblue'><span style='font-weight:bold'>%d.</span></font> %s</div>",
				i+1, meaningToHTML(def.Meaning)))
			for _, ex := range def.Examples {
				sentence := formatTextWithBoldMarkers(html.EscapeString(ex.Sentence))
				translation := formatTextWithBoldMarkers(html.EscapeString(ex.Translation))
				sb.WriteString(fmt.Sprintf("<div style='margin-left:3em'>%s — %s</div>", sentence, translation))
			}
		}
	}

	if w.Derivation != "" {
		sb.WriteString(fmt.Sprintf("<p>Derivation: %s</p>", html.EscapeString(w.Derivation)))
	}

	if len(w.Synonyms) > 0 {
		sb.WriteString("<h3>Synonyms:</h3>")
		for _, synonym := range w.Synonyms {
			sb.WriteString(fmt.Sprintf("<div style='margin-left:1em'>%s</div>", html.EscapeString(synonym)))
		}
	}

	sb.WriteString("</div>")
	return sb.String()
}

// CallConvertPhase02ToPhase03 reads all Phase 02 JSON files and converts their
// values into HTML format, outputting Phase 03 files. HTML-format dicts are
// copied as-is. Plain and JSON formats are converted to HTML.
func CallConvertPhase02ToPhase03() {
	srcDir := "content/phase-02-json-data"
	distDir := "content/phase-03-html-data"

	if err := os.MkdirAll(distDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create output directory: %v", err))
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to read source directory: %v", err))
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(srcDir, entry.Name())
		distPath := filepath.Join(distDir, entry.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading %s: %s\n", filePath, err)
			continue
		}

		// Detect format from JSON
		var formatDetector struct {
			Format modals.DictFormat `json:"format"`
		}
		if err := json.Unmarshal(data, &formatDetector); err != nil {
			fmt.Printf("Error parsing format from %s: %s\n", filePath, err)
			continue
		}

		fmt.Printf("Converting to Phase 03: %s (format: %d)\n", entry.Name(), formatDetector.Format)

		switch formatDetector.Format {
		case modals.DictFormatHTML:
			// Already HTML — pass values through as-is
			var dictObj modals.DictObjectPlainText
			if err := json.Unmarshal(data, &dictObj); err != nil {
				fmt.Printf("Error parsing %s: %s\n", filePath, err)
				continue
			}
			htmlDict := modals.NewDictObjectHTML(dictObj.Title, dictObj.Id, dictObj.FromLang, dictObj.ToLang)
			for key, values := range dictObj.WordsToPlainTextMap {
				htmlDict.WordsToHtmlMap[key] = values
			}
			if err := utils.SaveDictToJSON(distPath, htmlDict); err != nil {
				panic(err)
			}

		case modals.DictFormatPlain:
			var dictObj modals.DictObjectPlainText
			if err := json.Unmarshal(data, &dictObj); err != nil {
				fmt.Printf("Error parsing %s: %s\n", filePath, err)
				continue
			}
			htmlDict := modals.NewDictObjectHTML(dictObj.Title, dictObj.Id, dictObj.FromLang, dictObj.ToLang)
			for key, values := range dictObj.WordsToPlainTextMap {
				htmlValues := make([]string, len(values))
				for i, val := range values {
					htmlValues[i] = meaningToHTML(val)
				}
				htmlDict.WordsToHtmlMap[key] = htmlValues
			}
			if err := utils.SaveDictToJSON(distPath, htmlDict); err != nil {
				panic(err)
			}

		case modals.DictFormatJSON:
			var dictObj modals.DictObjectJsonObj
			if err := json.Unmarshal(data, &dictObj); err != nil {
				fmt.Printf("Error parsing %s: %s\n", filePath, err)
				continue
			}
			htmlDict := modals.NewDictObjectHTML(dictObj.Title, dictObj.Id, dictObj.FromLang, dictObj.ToLang)
			for key, wordObj := range dictObj.WordsToJsonObjMap {
				htmlDict.WordsToHtmlMap[key] = []string{wordObjectToHTML(key, wordObj)}
			}
			if err := utils.SaveDictToJSON(distPath, htmlDict); err != nil {
				panic(err)
			}

		default:
			fmt.Printf("Unknown format %d in %s, skipping\n", formatDetector.Format, filePath)
		}
	}

	fmt.Println("Phase 02 → Phase 03 conversion complete.")
}
