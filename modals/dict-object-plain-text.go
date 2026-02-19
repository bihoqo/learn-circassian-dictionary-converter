package modals

type DictFormat int

// 2. Declare the constants using iota
const (
	DictFormatUnknown DictFormat = iota // 0
	DictFormatHTML                      // 1
	DictFormatJSON                      // 2
	DictFormatPlain                     // 3
)

type DictObjectPlainText struct {
	Title               string              `json:"title"`
	Id                  int                 `json:"id"`
	WordsToPlainTextMap map[string][]string `json:"words_to_plain_text_map"`
	FromLang            string              `json:"from_lang"`
	ToLang              string              `json:"to_lang"`
	Format              DictFormat          `json:"format"`
}

func NewDictObjectPlainText(title string, id int, fromLang string, toLang string, format DictFormat) *DictObjectPlainText {
	return &DictObjectPlainText{
		Title:               title,
		Id:                  id,
		WordsToPlainTextMap: make(map[string][]string),
		FromLang:            fromLang,
		ToLang:              toLang,
		Format:              format,
	}
}
