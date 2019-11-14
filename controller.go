package apollo

type Controller interface {
	Process(m Model) View
}
