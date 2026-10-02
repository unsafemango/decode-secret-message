package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const url = "https://example.com"

type point struct {
	x, y int
	ch   string
}

var tagRe = regexp.MustCompile(`<[^>]*>`)

// Clean each cell
func cellText(inner string) string {
	s := tagRe.ReplaceAllString(inner, "")
	s = html.UnescapeString(s)

	return strings.TrimSpace(s)
}
func printSecretMessage(url string) error {
	// Fetch the doccument from the url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Read the downloaded text
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Pull out the table cells
	var cellRe = regexp.MustCompile(`(?s)<td[^>]*>(.*?)</td>`)

	var cells []string
	for _, m := range cellRe.FindAllStringSubmatch(string(data), -1) {
		cells = append(cells, cellText(m[1]))
	}

	var points []point
	var maxX, maxY = 0, 0
	// Read the cells three at a time
	for i := 0; i < len(cells); i += 3 {
		x, errX := strconv.Atoi(cells[i])
		y, errY := strconv.Atoi(cells[i+2])

		if errX != nil || errY != nil {
			continue
		}

		points = append(points, point{x, y, cells[i+1]})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	// Initialize an empty grid of spaces
	grid := make([][]rune, maxY+1)
	for y := range grid {
		grid[y] = make([]rune, maxX+1)
		for x := range grid[y] {
			grid[y][x] = ' '
		}
	}

	// Drop each character into the grid
	for _, p := range points {
		r := []rune(p.ch)
		if len(r) > 0 {
			grid[p.y][p.x] = r[0]
		}
	}

	// Print the grid
	for y := maxY; y >= 0; y-- {
		fmt.Println(string(grid[y]))
	}
	return nil
}

func main() {
	if err := printSecretMessage(url); err != nil {
		fmt.Println("Error:", err)
	}
}
