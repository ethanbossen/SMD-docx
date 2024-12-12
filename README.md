# SMD-docx 

SMD-docx is a simple Go application that watches a specified folder for the creation of .docx, .doc and .pptx files. I use it on my downloads folder for school. It will watch the directroy and upon detecting a file of the specific types it will convert it into a PDF using LibreOffice's CLI and remove the original file from the system.

## Features
- **File Monitoring**: continuously watch a speciifed fold for new .docx, .doc or .pptx files
- **Automated Conversion**: Converts detected files to PDF format using LibreOFfice in headless mode.
- **File Cleanup**: deletes the original document after conversion

## Requirements
- **Go**: Make sure you have Go installed you can install go from [golang.org/dl/](url).
- **LibreOffice**: This program uses LibreOffice's CLI tools to convert docs to PDF you can install LibreOffice from here: [www.libreoffice.org/download/download/](url)
- **fsnotify**: This program also uses the fsnotify package to watch for file events. It's incldued as a dependency and will be installed with ```go mod tidy```

## Installation
1. Clone this repository to your local machine:
```git clone https://github.com/<your-username>/SMD-docx.git```
```cd SMD-docx```
2. Install required dependencies:
```go mod tidy```
3. Ensure LibreOffice is installed and accessible via your system's PATH

## Usage
1. Change the specified directory to the one you want to monitor
2. Run the program with: ```go run main.go```
3. Download a .docx, .doc, or .pptx
4. Document is converted and saved to the directory
