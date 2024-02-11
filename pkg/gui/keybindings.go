package gui

import (
	"github.com/danvergara/gocui"
)

// keyBinding struct used to defines the multiple actions defined to interact with dblab.
type keyBinding struct {
	view    string
	key     interface{}
	mod     gocui.Modifier
	handler func(*gocui.Gui, *gocui.View) error
}

// initialKeyBindings returns an slice with the standard key bindings.
func initialKeyBindings() []keyBinding {

	bindings := []keyBinding{
		{
			view:    "query",
			key:     gocui.KeyCtrlP,
			mod:     gocui.ModNone,
			handler: setViewOnTop("query", "rows"),
		},
		{
			view:    "rows",
			key:     gocui.KeyCtrlP,
			mod:     gocui.ModNone,
			handler: setViewOnTop("rows", "tables"),
		},
		{
			view:    "tables",
			key:     gocui.KeyCtrlP,
			mod:     gocui.ModNone,
			handler: setViewOnTop("tables", "query"),
		},
		{
			view:    "structure",
			key:     gocui.KeyCtrlH,
			mod:     gocui.ModNone,
			handler: nextView("structure", "tables"),
		},
		{
			view:    "structure",
			key:     gocui.KeyCtrlK,
			mod:     gocui.ModNone,
			handler: nextView("structure", "query"),
		},
		{
			view:    "constraints",
			key:     gocui.KeyCtrlH,
			mod:     gocui.ModNone,
			handler: nextView("constraints", "tables"),
		},
		{
			view:    "constraints",
			key:     gocui.KeyCtrlK,
			mod:     gocui.ModNone,
			handler: nextView("constraints", "query"),
		},
		{
			view:    "indexes",
			key:     gocui.KeyCtrlH,
			mod:     gocui.ModNone,
			handler: nextView("indexes", "tables"),
		},
		{
			view:    "indexes",
			key:     gocui.KeyCtrlK,
			mod:     gocui.ModNone,
			handler: nextView("indexes", "query"),
		},
		{
			view:    "constraints",
			key:     gocui.KeyCtrlF,
			mod:     gocui.ModNone,
			handler: setViewOnTop("constraints", "rows"),
		},
		{
			view:    "rows",
			key:     gocui.KeyCtrlF,
			mod:     gocui.ModNone,
			handler: setViewOnTop("rows", "constraints"),
		},
		{
			view:    "structure",
			key:     gocui.KeyCtrlS,
			mod:     gocui.ModNone,
			handler: setViewOnTop("structure", "rows"),
		},
		{
			view:    "rows",
			key:     gocui.KeyCtrlS,
			mod:     gocui.ModNone,
			handler: setViewOnTop("rows", "structure"),
		},
		{
			view:    "indexes",
			key:     gocui.KeyCtrlI,
			mod:     gocui.ModNone,
			handler: setViewOnTop("indexes", "rows"),
		},
		{
			view:    "rows",
			key:     gocui.KeyCtrlI,
			mod:     gocui.ModNone,
			handler: setViewOnTop("rows", "indexes"),
		},
		{
			view:    "navigation",
			key:     gocui.MouseLeft,
			mod:     gocui.ModNone,
			handler: navigation,
		},
		{
			view:    "",
			key:     gocui.KeyCtrlC,
			mod:     gocui.ModNone,
			handler: quit,
		},
	}

	// output views navigation.
	// for _, viewName := range []string{"rows", "structure", "constraints"} {
	// 	bindings = append(bindings, []keyBinding{
	// 		{view: viewName, key: 'k', mod: gocui.ModNone, handler: moveCursorVertically("up")},
	// 		{view: viewName, key: 'j', mod: gocui.ModNone, handler: moveCursorVertically("down")},
	// 		{view: viewName, key: 'l', mod: gocui.ModNone, handler: moveCursorHorizontally("right")},
	// 		{view: viewName, key: 'h', mod: gocui.ModNone, handler: moveCursorHorizontally("left")},
	// 	}...)
	// }

	// arrow keys navigation.
	for _, viewName := range []string{"rows", "structure", "constraints"} {
		bindings = append(bindings, []keyBinding{
			{view: viewName, key: gocui.KeyArrowUp, mod: gocui.ModNone, handler: moveCursorVertically("up")},
			{view: viewName, key: gocui.KeyArrowDown, mod: gocui.ModNone, handler: moveCursorVertically("down")},
			{view: viewName, key: gocui.KeyArrowRight, mod: gocui.ModNone, handler: moveCursorHorizontally("right")},
			{view: viewName, key: gocui.KeyArrowLeft, mod: gocui.ModNone, handler: moveCursorHorizontally("left")},
		}...)
	}

	// defines the navigation on the "tables" view.
	bindings = append(bindings, []keyBinding{
		{view: "tables", key: 'k', mod: gocui.ModNone, handler: moveCursorVertically("up")},
		{view: "tables", key: gocui.KeyArrowUp, mod: gocui.ModNone, handler: moveCursorVertically("up")},
		{view: "tables", key: 'j', mod: gocui.ModNone, handler: moveCursorVertically("down")},
		{view: "tables", key: gocui.KeyArrowDown, mod: gocui.ModNone, handler: moveCursorVertically("down")},
	}...)

	return bindings
}
func (gui *Gui) keybindings() error {
	for _, k := range initialKeyBindings() {
		if err := gui.g.SetKeybinding(k.view, k.key, k.mod, k.handler); err != nil {
			return err
		}
	}

	// SQL helpers
	if err := gui.g.SetKeybinding("query", gocui.KeyCtrlSlash, gocui.ModNone, gui.inputQuery()); err != nil {
		return err
	}

	if err := gui.g.SetKeybinding("tables", gocui.KeyEnter, gocui.ModNone, gui.metadata); err != nil {
		return err
	}
	if err := gui.g.SetKeybinding("tables", gocui.MouseLeft, gocui.ModNone, gui.metadata); err != nil {
		return err
	}
	// SQL pagination.
	if err := gui.g.SetKeybinding("next", gocui.MouseLeft, gocui.ModNone, gui.nextPage); err != nil {
		return err
	}

	if err := gui.g.SetKeybinding("back", gocui.MouseLeft, gocui.ModNone, gui.previousPage); err != nil {
		return err
	}
	if err := gui.g.SetKeybinding("tables", gocui.MouseWheelUp, gocui.ModNone, scrollUp); err != nil {
		return err
	}
	if err := gui.g.SetKeybinding("tables", gocui.MouseWheelDown, gocui.ModNone, scrollDown); err != nil {
		return err
	}
	if err := gui.g.SetKeybinding("rows", gocui.MouseWheelUp, gocui.ModNone, scrollUp); err != nil {
		return err
	}
	if err := gui.g.SetKeybinding("rows", gocui.MouseWheelDown, gocui.ModNone, scrollDown); err != nil {
		return err
	}
	return nil
}

func scrollUp(g *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	v.SetOrigin(ox, max(0, oy-1)) // Scroll up
	return nil
}

func scrollDown(g *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	v.SetOrigin(ox, oy+1) // Scroll down
	return nil
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
