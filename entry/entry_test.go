package entry

import (
	"fmt"
	"testing"
)

func TestDesktopEntry_FillDefault(t *testing.T) {
	de := DesktopEntry{}
	de.FillDefault()

	if de.Type != "Application" {
		t.Fail()
	}

	if de.Version != SpecsVersion {
		t.Fail()
	}
}

func TestDesktopEntry_StringEmpty(t *testing.T) {
	de := DesktopEntry{}
	de.FillDefault()
	fmt.Print(de.String())
	if de.String() != fmt.Sprintf(`[Desktop Entry]
Type=Application
Version=%v
`, SpecsVersion) {
		t.Fail()
	}
}

func TestDesktopEntry_String(t *testing.T) {
	de := DesktopEntry{
		Name:        "Mozilla Firefox",
		GenericName: "Web browser",
		Comment:     "Software to browse the web",
		Categories:  "Network;WebBrowser",
		Exec:        "/opt/firefox/firefox-bin",
		Icon:        "firefox",
	}
	de.FillDefault()

	if de.String() != fmt.Sprintf(`[Desktop Entry]
Type=Application
Version=%v
Name=Mozilla Firefox
GenericName=Web browser
Comment=Software to browse the web
Categories=Network;WebBrowser
Exec=/opt/firefox/firefox-bin
Icon=firefox
`, SpecsVersion) {
		t.Fail()
	}
}
