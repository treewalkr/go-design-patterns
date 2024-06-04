package abstractfactory

import "fmt"

// UI Component interfaces
type Window interface {
	Render() string
	SetTitle(title string)
	Maximize()
	Minimize()
}

type Button interface {
	Render() string
	SetLabel(label string)
	HandleClick()
}

type Menu interface {
	Render() string
	AddMenuItem(item string)
	SelectItem(index int)
}

// UI Factory interface
type UIFactory interface {
	CreateWindow() Window
	CreateButton() Button
	CreateMenu() Menu
}

// Windows UI Components
type WindowsWindow struct {
	title       string
	isMaximized bool
}

func (w *WindowsWindow) Render() string {
	state := "Normal"
	if w.isMaximized {
		state = "Maximized"
	}
	return fmt.Sprintf("Windows Window [%s] - Title: %s", state, w.title)
}

func (w *WindowsWindow) SetTitle(title string) {
	w.title = title
}

func (w *WindowsWindow) Maximize() {
	w.isMaximized = true
	fmt.Println("Windows window maximized with Windows animation")
}

func (w *WindowsWindow) Minimize() {
	w.isMaximized = false
	fmt.Println("Windows window minimized to taskbar")
}

type WindowsButton struct {
	label string
}

func (b *WindowsButton) Render() string {
	return fmt.Sprintf("Windows Button [%s] with gray background", b.label)
}

func (b *WindowsButton) SetLabel(label string) {
	b.label = label
}

func (b *WindowsButton) HandleClick() {
	fmt.Printf("Windows button '%s' clicked with ripple effect\n", b.label)
}

type WindowsMenu struct {
	items []string
}

func (m *WindowsMenu) Render() string {
	return fmt.Sprintf("Windows Menu Bar with items: %v", m.items)
}

func (m *WindowsMenu) AddMenuItem(item string) {
	m.items = append(m.items, item)
}

func (m *WindowsMenu) SelectItem(index int) {
	if index < len(m.items) {
		fmt.Printf("Selected Windows menu item: %s with highlight effect\n", m.items[index])
	}
}

// macOS UI Components
type MacWindow struct {
	title       string
	isMaximized bool
}

func (w *MacWindow) Render() string {
	state := "Normal"
	if w.isMaximized {
		state = "Maximized"
	}
	return fmt.Sprintf("macOS Window [%s] - Title: %s with traffic light buttons", w.title, state)
}

func (w *MacWindow) SetTitle(title string) {
	w.title = title
}

func (w *MacWindow) Maximize() {
	w.isMaximized = true
	fmt.Println("macOS window maximized with zoom animation")
}

func (w *MacWindow) Minimize() {
	w.isMaximized = false
	fmt.Println("macOS window minimized with genie effect")
}

type MacButton struct {
	label string
}

func (b *MacButton) Render() string {
	return fmt.Sprintf("macOS Button [%s] with gradient background", b.label)
}

func (b *MacButton) SetLabel(label string) {
	b.label = label
}

func (b *MacButton) HandleClick() {
	fmt.Printf("macOS button '%s' clicked with smooth animation\n", b.label)
}

type MacMenu struct {
	items []string
}

func (m *MacMenu) Render() string {
	return fmt.Sprintf("macOS Menu Bar (top of screen) with items: %v", m.items)
}

func (m *MacMenu) AddMenuItem(item string) {
	m.items = append(m.items, item)
}

func (m *MacMenu) SelectItem(index int) {
	if index < len(m.items) {
		fmt.Printf("Selected macOS menu item: %s with smooth dropdown\n", m.items[index])
	}
}

// Concrete Factories
type WindowsUIFactory struct{}

func (f *WindowsUIFactory) CreateWindow() Window {
	return &WindowsWindow{}
}

func (f *WindowsUIFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (f *WindowsUIFactory) CreateMenu() Menu {
	return &WindowsMenu{items: make([]string, 0)}
}

type MacUIFactory struct{}

func (f *MacUIFactory) CreateWindow() Window {
	return &MacWindow{}
}

func (f *MacUIFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacUIFactory) CreateMenu() Menu {
	return &MacMenu{items: make([]string, 0)}
}

// Helper function to create application window
func CreateAppWindow(factory UIFactory) {
	// Create window and components
	window := factory.CreateWindow()
	button := factory.CreateButton()
	menu := factory.CreateMenu()

	// Configure window
	window.SetTitle("Cross-Platform App")

	// Configure button
	button.SetLabel("Click Me")

	// Configure menu
	menu.AddMenuItem("File")
	menu.AddMenuItem("Edit")
	menu.AddMenuItem("Help")

	// Render UI
	fmt.Println(window.Render())
	fmt.Println(button.Render())
	fmt.Println(menu.Render())

	// Simulate user interactions
	window.Maximize()
	button.HandleClick()
	menu.SelectItem(0)
	window.Minimize()
}

func Example() {
	// Create Windows version
	fmt.Println("=== Creating Windows Application ===")
	windowsFactory := &WindowsUIFactory{}
	CreateAppWindow(windowsFactory)

	fmt.Println("\n=== Creating macOS Application ===")
	macFactory := &MacUIFactory{}
	CreateAppWindow(macFactory)
}
