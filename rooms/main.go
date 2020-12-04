/*
Rooms diagram
https://lucid.app/lucidchart/invitations/accept/a566a6b2-36d2-4fe8-b11d-9d6d99af9add
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

var flags map[string]bool = map[string]bool{
	"check_pockets":  false,
	"find_safe_code": false,
}

type room struct {
	name        string
	description string
	moves       []*move
	actions     []*action
}

type move struct {
	cmd         string
	description string
	nextRoom    *room
}

type action struct {
	cmd         string
	description string
}

func (room *room) addMove(cmd string, description string, nextRoom *room) {
	move := &move{cmd, description, nextRoom}
	room.moves = append(room.moves, move)
}

func (room *room) addAction(cmd string, description string) {
	action := &action{cmd, description}
	room.actions = append(room.actions, action)
}

func (room *room) removeMove(cmd string) {
	var newMoves []*move
	for _, move := range room.moves {
		if move.cmd != cmd {
			newMoves = append(newMoves, move)
		}
	}
	room.moves = newMoves
}

func (room *room) removeAction(cmd string) {
	var newActions []*action
	for _, action := range room.actions {
		if action.cmd != cmd {
			newActions = append(newActions, action)
		}
	}
	room.actions = newActions
}

func (room *room) executeMoveCmd(cmd string) *room {
	for _, move := range room.moves {
		if strings.ToLower(move.cmd) == strings.ToLower(cmd) {
			return move.nextRoom
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return room
}

/*func (room *room) executeActionCmd(cmd string) {
	for _, action := range room.actions {
		if strings.ToLower(action.cmd) == strings.ToLower(cmd) {
			fmt.Println(action.description)
			room.removeAction(cmd)
		}
	}
	fmt.Println("Sorry, I didn't understand that.")

}*/

func (room *room) render() {
	fmt.Printf("Вы входите в комнату: %s.\n", room.name)
	fmt.Println(room.description)
	if room.moves != nil {
		fmt.Println("Движения:")
		for _, move := range room.moves {
			fmt.Printf("%s: %s\n", move.cmd, move.description)
		}
	}
	if room.actions != nil {
		fmt.Println("Действия:")
		for _, action := range room.actions {
			fmt.Printf("%s: %s\n", action.cmd, action.description)
		}
	}
}

func play(room *room) {
	room.render()
	if room.moves != nil {
		scanner.Scan()
		fmt.Println()
		play(room.executeMoveCmd(scanner.Text()))
	}
}

func main() {
	printHeading()
	scanner = bufio.NewScanner(os.Stdin)

	hall := room{name: "Холл", description: "Вы обнаруживаете себя в комнате. Сзади - входная дверь."}

	darkRoom := room{name: "Темная комната", description: "Темная комната. Вы ничего не видите."}

	darkRoomLit := room{name: "Темная комната освещена", description: "Темная комната теперь освещена."}

	grue := room{name: "Чудище", description: "Вас съедает чудище из темноты."}

	trap := room{name: "Ловушка", description: "Это ловушка."}

	treasure := room{name: "Сокровищница", description: "Комната, полная сокровищ."}

	hall.addMove("N", "На север", &darkRoom)
	hall.addMove("S", "На юг", &darkRoom)
	hall.addMove("E", "На восток", &trap)
	hall.addAction("P", "Покопаться по карманам")

	darkRoom.addMove("S", "Назад на юг", &grue)
	darkRoom.addMove("O", "Включить фонарь", &darkRoomLit)

	darkRoomLit.addMove("N", "На север", &treasure)
	darkRoomLit.addMove("S", "На юг", &hall)

	play(&hall)

	fmt.Println("The End.")
}

func printHeading() {
	fmt.Print(`
 ____
|  _ \    ___     ___    _ __ ___    ___
| |_) |  / _ \   / _ \  | '_ ' _ \  / __|
|  _ <  | (_) | | (_) | | | | | | | \__ \
|_| \_\  \___/   \___/  |_| |_| |_| |___/
`)

	fmt.Print(`
                Rooms v1.0                
  Horror game by Akim Gubanov (c) 2020`, "\n\n")

	return
}
