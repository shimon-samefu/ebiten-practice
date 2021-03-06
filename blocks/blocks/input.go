// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package blocks

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input manages the input state including gamepads and keyboards.
type Input struct {
	gamepadIDs                 []ebiten.GamepadID
}

// GamepadIDButtonPressed returns a gamepad ID where at least one button is pressed.
// If no button is pressed, GamepadIDButtonPressed returns -1.
func (i *Input) GamepadIDButtonPressed() ebiten.GamepadID {
	return -1
}

func (i *Input) Update() {
	return
}

func (i *Input) IsRotateRightJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyX) {
		return true
	}
	return false
}

func (i *Input) IsRotateLeftJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		return true
	}
	return false
}

func (i *Input) StateForLeft() int {
	if v := inpututil.KeyPressDuration(ebiten.KeyArrowLeft); 0 < v {
		return v
	}
	return 0
}

func (i *Input) StateForRight() int {
	if v := inpututil.KeyPressDuration(ebiten.KeyArrowRight); 0 < v {
		return v
	}
	return 0
}

func (i *Input) StateForDown() int {
	if v := inpututil.KeyPressDuration(ebiten.KeyArrowDown); 0 < v {
		return v
	}
	return 0
}
