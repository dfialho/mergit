package ui

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"mergit/pdf"
)

const maxPDFs = 100

type PDFMergerApp struct {
	app           fyne.App
	window        fyne.Window
	pdfFiles      []string
	fileList      *widget.List
	mergeButton   *widget.Button
	addButton     *widget.Button
	clearButton   *widget.Button
	statusLabel   *widget.Label
	dragStartIdx  int
	isDragging    bool
}

// NewPDFMergerApp creates and initializes the application
func NewPDFMergerApp() *PDFMergerApp {
	a := app.New()
	w := a.NewWindow("mergit - PDF Merger")

	pdfApp := &PDFMergerApp{
		app:          a,
		window:       w,
		pdfFiles:     []string{},
		dragStartIdx: -1,
		isDragging:   false,
	}

	pdfApp.setupUI()
	return pdfApp
}

// setupUI creates and configures the user interface
func (p *PDFMergerApp) setupUI() {
	// Status label showing file count
	p.statusLabel = widget.NewLabel("Drop PDF files here or click 'Add PDF'")
	p.statusLabel.Alignment = fyne.TextAlignCenter

	// Create file list widget with drag-and-drop reordering
	p.fileList = widget.NewList(
		func() int {
			return len(p.pdfFiles)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(theme.DocumentIcon())
			label := widget.NewLabel("Template")
			removeBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
			removeBtn.Importance = widget.LowImportance

			return container.NewBorder(nil, nil, icon, removeBtn, label)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			border := item.(*fyne.Container)
			label := border.Objects[0].(*widget.Label)
			removeBtn := border.Objects[2].(*widget.Button)

			filename := filepath.Base(p.pdfFiles[id])
			label.SetText(filename)

			// Set up remove button for this specific item
			removeBtn.OnTapped = func() {
				p.removeFile(id)
			}
		},
	)

	// Add drag and drop support for reordering
	p.fileList.OnSelected = func(id widget.ListItemID) {
		p.dragStartIdx = id
		p.isDragging = true
	}

	p.fileList.OnUnselected = func(id widget.ListItemID) {
		if p.isDragging && p.dragStartIdx != -1 && p.dragStartIdx != id {
			// Reorder the files
			p.reorderFiles(p.dragStartIdx, id)
		}
		p.isDragging = false
		p.dragStartIdx = -1
	}

	// Buttons
	p.addButton = widget.NewButtonWithIcon("Add PDF", theme.FileIcon(), func() {
		p.showAddFileDialog()
	})

	p.clearButton = widget.NewButton("Clear All", func() {
		p.clearAllFiles()
	})

	p.mergeButton = widget.NewButtonWithIcon("Merge PDFs", theme.DocumentSaveIcon(), func() {
		p.startMerge()
	})
	p.mergeButton.Importance = widget.HighImportance
	p.mergeButton.Disable() // Initially disabled

	// Button container
	buttonBox := container.NewHBox(
		p.addButton,
		p.clearButton,
		p.mergeButton,
	)

	// Main layout
	content := container.NewBorder(
		p.statusLabel,   // top
		buttonBox,       // bottom
		nil,             // left
		nil,             // right
		p.fileList,      // center
	)

	// Set up drag and drop for the window
	p.window.SetContent(content)
	p.window.Resize(fyne.NewSize(800, 600))
	p.window.SetOnDropped(func(_ fyne.Position, uris []fyne.URI) {
		p.handleDroppedFiles(uris)
	})

	p.updateUI()
}

// showAddFileDialog opens a file picker to add PDF files
func (p *PDFMergerApp) showAddFileDialog() {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, p.window)
			return
		}
		if reader == nil {
			return // User cancelled
		}
		defer reader.Close()

		uri := reader.URI()
		p.addFile(uri.Path())
	}, p.window)

	fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
	fd.Show()
}

// handleDroppedFiles processes files dropped onto the window
func (p *PDFMergerApp) handleDroppedFiles(uris []fyne.URI) {
	for _, uri := range uris {
		path := uri.Path()
		p.addFile(path)
	}
}

