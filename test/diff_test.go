package test

import (
	"cli-text-compare/pkg"
	"os"
	"strings"
	"testing"
	"time"
)

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	return string(data)
}

func TestDiff(t *testing.T) {
	testCases := []struct {
		name            string
		oldFile         string
		newFile         string
		expectAllMatch  bool
		expectAllInsert bool
		expectAllDelete bool
	}{
		{
			name:    "empty to empty",
			oldFile: "./emptyOld.txt",
			newFile: "./emptyNew.txt",
		},
		{
			name:            "empty to short",
			oldFile:         "./emptyOld.txt",
			newFile:         "./shortNew.txt",
			expectAllInsert: true,
		},
		{
			name:            "empty to long",
			oldFile:         "./emptyOld.txt",
			newFile:         "./longNew.txt",
			expectAllInsert: true,
		},
		{
			name:            "empty to weird",
			oldFile:         "./emptyOld.txt",
			newFile:         "./weirdNew.txt",
			expectAllInsert: true,
		},
		{
			name:            "short to empty",
			oldFile:         "./shortOld.txt",
			newFile:         "./emptyNew.txt",
			expectAllDelete: true,
		},
		{
			name:            "long to empty",
			oldFile:         "./longOld.txt",
			newFile:         "./emptyNew.txt",
			expectAllDelete: true,
		},
		{
			name:            "weird to empty",
			oldFile:         "./weirdOld.txt",
			newFile:         "./emptyNew.txt",
			expectAllDelete: true,
		},
		{
			name:           "identical to identical",
			oldFile:        "./identicalOld.txt",
			newFile:        "./identicalNew.txt",
			expectAllMatch: true,
		},
		{
			name:    "short to short",
			oldFile: "./shortOld.txt",
			newFile: "./shortNew.txt",
		},
		{
			name:    "long to long",
			oldFile: "./longOld.txt",
			newFile: "./longNew.txt",
		},
		{
			name:    "weird to weird",
			oldFile: "./weirdOld.txt",
			newFile: "./weirdNew.txt",
		},
		{
			name:    "short to long",
			oldFile: "./shortOld.txt",
			newFile: "./longNew.txt",
		},
		{
			name:    "long to short",
			oldFile: "./longOld.txt",
			newFile: "./shortNew.txt",
		},
		{
			name:    "short to weird",
			oldFile: "./shortOld.txt",
			newFile: "./weirdNew.txt",
		},
		{
			name:    "weird to short",
			oldFile: "./weirdOld.txt",
			newFile: "./shortNew.txt",
		},
		{
			name:    "weird to long",
			oldFile: "./weirdOld.txt",
			newFile: "./longNew.txt",
		},
		{
			name:    "long to weird",
			oldFile: "./longOld.txt",
			newFile: "./weirdNew.txt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldContent := readFile(t, tc.oldFile)
			newContent := readFile(t, tc.newFile)

			oldChars := strings.Split(oldContent, "")
			newChars := strings.Split(newContent, "")

			// Filter out the empty string that Split produces for empty input.
			if oldContent == "" {
				oldChars = nil
			}
			if newContent == "" {
				newChars = nil
			}

			start := time.Now()
			actions := pkg.Diff(oldContent, newContent)
			elapsed := time.Since(start)
			t.Logf("Diff took %v ms", elapsed.Milliseconds())

			// Universal Levenshtein invariant:
			//   matches + substitutes + deletes == len(oldChars)
			//   matches + substitutes + inserts == len(newChars)
			var nMatch, nInsert, nDelete, nSub int
			for _, a := range actions {
				switch a {
				case pkg.ActionMatch:
					nMatch++
				case pkg.ActionInsert:
					nInsert++
				case pkg.ActionDelete:
					nDelete++
				case pkg.ActionSubstitute:
					nSub++
				}
			}
			if got := nMatch + nSub + nDelete; got != len(oldChars) {
				t.Errorf("matches+substitutes+deletes = %d, want len(old) = %d", got, len(oldChars))
			}
			if got := nMatch + nSub + nInsert; got != len(newChars) {
				t.Errorf("matches+substitutes+inserts = %d, want len(new) = %d", got, len(newChars))
			}

			// All-match: identical content produces no edits.
			if tc.expectAllMatch {
				if nInsert+nDelete+nSub > 0 {
					t.Errorf("expected all matches but got %d inserts, %d deletes, %d substitutes",
						nInsert, nDelete, nSub)
				}
				if nMatch != len(oldChars) {
					t.Errorf("got %d matches, want %d", nMatch, len(oldChars))
				}
			}

			// All-insert: empty old, every new char must be an insert.
			if tc.expectAllInsert {
				if nMatch+nDelete+nSub > 0 {
					t.Errorf("expected all inserts but got %d matches, %d deletes, %d substitutes",
						nMatch, nDelete, nSub)
				}
				if nInsert != len(newChars) {
					t.Errorf("got %d inserts, want %d", nInsert, len(newChars))
				}
			}

			// All-delete: empty new, every old char must be a delete.
			if tc.expectAllDelete {
				if nMatch+nInsert+nSub > 0 {
					t.Errorf("expected all deletes but got %d matches, %d inserts, %d substitutes",
						nMatch, nInsert, nSub)
				}
				if nDelete != len(oldChars) {
					t.Errorf("got %d deletes, want %d", nDelete, len(oldChars))
				}
			}

			// For non-identical, non-empty pairs: verify the diff actually contains edits.
			if !tc.expectAllMatch && !tc.expectAllInsert && !tc.expectAllDelete {
				if len(oldChars) > 0 && len(newChars) > 0 && nInsert+nDelete+nSub == 0 {
					t.Errorf("expected some edits between different files but got all matches")
				}
			}
		})
	}
}
