package code

import (
	"encoding/json"
	"fmt"
	"learn-circassian-helper/modals"
	"learn-circassian-helper/utils"
	"regexp"
	"strings"
)

// CallConvertPhase01ToPhase02 orchestrates the conversion of raw dictionary data (Phase 1)
// into standardized JSON formats (Phase 2).
func CallConvertPhase01ToPhase02() {
	// Standard HTML Dictionaries
	ConvertStandardHTML("00-Ady-Ady_AIG.json", modals.NewDictObjectPlainText("Адыгабзэм изэхэф гущы1алъ (2006)", 0, "Ady", "Ru", modals.DictFormatHTML))
	ConvertStandardHTML("01-Ady-Ady_AP.json", modals.NewDictObjectPlainText("Адыгэ-урыс псалъалъэ (2012)", 1, "Kbd", "Ru", modals.DictFormatHTML))

	// Arabic HTML (Specific cleanup)
	ConvertArabicHTML("02-Ady-Ara.json", modals.NewDictObjectPlainText("Adel Abdulsalam Lash", 2, "Ady", "Ar", modals.DictFormatPlain))

	// Simple JSON Dictionaries
	ConvertSimpleJSON("03-Ady-En.json", modals.NewDictObjectJsonObj("Адыгэбзэ-инджылыбзэ гущы1алъэ", 3, "Ady", "En", modals.DictFormatJSON))

	// Rich JSON Dictionaries (With Examples/Cognates)
	ConvertRichJSON("04-Ady-En_Adam.json", modals.NewDictObjectJsonObj("Adam Shagash's Adyghe to English Dictionary (2020)", 4, "Ady", "En", modals.DictFormatJSON))

	// More Standard HTML
	ConvertStandardHTML("05-Ady-Rus_Qarden.json", modals.NewDictObjectPlainText("Къардэн (1957)", 5, "Kbd", "Ru", modals.DictFormatHTML))
	ConvertStandardHTML("06-Ady-Rus_Sherdjes.json", modals.NewDictObjectPlainText("Шэрджэс Алий - Яхуэмыфащэу лъэныкъуэ едгъэза псалъэхэр (2009)", 6, "Kbd", "Ru", modals.DictFormatHTML))
	ConvertStandardHTML("07-Ady-Rus_Tharkaho.json", modals.NewDictObjectPlainText("Тхьаркъуахъо (1991)", 7, "Ady", "Ru", modals.DictFormatHTML))

	// Multi-Key HTML (Huvaj)
	ConvertMultiKeyHTML("08-Ady-Tur_Huvaj.json", modals.NewDictObjectPlainText("Хъуажь - Circassian to Turkish (2007)", 8, "Ady/Kbd", "Tr", modals.DictFormatHTML))

	// Simple JSONs
	ConvertSimpleJSON("09.En-Ady.json", modals.NewDictObjectJsonObj("Simple English to Adyghe dictionary", 9, "En", "Ady", modals.DictFormatJSON))
	ConvertSimpleJSON("10-En-Ady_Adam.json", modals.NewDictObjectJsonObj("Adam Shagash's English to Adyghe Dictionary (2020)", 10, "En", "Ady", modals.DictFormatJSON))
	ConvertSimpleJSON("11-En-Kbd-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's English to Kabardian dictionary", 11, "En", "Kbd", modals.DictFormatJSON))
	ConvertSimpleJSON("12-En-Kbd-Ziwar.json", modals.NewDictObjectJsonObj("Ziwar Gish's English to Kabardian dictionary", 12, "En", "Kbd", modals.DictFormatJSON))
	ConvertSimpleJSON("13-Kbd-Ar-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Kabardian to Arabic dictionary", 13, "Ar", "Kbd", modals.DictFormatJSON))

	// Rich JSON
	ConvertRichJSON("14-Kbd-En-2-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Kabardian to English dictionary 2", 14, "Kbd", "En", modals.DictFormatJSON))

	// Simple JSONs
	ConvertSimpleJSON("15-Kbd-En-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Kabardian to English dictionary", 15, "Kbd", "En", modals.DictFormatJSON))
	ConvertSimpleJSON("16-Kbd-En-Ziwar.json", modals.NewDictObjectJsonObj("Ziwar Gish's Kabardian to English dictionary", 16, "Kbd", "En", modals.DictFormatJSON))
	ConvertSimpleJSON("17-Kbd-En_Amjad.json", modals.NewDictObjectJsonObj("Amjad Jaimoukha's Kabardian to English dictionary", 17, "Kbd", "En", modals.DictFormatJSON))
	ConvertSimpleJSON("18-Kbd-Ru&En.json", modals.NewDictObjectJsonObj("Kabardian to Russian & English", 18, "Kbd", "En", modals.DictFormatJSON))

	// Rich JSON
	ConvertRichJSON("19-Kbd-Ru-2-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Kabardian to Russian dictionary 2", 19, "Kbd", "Ru", modals.DictFormatJSON))

	// Simple JSONs
	ConvertSimpleJSON("20-Kbd-Ru-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Kabardian to Russian dictionary", 20, "Kbd", "Ru", modals.DictFormatJSON))
	ConvertSimpleJSON("21-Kbd-Tu-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Kabardian to Turkish dictionary", 21, "Kbd", "Tr", modals.DictFormatJSON))
	ConvertSimpleJSON("22-Ru-Kbd-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Russian to Kabardian dictionary", 22, "Ru", "Kbd", modals.DictFormatJSON))

	// Standard HTML
	ConvertStandardHTML("23-Rus-Ady_Blaghoj.json", modals.NewDictObjectPlainText("Блэгъожъ (1991)", 23, "Ru", "Ady", modals.DictFormatHTML))
	ConvertStandardHTML("24-Rus-Ady_UAG.json", modals.NewDictObjectPlainText("Одэжьдэкъо (1960)", 24, "Ru", "Ady", modals.DictFormatHTML))
	ConvertStandardHTML("25-Rus-Ady_UASP.json", modals.NewDictObjectPlainText("Урыс-адыгэ школ псалъалъэ (1991)", 25, "Ru", "Kbd", modals.DictFormatHTML))

	// Simple JSON
	ConvertSimpleJSON("26-Tu-Kbd-Jonty.json", modals.NewDictObjectJsonObj("Jonty Yamisha's Turkish to Kabardian dictionary", 26, "Tr", "Kbd", modals.DictFormatJSON))

	// Standard HTML
	ConvertStandardHTML("27-Tur-Ady_Abaze.json", modals.NewDictObjectPlainText("Ибрагим Алхаз Абазэ (2005)", 27, "Tr", "Kbd", modals.DictFormatHTML))
	ConvertStandardHTML("28-Tur-Ady_Huvaj.json", modals.NewDictObjectPlainText("Хъуажь - Turkish to Circassian (2007)", 28, "Tr", "Ady/Kbd", modals.DictFormatHTML))
	ConvertStandardHTML("29-Tur-Ady_Teshu.json", modals.NewDictObjectPlainText("Т1эшъу (1991)", 29, "Tr", "Ady", modals.DictFormatHTML))

	// Plain Text Dictionaries
	ConvertThreeVolumes("30-Ady-Rus_ThreeVolumes.txt", modals.NewDictObjectPlainText("Адыгабзэм изэхэф гущы1алъ томищ мэхъу (2011)", 30, "Ady", "Ru", modals.DictFormatPlain))
	ConvertTurkishAdyghe("31-Tu-Ady_Hilmi.txt", modals.NewDictObjectPlainText("Ацумыжъ Хилми (2013)", 31, "Tr", "Ady", modals.DictFormatPlain))
	ConvertSingleLineRusKbd("32-Rus-Kbd_Nalchik_2013.txt", modals.NewDictObjectPlainText("Еджап1эм папщ1э урыс-адыгэ псалъалъэ (2013)", 32, "Ru", "Kbd", modals.DictFormatPlain))
	ConvertAdyRus1960("33-Ady-Rus-1960.txt", modals.NewDictObjectPlainText("Адыгабзэм изэхэф гущы1алъ жъы (1960)", 33, "Ady", "Ru", modals.DictFormatPlain))
	ConvertSingleLineKbdRu("34-Kbd-Ru-2008.txt", modals.NewDictObjectPlainText("адыгэ-урыс псалъалъэ (2008)", 34, "Kbd", "Ru", modals.DictFormatPlain))
}

