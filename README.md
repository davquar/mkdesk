# mkdesk

Simple interactive CLI utility to quickly create **Desktop Entries** according to (a tiny tiny subset of) the Freedesktop specifications.

## Usage (interactive)

```shell
$ ./mkdesk
Create your desktop entry

Field:  Name (e.g. Mozilla Firefox)
Value:  MarkText

Field:  GenericName (e.g. Web Browser)
Value:  Markdown editor

Field:  Comment (e.g. Software to browse the web)
Value:     

Field:  Categories (e.g. Network;WebBrowser)
Value:  Office

Field:  Exec (e.g. /path/to/firefox)
Value:  /opt/marktext.AppImage                

Field:  Icon (e.g. /path/to/firefox/icon.png)
Value:  /opt/.icons/marktext.png

Desktop entry saved.
```

The desktop entry will be saved in `~/.local/share/applications`.

## Building

```shell
go build mkdesk
```

## License

MIT