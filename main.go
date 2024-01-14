package main

import (
	"flag"
	"fmt"
	"mkdesk/entry"
	"mkdesk/utils"
	"os"
	"path"
)

const LocalEntriesDir = ".local/share/applications/"

func main() {
	de := entry.DesktopEntry{}
	dryRun := false

	flag.StringVar(&de.Name, "name", "", "Name")
	flag.StringVar(&de.GenericName, "generic-name", "", "Generic name")
	flag.StringVar(&de.Comment, "comment", "", "Comment")
	flag.StringVar(&de.Categories, "categories", "", "Semicolon-separated list of categories")
	flag.StringVar(&de.Exec, "exec", "", "Executable path")
	flag.StringVar(&de.Icon, "icon", "", "Icon path")
	flag.BoolVar(&dryRun, "dry-run", false, "Just print the final desktop entry, without saving it")
	flag.Parse()

	de.FillDefault()

	err := de.InteractiveFill(true)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error while filling the desktop entry: %v\n", err)
		os.Exit(1)
	}

	if dryRun {
		fmt.Println()
		fmt.Println(de.String())
		os.Exit(0)
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