// addFile adds a PDF file to the list
func (p *PDFMergerApp) addFile(path string) {
	// Check max limit
	if len(p.pdfFiles) >= maxPDFs {
		dialog.ShowError(fmt.Errorf("maximum %d PDF files allowed", maxPDFs), p.window)
		return
	}

	// Validate PDF
	if err := pdf.ValidatePDF(path); err != nil {
		dialog.ShowError(err, p.window)
		return
	}

	// Add to list
	p.pdfFiles = append(p.pdfFiles, path)
	p.updateUI()
}

// removeFile removes a file from the list
func (p *PDFMergerApp) removeFile(index int) {
	if index < 0 || index >= len(p.pdfFiles) {
		return
	}

	p.pdfFiles = append(p.pdfFiles[:index], p.pdfFiles[index+1:]...)
	p.updateUI()
}

// clearAllFiles removes all files from the list
func (p *PDFMergerApp) clearAllFiles() {
	p.pdfFiles = []string{}
	p.updateUI()
}

// reorderFiles moves a file from one position to another
func (p *PDFMergerApp) reorderFiles(from, to int) {
	if from == to || from < 0 || to < 0 || from >= len(p.pdfFiles) || to >= len(p.pdfFiles) {
		return
	}

	// Remove from old position
	file := p.pdfFiles[from]
	p.pdfFiles = append(p.pdfFiles[:from], p.pdfFiles[from+1:]...)

	// Insert at new position
	if to > from {
		to-- // Adjust index because we removed an element
	}
	p.pdfFiles = append(p.pdfFiles[:to], append([]string{file}, p.pdfFiles[to:]...)...)

	p.updateUI()
}

// startMerge begins the PDF merging process
func (p *PDFMergerApp) startMerge() {
	if len(p.pdfFiles) < 2 {
		dialog.ShowInformation("Not Enough Files", "Please add at least 2 PDF files to merge.", p.window)
		return
	}

	// Show save dialog with default filename
	defaultName := p.getDefaultFilename()

	fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, p.window)
			return
		}
		if writer == nil {
			return // User cancelled
		}
		defer writer.Close()

		outputPath := writer.URI().Path()
		p.performMerge(outputPath)
	}, p.window)

	fd.SetFileName(defaultName)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
	fd.Show()
}

// performMerge executes the actual merge operation
func (p *PDFMergerApp) performMerge(outputPath string) {
	// Create progress dialog
	progressBar := widget.NewProgressBarInfinite()
	progressDialog := dialog.NewCustom("Merging PDFs", "Please wait...", progressBar, p.window)
	progressDialog.Show()

	// Perform merge in a goroutine to keep UI responsive
	go func() {
		err := pdf.MergePDFs(p.pdfFiles, outputPath, nil)

		// Close progress dialog
		progressDialog.Hide()

		// Show result on main thread
		if err != nil {
			dialog.ShowError(err, p.window)
		} else {
			// Show success dialog with file path
			successMsg := fmt.Sprintf("PDF merged successfully!\n\nSaved to:\n%s", outputPath)
			dialog.ShowInformation("Success", successMsg, p.window)

			// Clear the list (reset to beginning)
			p.clearAllFiles()
		}
	}()
}

// getDefaultFilename generates a default filename with today's date
func (p *PDFMergerApp) getDefaultFilename() string {
	today := time.Now().Format("2006-01-02")
	return fmt.Sprintf("merged_%s.pdf", today)
}

// updateUI refreshes the UI state
func (p *PDFMergerApp) updateUI() {
	// Update status label
	count := len(p.pdfFiles)
	if count == 0 {
		p.statusLabel.SetText("Drop PDF files here or click 'Add PDF'")
	} else {
		p.statusLabel.SetText(fmt.Sprintf("%d PDF%s added (max %d)", count, pluralize(count), maxPDFs))
	}

	// Enable/disable merge button (requires at least 2 PDFs)
	if count >= 2 {
		p.mergeButton.Enable()
	} else {
		p.mergeButton.Disable()
	}

	// Refresh file list
	p.fileList.Refresh()
}

// pluralize returns "s" for counts != 1, empty string otherwise
func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

// Run starts the application
func (p *PDFMergerApp) Run() {
	p.window.ShowAndRun()
}
