// forked from: https://github.com/knipferrc/teacup/blob/main/icons/icons.go

package icons

func GetIcon(ext string) (icon, color string) {
	var i *IconInfo

	i, _ = IconSet[ext]

	return i.GetGlyph(), i.GetColor(1)
}
