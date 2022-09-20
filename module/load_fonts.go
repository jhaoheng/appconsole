package module

import (
	"os"
	"strings"

	"github.com/flopp/go-findfont"
	"github.com/sirupsen/logrus"
)

var ChineseFonts = map[string]bool{
	"STHeiti Light.ttc": true,
	"simkai.ttf":        true,
	"simhei.ttf":        true,
}

var JapaneseFonts = map[string]bool{}

// for chinese word
func LoadFont() {
	defaultFontPath := "resources/fonts/STHeiti Light.ttc"

	// load from system
	for _, path := range findfont.List() {
		components := strings.Split(path, "/")
		key := components[len(components)-1]

		if _, ok := ChineseFonts[key]; ok {
			defaultFontPath = path
			break
		} else if _, ok := JapaneseFonts[key]; ok {
			defaultFontPath = path
			break
		}
	}

	// _, err := os.ReadFile(defaultFontPath)
	// if err != nil {
	// 	panic(err)
	// }

	logrus.Infof("Load font setting from: %v\n", defaultFontPath)

	os.Setenv("FYNE_FONT", defaultFontPath)
}
