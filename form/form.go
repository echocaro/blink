package form

import (
	"github.com/charmbracelet/huh"
)

// var Confirm string // true if user selects affirmative case

type FormModel struct {
	Form    *huh.Form
	Confirm string
}

func CreateForm() FormModel {
	charm := huh.ThemeCharm()
	formModel := FormModel{}

	formModel.Form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Are you ready to get started?").
				Options(
					huh.NewOption[string]("Yes!", "yes"),
					huh.NewOption[string]("Yes, but run in the background.", "background"),
					huh.NewOption[string]("No.", "no"),
				).
				Value(&formModel.Confirm),
		),
	).WithTheme(charm)

	return formModel
}
