package entry

import (
	"fmt"
	"mkdesk/utils"
	"reflect"
	"strings"
)

const (
	SpecsVersion = "1.1"
)

type DesktopEntry struct {
	Type        string `hint:"Application"`
	Version     string `hint:"1.1"`
	Name        string `hint:"Mozilla Firefox"`
	GenericName string `hint:"Web Browser"`
	Comment     string `hint:"Software to browse the web"`
	Categories  string `hint:"Network;WebBrowser"`
	Exec        string `hint:"/path/to/firefox"`
	Icon        string `hint:"/path/to/firefox/icon.png"`
}

// FillDefault prepares DesktopEntry by filling some default values.
func (de *DesktopEntry) FillDefault() {
	de.Type = "Application"
	de.Version = SpecsVersion
}

// InteractiveFill starts an interactive session to ask the user to
// set the value of each field of DesktopEntry.
func (de *DesktopEntry) InteractiveFill(skipAlreadyFilled bool) error {
	ref := reflect.ValueOf(*de)

	fmt.Println("Create your desktop entry")
	for idx := 0; idx < ref.NumField(); idx++ {
		if skipAlreadyFilled && len(ref.Field(idx).String()) != 0 {
			continue
		}

		if idx > 0 {
			fmt.Println()
		}

		key := ref.Type().Field(idx).Name
		hint := ref.Type().Field(idx).Tag.Get("hint")

		fmt.Printf("Field:\t%s (e.g. %s)\n", key, hint)
		fmt.Print("Value:\t")

		value, err := utils.StringInput()
		if err != nil {
			return err
		}

		reflect.Indirect(reflect.ValueOf(de)).Field(idx).SetString(value)
	}

	return nil
}

// String returns the string representation of DesktopEntry.
func (de *DesktopEntry) String() string {
	sb := strings.Builder{}
	ref := reflect.ValueOf(*de)

	sb.WriteString("[Desktop Entry]\n")

	for idx := 0; idx < ref.NumField(); idx++ {
		key := ref.Type().Field(idx).Name
		value := ref.Field(idx).String()

		if value != "" {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		}
	}

	return sb.String()
}
