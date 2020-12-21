package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/urfave/cli/v2"
)

var IMailCMD = cli.Command{
	Name:    "imail",
	Aliases: []string{"im"},
	Usage:   "发送邮件",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for this",
		},
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	},
	Action: func(c *cli.Context) error {
		IMail()
		return nil
	},
}

// const NumGoroutines = 10

// var (
// 	done = make(chan struct{})
// 	wg   sync.WaitGroup

// 	mu       sync.Mutex // protects ctr
// 	ctr      = 0
// )

// // IMail send emails
// func IMail_() {
// 	g, err := gocui.NewGui(gocui.OutputNormal)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer g.Close()

// 	g.SetManagerFunc(layout)

// 	if err := keybindings(g); err != nil {
// 		log.Panicln(err)
// 	}

// 	for i := 0; i < NumGoroutines; i++ {
// 		wg.Add(1)
// 		go counter(g)
// 	}

// 	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
// 		log.Panicln(err)
// 	}

// 	wg.Wait()
// }

// func layout(g *gocui.Gui) error {
// 	if v, err := g.SetView("ctr", 2, 2, 12, 4); err != nil {
// 		if err != gocui.ErrUnknownView {
// 			return err
// 		}
// 		fmt.Fprintln(v, "0")
// 	}
// 	return nil
// }

// func keybindings(g *gocui.Gui) error {
// 	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func quit(g *gocui.Gui, v *gocui.View) error {
// 	close(done)
// 	return gocui.ErrQuit
// }

// func counter(g *gocui.Gui) {
// 	defer wg.Done()

// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case <-time.After(500 * time.Millisecond):
// 			mu.Lock()
// 			n := ctr
// 			ctr++
// 			mu.Unlock()

// 			g.Update(func(g *gocui.Gui) error {
// 				v, err := g.View("ctr")
// 				if err != nil {
// 					return err
// 				}
// 				v.Clear()
// 				fmt.Fprintln(v, n)
// 				return nil
// 			})
// 		}
// 	}
// }

/////////////

// const delta = 0.2

// type HelpWidget struct {
// 	name string
// 	x, y int
// 	w, h int
// 	body string
// }

// func NewHelpWidget(name string, x, y int, body string) *HelpWidget {
// 	lines := strings.Split(body, "\n")

// 	w := 0
// 	for _, l := range lines {
// 		if len(l) > w {
// 			w = len(l)
// 		}
// 	}
// 	h := len(lines) + 1
// 	w = w + 1

// 	return &HelpWidget{name: name, x: x, y: y, w: w, h: h, body: body}
// }

// func (w *HelpWidget) Layout(g *gocui.Gui) error {
// 	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
// 	if err != nil {
// 		if err != gocui.ErrUnknownView {
// 			return err
// 		}
// 		fmt.Fprint(v, w.body)
// 	}
// 	return nil
// }

// type StatusbarWidget struct {
// 	name string
// 	x, y int
// 	w    int
// 	val  float64
// }

// func NewStatusbarWidget(name string, x, y, w int) *StatusbarWidget {
// 	return &StatusbarWidget{name: name, x: x, y: y, w: w}
// }

// func (w *StatusbarWidget) SetVal(val float64) error {
// 	if val < 0 || val > 1 {
// 		return errors.New("invalid value")
// 	}
// 	w.val = val
// 	return nil
// }

// func (w *StatusbarWidget) Val() float64 {
// 	return w.val
// }

// func (w *StatusbarWidget) Layout(g *gocui.Gui) error {
// 	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2)
// 	if err != nil && err != gocui.ErrUnknownView {
// 		return err
// 	}
// 	v.Clear()

// 	rep := int(w.val * float64(w.w-1))
// 	fmt.Fprint(v, strings.Repeat("▒", rep))
// 	return nil
// }

// type ButtonWidget struct {
// 	name    string
// 	x, y    int
// 	w       int
// 	label   string
// 	handler func(g *gocui.Gui, v *gocui.View) error
// }