// formatNumberDots replaces " N." with "\n\tN." for numbered sub-definitions.
func formatNumberDots(s string) string {
	for i := 1; i <= 9; i++ {
		s = strings.ReplaceAll(s, fmt.Sprintf(" %d.", i), fmt.Sprintf("\n\t%d.", i))
	}
	return s
}

// formatNumberDotsAndParens replaces " N." with "\n\tN." and " N)" with "\n\t\tN)".
func formatNumberDotsAndParens(s string) string {
	for i := 1; i <= 9; i++ {
		s = strings.ReplaceAll(s, fmt.Sprintf(" %d.", i), fmt.Sprintf("\n\t%d.", i))
		s = strings.ReplaceAll(s, fmt.Sprintf(" %d)", i), fmt.Sprintf("\n\t\t%d)", i))
	}
	return s
}

// formatNumberDotsStartAware handles " N." → "\n\tN." and also handles lines
// that start with a number (no leading space) for the Adyghe explanatory dicts.
func formatNumberDotsStartAware(s string) string {
	s = formatNumberDots(s)
	if utils.StartsWithNumber(s) {
		for i := 1; i <= 9; i++ {
			s = strings.ReplaceAll(s, fmt.Sprintf("%d.", i), fmt.Sprintf("\n\t%d.", i))
		}
	}
	return s
}

