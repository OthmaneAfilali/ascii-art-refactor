package generator

import (
	"ascii-art/pkg/fileMgr"
	"fmt"
	"strings"
)

var style = make(map[rune][]string)

// GenArt generate ascii art string by combining [][8]artStr into one string.
func GenArt(txt, styleNm string) string {
	getStyle(styleNm)
	for _, rn := range txt {
		if rn < 32 || rn > 127 {
			fmt.Println("Character not an ASCII char:", string(rn))
			return ""
		}
	}
	txtLns := strings.Split(txt, "\\n")
	return genArtStrs(txtLns)
}

// getStyle read <styleName>.txt and store the ascii art runes in a map[rune][]string.
func getStyle(styleNm string) {
	rawStyle := strings.Split(fileMgr.ReadFile("./assets/"+styleNm+".txt"), "\n")
	for i := 1; i < len(rawStyle); i = i + 9 {
		curChar := rune(32 + i/9)
		style[curChar] = rawStyle[i : i+8]
	}
}

// genArtStr generate [8]string of ascii art for each txtLn.
func genArtStrs(txtLns []string) string {
	var builder strings.Builder
	for i, txtLn := range txtLns {
		artStrs := [8]string{}
		if txtLn == "" {
			if i > 0 {
				builder.WriteString("\n")
			}
			continue
		}
		for _, rn := range txtLn {
			for j := 0; j < 8; j++ {
				artStrs[j] += style[rn][j]
			}
		}
		for j := 0; j < 8; j++ {
			builder.WriteString(artStrs[j])
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
