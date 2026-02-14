package ui

import (
	"fmt"
	"path/filepath"
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
	app         fyne.App
	window      fyne.Window
	pdfFiles    []string
	fileList    *widget.List
	mergeButton *widget.Button
	addButton   *widget.Button
	clearButton *widget.Button
	statusLabel *widget.Label
}

// NewPDFMergerApp creates and initializes the application
func NewPDFMergerApp() *PDFMergerApp {
	a := app.New()
	w := a.NewWindow("mergit - PDF Merger")

	pdfApp := &PDFMergerApp{
		app:      a,
		window:   w,
		pdfFiles: []string{},
	}

	pdfApp.setupUI()
	return pdfApp
}

// setupUI creates and configures the user interface
func (p *PDFMergerApp) setupUI() {
	// Status label showing file count
	p.statusLabel = widget.NewLabel("Drop PDF files here or click 'Add PDF'")
	p.statusLabel.Alignment = fyne.TextAlignCenter

	// Create file list widget with up/down buttons for reordering
	p.fileList = widget.NewList(
		func() int {
			return len(p.pdfFiles)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(theme.DocumentIcon())
			label := widget.NewLabel("Template")

			upBtn := widget.NewButtonWithIcon("", theme.MoveUpIcon(), nil)
			upBtn.Importance = widget.LowImportance

			downBtn := widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil)
			downBtn.Importance = widget.LowImportance

			removeBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
			removeBtn.Importance = widget.LowImportance

			buttons := container.NewHBox(upBtn, downBtn, removeBtn)
			return container.NewBorder(nil, nil, icon, buttons, label)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			border := item.(*fyne.Container)
			label := border.Objects[0].(*widget.Label)
			buttonsBox := border.Objects[2].(*fyne.Container)

			upBtn := buttonsBox.Objects[0].(*widget.Button)
			downBtn := buttonsBox.Objects[1].(*widget.Button)
			removeBtn := buttonsBox.Objects[2].(*widget.Button)

			filename := filepath.Base(p.pdfFiles[id])
			label.SetText(filename)

			// Set up up button
			upBtn.OnTapped = func() {
				p.moveUp(id)
			}
			// Disable up button for first item
			if id == 0 {
				upBtn.Disable()
			} else {
				upBtn.Enable()
			}

			// Set up down button
			downBtn.OnTapped = func() {
				p.moveDown(id)
			}
			// Disable down button for last item
			if id == len(p.pdfFiles)-1 {
				downBtn.Disable()
			} else {
				downBtn.Enable()
			}

			// Set up remove button
			removeBtn.OnTapped = func() {
				p.removeFile(id)
			}
		},
	)

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
		p.statusLabel, // top
		buttonBox,     // bottom
		nil,           // left
		nil,           // right
		p.fileList,    // center
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

// moveUp moves a file up one position in the list
func (p *PDFMergerApp) moveUp(index int) {
	if index <= 0 || index >= len(p.pdfFiles) {
		return
	}
	// Swap with previous item
	p.pdfFiles[index], p.pdfFiles[index-1] = p.pdfFiles[index-1], p.pdfFiles[index]
	p.updateUI()
}

// moveDown moves a file down one position in the list
func (p *PDFMergerApp) moveDown(index int) {
	if index < 0 || index >= len(p.pdfFiles)-1 {
		return
	}
	// Swap with next item
	p.pdfFiles[index], p.pdfFiles[index+1] = p.pdfFiles[index+1], p.pdfFiles[index]
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
