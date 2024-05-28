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
					huh.NewOption[string]("Yes!", "yes"),
					huh.NewOption[string]("No.", "no"),
				).
				Value(&Confirm),
		),
	).WithTheme(charm)

	return FormModel{Form: form}
}
