# OCR-in-GO-Fiber

This project is a simple web application built with Go and the Fiber framework that performs Optical Character Recognition (OCR) on uploaded images using the `gosseract` library.

## Features

- Upload an image file
- Perform OCR on the uploaded image
- Return the extracted text as a JSON response

## Requirements

- Go 1.20 or later
- Fiber v2.38.0
- gosseract v2.3.0

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/kishyassin/OCR-in-GO-Fiber.git
    cd OCR-in-GO-Fiber
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:

    ```sh
    go run main.go
    ```

2. The application will start on port 3000. You can upload an image file to the `/upload` endpoint using a tool like `curl` or Postman.

    Example using `curl`:

    ```sh
    curl -X POST http://localhost:3000/upload -F "file=@path/to/your/image.png"
    ```

3. The server will respond with the extracted text in JSON format.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.