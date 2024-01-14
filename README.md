# mkdesk

Simple interactive CLI utility to quickly create **Desktop Entries** according to (a tiny tiny subset of) the Freedesktop specifications.

## Usage

mkdesk supports both **flags** and **interactive** mode. The two modalities can be used **simultaneously**: the interactive mode will just skip arguments supplied via flags and ask you the remaining ones.

### Supported flags

```
--name           Name
--generic-name   Generic name
--comment        Comment
--exec           Executable path
--icon           Icon path
--categories     Semicolon-separated list of categories
--dry-run        Just print the final desktop entry, without saving it
```

The desktop entry will be saved in `~/.local/share/applications`.

## Building

```shell
go build mkdesk
```

## License

MIT
