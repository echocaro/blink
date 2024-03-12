package form

import (
	"github.com/charmbracelet/huh"
)

var Confirm bool // true if user selects affirmative case

type FormModel struct {
	Form *huh.Form
}

func CreateForm() FormModel {
	charm := huh.ThemeCharm()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Ready to get started?").
				Affirmative("Yes! 🙌").
				Negative("No. 🙅").
				Value(&Confirm),
		),
	).WithTheme(charm)

	return FormModel{form}
}
