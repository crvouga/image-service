package static

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var ErrFileNotFound = errors.New("file not found")
var ErrInvalidSuffix = errors.New("invalid suffix")

var whitelistSuffix = []string{".js", ".css", ".html"}

func hasValidSuffix(requestPath string) bool {
	for _, suffix := range whitelistSuffix {
		if strings.HasSuffix(requestPath, suffix) {
			return true
		}
	}
	return false
}

func ServeStaticAssets(w http.ResponseWriter, r *http.Request) error {
	requestPath := r.URL.Path

	log.Println("requestPath " + requestPath)

	if !hasValidSuffix(requestPath) {
		return ErrInvalidSuffix
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(currentDir, requestPath)

	info, err := os.Stat(filePath)
	if err != nil {
		return ErrFileNotFound
	}

	if info.IsDir() {
		return ErrFileNotFound
	}

	http.ServeFile(w, r, filePath)
	return nil
}

func GetSiblingPath(filename string) string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(file), filename)
}
