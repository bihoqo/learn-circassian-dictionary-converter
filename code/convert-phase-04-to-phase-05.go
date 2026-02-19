package code

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"learn-circassian-helper/modals"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// CallConvertPhase04ToPhase05 reads the merged JSON database and dictionaries
// metadata from Phase 04 and writes them into a SQLite database with two tables:
//   - "dictionaries": one row per dictionary source (id, title, from_lang, to_lang)
//   - "words": one row per word, entries stored as JSON array of {id, html} objects
//
// This normalization avoids repeating dictionary titles and language info in every
// entry, reducing database size significantly.
func CallConvertPhase04ToPhase05() {
	srcDir := "content/phase-04-merged-database"
	mergedPath := filepath.Join(srcDir, "merged-database.json")
	dictsPath := filepath.Join(srcDir, "dictionaries.json")
	distDir := "content/phase-05-sqlite"
	distPath := filepath.Join(distDir, "dictionary.db")

	if err := os.MkdirAll(distDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create output directory: %v", err))
	}

	// Remove existing DB so we start fresh
	os.Remove(distPath)

	mergedData, err := os.ReadFile(mergedPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read %s: %v", mergedPath, err))
	}

	dictsData, err := os.ReadFile(dictsPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read %s: %v", dictsPath, err))
	}

	var merged map[string][]modals.MergedDictEntry
	if err := json.Unmarshal(mergedData, &merged); err != nil {
		panic(fmt.Sprintf("Failed to parse merged JSON: %v", err))
	}

	var dictionaries []modals.DictionaryInfo
	if err := json.Unmarshal(dictsData, &dictionaries); err != nil {
		panic(fmt.Sprintf("Failed to parse dictionaries JSON: %v", err))
	}

	db, err := sql.Open("sqlite", distPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open SQLite: %v", err))
	}
	defer db.Close()

	// Create tables: dictionaries for metadata, words for the actual entries
	_, err = db.Exec(`
		CREATE TABLE dictionaries (
			id INTEGER PRIMARY KEY NOT NULL,
			title TEXT NOT NULL,
			from_lang TEXT NOT NULL,
			to_lang TEXT NOT NULL
		);
		CREATE TABLE words (
			word TEXT PRIMARY KEY NOT NULL,
			entries TEXT NOT NULL
		);
		CREATE INDEX idx_word ON words(word);
	`)
	if err != nil {
		panic(fmt.Sprintf("Failed to create tables: %v", err))
	}

	tx, err := db.Begin()
	if err != nil {
		panic(fmt.Sprintf("Failed to begin transaction: %v", err))
	}

	// Insert dictionaries
	dictStmt, err := tx.Prepare("INSERT INTO dictionaries (id, title, from_lang, to_lang) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(fmt.Sprintf("Failed to prepare dictionaries statement: %v", err))
	}
	defer dictStmt.Close()

	for _, d := range dictionaries {
		if _, err := dictStmt.Exec(d.Id, d.Title, d.FromLang, d.ToLang); err != nil {
			fmt.Printf("Error inserting dictionary %d (%s): %v\n", d.Id, d.Title, err)
			continue
		}
	}

	// Insert words
	wordStmt, err := tx.Prepare("INSERT INTO words (word, entries) VALUES (?, ?)")
	if err != nil {
		panic(fmt.Sprintf("Failed to prepare words statement: %v", err))
	}
	defer wordStmt.Close()

	count := 0
	for word, entries := range merged {
		entriesJSON, err := json.Marshal(entries)
		if err != nil {
			fmt.Printf("Error marshaling entries for %q: %v\n", word, err)
			continue
		}
		if _, err := wordStmt.Exec(word, string(entriesJSON)); err != nil {
			fmt.Printf("Error inserting %q: %v\n", word, err)
			continue
		}
		count++
	}

	if err := tx.Commit(); err != nil {
		panic(fmt.Sprintf("Failed to commit: %v", err))
	}

	fmt.Printf("Phase 04 → Phase 05 complete. SQLite DB: %s (%d words, %d dictionaries)\n", distPath, count, len(dictionaries))
}
