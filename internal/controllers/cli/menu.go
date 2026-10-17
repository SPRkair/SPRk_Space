package cli

import (
	"fmt"
)

type Menu struct {
	current *screen
	screens []*screen
	exit    bool
}

type screen struct {
	name  string
	items []item
}

type item struct {
	submenu *screen
	caption string
	action  func()
}

func newMenu(h *handlers) *Menu {
	fakeSerice := &screen{
		name: "FakeService",
		items: []item{
			{caption: "Show users list", action: func() { h.ShowUsersList() }},
			{caption: "Show user by id", action: func() { h.GetUserById() }},
			{caption: "Update user", action: func() { h.UpdateUser() }},
		},
	}
	services := &screen{
		name: "Services",
		items: []item{
			{caption: "FakeService", submenu: fakeSerice},
			{caption: "UserService", action: func() { /* ... */ }},
		},
	}
	logs := &screen{
		name: "Log",
		items: []item{
			{caption: "Show last 10 records", action: func() { /* ... */ }},
			{caption: "Clear", action: func() { /* ... */ }},
		},
	}
	settings := &screen{
		name:  "Settings",
		items: []item{},
	}
	main := &screen{
		name: "Main",
		items: []item{
			{caption: "Services", submenu: services},
			{caption: "Log", submenu: logs},
			{caption: "Settings", submenu: settings},
		},
	}

	return &Menu{current: main}
}

func (m *Menu) show() {
	m.clearScreen()
	fmt.Println("== " + m.current.name + " ==")

	for i, it := range m.current.items {
		fmt.Printf("%d. %s\n", i+1, it.caption)
	}

	if len(m.screens) == 0 {
		fmt.Println("0. Exit")
	} else {
		fmt.Println("0. Back")
	}
}

func (m *Menu) clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func (m *Menu) Select(n int) {
	if n == 0 {
		if len(m.screens) == 0 {
			m.exit = true
			return
		}
		m.current = m.screens[len(m.screens)-1]
		m.screens = m.screens[:len(m.screens)-1]
		return
	}
	if n < 1 || n > len(m.current.items) {
		return
	}
	it := m.current.items[n-1]
	if it.submenu != nil {
		m.screens = append(m.screens, m.current)
		m.current = it.submenu
	} else if it.action != nil {
		it.action()
	}
}
