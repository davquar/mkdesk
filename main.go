package main

import (
	"fmt"
	"mkdesk/entry"
	"mkdesk/utils"
	"os"
	"path"
)

const LocalEntriesDir = ".local/share/applications/"

func main() {
	de := entry.DesktopEntry{}
	de.FillDefault()

	err := de.InteractiveFill(true)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error while filling the desktop entry: %v\n", err)
		os.Exit(1)
	}

	baseDir, err := os.UserHomeDir()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error while retrieving the home directory: %v", err)
		os.Exit(1)
	}

	err = SaveDesktopEntry(de, baseDir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error while saving the desktop entry: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Desktop entry saved.")
}

// SaveDesktopEntry writes the string representation of the given SaveDesktopEntry
// in the appropriate path, using a cleaned up SaveDesktopEntry.Name as filename.
func SaveDesktopEntry(e entry.DesktopEntry, baseDir string) error {
	fileName, err := utils.CleanFileName(e.Name)
	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(baseDir, LocalEntriesDir, fileName))
	if err != nil {
		return err
	}

	_, err = file.WriteString(e.String())
	if err != nil {
		return err
	}

	return nil
}
