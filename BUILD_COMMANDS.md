# Build Commands for mergit

## Prerequisites Installation

### Install Go (Required)

```bash
# Download and install Go 1.21.6 or higher
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz

# Add to PATH
export PATH=$PATH:/usr/local/go/bin

# Add to ~/.bashrc or ~/.zshrc for persistence
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

# Verify installation
go version
```

### Install System Dependencies

```bash
# For Linux builds (required)
sudo apt-get update
sudo apt-get install -y gcc libgl1-mesa-dev xorg-dev

# For Windows cross-compilation (optional)
sudo apt-get install -y gcc-mingw-w64
```

---

## Build Instructions

### 1. Download Go Dependencies

**First time only:**
```bash
cd /home/dfialho/Projects/mergit
go mod download
```

This downloads:
- `fyne.io/fyne/v2` (GUI framework)
- `github.com/pdfcpu/pdfcpu` (PDF library)
- All transitive dependencies

---

### 2. Build the Application

#### Option A: Using Build Scripts (Recommended)

**Linux binary:**
```bash
chmod +x build/build-linux.sh
./build/build-linux.sh
```

**Windows binary:**
```bash
chmod +x build/build-windows.sh
./build/build-windows.sh
```

**Both platforms:**
```bash
chmod +x build/build-all.sh
./build/build-all.sh
```

#### Option B: Manual Build Commands

**Linux binary:**
```bash
go build -ldflags="-s -w" -o build/mergit
```

**Windows binary (cross-compile from Linux):**
```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
  go build -ldflags="-s -w -H windowsgui" -o build/mergit.exe
```

---

## Build Flags Explained

### Linux Build
```bash
go build -ldflags="-s -w" -o build/mergit
```

- `go build` - Compile the Go program
- `-ldflags="-s -w"` - Strip debug symbols (reduces binary size)
  - `-s` - Omit symbol table
  - `-w` - Omit DWARF debug info
- `-o build/mergit` - Output file location

### Windows Build
```bash
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
  go build -ldflags="-s -w -H windowsgui" -o build/mergit.exe
```

- `GOOS=windows` - Target operating system
- `GOARCH=amd64` - Target architecture (64-bit)
- `CGO_ENABLED=1` - Enable C compiler (required for Fyne)
- `CC=x86_64-w64-mingw32-gcc` - Use mingw cross-compiler
- `-ldflags="-s -w -H windowsgui"` - Strip debug + hide console window
  - `-H windowsgui` - Windows GUI application (no console)
- `-o build/mergit.exe` - Output file location

---

## Running the Application

### On Linux
```bash
./build/mergit
```

### On Windows
Transfer `build/mergit.exe` to Windows and:
- Double-click the file, OR
- Run from Command Prompt: `mergit.exe`

---

## Expected Output

### After successful Linux build:
```
Building mergit for Linux...
✓ Linux binary created: build/mergit
-rwxr-xr-x 1 user user 19M Feb 14 10:30 build/mergit

To run: ./build/mergit
```

### After successful Windows build:
```
Building mergit for Windows...
✓ Windows binary created: build/mergit.exe
-rwxr-xr-x 1 user user 21M Feb 14 10:31 build/mergit.exe

To run on Windows: Double-click mergit.exe
```

---

## Common Build Issues

### Issue: "go: command not found"
**Solution:**
```bash
# Install Go (see Prerequisites section above)
# Then add to PATH:
export PATH=$PATH:/usr/local/go/bin
```

### Issue: "cannot find package fyne.io/fyne/v2"
**Solution:**
```bash
# Download dependencies first
go mod download
```

### Issue: "x86_64-w64-mingw32-gcc: command not found"
**Solution:**
```bash
# Install mingw for Windows cross-compilation
sudo apt-get install gcc-mingw-w64
```

### Issue: "undefined reference to `XOpenDisplay`"
**Solution:**
```bash
# Install X11 development libraries
sudo apt-get install libgl1-mesa-dev xorg-dev
```

### Issue: Build succeeds but binary won't run
**Solution:**
```bash
# Add execute permission
chmod +x build/mergit
```

---

## Rebuild After Code Changes

If you modify the source code:

```bash
# Rebuild for Linux
./build/build-linux.sh

# Or rebuild for both platforms
./build/build-all.sh
```

No need to run `go mod download` again unless you change `go.mod`.

---

## Clean Build

To start fresh:

```bash
# Remove old binaries
rm -f build/mergit build/mergit.exe

# Clean Go build cache (optional)
go clean -cache

# Rebuild
./build/build-all.sh
```

---

## Testing the Build

### Quick Test
```bash
# Build for Linux
./build/build-linux.sh

# Run the application
./build/mergit

# You should see the GUI window open
```

### Full Test
1. Build the application
2. Run it: `./build/mergit`
3. Add 2-3 test PDF files
4. Reorder them
5. Click "Merge PDFs"
6. Choose a save location
7. Verify the merged PDF opens correctly

---

## Distribution

### Linux
Simply share the `build/mergit` file. Users can:
```bash
chmod +x mergit
./mergit
```

### Windows
Simply share the `build/mergit.exe` file. Users can:
- Double-click to run
- No installation needed
- First run may show Windows security prompt (normal for unsigned executables)

---

## Binary Size Optimization

Already optimized with `-ldflags="-s -w"`.

Current sizes:
- **Linux:** ~18-22 MB
- **Windows:** ~20-25 MB

These include all dependencies statically linked. Cannot be reduced further without removing functionality.

---

## Development Workflow

### During Development
```bash
# Quick build and run (Linux)
go run main.go

# Or build and run
go build -o build/mergit && ./build/mergit
```

### For Release
```bash
# Full optimized build for both platforms
./build/build-all.sh
```

---

## Summary - Copy/Paste Quick Start

```bash
# Install Go
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Install dependencies
sudo apt-get install -y gcc libgl1-mesa-dev xorg-dev gcc-mingw-w64

# Navigate to project
cd /home/dfialho/Projects/mergit

# Download Go packages
go mod download

# Build for both platforms
chmod +x build/build-all.sh
./build/build-all.sh

# Run on Linux
./build/mergit
```

That's it! 🚀
