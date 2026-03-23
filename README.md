# go-html-zip-serve

Go HTTP server that serves HTML content from ZIP files.

## Features

- Serve HTML documentation directly from ZIP archives
- Home page with list of available ZIPs
- Auto-detection of index.html when accessing directories
- PicoCSS interface (automatic light/dark theme)
- Direct ZIP reading without extraction
- Path traversal protection

## Installation

```bash
go build -o go-html-zip-serve.exe
```

## Usage

1. **Start the server:**
   ```bash
   ./go-html-zip-serve
   ```

2. **Place ZIP files in the `http/` folder:**
   ```
   http/
   └── example.zip    → accessible at http://localhost:4000/example/
   ```

3. **Access in browser:**
   - `http://localhost:4000/` → Lists all ZIPs
   - `http://localhost:4000/example/` → Serves index.html from ZIP
   - `http://localhost:4000/example/path/file.html` → Specific file

## File Structure

```
go-html-zip-serve/
├── config.json          # Configuration (port, directory)
├── http/                # Place .zip files here
│   ├── example.zip
│   └── project.zip
├── static/              # Static files (CSS)
│   └── pico.min.css
├── main.go
└── go-html-zip-serve.exe
```

## Configuration

Edit `config.json`:

```json
{
  "port": ":4000",
  "httpDir": "http"
}
```

| Option    | Description            | Default |
|-----------|------------------------|---------|
| `port`    | HTTP server port       | `:4000` |
| `httpDir` | Directory with ZIPs    | `http`  |

**Examples of `httpDir`:**
- `"http"` → local folder (default)
- `"documents"` → another local folder
- `"C:/zips"` → absolute path
- `"../shared/zips"` → external folder

## Supported MIME Types

- HTML, CSS, JavaScript
- JSON, XML
- PNG, JPG, GIF, SVG, ICO
- Fonts (WOFF, WOFF2, TTF)

## Development

```bash
# Build
go build -o go-html-zip-serve.exe

# Run
./go-html-zip-serve.exe

# Test
curl http://localhost:4000/
```
