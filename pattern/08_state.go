package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// State
type MobileAlertState interface {
	Alert(ctx *MobilePhoneNotification)
}

// Context
type MobilePhoneNotification struct {
	state MobileAlertState
}

func NewMobilePhoneNotification() *MobilePhoneNotification {
	return &MobilePhoneNotification{
		state: NewSoundState(),
	}
}

func (ctx *MobilePhoneNotification) SetState(state MobileAlertState) {
	ctx.state = state
}

func (ctx *MobilePhoneNotification) Alert() {
	ctx.state.Alert(ctx)
}

// Concrete states
type SoundState struct{}

func NewSoundState() *SoundState {
	return &SoundState{}
}

func (ss *SoundState) Alert(ctx *MobilePhoneNotification) {
	fmt.Println("*Ringtone starts playing*")
}

type VibrationState struct{}

func NewVibrationState() *VibrationState {
	return &VibrationState{}
}

func (vs *VibrationState) Alert(ctx *MobilePhoneNotification) {
	fmt.Println("*Vibrating*")
}

type SilentState struct{}

func NewSilentState() *SilentState {
	return &SilentState{}
}

func (ss *SilentState) Alert(ctx *MobilePhoneNotification) {
	fmt.Println("...")
}

func main() {
	notify := NewMobilePhoneNotification()
	notify.Alert()
	notify.Alert()
	notify.SetState(NewVibrationState())
	notify.Alert()
	notify.SetState(NewSilentState())
	notify.Alert()
	notify.Alert()
}
