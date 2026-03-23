# go-html-zip-serve

Servidor HTTP em Go que serve conteúdo HTML a partir de arquivos ZIP.

## Funcionalidades

- Serve documentação HTML diretamente de arquivos ZIP
- Página inicial com lista de zips disponíveis
- Auto-detecção de index.html ao acessar diretórios
- Interface com PicoCSS (tema claro/escuro automático)
- Leitura direta do ZIP sem extração
- Proteção contra path traversal

## Instalação

```bash
go build -o go-html-zip-serve.exe
```

## Uso

1. **Executar o servidor:**
   ```bash
   ./go-html-zip-serve
   ```

2. **Colocar arquivos ZIP na pasta `http/`:**
   ```
   http/
   └── exemplo.zip    → acessível em http://localhost:4000/exemplo/
   ```

3. **Acessar no browser:**
   - `http://localhost:4000/` → Lista todos os zips
   - `http://localhost:4000/exemplo/` → Serve index.html do zip
   - `http://localhost:4000/exemplo/path/arquivo.html` → Arquivo específico

## Estrutura de Arquivos

```
go-html-zip-serve/
├── config.json          # Configuração (porta, pasta)
├── http/                # Colocar arquivos .zip aqui
│   ├── exemplo.zip
│   └── projeto.zip
├── static/              # Arquivos estáticos (CSS)
│   └── pico.min.css
├── main.go
└── go-html-zip-serve.exe
```

## Configuração

Editar `config.json`:

```json
{
  "port": ":4000",
  "httpDir": "http"
}
```

| Opção     | Descrição              | Default |
|-----------|------------------------|---------|
| `port`    | Porta do servidor HTTP | `:4000` |
| `httpDir` | Pasta com os ZIPs      | `http`  |

**Exemplos de `httpDir`:**
- `"http"` → pasta local (default)
- `"documentos"` → outra pasta local
- `"C:/zips"` → caminho absoluto
- `"../shared/zips"` → pasta externa

## MIME Types Suportados

- HTML, CSS, JavaScript
- JSON, XML
- PNG, JPG, GIF, SVG, ICO
- Fontes (WOFF, WOFF2, TTF)

## Desenvolvimento

```bash
# Build
go build -o go-html-zip-serve.exe

# Run
./go-html-zip-serve.exe

# Testar
curl http://localhost:4000/
```
