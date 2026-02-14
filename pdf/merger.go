package pdf

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	return nil
}

// MergePDFs combines multiple PDF files into a single output file using Ghostscript
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

	// Use Ghostscript to merge PDFs - it's very robust with problematic PDFs
	// Build the Ghostscript command
	args := []string{
		"-dBATCH",                    // Exit after processing
		"-dNOPAUSE",                  // Don't pause between pages
		"-q",                         // Quiet mode
		"-sDEVICE=pdfwrite",          // Output device
		"-dPDFSETTINGS=/default",     // Default quality settings
		"-sOutputFile=" + outputPath, // Output file
	}

	// Add all input files
	args = append(args, inputPaths...)

	// Execute Ghostscript
	cmd := exec.Command("gs", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Clean up partial output file if it exists
		os.Remove(outputPath)
		return fmt.Errorf("failed to merge PDFs: %v\nOutput: %s", err, string(output))
	}

	// Verify output file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return fmt.Errorf("merge completed but output file was not created")
	}

	// Call progress callback after completion
	if progressCallback != nil {
		progressCallback()
	}

	return nil
}
