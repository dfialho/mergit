# Mergit 🔀

A client-side PDF merger that runs entirely in your browser. No uploads, no servers, 100% private.

## ✨ Features

- 🚀 **No Installation** - Just open in your browser
- 🔒 **100% Private** - All processing happens locally, files never leave your device
- 📱 **Mobile Friendly** - Works on phones and tablets
- 🎨 **Modern Design** - Beautiful, gradient-based UI
- ⚡ **Fast** - Client-side processing using pdf-lib
- 🌐 **Cross-Platform** - Works on any modern browser (Chrome, Firefox, Safari, Edge)
- 💾 **Offline Capable** - Can work offline after first load

### Core Functionality

✅ Add PDFs via drag-and-drop or file picker
✅ Reorder PDFs with up/down arrow buttons
✅ Remove individual PDFs
✅ Clear all PDFs
✅ Merge 2-100 PDFs into a single file 
✅ Custom filename on download
✅ Default filename with today's date 
✅ Auto-clear after merge
✅ Error handling with user-friendly messages 

## 🚀 Quick Start

### Option 1: Open Locally

1. Download `index.html`
2. Double-click to open in your browser
3. Start merging PDFs!

### Option 2: Run with Local Server

```bash
# Using Python 3
python3 -m http.server 8000

# Using Python 2
python -m SimpleHTTPServer 8000

# Using Node.js
npx serve

# Using PHP
php -S localhost:8000
```

Then open: `http://localhost:8000`

## 🛠️ How It Works

### Architecture

```
User drops PDF files
      ↓
Validate file type (.pdf)
      ↓
Add to in-memory array
      ↓
User reorders/removes files
      ↓
Click "Merge PDFs"
      ↓
Load each PDF with pdf-lib
      ↓
Copy all pages to new PDF
      ↓
Generate merged PDF bytes
      ↓
Prompt for filename
      ↓
Download to user's device
      ↓
Clear list (ready for next merge)
```

### Privacy & Security

🔒 **All processing is client-side:**
- PDFs are loaded into browser memory only
- No files are uploaded to any server
- No tracking or analytics
- No external dependencies except pdf-lib

🌐 **Can work offline:**
- After first load, pdf-lib is cached by the browser
- The HTML file can be saved and used offline

## 🎨 Features Overview

### File Management
- **Add PDFs**: Drag-and-drop or click "Add PDF" button
- **Reorder**: Use ↑ and ↓ arrow buttons
- **Remove**: Click ✕ button next to any file
- **Clear All**: Remove all files at once

### Merging
- **Minimum**: 2 PDFs required
- **Maximum**: 100 PDFs allowed
- **Process**: Client-side, no upload needed
- **Speed**: Depends on PDF sizes and device performance

### Download
- **Prompt**: User is asked for filename
- **Default**: `merged_YYYY-MM-DD.pdf` (today's date)
- **Format**: Standard PDF format
- **Location**: Browser's default download folder


## 🐛 Troubleshooting

### "Not a PDF file" Error
- Make sure the file has `.pdf` extension
- Verify the file isn't corrupted
- Try opening it in a PDF viewer first

### Merge Fails
- **Large files**: Browser may run out of memory with very large PDFs (100+ MB)
- **Corrupted PDFs**: Try re-downloading the PDF
- **Password-protected PDFs**: Not currently supported

### Slow Performance
- **Large PDFs**: Processing 100+ page PDFs may take time
- **Many files**: Merging 50+ files will be slower
- **Old devices**: Use a faster device or fewer files

### Download Doesn't Start
- Check browser's download permissions
- Try a different browser
- Disable browser extensions that block downloads


## 📊 File Size Limits

- **Individual PDF**: No hard limit, but 100MB+ may be slow
- **Total merge**: Limited by browser memory (~2GB typically)
- **Number of files**: Maximum 100 PDFs per merge


## 🔧 Customization

### Change Theme Colors

Edit the CSS `:root` variables:

```css
:root {
    --primary: #6366f1;        /* Main color */
    --primary-dark: #4f46e5;   /* Darker shade */
    --bg: #0f172a;             /* Background */
    --text: #f1f5f9;           /* Text color */
}
```

### Change Max PDF Limit

Edit the JavaScript constant:

```javascript
const state = {
    maxPDFs: 100  // Change to desired limit
};
```

### Bundle pdf-lib Inline

For offline use, download pdf-lib and embed it:

```html
<script>
    // Paste entire pdf-lib.min.js content here
</script>
```

This makes the HTML file larger (~520KB) but fully offline-capable.

---

## 📝 License

MIT License - Free for personal and commercial use

---

## 🙏 Credits

- **pdf-lib** - [github.com/Hopding/pdf-lib](https://github.com/Hopding/pdf-lib)
- **Design** - Custom modern gradient theme

---

## 🚀 Future Enhancements (Ideas)

Possible features for future versions:
- [ ] PDF preview thumbnails
- [ ] Page range selection
- [ ] Dark/light mode toggle
- [ ] PWA support (install as app)
- [ ] Batch processing (save merge configs)
- [ ] Password-protected PDF support
- [ ] PDF compression options

