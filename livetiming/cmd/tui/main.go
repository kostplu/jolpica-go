package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/kostplu/jolpica-go/livetiming"
)

func main() {
	client := livetiming.NewClient(livetiming.WithCache("/tmp/livetiming.db", 24*time.Hour))
	p := tea.NewProgram(
		NewModel(client),
	)
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
