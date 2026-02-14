# mergit - PDF Merger

A simple, standalone PDF merger application for Linux and Windows. Drag, drop, reorder, and merge multiple PDF files into a single document.

## Features

✅ **Single binary** - No installation required, just run the executable  
✅ **Cross-platform** - Works on Linux and Windows  
✅ **Drag-and-drop interface** - Add PDF files by dragging them into the app  
✅ **Reorder PDFs** - Click and drag to reorder files in the list  
✅ **Simple workflow** - Add files, order them, click merge, choose destination  
✅ **Progress indication** - Visual feedback during merge operation  
✅ **Auto-reset** - Automatically clears after successful merge  
✅ **File validation** - Ensures only valid PDFs are added  
✅ **Maximum 100 PDFs** - Supports merging up to 100 documents  

## Download

Pre-built binaries will be available in the `build/` directory after compilation:
- **Linux:** `build/mergit`
- **Windows:** `build/mergit.exe`

## Usage

### Running the Application

**On Linux:**
```bash
./build/mergit
```

**On Windows:**
Simply double-click `mergit.exe`

### How to Merge PDFs

1. **Launch mergit** - Open the application
2. **Add PDF files** - Either:
   - Drag and drop PDF files from your file manager into the window
   - Click the "Add PDF" button to browse for files
3. **Reorder files** (optional) - Click and drag files in the list to change their order
4. **Remove files** (optional) - Click the [X] button next to any file to remove it
5. **Merge** - Click the "Merge PDFs" button (enabled when you have 2+ PDFs)
6. **Choose destination** - Select where to save your merged PDF
   - Default filename: `merged_YYYY-MM-DD.pdf` (today's date)
7. **Done!** - A success dialog will show the saved file path
   - The list automatically clears, ready for the next merge

## Building from Source

### Prerequisites

#### Required for Linux builds:
```bash
# On Ubuntu/Debian:
sudo apt-get update
sudo apt-get install -y golang gcc libgl1-mesa-dev xorg-dev

# Verify Go version (1.21 or higher required):
go version
```

#### Additional for Windows cross-compilation:
```bash
# On Ubuntu/Debian:
sudo apt-get install -y gcc-mingw-w64
```

### Install Dependencies

```bash
# Navigate to the project directory
cd mergit

# Download Go dependencies
go mod download
```

### Build Commands

#### Build for Linux:
```bash
# Using the build script (recommended):
chmod +x build/build-linux.sh
./build/build-linux.sh

# Or manually:
go build -ldflags="-s -w" -o build/mergit
```

#### Build for Windows (from Linux):
```bash
# Using the build script (recommended):
chmod +x build/build-windows.sh
./build/build-windows.sh

# Or manually:
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
  go build -ldflags="-s -w -H windowsgui" -o build/mergit.exe
```

#### Build for both platforms:
```bash
chmod +x build/build-all.sh
./build/build-all.sh
```

### Binary Sizes

- **Linux:** ~18-22 MB (single binary, no dependencies)
- **Windows:** ~20-25 MB (single .exe, no dependencies)

All dependencies are statically linked into the binary.

## Project Structure

```
mergit/
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── pdf/
│   └── merger.go        # PDF merging business logic
├── ui/
│   └── app.go           # Fyne UI implementation
├── build/
│   ├── build-linux.sh   # Linux build script
│   ├── build-windows.sh # Windows build script
│   ├── build-all.sh     # Build both platforms
│   ├── mergit           # Linux binary (after build)
│   └── mergit.exe       # Windows binary (after build)
└── README.md            # This file
```

## Technology Stack

- **Language:** Go 1.21+
- **GUI Framework:** [Fyne v2.4+](https://fyne.io/) - Cross-platform native GUI
- **PDF Library:** [pdfcpu v0.6+](https://github.com/pdfcpu/pdfcpu) - PDF manipulation (Apache 2.0 license)

## Limitations

- Maximum 100 PDF files per merge operation
- Requires at least 2 PDF files to perform a merge
- Does not support password-protected PDFs
- No page-level manipulation (merges entire documents)

## Troubleshooting

### Build Issues

**"command not found: go"**
- Go is not installed. Follow the prerequisites section to install Go 1.21 or higher.

**"x86_64-w64-mingw32-gcc: command not found"**
- Install mingw-w64: `sudo apt-get install gcc-mingw-w64`

**"cannot find package"**
- Run `go mod download` to download dependencies

### Runtime Issues

**"Invalid or corrupted PDF"**
- The file may be damaged or password-protected
- Try opening the PDF in a PDF reader to verify it's valid

**Application won't start on Linux**
- Make sure the binary has execute permissions: `chmod +x build/mergit`
- Verify required libraries are installed (see prerequisites)

**Application won't start on Windows**
- Windows Defender may block the executable initially
- Right-click → Properties → Unblock if prompted

## License

This project uses the following open-source libraries:
- **Fyne** - BSD 3-Clause License
- **pdfcpu** - Apache License 2.0

Feel free to use, modify, and distribute this application.

## Contributing

This is a simple standalone tool. If you find bugs or want to add features:
1. Fork the repository
2. Make your changes
3. Test on both Linux and Windows
4. Submit a pull request

## Credits

- Built with [Go](https://golang.org/)
- UI powered by [Fyne](https://fyne.io/)
- PDF merging by [pdfcpu](https://github.com/pdfcpu/pdfcpu)
