# Simple File Manager with GO

A simple file management with Go Lang, designed for managing files on a server. The API allows users to upload, download, rename, and delete files, and can be easily customized using environment variables.

## Features

- **File Listing**: List all available files on the server.
- **Download Files**: Download single files or all files in a ZIP archive.
- **Rename Files**: Rename an existing file on the server.
- **Delete Files**: Delete a file from the server.
- **Configurable**: Easily configurable using environment variables for settings like the directory path and server port.

## Requirements

- Go 1.18+ (for building and running the API)
- Access to a local or remote server to store the files

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ialaminpro/go-simple-file-manager.git
   cd go-simple-file-manager
   ```

2. Set up environment variables:

   Create a `.env` file in the root of the project and add the following variables:

   ```
   PORT=:8080
   FILE_DIRECTORY=/path/to/your/files
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

   The server will now be running at `http://localhost:8080`.

## API Endpoints

- **GET /files**: List all files in the directory.
- **POST /download-all**: Download all files in a ZIP archive.
- **DELETE /delete**: Delete a specified file.
- **PUT /rename**: Rename an existing file.

### Example API Requests

#### List Files
```bash
GET http://localhost:8080/files
```

#### Delete File
```bash
DELETE http://localhost:8080/delete
{
  "fileName": "example.txt"
}
```

#### Rename File
```bash
PUT http://localhost:8080/rename
{
  "oldName": "oldname.txt",
  "newName": "newname.txt"
}
```

## Environment Variables

- **PORT**: The port number where the server will run. Default is `8080`.
- **FILE_DIRECTORY**: The directory path where the files are stored. You can change this to the directory where you want to manage your files.
