# SMD-docx

SMD-docx is a simple Go application that watches a specified folder for the creation of .docx, .doc, and .pptx files. Upon detecting a new file, it automatically converts the file to PDF using LibreOffice’s CLI, and then removes the original file from the system.

### Features
	•	File Monitoring: Continuously watches a specified folder (Downloads folder by default) for new .docx, .doc, or .pptx files.
	•	Automated Conversion: Converts detected files to PDF format using LibreOffice in headless mode.
	•	File Cleanup: Deletes the original document after conversion to save space.

### Requirements
	•	Go: Make sure you have Go installed. You can install Go from https://golang.org/dl/.
	•	LibreOffice: The program uses LibreOffice’s command-line tools to convert documents to PDF. Install LibreOffice from here.
	•	fsnotify: The program uses the fsnotify package to watch for file events. It is included as a dependency in the Go code.

### Installation
	1.	Clone this repository to your local machine:

```git clone https://github.com/<your-username>/SMD-docx.git
cd SMD-docx```


	2.	Install the required dependencies:

```go mod tidy```


	3.	Ensure that LibreOffice is installed and accessible via your system’s PATH.

### Usage

	1.	Run the program: This application will monitor the Downloads folder for new .docx, .doc, or .pptx files and convert them to PDFs.

```go run main.go```

	2.	File Detection: Once the program is running, it will continuously check the folder for new files. Upon detecting a new file with the .docx, .doc, or .pptx extension, it will convert it to PDF.
	3.	Conversion and Cleanup: The original document is removed after the conversion is successful.
	4.	Custom Folder: If you’d like to monitor a different folder, modify the downloadsFolder variable in the code to the desired path.

### Example

```File Detected: Begin Convert.
Converted /home/user/Downloads/example.docx to /home/user/Downloads/example.pdf See ya.```

