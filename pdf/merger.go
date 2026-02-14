package pdf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// ValidatePDF checks if a file is a valid PDF
func ValidatePDF(path string) error {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", path)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".pdf" {
		return fmt.Errorf("file is not a PDF: %s", filepath.Base(path))
	}

	// Validate PDF structure using pdfcpu
	if err := api.ValidateFile(path, nil); err != nil {
		return fmt.Errorf("invalid or corrupted PDF: %s", filepath.Base(path))
	}

	return nil
}

// MergePDFs combines multiple PDF files into a single output file
// progressCallback is called periodically to update UI (can be nil)
func MergePDFs(inputPaths []string, outputPath string, progressCallback func()) error {
	if len(inputPaths) < 2 {
		return fmt.Errorf("at least 2 PDF files are required for merging")
	}

	if len(inputPaths) > 100 {
		return fmt.Errorf("maximum 100 PDF files allowed")
	}

	// Validate all input files first
	for _, path := range inputPaths {
		if err := ValidatePDF(path); err != nil {
			return err
		}
	}

	// Call progress callback before starting
	if progressCallback != nil {
		progressCallback()
	}

	// Create default configuration for pdfcpu
	conf := model.NewDefaultConfiguration()

	// Perform the merge operation
	// The 'false' parameter means: don't create divider pages between merged PDFs
	if err := api.MergeCreateFile(inputPaths, outputPath, false, conf); err != nil {
		return fmt.Errorf("failed to merge PDFs: %v", err)
	}

	// Call progress callback after completion
	if progressCallback != nil {
		progressCallback()
	}

	return nil
}
