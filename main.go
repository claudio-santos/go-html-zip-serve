// go-html-zip-serve - Servidor HTTP que serve conteúdo HTML a partir de arquivos ZIP
package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// config armazena a configuração do servidor
type config struct {
	Port    string `json:"port"`
	HTTPDir string `json:"httpDir"`
}

// cfg é a configuração global com valores default
var cfg = &config{
	Port:    ":4000",
	HTTPDir: "http",
}

func main() {
	loadConfig()
	os.MkdirAll(cfg.HTTPDir, 0755)
	os.MkdirAll("static", 0755)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)

	fmt.Printf("Servidor em http://localhost%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}

// loadConfig lê o config.json e atualiza a configuração global
func loadConfig() {
	f, err := os.Open("config.json")
	if err != nil {
		return
	}
	defer f.Close()
	json.NewDecoder(f).Decode(cfg)
}

// handler despacha as requisições para serveIndex ou serveZip
func handler(w http.ResponseWriter, r *http.Request) {
	p := strings.Trim(r.URL.Path, "/")
	if p == "" {
		serveIndex(w)
		return
	}
	serveZip(w, p)
}

// serveIndex gera a página HTML com lista de zips disponíveis
func serveIndex(w http.ResponseWriter) {
	entries, _ := os.ReadDir(cfg.HTTPDir)
	var zips []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(strings.ToLower(e.Name()), ".zip") {
			zips = append(zips, strings.TrimSuffix(e.Name(), filepath.Ext(e.Name())))
		}
	}
	sort.Strings(zips)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<!DOCTYPE html><html lang="pt"><head><meta charset="UTF-8"><title>go-html-zip-serve</title><link rel="stylesheet" href="/static/pico.min.css"></head><body><main class="container"><h1>Documentos</h1><ul>`)
	if len(zips) == 0 {
		fmt.Fprintf(w, `<li><em>Nenhum zip em %s/</em></li>`, cfg.HTTPDir)
	} else {
		for _, z := range zips {
			fmt.Fprintf(w, `<li><a href="/%s/">%s</a></li>`, z, z)
		}
	}
	fmt.Fprint(w, `</ul></main></body></html>`)
}

// serveZip serve um arquivo específico de dentro de um zip
func serveZip(w http.ResponseWriter, p string) {
	parts := strings.Split(p, "/")
	zipName := parts[0]
	filePath := strings.Join(parts[1:], "/")

	if filePath == "" || strings.HasSuffix(p, "/") {
		filePath = "index.html"
	}

	filePath = path.Clean(filePath)
	if strings.HasPrefix(filePath, "..") {
		http.Error(w, "Acesso negado", http.StatusForbidden)
		return
	}

	r, err := zip.OpenReader(filepath.Join(cfg.HTTPDir, zipName+".zip"))
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Zip não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		if filepath.ToSlash(f.Name) == filePath {
			rc, err := f.Open()
			if err != nil {
				http.Error(w, "Erro", http.StatusInternalServerError)
				return
			}
			defer rc.Close()
			w.Header().Set("Content-Type", mimeByExt(filePath))
			io.Copy(w, rc)
			return
		}
	}
	http.Error(w, "Arquivo não encontrado", http.StatusNotFound)
}

// mimeByExt retorna o MIME type baseado na extensão do arquivo
func mimeByExt(p string) string {
	if t := mime.TypeByExtension(strings.ToLower(filepath.Ext(p))); t != "" {
		return t
	}
	switch strings.ToLower(filepath.Ext(p)) {
	case ".html", ".htm":
		return "text/html; charset=utf-8"
	case ".css":
		return "text/css; charset=utf-8"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}
