# Project Instructions

## Keeping Documentation Up to Date

Whenever you make changes that affect the project's structure, conventions, or behavior, you **MUST** also update:

1. **`CLAUDE.md`** — This file (AI instructions for Claude)
2. **`GEMINI.md`** — AI instructions for Gemini (must stay identical to CLAUDE.md)
3. **`README.md`** — The project README

Keep all three files in sync. If you add a new phase, new converter, new data model, new utility, or change any convention documented here, reflect it in all three files.

## About This Project

This project converts Circassian (Adyghe/Kabardian) dictionaries from various raw formats through a 5-phase pipeline into a unified SQLite database. The dictionaries come from multiple sources with different languages (Russian, Turkish, English, Arabic) paired with Circassian.

### Pipeline Phases

| Phase | Description | Code |
|-------|-------------|------|
| 01 → 02 | Parse raw data (HTML, JSON, plain text) into standardized JSON | `convert-phase-01-to-phase-02.go` |
| 02 → 03 | Convert definitions to HTML with formatting (bold, indentation) | `convert-phase-02-to-phase-03.go` |
| 03 → 04 | Merge all dictionaries into a single key→entries JSON database | `convert-phase-03-to-phase-04.go` |
| 04 → 05 | Write merged database to SQLite for efficient lookups | `convert-phase-04-to-phase-05.go` |

### Project Structure

```
main.go                           — Entry point, runs all phases sequentially
code/
  convert-phase-01-to-phase-02.go — Raw → standardized JSON converters
  convert-phase-02-to-phase-03.go — JSON → HTML-enriched JSON
  convert-phase-03-to-phase-04.go — Merge all dictionaries into one DB
  convert-phase-04-to-phase-05.go — Merged JSON → SQLite
modals/
  dict-object-plain-text.go       — DictObjectPlainText type + DictFormat enum
  dict-object-json-obj.go         — DictObjectJsonObj type (WordObject with examples/cognates)
  dict-object-html.go             — DictObjectHTML type + MergedDictEntry + DictionaryInfo
utils/
  text.go                         — Text utilities (palochka, casing, etc.)
  files.go                        — File I/O (ReadFileLineByLine, SaveDictToJSON)
python_scripts/
  process_data.py                 — Auxiliary Python processing
content/
  raw-data-samples/               — Small excerpts (OK to read)
  phase-01-raw-data/              — Original dictionary files (DO NOT read)
  phase-02-json-data/             — Standardized JSON output (DO NOT read)
  phase-03-html-data/             — HTML-enriched JSON output (DO NOT read)
  phase-04-merged-database/       — Single merged JSON database (DO NOT read)
  phase-05-sqlite/                — Final SQLite database (DO NOT read)
```

### Data Models

- **`DictObjectPlainText`** (`map[string][]string`) — Used for Phase 01→02 when source is HTML or plain text. Key is the headword, value is a list of definition strings.
- **`DictObjectJsonObj`** (`map[string]*WordObject`) — Used for Phase 01→02 when source is rich JSON. WordObject contains definitions, examples, cognates, synonyms, derivation, redirect.
- **`DictObjectHTML`** (`map[string][]string`) — Phase 03 output. Key is headword, value is list of HTML-formatted definition strings.
- **`MergedDictEntry`** — Phase 04 word entry containing only `id` (dictionary ID) and `html` (formatted content). Dictionary metadata (title, languages) is stored separately.
- **`DictionaryInfo`** — Dictionary metadata: `id`, `title`, `from_lang`, `to_lang`. Stored in `dictionaries.json` (Phase 04) and the `dictionaries` SQLite table (Phase 05).

### SQLite Schema (Two Tables)

The final SQLite database uses two tables to avoid repeating dictionary titles and language info in every word entry:

- **`dictionaries`** — One row per dictionary source. Columns: `id` (INTEGER PRIMARY KEY), `title` (TEXT), `from_lang` (TEXT), `to_lang` (TEXT).
- **`words`** — One row per word. Columns: `word` (TEXT PRIMARY KEY), `entries` (TEXT — JSON array of `{id, html}` objects).

To get a word's full entry with dictionary titles, join the two tables by matching each entry's `id` to `dictionaries.id`.

## Data Files — Reading Rules

- **`content/raw-data-samples/`** — You **CAN** freely read these files (names and content). They are small sample excerpts meant to help you understand each dictionary's format.
- **All other folders inside `content/`** (e.g., `content/phase-01-raw-data/`, `content/phase-02-json-data/`) — **NEVER read the content** of these files. They are 1-10MB+ and reading them will waste tokens. You may read the **file names** only (e.g., via `ls`). The user will provide relevant excerpts when needed.

## Palochka (Polachka) — The "1" Convention

Circassian languages use a "palochka" letter that looks like Latin "I", "i", "l", Turkish "ı"/"İ", Cyrillic "Ӏ"/"ӏ", or Ukrainian "І"/"і". This causes encoding chaos across sources.

**Rule:** For all Circassian text, convert ALL palochka-looking characters into the digit **"1"**. This is done by `utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords()` in `utils/text.go`. It identifies Circassian words by detecting a Cyrillic letter adjacent to a palochka-looking character.

### Turkish "i" Preservation

Be **very careful** with dictionaries involving Turkish (e.g., `Tu-Ady`, `Tur-Ady`, `Tu-Kbd`). Turkish has its own "i" and "İ" letters that must NOT be converted to "1". The `ConvertAllPolachkaLookingLettersTo1InCircassianWords` function handles this correctly because it only converts when a palochka-looking character is adjacent to a Cyrillic letter. Turkish text uses Latin letters, so Turkish "i" next to Latin letters is preserved.

For Turkish dictionaries: do NOT apply polachka conversion to the dictionary key (the Turkish word). Only apply it to the value/definition (which contains Circassian text).

## Dictionary Key "/" Splitting

Some dictionaries have keys (headwords) containing "/". This means the definition applies to multiple word forms:

- `X / Y` where X and Y are **different** (after lowercasing) — duplicate the definition for both keys: `x: {...}` and `y: {...}`
- `X / X` where both sides are **identical** (after lowercasing) — no duplication needed, use one key
- `X / Y ZZZZZ` where ZZZZZ is trailing text — create `x ZZZZZ: {...}` and `y ZZZZZ: {...}` (if X !== Y after lowercasing)

Always lowercase the key (headword).

## Duplicate Keys — Merge, Never Override

When the same key (headword) appears more than once in a dictionary, **append** the new definitions to the existing entry. Never overwrite a previous entry with a new one.

- **`DictObjectSimple`** (`WordsMap map[string][]string`): Use `append()` to add values to the existing slice.
- **`DictObjectFull`** (`Words map[string]*WordObject`): If the key already exists, merge the new `WordObject`'s definitions, cognates, and synonyms into the existing one. Do not replace the existing `*WordObject`.

## Code Patterns

- Converter functions live in `code/convert-phase-01-to-phase-02.go`
- Each dictionary format gets its own converter function
- All converters are registered in `CallConvertPhase01ToPhase02()`
- Use `utils.ReadFileLineByLine()` for line-by-line processing
- Use `utils.SaveDictToJSON()` for output
- Use `modals.DictObjectSimple` for simple key→definitions maps
- Use `modals.DictObjectFull` for rich entries with examples/cognates
- Helper functions (text utilities) go in `utils/text.go`

## Running the Project

```bash
go run main.go
```

This runs all five phases in sequence. Final output: `content/phase-05-sqlite/dictionary.db`.

Requires Go 1.25+. Dependencies are managed via `go.mod` (SQLite via `modernc.org/sqlite` — no CGO required).
