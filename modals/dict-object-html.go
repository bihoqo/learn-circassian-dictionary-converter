package modals

type DictObjectHTML struct {
	Title          string              `json:"title"`
	Id             int                 `json:"id"`
	WordsToHtmlMap map[string][]string `json:"words_to_html_map"`
	FromLang       string              `json:"from_lang"`
	ToLang         string              `json:"to_lang"`
}

func NewDictObjectHTML(title string, id int, fromLang string, toLang string) *DictObjectHTML {
	return &DictObjectHTML{
		Title:          title,
		Id:             id,
		WordsToHtmlMap: make(map[string][]string),
		FromLang:       fromLang,
		ToLang:         toLang,
	}
}

type MergedDictEntry struct {
	Id   int    `json:"id"`
	Html string `json:"html"`
}

type DictionaryInfo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	FromLang string `json:"from_lang"`
	ToLang   string `json:"to_lang"`
}
