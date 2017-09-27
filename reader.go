package main

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// Read Does stuff
func Read(url string, board [][]rune) {
	//Reads
	board[0][0] = 'X'
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "piece"
			if isAnchor {
				var name rune
				for _, a := range t.Attr {
					if a.Key == "class" {
						atts := strings.Split(a.Val, " ")
						name = piece(atts[1], atts[0])
					}
					if a.Key == "style" {
						re := regexp.MustCompile("[0-9 + .]+")
						locs := re.FindAllString(string(a.Val), -1)
						x, _ := strconv.ParseFloat(locs[0], 32)
						xint := int(x / 12.5)
						y, _ := strconv.ParseFloat(locs[1], 32)
						yint := int(y / 12.5)
						board[xint][yint] = name
					}
				}
			}
		}
	}
}

func piece(name string, color string) rune {
	var code rune
	switch name {
	case "king":
		{
			code = 'K'
		}
	case "queen":
		{
			code = 'Q'
		}
	case "pawn":
		{
			code = 'P'
		}
	case "knight":
		{
			code = 'N'
		}
	case "bishop":
		{
			code = 'B'
		}
	case "rook":
		{
			code = 'R'
		}
	default:
		code = '-'
	}
	if color == "black" {
		code += 32
	}
	return code
}
