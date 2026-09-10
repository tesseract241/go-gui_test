package main
import (
	_ "embed"
	"github.com/go-gui-org/go-gui/gui"
	"github.com/go-gui-org/go-gui/gui/backend"
	"fmt"
	"strings"
	"time"
)

type State struct {
	text string
	c chan string
	i int
	l int
}

//go:embed text.txt
var source_text string

func generate_text(state *State, lines []string){
	for {
		state.c<-lines[state.i]
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	lines := strings.Split(source_text, string('\n'))
	fmt.Print(len(lines))
	state := &State{c : make(chan string, 1), l : len(lines)}
	go generate_text(state, lines)
	w := gui.NewWindow(gui.WindowCfg{
		State: state, Title: "Text Test", Width: 600, Height: 100, OnInit: func(w *gui.Window) {
			w.UpdateView(mainView)
		}, 
	})
	backend.Run(w)
}

func mainView(w *gui.Window) gui.View {
	state := gui.State[State](w)
	select {
		case text := <- state.c :
			state.text = text
			state.i++
			state.i%=state.l
		default :
	}
	return gui.Column(gui.ContainerCfg{
		Sizing: gui.FillFill,
		HAlign: gui.HAlignCenter,
		VAlign: gui.VAlignMiddle,
		Content: []gui.View{
			gui.Label(state.text, gui.CurrentTheme().N1),
		},
	})
}
