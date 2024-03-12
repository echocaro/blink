package form

import (
	"github.com/charmbracelet/huh"
)

var Confirm string // true if user selects affirmative case

type FormModel struct {
	Form *huh.Form
}

func CreateForm() FormModel {
	charm := huh.ThemeCharm()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Are you ready to get started?").
				Options(
					huh.NewOption("Yes!", "yes"),
					huh.NewOption("Yes, but run in the background.", "background"),
					huh.NewOption("No.", "no"),
				),
			// huh.NewConfirm().
			// 	Title("Ready to get started?").
			// 	Affirmative("Yes! 🙌").
			// 	Negative("No. 🙅").
			// 	Value(&Confirm),
		),
	).WithTheme(charm)

	return FormModel{form}
}
