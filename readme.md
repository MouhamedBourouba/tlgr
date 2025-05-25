# TLGR

A fast tldr client written in Go that provides quick access to simplified command examples and explanations.

## About

TLGR is a lightweight, fast alternative to the traditional `man` pages, providing concise and practical examples for common command-line tools. Instead of reading through lengthy manual pages, get straight to the point with real-world usage examples.

## Features

- 🚀 **Fast**: Written in Go for optimal performance
- 💾 **Local Caching**: Cache pages locally for offline access
- 🔄 **Auto-updates**: Keep your command database up to date
- 🖥️ **Cross-platform**: Support for Linux, macOS, Windows, Android, FreeBSD, and NetBSD
- 📋 **List Commands**: Browse all available commands in your cache
- 🧹 **Cache Management**: Clear cache when needed

### Options

| Option | Description |
|--------|-------------|
| `-help` | Print help information |
| `-version` | Print the version |
| `-update` | Update the local cache with latest pages |
| `-clear-cache` | Clear all cached pages |
| `-list` | List all available commands in cache |
| `-platform` | Override the operating system (linux, macos, windows, android, freebsd, netbsd) |

### Examples

```bash
# Get examples for git commit
tlgr git commit

# Update your local cache
tlgr -update

# List all available commands
tlgr -list

# Get examples for a specific platform
tlgr -platform macos docker

# Clear cache and start fresh
tlgr -clear-cache
```

## First Run

On your first run, use the update command to populate your local cache:
```bash
tlgr -update
```

This will download the latest tldr pages and cache them locally for fast access.

## Cache Management

TLGR maintains a local cache of tldr pages for faster access. You can manage this cache using:

- **Update cache**: `tlgr -update` - Downloads the latest pages
- **Clear cache**: `tlgr -clear-cache` - Removes all cached pages
- **List cached commands**: `tlgr -list` - Shows all available commands

## Platform Support

TLGR supports multiple platforms. By default, it uses your current operating system, but you can override this:

- `linux` (default on Linux)
- `macos` 
- `windows`
- `android`
- `freebsd`
- `netbsd`

Example:
```bash
tlgr -platform windows powershell
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Build: `go build`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

⭐ If you find TLGR useful, please consider giving it a star on GitHub!
