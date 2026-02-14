# How to Use mergit

## Application Interface

```
┌──────────────────────────────────────────────────┐
│  mergit - PDF Merger                      [_][□][X]│
├──────────────────────────────────────────────────┤
│                                                  │
│   Drop PDF files here or click 'Add PDF'        │
│                                                  │
│  ┌────────────────────────────────────────────┐ │
│  │                                            │ │
│  │  📄 invoice_2024.pdf              [X]     │ │
│  │  📄 report_Q4.pdf                 [X]     │ │
│  │  📄 meeting_notes.pdf             [X]     │ │
│  │  📄 summary.pdf                   [X]     │ │
│  │                                            │ │
│  │  4 PDFs added (max 100)                   │ │
│  │                                            │ │
│  └────────────────────────────────────────────┘ │
│                                                  │
│  [Add PDF]  [Clear All]  [Merge PDFs]           │
│                                                  │
└──────────────────────────────────────────────────┘
```

## Step-by-Step Guide

### 1. Launch the Application

**Linux:**
```bash
./build/mergit
```

**Windows:**
Double-click `mergit.exe`

---

### 2. Add PDF Files

You have two options:

#### Option A: Drag and Drop
1. Open your file manager (Nautilus, Dolphin, Windows Explorer, etc.)
2. Navigate to your PDF files
3. Drag one or more PDF files into the mergit window
4. Files will appear in the list

#### Option B: File Picker
1. Click the **"Add PDF"** button
2. Browse to your PDF file
3. Select the file and click Open
4. File will appear in the list

**Notes:**
- Only `.pdf` files are accepted
- Maximum 100 files can be added
- Invalid or corrupted PDFs will show an error

---

### 3. Reorder Files (Optional)

The order of files in the list determines their order in the final merged PDF.

**To reorder:**
1. Click on a PDF in the list (it will be highlighted)
2. While holding the mouse button, drag it up or down
3. Release at the desired position
4. The file will move to the new position

**Example:**
```
Before reordering:          After dragging "report" to top:
1. invoice.pdf              1. report.pdf
2. report.pdf          →    2. invoice.pdf
3. summary.pdf              3. summary.pdf
```

---

### 4. Remove Files (Optional)

**To remove a single file:**
- Click the **[X]** button next to the file you want to remove

**To remove all files:**
- Click the **"Clear All"** button at the bottom

---

### 5. Merge PDFs

Once you have **at least 2 PDFs** in the list:

1. Click the **"Merge PDFs"** button (it's highlighted and enabled)
2. A file save dialog will appear
3. Choose where to save the merged PDF:
   - Default filename: `merged_2026-02-14.pdf` (today's date)
   - You can change the name or location
4. Click **Save**

---

### 6. Wait for Merge

- A progress dialog will appear showing a spinning animation
- Message: "Merging PDFs... Please wait..."
- The merge happens in the background

**Note:** For large files or many PDFs, this may take a few seconds.

---

### 7. Success!

After successful merge:

1. A **success dialog** appears showing:
   ```
   PDF merged successfully!
   
   Saved to:
   /home/user/Documents/merged_2026-02-14.pdf
   ```

2. Click **OK** to close the dialog

3. The file list **automatically clears**

4. The app is ready for your next merge!

---

## Common Workflows

### Simple Merge (2 PDFs)
1. Launch mergit
2. Drag `contract.pdf` and `addendum.pdf` into the window
3. Click "Merge PDFs"
4. Save as `contract_complete.pdf`
5. Done!

### Complex Merge (Multiple PDFs in specific order)
1. Launch mergit
2. Add all PDFs via drag-and-drop or "Add PDF" button
3. Reorder PDFs by dragging them:
   - Title page first
   - Table of contents second
   - Chapters in order
   - Appendix last
4. Click "Merge PDFs"
5. Save with descriptive name
6. Done!

### Batch Processing
1. Merge first set of PDFs
2. After success dialog, list clears automatically
3. Immediately add next set of PDFs
4. Merge again
5. Repeat as needed!

---

## Tips & Tricks

✅ **Check file order carefully** - The visual order in the list is exactly how pages will appear in the merged PDF

✅ **Use descriptive filenames** - When saving, use names like `report_2024_final.pdf` instead of `merged.pdf`

✅ **Test with small files first** - If you're merging many large PDFs, test with a few small ones first

✅ **Keep original files** - mergit creates a new merged PDF without modifying your original files

✅ **Remove duplicates** - If you accidentally add a file twice, use the [X] button to remove one

✅ **Date in filename** - The default filename includes today's date for easy organization

---

## Error Messages

| Error | Meaning | Solution |
|-------|---------|----------|
| "Maximum 100 PDF files allowed" | You tried to add more than 100 files | Remove some files or merge in batches |
| "File is not a PDF" | The file doesn't have a .pdf extension | Only PDF files are supported |
| "Invalid or corrupted PDF" | The PDF file is damaged | Try opening the file in a PDF reader to verify it works |
| "Please add at least 2 PDF files" | You tried to merge with 0 or 1 file | Add at least 2 PDFs before merging |
| "File does not exist" | The file was moved or deleted | Re-add the file from its current location |

---

## Keyboard Shortcuts

Currently, mergit is primarily mouse-driven. Keyboard shortcuts may be added in future versions.

**Current functionality:**
- Click to select files
- Drag to reorder
- Click buttons to perform actions

---

## Platform-Specific Notes

### Linux
- Application respects your system theme (light/dark)
- File dialogs use native GTK dialogs
- Drag-and-drop works from any file manager

### Windows
- Application uses Windows native look
- File dialogs are Windows standard
- No console window appears (GUI only)
- May need to "Unblock" the .exe on first run (Windows security feature)

---

## Questions?

- **Can I merge password-protected PDFs?** No, currently not supported
- **Does this modify my original files?** No, it creates a new merged PDF
- **Can I select specific pages?** Not in the current version (merges entire documents)
- **Is there a file size limit?** No specific limit, but very large files may take longer
- **Can I undo a merge?** No, but your original files are untouched
- **Where is the merged file saved?** Wherever you choose in the save dialog

---

Enjoy using **mergit**! 🎉