// ConvertMultiKeyHTML handles dictionaries where a single key might contain multiple
// variants split by slashes (e.g., "WordA / WordB Suffix"). It expands these
// into separate entries pointing to the same definition.
func ConvertMultiKeyHTML(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", fileName)
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Multi-Key HTML): %s\n", srcFile)

	// Nested helper to expand keys
	expandKey := func(key string) []string {
		parts := strings.Split(key, "/")

		// If no slash, return original
		if len(parts) == 1 {
			return []string{key}
		}

		// Trim whitespace from all parts
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		// Logic for 2 parts: check for common prefixes or suffixes
		if len(parts) == 2 {
			p1Fields := strings.Fields(parts[0])
			p2Fields := strings.Fields(parts[1])

			// Case 1: "Prefix WordA / WordB" (Common Prefix)
			if len(p1Fields) > 1 && len(p2Fields) == 1 {
				prefix := strings.Join(p1Fields[:len(p1Fields)-1], " ")
				variant1 := parts[0]
				variant2 := prefix + " " + parts[1]
				return []string{variant1, variant2}
			}

			// Case 2: "WordA / WordB Suffix" (Common Suffix)
			if len(p1Fields) == 1 && len(p2Fields) > 1 {
				suffix := strings.Join(p2Fields[1:], " ")
				variant1 := parts[0] + " " + suffix
				variant2 := parts[1]
				return []string{variant1, variant2}
			}
		}

		// Fallback for 3+ parts or non-matching 2-part patterns
		return parts
	}

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		if line == "" || line == "{" || line == "}" {
			return nil
		}

		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Invalid line %d: %s", index, line))
			return nil
		}

		rawKey := strings.Trim(strings.TrimSpace(split[0]), "\"")
		if rawKey == "" {
			return nil
		}

		// Use the nested expansion function
		expandedKeys := expandKey(rawKey)

		value := strings.TrimSuffix(strings.TrimSpace(split[1]), ",")
		value = strings.Trim(value, "\"")
		value = strings.ReplaceAll(value, "\\\"", "\"")
		value = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(value)

		for _, expandedKey := range expandedKeys {
			key := strings.ToLower(expandedKey)
			key = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(key)

			if _, exists := dictObj.WordsToPlainTextMap[key]; !exists {
				dictObj.WordsToPlainTextMap[key] = make([]string, 0)
			}
			dictObj.WordsToPlainTextMap[key] = append(dictObj.WordsToPlainTextMap[key], value)
		}

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertStandardHTML processes standard dictionaries where the value is an HTML string.
// It performs standard Circassian letter cleaning on both Keys and Values.
func ConvertStandardHTML(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", fileName)
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Standard HTML): %s\n", srcFile)

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		if line == "" || line == "{" || line == "}" {
			return nil
		}

		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Invalid line %d: %s", index, line))
			return nil
		}

		key := strings.Trim(strings.TrimSpace(split[0]), "\"")
		if key == "" {
			return nil
		}
		key = strings.ToLower(key)
		key = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(key)

		value := strings.TrimSuffix(strings.TrimSpace(split[1]), ",")
		value = strings.Trim(value, "\"")
		value = strings.ReplaceAll(value, "\\\"", "\"")
		value = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(value)

		if _, exists := dictObj.WordsToPlainTextMap[key]; !exists {
			dictObj.WordsToPlainTextMap[key] = make([]string, 0)
		}
		dictObj.WordsToPlainTextMap[key] = append(dictObj.WordsToPlainTextMap[key], value)

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertArabicHTML processes dictionaries with Arabic content.
// It includes specific logic to strip HTML tags and collapse whitespace for cleaner plain text.
func ConvertArabicHTML(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", fileName)
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Arabic Plain): %s\n", srcFile)

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		if line == "" || line == "{" || line == "}" {
			return nil
		}

		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Invalid line %d: %s", index, line))
			return nil
		}

		key := strings.Trim(strings.TrimSpace(split[0]), "\"")
		if key == "" {
			return nil
		}
		key = strings.ToLower(key)
		key = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(key)

		value := strings.TrimSuffix(strings.TrimSpace(split[1]), ",")
		value = strings.Trim(value, "\"")

		// HTML Stripping & Whitespace Cleaning
		value = strings.ReplaceAll(value, "<div style=\\\"margin-left:1em\\\">", " ")
		value = strings.ReplaceAll(value, "</div>", " ")
		value = strings.ReplaceAll(value, "\\\"", "'")
		value = strings.Join(strings.Fields(value), " ")

		if value == "" {
			return nil
		}

		if _, exists := dictObj.WordsToPlainTextMap[key]; !exists {
			dictObj.WordsToPlainTextMap[key] = make([]string, 0)
		}

		// Prevent exact duplicates
		isDup := false
		for _, existingVal := range dictObj.WordsToPlainTextMap[key] {
			if existingVal == value {
				isDup = true
				break
			}
		}

		if !isDup {
			dictObj.WordsToPlainTextMap[key] = append(dictObj.WordsToPlainTextMap[key], value)
		}

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertSimpleJSON processes JSONs with basic structure (Type, Definitions, Links).
// It handles cleaning and formatting of "Clarity" fields in links.
func ConvertSimpleJSON(fileName string, dictObj *modals.DictObjectJsonObj) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", fileName)
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Simple JSON): %s\n", srcFile)

	type RawLink struct {
		Word    string `json:"word"`
		Clarity string `json:"clarity"`
	}
	type RawDef struct {
		Meaning string `json:"meaning"`
	}
	type RawEntry struct {
		Type        string    `json:"type"`
		Definitions []RawDef  `json:"definitions"`
		Links       []RawLink `json:"links"`
	}

	cleanText := func(text string) string {
		text = strings.TrimSpace(text)
		isCircassianTarget := strings.ToLower(dictObj.ToLang) == "ady" || strings.ToLower(dictObj.ToLang) == "kbd"
		isSpecialMixedFile := strings.Contains(fileName, "18-Kbd-Ru&En.json")

		if isCircassianTarget || isSpecialMixedFile {
			text = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(text)
		}
		return text
	}

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		if line == "" || line == "{" || line == "}" {
			return nil
		}
		line = strings.TrimSuffix(line, ",")

		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Invalid line %d: %s", index, line))
			return nil
		}

		rawKey := strings.TrimSpace(split[0])
		rawKey = strings.Trim(rawKey, "\"")
		if strings.ToLower(dictObj.FromLang) == "ady" || strings.ToLower(dictObj.FromLang) == "kbd" {
			rawKey = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(rawKey)
		}
		key := strings.ToLower(rawKey)

		rawValueStr := strings.TrimSpace(split[1])

		// Sanitize JSON string boundaries
		firstBrace := strings.Index(rawValueStr, "{")
		lastBrace := strings.LastIndex(rawValueStr, "}")

		if firstBrace != -1 && lastBrace != -1 && lastBrace > firstBrace {
			rawValueStr = rawValueStr[firstBrace : lastBrace+1]
		} else {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("No JSON object found in line %d: %s", index, line))
			return nil
		}

		var rawEntry RawEntry
		if err := json.Unmarshal([]byte(rawValueStr), &rawEntry); err != nil {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("JSON Parse Error line %d: %v", index, err))
			return nil
		}

		wordObj := modals.NewWordObject(rawEntry.Type)

		for _, def := range rawEntry.Definitions {
			if def.Meaning != "" {
				cleanedMeaning := cleanText(def.Meaning)
				wordObj.AddDefinition(cleanedMeaning, nil)
			}
		}

		for _, link := range rawEntry.Links {
			if link.Word != "" {
				cleanedWord := cleanText(link.Word)
				finalMeaning := cleanedWord
				if link.Clarity != "" {
					finalMeaning = fmt.Sprintf("%s (%s)", cleanedWord, strings.TrimSpace(link.Clarity))
				}
				wordObj.AddDefinition(finalMeaning, nil)
			}
		}

		if len(wordObj.Definitions) > 0 {
			if dictObj.WordsToJsonObjMap == nil {
				dictObj.WordsToJsonObjMap = make(map[string]*modals.WordObject)
			}
			if existing, exists := dictObj.WordsToJsonObjMap[key]; exists {
				existing.Definitions = append(existing.Definitions, wordObj.Definitions...)
				existing.Cognates = append(existing.Cognates, wordObj.Cognates...)
				existing.Synonyms = append(existing.Synonyms, wordObj.Synonyms...)
			} else {
				dictObj.WordsToJsonObjMap[key] = wordObj
			}
		}

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertRichJSON processes complex JSON structures containing definitions, examples,
// synonyms, and cognates. It handles cleaning of pipe characters '|' in examples.
func ConvertRichJSON(fileName string, dictObj *modals.DictObjectJsonObj) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", fileName)
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Rich JSON): %s\n", srcFile)

	type RawExample struct {
		Sentence    string `json:"sentence"`
		Translation string `json:"translation"`
	}
	type RawDef struct {
		Meaning  string       `json:"meaning"`
		Examples []RawExample `json:"examples"`
	}
	type RawSynonym struct {
		Word    string `json:"word"`
		Clarity string `json:"clarity"`
	}
	type RawEntry struct {
		Type        string       `json:"type"`
		Shapsug     string       `json:"shapsug"`
		Kabardian   string       `json:"kabardian"`
		Synonyms    []RawSynonym `json:"synonyms"`
		Definitions []RawDef     `json:"definitions"`
	}

	cleanContent := func(text string) string {
		if text == "" {
			return ""
		}
		text = strings.ReplaceAll(text, "|", "")
		text = strings.TrimSpace(text)
		text = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(text)
		return text
	}

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		if line == "" || line == "{" || line == "}" {
			return nil
		}
		line = strings.TrimSuffix(line, ",")

		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Invalid line %d: %s", index, line))
			return nil
		}

		rawKey := strings.TrimSpace(split[0])
		rawKey = strings.Trim(rawKey, "\"")

		if strings.ToLower(dictObj.FromLang) == "ady" || strings.ToLower(dictObj.FromLang) == "kbd" {
			rawKey = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(rawKey)
		}
		key := strings.ToLower(rawKey)

		rawValueStr := strings.TrimSpace(split[1])
		var rawEntry RawEntry

		if err := json.Unmarshal([]byte(rawValueStr), &rawEntry); err != nil {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("JSON Parse Error line %d: %v", index, err))
			return nil
		}

		wordObj := modals.NewWordObject(rawEntry.Type)

		// A. Cognates
		if rawEntry.Shapsug != "" {
			wordObj.AddCognate("shapsug", cleanContent(rawEntry.Shapsug))
		}
		if rawEntry.Kabardian != "" {
			wordObj.AddCognate("kabardian", cleanContent(rawEntry.Kabardian))
		}

		// B. Synonyms
		for _, syn := range rawEntry.Synonyms {
			if syn.Word != "" {
				wordObj.AddSynonym(cleanContent(syn.Word), cleanContent(syn.Clarity))
			}
		}

		// C. Definitions & Examples
		for _, def := range rawEntry.Definitions {
			if def.Meaning == "" && len(def.Examples) == 0 {
				continue
			}

			var modalExamples []modals.Example
			for _, ex := range def.Examples {
				if ex.Sentence != "" || ex.Translation != "" {
					modalExamples = append(modalExamples, modals.Example{
						Sentence:    cleanContent(ex.Sentence),
						Translation: cleanContent(ex.Translation),
					})
				}
			}

			cleanedMeaning := cleanContent(def.Meaning)
			wordObj.AddDefinition(cleanedMeaning, modalExamples)
		}

		if len(wordObj.Definitions) > 0 || len(wordObj.Synonyms) > 0 || len(wordObj.Cognates) > 0 {
			if dictObj.WordsToJsonObjMap == nil {
				dictObj.WordsToJsonObjMap = make(map[string]*modals.WordObject)
			}
			if existing, exists := dictObj.WordsToJsonObjMap[key]; exists {
				existing.Definitions = append(existing.Definitions, wordObj.Definitions...)
				existing.Cognates = append(existing.Cognates, wordObj.Cognates...)
				existing.Synonyms = append(existing.Synonyms, wordObj.Synonyms...)
			} else {
				dictObj.WordsToJsonObjMap[key] = wordObj
			}
		}

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertThreeVolumes processes the three-volume Adyghe explanatory dictionary (plain text).
// Each entry starts with a fully-capitalized Cyrillic word. Continuation lines are appended
// to the current entry. Numbered sub-definitions (1., 2., ...) get newline+tab formatting.
func ConvertThreeVolumes(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", strings.TrimSuffix(fileName, ".txt")+".json")
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Three Volumes): %s\n", srcFile)

	var currentKey string
	var currentValue strings.Builder

	flushEntry := func() {
		if currentKey == "" {
			return
		}
		spelling := strings.ToLower(currentKey)
		spelling = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(spelling)

		value := utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(currentValue.String())

		if _, exists := dictObj.WordsToPlainTextMap[spelling]; !exists {
			dictObj.WordsToPlainTextMap[spelling] = make([]string, 0)
		}
		dictObj.WordsToPlainTextMap[spelling] = append(dictObj.WordsToPlainTextMap[spelling], value)
	}

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		trimmedLine := strings.TrimSpace(line)
		trimmedLine = utils.StripZeroWidthChars(trimmedLine)

		// Pre-apply polachka conversion so that Latin I/i (palochka) becomes "1"
		// before we check capitalization. This ensures headwords like "ЗЕФЭГЪОШIУ"
		// are correctly detected (the I becomes 1, a digit ignored by IsFullyCapitalized).
		normalizedLine := utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(trimmedLine)
		words := strings.Fields(normalizedLine)

		if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Empty line %d", index))
			return nil
		}

		trimmedLine = formatNumberDotsStartAware(trimmedLine)

		firstWord := utils.RemoveSuffixes(words[0])

		if utils.IsFullyCapitalized(firstWord) && !utils.StartsWithSpecialCharacter(firstWord) {
			flushEntry()
			currentKey = firstWord
			currentValue.Reset()
			currentValue.WriteString(trimmedLine)
		} else {
			if currentValue.Len() > 0 {
				currentValue.WriteString(" ")
			}
			currentValue.WriteString(trimmedLine)
		}

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	flushEntry()

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertTurkishAdyghe processes the Turkish-Adyghe dictionary by Hilmi (plain text).
// Keys are fully-capitalized Latin/Turkish words. Values contain Circassian text.
// Turkish "i" in keys is preserved (no polachka conversion on keys).
// Polachka conversion is only applied to values (which contain Circassian text).
func ConvertTurkishAdyghe(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", strings.TrimSuffix(fileName, ".txt")+".json")
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Turkish-Adyghe): %s\n", srcFile)

	var currentKey string
	var currentValue strings.Builder

	flushEntry := func() {
		if currentKey == "" {
			return
		}
		// Turkish key: no polachka conversion, just lowercase and trim colons
		spelling := strings.ToLower(currentKey)
		spelling = strings.Trim(spelling, ":")

		// Polachka conversion only on the value (Circassian text)
		value := utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(currentValue.String())

		if _, exists := dictObj.WordsToPlainTextMap[spelling]; !exists {
			dictObj.WordsToPlainTextMap[spelling] = make([]string, 0)
		}
		dictObj.WordsToPlainTextMap[spelling] = append(dictObj.WordsToPlainTextMap[spelling], value)
	}

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		trimmedLine := strings.TrimSpace(line)
		trimmedLine = utils.StripZeroWidthChars(trimmedLine)
		words := strings.Fields(trimmedLine)

		trimmedLine = formatNumberDotsAndParens(trimmedLine)

		if strings.TrimSpace(trimmedLine) == "" {
			return nil
		} else if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Empty line %d", index))
		} else if len(words) == 1 && len(words[0]) == 3 && strings.Contains(words[0], "-") {
			// Section headers like "A-B" — skip
			invalidLinesList = append(invalidLinesList, trimmedLine)
		} else if len(words) > 0 && utils.IsFullyCapitalized(words[0]) && !utils.StartsWithNumber(words[0]) && !utils.StartsWithSpecialCharacter(words[0]) {
			flushEntry()
			currentKey = utils.RemoveSuffixes(words[0])
			currentValue.Reset()
			currentValue.WriteString(trimmedLine)
		} else {
			if currentValue.Len() > 0 {
				currentValue.WriteString(" ")
			}
			currentValue.WriteString(trimmedLine)
		}

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	flushEntry()

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertSingleLineRusKbd processes the Russian-Kabardian school dictionary (Nalchik 2013).
// Each line is a single entry. The first word is the Russian key.
// Polachka conversion is applied to the value (Kabardian content) but not the Russian key.
func ConvertSingleLineRusKbd(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", strings.TrimSuffix(fileName, ".txt")+".json")
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Rus-Kbd Single Line): %s\n", srcFile)

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		line = utils.StripZeroWidthChars(line)

		words := strings.Fields(line)
		if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Empty line %d", index))
			return nil
		}

		// Russian key: no polachka conversion
		spelling := strings.ToLower(words[0])

		value := formatNumberDotsAndParens(line)
		value = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(value)

		if _, exists := dictObj.WordsToPlainTextMap[spelling]; !exists {
			dictObj.WordsToPlainTextMap[spelling] = make([]string, 0)
		}
		dictObj.WordsToPlainTextMap[spelling] = append(dictObj.WordsToPlainTextMap[spelling], value)

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertAdyRus1960 processes the 1960 Adyghe explanatory dictionary (plain text).
// Similar to ConvertThreeVolumes but includes OCR space-removal between uppercase
// Cyrillic letters and zero-width character stripping. Keys may contain digits (e.g., "А1").
func ConvertAdyRus1960(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", strings.TrimSuffix(fileName, ".txt")+".json")
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Ady-Rus 1960): %s\n", srcFile)

	// removeFirstWordSpaces collapses spaces between uppercase Cyrillic letters and 'I'.
	// Fixes OCR artifacts where "АБАДЗЭ" was scanned as "А Б А Д З Э".
	reCyrillicSpaces := regexp.MustCompile(`([А-ЯЁI])\s+([А-ЯЁI])`)
	removeFirstWordSpaces := func(line string) string {
		return reCyrillicSpaces.ReplaceAllString(line, "$1$2")
	}

	var currentKey string
	var currentValue strings.Builder

	flushEntry := func() {
		if currentKey == "" {
			return
		}
		spelling := strings.ToLower(currentKey)
		spelling = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(spelling)

		value := utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(currentValue.String())

		if _, exists := dictObj.WordsToPlainTextMap[spelling]; !exists {
			dictObj.WordsToPlainTextMap[spelling] = make([]string, 0)
		}
		dictObj.WordsToPlainTextMap[spelling] = append(dictObj.WordsToPlainTextMap[spelling], value)
	}

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		line = utils.StripZeroWidthChars(line)
		line = removeFirstWordSpaces(line)
		trimmedLine := strings.TrimSpace(line)

		words := strings.Fields(trimmedLine)

		trimmedLine = formatNumberDotsStartAware(trimmedLine)

		if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Empty line %d", index))
		} else if utils.IsFullyCapitalized(words[0]) && !utils.StartsWithSpecialCharacter(words[0]) {
			flushEntry()
			currentKey = utils.RemoveSuffixes(words[0])
			currentValue.Reset()
			currentValue.WriteString(trimmedLine)
		} else {
			if currentValue.Len() > 0 {
				currentValue.WriteString(" ")
			}
			currentValue.WriteString(trimmedLine)
		}

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	flushEntry()

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}

