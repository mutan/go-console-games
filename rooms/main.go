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
	choices     []*choice
}

type choice struct {
	cmd         string
	description string
	nextRoom    *room
}

func (room *room) addChoice(cmd string, description string, nextRoom *room) {
	choice := &choice{cmd, description, nextRoom}
	room.choices = append(room.choices, choice)
}

func (room *room) removeChoice(cmd string) {
	var newChoices []*choice
	for _, choice := range room.choices {
		if choice.cmd != cmd {
			newChoices = append(newChoices, choice)
		}
	}
	room.choices = newChoices
}

func (room *room) render() {
	fmt.Printf("Вы входите в комнату: %s.\n", room.name)
	fmt.Println(room.description)
	if room.choices != nil {
		for _, choice := range room.choices {
			fmt.Printf("%s: %s\n", choice.cmd, choice.description)
		}
	}
}

func (room *room) executeCmd(cmd string) *room {
	for _, choice := range room.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextRoom
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return room
}

func play(room *room) {
	room.render()
	if room.choices != nil {
		scanner.Scan()
		fmt.Println()
		play(room.executeCmd(scanner.Text()))
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

	hall.addChoice("N", "На север", &darkRoom)
	hall.addChoice("S", "На юг", &darkRoom)
	hall.addChoice("E", "На восток", &trap)

	darkRoom.addChoice("S", "Назад на юг", &grue)
	darkRoom.addChoice("O", "Включить фонарь", &darkRoomLit)

	darkRoomLit.addChoice("N", "На север", &treasure)
	darkRoomLit.addChoice("S", "На юг", &hall)

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
