# TLGR

A fast tldr client written in Go that provides quick access to simplified command examples and explanations.

## Why
Why not

## Features
- 🚀 **Fast**: Written in Go for optimal performance
- 💾 **Local Caching**: Cache pages locally for offline access
- 🔄 **Auto-updates**: Keep your command database up to date
- 🖥️ **Cross-platform**: Any go supported platform really

## Flags

| Option | Description |
|--------|-------------|
| `-version` | Print the version |
| `-update` | Update the local cache with latest pages |
| `-clear-cache` | Clear all cached pages |
| `-list` | List all available commands in cache |
| `-platform` | Override the operating system (linux, macos, windows, android, freebsd, netbsd) |

## Platform Support
TLGR supports multiple platforms. By default, it uses your current operating system, but you can override this:

- `linux` (default on Linux)
- `macos` 
- `windows`
- `android`
- `freebsd`
- `netbsd`

## Examples:
```bash
tlgr -update
tlgr -clear-cache
tlgr -list
tlgr git commit
tlgr -platform macos docker
tlgr -platform windows powershell
```

### Development
1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Build: `go build`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
