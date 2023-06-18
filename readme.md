# QR Generator

A Golang app to generate QR code.

## Installation
 1. Clone the repository
 ```sh
    git clone https://github.com/peekeah/qr-code-generator 
 ```

 2. Go to the project folder
 ```sh
    cd qr-code-generator
 ```

 3. Build the app
 ```sh
    go build .
 ```

4. Run the app
```sh
    ./qr-code-generator.exe
```


## Usage

### Endpoints

#### POST /

Generate QR for the given text.

**Request**

- Method: POST
- Path: `/`
- Body: `JSON 
    {
        "text": "www.google.com"
    }
`

**Response**

- Status: 200 OK
- Body: png QR