// ConvertSingleLineKbdRu processes the 2008 Kabardian-Russian dictionary (plain text).
// Each line is a single entry. The first word is the Kabardian key.
// Polachka conversion is applied to both key and value.
func ConvertSingleLineKbdRu(fileName string, dictObj *modals.DictObjectPlainText) {
	srcFile := fmt.Sprintf("content/phase-01-raw-data/%s", fileName)
	distFile := fmt.Sprintf("content/phase-02-json-data/%s", strings.TrimSuffix(fileName, ".txt")+".json")
	invalidLinesList := make([]string, 0)

	fmt.Printf("Starting conversion (Kbd-Ru Single Line): %s\n", srcFile)

	err := utils.ReadFileLineByLine(srcFile, func(line string, index int) error {
		line = strings.TrimSpace(line)
		line = utils.StripZeroWidthChars(line)

		words := strings.Fields(line)
		if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, fmt.Sprintf("Empty line %d", index))
			return nil
		}

		// Kabardian key: apply polachka conversion
		spelling := strings.ToLower(words[0])
		spelling = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(spelling)

		value := formatNumberDotsAndParens(line)
		value = utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords(value)

		if _, exists := dictObj.WordsToPlainTextMap[spelling]; !exists {
			dictObj.WordsToPlainTextMap[spelling] = make([]string, 0)
		}
		dictObj.WordsToPlainTextMap[spelling] = append(dictObj.WordsToPlainTextMap[spelling], value)

		if index%1000 == 0 {
			fmt.Printf("Processed line %d...\n", index)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(invalidLinesList) > 0 {
		fmt.Printf("\n--Invalid lines in %s:--\n", srcFile)
		for idx, line := range invalidLinesList {
			fmt.Printf("%d. %s\n", idx, line)
		}
	}

	err = utils.SaveDictToJSON(distFile, dictObj)
	if err != nil {
		panic(err)
	}
}
