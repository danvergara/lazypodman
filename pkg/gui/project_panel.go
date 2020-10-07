package gui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
)

func (gui *Gui) refreshProjectView() error {
	v, err := gui.g.View("project")
	if err != nil {
		return err
	}

	gui.g.Update(func(g *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, gui.ProjectName)
		return nil
	})

	return nil
}
