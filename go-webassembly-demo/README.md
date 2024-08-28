# Go WebAssembly Demo
This is a demo project that showcases how to run Go code in the browser using WebAssembly.

## Prerequisites

Before getting started with this project, make sure you have the following software installed:

- Go: You can download and install Go from the official website: https://golang.org/dl/. Follow the installation instructions for your operating system.

- Visual Studio Code: You can download and install Visual Studio Code from the official website: https://code.visualstudio.com/. Follow the installation instructions for your operating system.

- Python: You can download and install Python from the official website: https://www.python.org/downloads/. Follow the installation instructions for your operating system.

Once you have installed these prerequisites, you are ready to proceed with the steps mentioned in the README.md file.

## Project Structure

The project has the following files:

- `src/main.go`: This file contains the Go code for the web assembly demo. It includes the necessary imports and defines the main function to run the web assembly code.

- `src/wasm_exec.js`: This file is a JavaScript wrapper that allows the Go web assembly code to be executed in the browser. It provides the necessary functions and utilities for running the web assembly code.

- `web/index.html`: This file is the HTML file that serves as the entry point for the web application. It includes the necessary script tags to load the web assembly code and the JavaScript file.

- `web/main.js`: This file contains the JavaScript code that interacts with the web assembly code. It initializes the web assembly module and defines functions to handle user interactions and display the results.

- `go.mod`: This file is the Go module file that specifies the dependencies for the project.

## Getting Started

To run the demo, follow these steps:

1. Clone the repository.

2. Build the Go code into WebAssembly using the following command:

   ```powershell
   $env:GOOS = "js"
   $env:GOARCH = "wasm"
   go build -o web/main.wasm src/main.go
   ```

3. Change to the `web` directory:

   ```powershell
   cd web
   ```

4. Start a local web server in the project directory using Python:

   ```powershell
   python -m http.server 8090
   ```

5. Open the browser and navigate to `http://localhost:8090`.

6. Interact with the web application and see the results.

Please note that you can use any text editor or IDE of your choice, such as Visual Studio Code, to work with this project.


