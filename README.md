# Circassian Dictionary Format Converter

A Go pipeline that converts Circassian (Adyghe/Kabardian) dictionaries from various raw formats into a unified SQLite database, passing through multiple standardization phases.

## What It Does

Takes 35+ bilingual dictionaries (Circassian paired with Russian, Turkish, English, or Arabic) from heterogeneous source formats and produces a single searchable SQLite database. The pipeline runs in five phases:

| Phase | Input | Output | Description |
|-------|-------|--------|-------------|
| 01 | Raw data (HTML, JSON, plain text) | — | Source dictionary files in their original formats |
| 02 | Phase 01 | Standardized JSON | Parse each dictionary into a uniform JSON structure |
| 03 | Phase 02 | HTML-enriched JSON | Convert definitions to HTML with formatting (bold, indentation) |
| 04 | Phase 03 | Merged JSON | Merge all dictionaries into a single key→entries database |
| 05 | Phase 04 | SQLite | Write the merged database to SQLite for efficient lookups |

## Supported Dictionaries

The converter handles 35 dictionaries covering multiple language pairs and formats:

- **Adyghe-Russian** (Qarden, Sherdjes, Tharkaho, Three Volumes, 1960)
- **Adyghe-English** (community, Adam Shagash)
- **Adyghe-Turkish** (Huvaj)
- **Adyghe-Arabic** (Adel Abdulsalam Lash)
- **Adyghe-Adyghe** (AIG, AP)
- **Kabardian-Russian** (Jonty, Nalchik 2013, 2008)
- **Kabardian-English** (Jonty, Ziwar, Amjad)
- **Kabardian-Turkish** (Jonty)
- **Kabardian-Arabic** (Jonty)
- **English-Adyghe / English-Kabardian** (various)
- **Russian-Adyghe / Russian-Kabardian** (various)
- **Turkish-Adyghe / Turkish-Kabardian** (various)

## Project Structure

```
.
├── main.go                         # Entry point — runs all phases sequentially
├── code/
│   ├── convert-phase-01-to-phase-02.go   # Raw → standardized JSON converters
│   ├── convert-phase-02-to-phase-03.go   # JSON → HTML-enriched JSON
│   ├── convert-phase-03-to-phase-04.go   # Merge all dictionaries into one DB
│   └── convert-phase-04-to-phase-05.go   # Merged JSON → SQLite
├── modals/
│   ├── dict-object-plain-text.go   # DictObjectPlainText (key → []string, plain/HTML source)
│   ├── dict-object-json-obj.go     # DictObjectJsonObj (key → WordObject with examples/cognates)
│   └── dict-object-html.go         # DictObjectHTML (key → []HTML string) + MergedDictEntry + DictionaryInfo
├── utils/
│   ├── text.go                     # Text utilities (palochka normalization, casing, etc.)
│   └── files.go                    # File I/O helpers (ReadFileLineByLine, SaveDictToJSON)
├── python_scripts/
│   └── process_data.py             # Auxiliary Python processing script
├── content/
│   ├── raw-data-samples/           # Small excerpts for understanding formats
│   ├── phase-01-raw-data/          # Original dictionary files
│   ├── phase-02-json-data/         # Standardized JSON output
│   ├── phase-03-html-data/         # HTML-enriched JSON output
│   ├── phase-04-merged-database/   # Single merged JSON database
│   └── phase-05-sqlite/            # Final SQLite database
├── CLAUDE.md                       # AI assistant instructions (Claude)
└── GEMINI.md                       # AI assistant instructions (Gemini)
```

## The Palochka Problem

Circassian languages use the "palochka" (Ӏ) character, which is visually identical to Latin "I", "l", Turkish "ı"/"İ", Cyrillic "І", and others. Different dictionary sources encode it inconsistently.

**Solution:** All palochka-looking characters in Circassian text are normalized to the digit `1`. This is handled by `utils.ConvertAllPolachkaLookingLettersTo1InCircassianWords()`, which detects Circassian words by looking for Cyrillic letters adjacent to palochka-like characters. This preserves Turkish "i" in Turkish text (since Turkish uses Latin script).

## SQLite Database Schema

The final SQLite database (`dictionary.db`) uses two tables instead of embedding all dictionary metadata in every word entry. This normalization avoids repeating the same dictionary title and language info thousands of times, reducing database size.

### `dictionaries` table
| Column | Type | Description |
|--------|------|-------------|
| `id` | INTEGER PRIMARY KEY | Dictionary ID (0–34) |
| `title` | TEXT | Dictionary name (e.g., "Тхьаркъуахъо (1991)") |
| `from_lang` | TEXT | Source language code (Ady, Kbd, Ru, En, Tr, Ar) |
| `to_lang` | TEXT | Target language code |

### `words` table
| Column | Type | Description |
|--------|------|-------------|
| `word` | TEXT PRIMARY KEY | Lowercased headword |
| `entries` | TEXT | JSON array of `{id, html}` objects |

Each object in the `entries` JSON array has:
- `id` — references `dictionaries.id` for the source dictionary
- `html` — the HTML-formatted definition content

To get a word's entries with full dictionary metadata, join the two tables by matching each entry's `id` to `dictionaries.id`.

## Running

```bash
go run main.go
```

This executes all five conversion phases in sequence. The final output is a SQLite database at `content/phase-05-sqlite/dictionary.db`.

## Requirements

- Go 1.25+
- Dependencies managed via `go.mod` (SQLite via `modernc.org/sqlite` — no CGO required)
