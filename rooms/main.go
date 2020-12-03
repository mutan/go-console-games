package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

type storyNode struct {
	text    string
	choices []*choice
}

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Printf("%s: %s\n", choice.cmd, choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return node
}

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	printHeading()
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: "Вы обнаруживаете себя в комнате. Сзади - входная дверь."}

	darkRoom := storyNode{text: "Темная комната. Вы ничего не видите."}

	darkRoomLit := storyNode{text: "Темная комната теперь освещена."}

	grue := storyNode{text: "Вас съедает чудище из темноты."}

	trap := storyNode{text: "Это ловушка."}

	treasure := storyNode{text: "Комната, полная сокровищ."}

	start.addChoice("N", "На север", &darkRoom)
	start.addChoice("S", "На юг", &darkRoom)
	start.addChoice("E", "На восток", &trap)

	darkRoom.addChoice("S", "Назад на юг", &grue)
	darkRoom.addChoice("O", "Включить фонарь", &darkRoomLit)

	darkRoomLit.addChoice("N", "На север", &treasure)
	darkRoomLit.addChoice("S", "На юг", &start)

	start.play()

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