// func NewButtonWidget(name string, x, y int, label string, handler func(g *gocui.Gui, v *gocui.View) error) *ButtonWidget {
// 	return &ButtonWidget{name: name, x: x, y: y, w: len(label) + 1, label: label, handler: handler}
// }

// func (w *ButtonWidget) Layout(g *gocui.Gui) error {
// 	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+2)
// 	if err != nil {
// 		if err != gocui.ErrUnknownView {
// 			return err
// 		}
// 		if _, err := g.SetCurrentView(w.name); err != nil {
// 			return err
// 		}
// 		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, w.handler); err != nil {
// 			return err
// 		}
// 		fmt.Fprint(v, w.label)
// 	}
// 	return nil
// }

// func IMail() {
// 	g, err := gocui.NewGui(gocui.OutputNormal)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer g.Close()

// 	g.Highlight = true
// 	g.SelFgColor = gocui.ColorRed

// 	help := NewHelpWidget("help", 1, 1, helpText)
// 	status := NewStatusbarWidget("status", 1, 7, 50)
// 	butdown := NewButtonWidget("butdown", 52, 7, "DOWN", statusDown(status))
// 	butup := NewButtonWidget("butup", 58, 7, "UP", statusUp(status))
// 	g.SetManager(help, status, butdown, butup)

// 	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
// 		log.Panicln(err)
// 	}
// 	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleButton); err != nil {
// 		log.Panicln(err)
// 	}

// 	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
// 		log.Panicln(err)
// 	}
// }

// func quit(g *gocui.Gui, v *gocui.View) error {
// 	return gocui.ErrQuit
// }

// func toggleButton(g *gocui.Gui, v *gocui.View) error {
// 	nextview := "butdown"
// 	if v != nil && v.Name() == "butdown" {
// 		nextview = "butup"
// 	}
// 	_, err := g.SetCurrentView(nextview)
// 	return err
// }

// func statusUp(status *StatusbarWidget) func(g *gocui.Gui, v *gocui.View) error {
// 	return func(g *gocui.Gui, v *gocui.View) error {
// 		return statusSet(status, delta)
// 	}
// }

// func statusDown(status *StatusbarWidget) func(g *gocui.Gui, v *gocui.View) error {
// 	return func(g *gocui.Gui, v *gocui.View) error {
// 		return statusSet(status, -delta)
// 	}
// }

// func statusSet(sw *StatusbarWidget, inc float64) error {
// 	val := sw.Val() + inc
// 	if val < 0 || val > 1 {
// 		return nil
// 	}
// 	return sw.SetVal(val)
// }

// const helpText = `KEYBINDINGS
// Tab: Move between buttons
// Enter: Push button
// ^C: Exit`

///////////////

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "side" {
		_, err := g.SetCurrentView("main")
		return err
	}
	_, err := g.SetCurrentView("side")
	return err
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func getLine(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}

func delMsg(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("side"); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("side", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyEnter, gocui.ModNone, getLine); err != nil {
		return err
	}
	if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, delMsg); err != nil {
		return err
	}

	if err := g.SetKeybinding("main", gocui.KeyCtrlS, gocui.ModNone, saveMain); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyCtrlW, gocui.ModNone, saveVisualMain); err != nil {
		return err
	}
	return nil
}

func saveMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	p := make([]byte, 5)
	v.Rewind()
	for {
		n, err := v.Read(p)
		if n > 0 {
			if _, err := f.Write(p[:n]); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func saveVisualMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	vb := v.ViewBuffer()
	if _, err := io.Copy(f, strings.NewReader(vb)); err != nil {
		return err
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, 30, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "收 件 人")
		fmt.Fprintln(v, "抄 送")
		fmt.Fprintln(v, "主 题")
		fmt.Fprintln(v, "发 件 人")
		fmt.Fprintln(v, "附 件")

		fmt.Fprint(v, "\rWill be")
		fmt.Fprint(v, "deleted\rItem 4\nItem 5")
	}
	if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		b, err := ioutil.ReadFile("./tools/Mark.Twain-Tom.Sawyer.txt")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(v, "%s", b)
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}

func IMail() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
