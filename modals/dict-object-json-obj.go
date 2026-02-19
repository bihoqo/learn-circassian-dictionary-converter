package modals

import "fmt"

type Cognate struct {
	Dialect string `json:"dialect,omitempty"`
	Word    string `json:"word,omitempty"`
}

type Example struct {
	Sentence    string `json:"sentence,omitempty"`
	Translation string `json:"translation,omitempty"` // Optional
}

type Definition struct {
	Meaning  string    `json:"meaning,omitempty"`
	Examples []Example `json:"examples,omitempty"` // Optional
}

type WordObject struct {
	Type        string       `json:"type,omitempty"` // e.g., "noun", "verb"
	Definitions []Definition `json:"definitions,omitempty"`

	// Optional fields
	Cognates   []Cognate `json:"cognates,omitempty"`
	Redirect   string    `json:"redirect,omitempty"`
	Derivation string    `json:"derivation,omitempty"`

	// Updated: Now a slice of strings
	Synonyms []string `json:"synonyms,omitempty"`
}

type DictObjectJsonObj struct {
	Title             string                 `json:"title"`
	Id                int                    `json:"id"`
	Format            DictFormat             `json:"format"`
	FromLang          string                 `json:"from_lang"`
	ToLang            string                 `json:"to_lang"`
	WordsToJsonObjMap map[string]*WordObject `json:"words_to_json_obj_map,omitempty"`
}

// --- Constructors ---

func NewDictObjectJsonObj(title string, id int, fromLang, toLang string, format DictFormat) *DictObjectJsonObj {
	return &DictObjectJsonObj{
		Title:             title,
		Id:                id,
		FromLang:          fromLang,
		ToLang:            toLang,
		Format:            format,
		WordsToJsonObjMap: make(map[string]*WordObject),
	}
}

func NewWordObject(wordType string) *WordObject {
	return &WordObject{
		Type: wordType,
	}
}

// --- Helper Methods ---

func (w *WordObject) AddCognate(language, spelling string) {
	w.Cognates = append(w.Cognates, Cognate{
		Word:    spelling,
		Dialect: language,
	})
}

// Updated: AddSynonym now formats the string directly
func (w *WordObject) AddSynonym(spelling, explanation string) {
	if explanation != "" {
		w.Synonyms = append(w.Synonyms, fmt.Sprintf("%s (%s)", spelling, explanation))
	} else {
		w.Synonyms = append(w.Synonyms, spelling)
	}
}

func (w *WordObject) AddDefinition(meaning string, examples []Example) {
	w.Definitions = append(w.Definitions, Definition{
		Meaning:  meaning,
		Examples: examples,
	})
}

func (d *Definition) AddExample(sentence, translation string) {
	d.Examples = append(d.Examples, Example{
		Sentence:    sentence,
		Translation: translation,
	})
}
