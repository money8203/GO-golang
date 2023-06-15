# GO Static Website
This is a simple Go web application that serves static HTML files and handles form submissions. It demonstrates how to create a basic web server in Go using the `net/http` package.

## Istallation
1. Clone the repository: `git clone https://github.com/money8203/your_username.git`
2. Navigate to the project directory: `cd go-server`
3. Build the project: `go build`
4. Run the application: `./go-server`
5. Open your web browser and visit http://localhost:8080 to access the static website.

## Features
- The main page `index.html` displays a simple heading.
- The form page `form.html` contains a form that allows users to enter their name and address.
- When the form is submitted, the server handles the `POST` request and prints the submitted values.
- The server also provides a `/hello` route that returns a "hello" message.

## Project Structure
- `main.go`: Contains the main Go application code.
- `static/`: Directory containing static HTML files.
