# Quick Start Guide for mergit

## Installation & Build Instructions

### Step 1: Install Go

**On Ubuntu/Debian:**
```bash
# Download and install Go 1.21+ from official source
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz

# Add to PATH (add to ~/.bashrc or ~/.zshrc for persistence)
export PATH=$PATH:/usr/local/go/bin

# Verify installation
go version
```

Or use your package manager (may have older version):
```bash
sudo apt-get update
sudo apt-get install golang-go
```

### Step 2: Install Build Dependencies

**For Linux builds:**
```bash
sudo apt-get update
sudo apt-get install -y gcc libgl1-mesa-dev xorg-dev
```

**For Windows cross-compilation (optional):**
```bash
sudo apt-get install -y gcc-mingw-w64
```

### Step 3: Download Dependencies

```bash
cd /home/dfialho/Projects/mergit
go mod download
```

This will download:
- Fyne v2.4+ (GUI framework)
- pdfcpu v0.6+ (PDF manipulation library)

### Step 4: Build the Application

**Option A: Build for Linux only**
```bash
chmod +x build/build-linux.sh
./build/build-linux.sh
```

**Option B: Build for Windows only**
```bash
chmod +x build/build-windows.sh
./build/build-windows.sh
```

**Option C: Build for both platforms**
```bash
chmod +x build/build-all.sh
./build/build-all.sh
```

### Step 5: Run the Application

**On Linux:**
```bash
./build/mergit
```

**On Windows:**
Transfer `build/mergit.exe` to a Windows machine and double-click it.

---

## Quick Commands Reference

```bash
# Navigate to project
cd /home/dfialho/Projects/mergit

# Download dependencies
go mod download

# Build for Linux
./build/build-linux.sh

# Build for Windows
./build/build-windows.sh

# Run on Linux
./build/mergit

# Manual build (Linux)
go build -ldflags="-s -w" -o build/mergit

# Manual build (Windows cross-compile)
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
  go build -ldflags="-s -w -H windowsgui" -o build/mergit.exe
```

---

## Troubleshooting

### "go: command not found"
- Go is not installed or not in PATH
- Install Go following Step 1 above
- Add `/usr/local/go/bin` to your PATH

### "cannot find package"
- Dependencies not downloaded
- Run: `go mod download`

### "x86_64-w64-mingw32-gcc: not found"
- mingw-w64 not installed (only needed for Windows builds)
- Run: `sudo apt-get install gcc-mingw-w64`

### Build fails with "undefined reference"
- Missing system libraries
- Run: `sudo apt-get install gcc libgl1-mesa-dev xorg-dev`

### Binary won't run
- Missing execute permission
- Run: `chmod +x build/mergit`

---

## File Size

After building, expect:
- **build/mergit** (Linux): ~18-22 MB
- **build/mergit.exe** (Windows): ~20-25 MB

These are standalone binaries with no external dependencies!
