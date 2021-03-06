package game

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/toms1441/chess-server/internal/board"
	"github.com/toms1441/chess-server/internal/model"
)

var gGame, _ = NewGame(cl1, cl2)
var specR, specW = io.Pipe()
var specCl = &Client{W: specW}

func TestTurns(t *testing.T) {
	resetPipe()

	cl1.g, cl2.g = nil, nil
	gGame, _ = NewGame(cl1, cl2)
	go gGame.SwitchTurn()

	time.Sleep(time.Millisecond * 10)

	select {
	case <-time.After(time.Millisecond * 100):
		t.Fatalf("timeout")
	case body := <-clientRead(rd1):
		t.Logf("\n%s", string(body))

		x := &model.TurnOrder{}
		u := &model.Order{}

		err := json.Unmarshal(body, u)
		if err != nil {
			t.Fatalf("json.Unmarshal: %s", err.Error())
		}

		err = json.Unmarshal(u.Data, x)
		if err != nil {
			t.Fatalf("json.Unmarshal: %s", err.Error())
		}

		if x.P1 != true {
			t.Fatalf("p1 is not true")
		}
	}
	// turn read
	<-clientRead(rd2)

	cherr := make(chan error)
	go func() {
		id := int8(23)
		dst := board.Point{7, 4}
		body, err := json.Marshal(model.MoveOrder{
			ID:  id,
			Dst: dst,
		})

		if err != nil {
			cherr <- fmt.Errorf("json.Marshal: %s", err.Error())
			return
		}

		c := model.Order{
			ID:   model.OrMove,
			Data: body,
		}

		err = cl2.Do(c)
		if err == nil {
			cherr <- fmt.Errorf("it's not 2 turn, yet it works .. :(")
			return
		}

		err = cl1.Do(c)
		if err != nil {
			cherr <- err
			return
		}
		//cherr <- nil
		close(cherr)
	}()

	<-clientRead(rd1)
	<-clientRead(rd2)

	select {
	case <-time.After(time.Millisecond * 100):
		t.Fatalf("timeout")
	case body := <-clientRead(rd1):
		t.Logf("\n%s", string(body))

		x := &model.TurnOrder{}
		u := &model.Order{}

		err := json.Unmarshal(body, u)
		if err != nil {
			t.Fatalf("json.Unmarshal: %s", err.Error())
		}

		err = json.Unmarshal(u.Data, x)
		if err != nil {
			t.Fatalf("json.Unmarshal: %s", err.Error())
		}

		if x.P1 {
			t.Fatalf("p1 is not false")
		}

		<-clientRead(rd2)

		select {
		case <-time.After(time.Millisecond * 100):
			t.Fatalf("timeout")
		case err := <-cherr:
			if err != nil {
				t.Fatalf("err: %s", err.Error())
			}
		}
	}
}

func TestGameAddSpectator(t *testing.T) {
	go func() {
		<-clientRead(rd2)
	}()
	cl1.LeaveGame()

	resetPipe()

	var err error
	gGame, err = NewGame(cl1, cl2)

	if err != nil {
		t.Fatalf("NewGame: %s", err.Error())
	}

	done := make(chan bool)
	go func() {
		<-clientRead(specR)
		<-clientRead(specR)

		close(done)
	}()
	gGame.AddSpectator(specCl)

	<-done
}

func TestGameSpectatorUpdate(t *testing.T) {
	go gGame.SwitchTurn()

	<-clientRead(rd1)
	<-clientRead(rd2)
	<-clientRead(specR)
}

func TestGameRmSpectator(t *testing.T) {
	gGame.RmSpectator(specCl)

	if _, ok := gGame.spectators[specCl]; ok {
		t.Fatalf("RmSpectator does not remove the spectator")
	}
}

func TestGameDone(t *testing.T) {

	resetPipe()

	cl1.g = nil
	cl2.g = nil

	gGame, _ = NewGame(cl1, cl2)

	go gGame.SwitchTurn()

	by1 := <-clientRead(rd1)
	by2 := <-clientRead(rd2)
	t.Log(string(by1), string(by2))

	done := false

	doMove := func(id int8, dst board.Point) {
		cl := gGame.cs[gGame.turn]

		p1, p2 := id, dst
		x, err := json.Marshal(model.MoveOrder{
			ID:  p1,
			Dst: p2,
		})
		if err != nil {
			t.Fatalf("json.Marshal: %s", err)
		}

		if !done {
			go func() {
				<-clientRead(rd1)
				<-clientRead(rd2)
				<-clientRead(rd1)
				<-clientRead(rd2)
			}()
		}

		err = cl.Do(model.Order{
			ID:   model.OrMove,
			Data: x,
		})
		if err != nil {
			t.Fatalf("cl.Do: %s", err.Error())
		}

	}
	t.Log("first move")
	doMove(21, board.Point{5, 5})
	t.Log("second move")
	doMove(12, board.Point{4, 3})
	t.Log("third move")
	doMove(22, board.Point{6, 4})
	done = true
	resetPipe()
	go func() {
		<-clientRead(rd1)
		<-clientRead(rd2)
		<-clientRead(rd1)
		<-clientRead(rd2)
		<-clientRead(rd2)
		<-clientRead(rd1)
	}()
	t.Log("fourth move")
	doMove(3, board.Point{7, 4})
	t.Log("after fourth move")

	/*
		R N B   K B N R
		P P P P   P P P

		        P
		            P Q
		          P
		P P P P P     P
		R N B Q K B N R
		t.Logf("\n%s", gGame.brd)
	*/

	lc1 := gGame.cs[true]
	lc2 := gGame.cs[false]
	if lc1 != nil || lc2 != nil {
		t.Fatalf("gGame cs: %v | %v", lc1, lc2)
	}

	t.Logf("\n%s", gGame.brd.String())

}
