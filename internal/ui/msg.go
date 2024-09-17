package ui

import (
	mathf32 "github.com/bloodmagesoftware/architect/internal/math/f32"
	"github.com/charmbracelet/log"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"sync"
)

type msgData struct {
	title   string
	message string
	buttons string
	done    chan int32
}

var (
	messages = make([]msgData, 0)
	msgMut   = sync.Mutex{}
)

func drawMsg() {
	if len(messages) == 0 {
		return
	}

	msgMut.Lock()
	defer msgMut.Unlock()
	m := messages[0]

	w := mathf32.Clamp(16*vw, 512, 95*vw)
	h := mathf32.Clamp(8*vh, 128, 75*vh)
	x := 50*vw - w/2
	y := 50*vh - h/2

	i := rg.MessageBox(rl.NewRectangle(x, y, w, h), m.title, m.message, m.buttons)
	if i != -1 {
		m.done <- i
		messages = messages[1:]
		close(m.done)
	}
}

func msgErr(err error) <-chan int32 {
	log.Error(err.Error())
	return msgSimple("Error", err.Error())
}

func msgSimple(title, content string) <-chan int32 {
	return msg(title, content, "OK")
}

func msg(title, content, buttons string) <-chan int32 {
	msgMut.Lock()
	defer msgMut.Unlock()

	done := make(chan int32)
	messages = append(messages, msgData{title: title, message: content, buttons: buttons, done: done})
	return done
}
