package main

import (
	"errors"
	"fmt"
)

type Volume byte
type Mute bool

type Setting struct {
	Name string
	Age  int
}

/* Originator */

type _Originator struct {
	Setting Setting
}

func (o *_Originator) GetSetting() {

}

func (o *_Originator) SetSettingsFromMemento(m _Memento) {
	o.Setting = m.state
}

func (o *_Originator) GetMemento() _Memento {
	return _Memento{state: o.Setting}
}

/* Memento */

type _Memento struct {
	state Setting
}

/* CareTaker */

type _CareTaker struct {
	mementos []_Memento
}

func (ct *_CareTaker) Push(m _Memento) {
	ct.mementos = append(ct.mementos, m)
}

func (ct *_CareTaker) Pop() (memento _Memento, err error) {
	mSize := len(ct.mementos)
	if mSize > 0 {
		memento = ct.mementos[mSize-1]
		ct.mementos = ct.mementos[0 : mSize-1]
	} else {
		memento = _Memento{}
		err = errors.New("No more mementos")
	}
	return
}

type MementoFacade struct {
	CurrentSetup _Originator
	careTaker    _CareTaker
}

func (mf *MementoFacade) SaveSetup() {
	mf.careTaker.Push(mf.CurrentSetup.GetMemento())
}

func (mf *MementoFacade) RestorePreviousSetup() {
	memento, err := mf.careTaker.Pop()
	if err == nil {
		mf.CurrentSetup.SetSettingsFromMemento(memento)
	}
}

func (mf *MementoFacade) PrintSetup() {
	fmt.Println(mf.CurrentSetup.GetMemento().state)
}

func main() {
	mf := MementoFacade{}
	setting := &mf.CurrentSetup.Setting
	setting.Age = 5
	setting.Name = "Daniel"
	mf.SaveSetup()
	mf.PrintSetup()
	setting = &mf.CurrentSetup.Setting
	setting.Age = 14
	setting.Name = "Juan"
	mf.PrintSetup()

	mf.RestorePreviousSetup()
	mf.PrintSetup()
}
