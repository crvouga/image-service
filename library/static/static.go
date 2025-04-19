package static

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var ErrFileNotFound = errors.New("file not found")
var ErrInvalidSuffix = errors.New("invalid suffix")

var whitelistSuffix = []string{".js", ".css", ".html", ".png", ".jpg", ".jpeg", ".gif", ".svg", ".ico", ".webp", ".bmp"}

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

// GetSiblingRelativePath returns a path relative to the project root directory
// instead of the absolute path that GetSiblingPath returns
func GetSiblingRelativePath(filename string) string {
	_, file, _, _ := runtime.Caller(1)
	absolutePath := filepath.Join(filepath.Dir(file), filename)

	// Get current working directory (project root)
	currentDir, err := os.Getwd()
	if err != nil {
		// Fall back to absolute path if we can't get working directory
		return absolutePath
	}

	// Convert absolute path to path relative to project root
	relativePath, err := filepath.Rel(currentDir, absolutePath)
	if err != nil {
		// Fall back to absolute path if we can't get relative path
		return absolutePath
	}

	// Ensure path uses forward slashes for consistency
	relativePath = filepath.ToSlash(relativePath)

	return "/" + relativePath
}
