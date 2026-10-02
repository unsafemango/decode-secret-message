# decode-secret-message

Go program that fetches a secret text, reads a table of `(x, character, y)` cells, and prints the characters on a 2D grid so the secret message is visible in the terminal.

## Requirements

- Go 1.25.1 or later

## Run

```bash
go run .
```

The program downloads the document at the URL in `main.go`, places each character at its coordinates, and prints the grid from top to bottom.
