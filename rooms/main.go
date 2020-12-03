package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

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

//type inventory

func (room *room) addChoice(cmd string, description string, nextRoom *room) {
	choice := &choice{cmd, description, nextRoom}
	room.choices = append(room.choices, choice)
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

func (room *room) play() {
	room.render()
	if room.choices != nil {
		scanner.Scan()
		fmt.Println()
		room.executeCmd(scanner.Text()).play()
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

	hall.play()

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
