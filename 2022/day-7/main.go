package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_SIZE = 100000
const MAX_UPDATE_SIZE = 30000000
const FILE_SYSTEM_SIZE = 70000000

type file struct {
	name string
	size int
}

type dir struct {
	name   string
	parent *dir
	dirs   []*dir
	files  []*file
	size   int
}

type tree struct {
	ref *dir
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	dir := &dir{
		name:   "/",
		parent: nil,
		dirs:   []*dir{},
		files:  []*file{},
		size:   0,
	}

	tree := &tree{
		ref: dir,
	}

	tree.createDirTree(input)
	return sumDirs(dir)
}

func part2(input string) int {
	dir := &dir{
		name:   "/",
		parent: nil,
		dirs:   []*dir{},
		files:  []*file{},
		size:   0,
	}

	tree := &tree{
		ref: dir,
	}

	tree.createDirTree(input)

	freeSpace := FILE_SYSTEM_SIZE - dir.size
	neededSpace := MAX_UPDATE_SIZE - freeSpace

	return findDirToDelete(dir, neededSpace)

}

func findDirToDelete(d *dir, neededSpace int) int {
	small := FILE_SYSTEM_SIZE

	if d.size > neededSpace && d.size < small {
		small = d.size
	}

	for _, dir := range d.dirs {
		if value := findDirToDelete(dir, neededSpace); value < small {
			small = value
		}
	}

	return small
}

func sumDirs(d *dir) int {
	var sum int
	if d.size <= MAX_SIZE {
		sum += d.size
	}

	for _, dir := range d.dirs {
		sum += sumDirs(dir)
	}
	return sum
}

func (t *tree) createDirTree(input string) {
	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			continue
		}
		if strings.Contains(line, "$") {
			t.processCommand(line)
			continue
		}
		if strings.Contains(line, "dir") {
			dir := &dir{
				name:   strings.Split(line, " ")[1],
				parent: t.ref,
				dirs:   []*dir{},
				files:  []*file{},
				size:   0,
			}
			t.ref.dirs = append(t.ref.dirs, dir)
			continue
		}
		splitted := strings.Split(line, " ")
		size, _ := strconv.Atoi(splitted[0])
		fileName := splitted[1]
		file := &file{
			name: fileName,
			size: size,
		}
		t.ref.files = append(t.ref.files, file)
		t.updateParentsSize(size)
	}
}

func (t *tree) processCommand(command string) {
	splitted := strings.Split(command, " ")
	if strings.Contains(command, "$ cd") {
		to := splitted[2]
		if to == ".." {
			t.ref = t.ref.parent
			return
		}
		t.ref = t.getDirByName(to)
	}
}

func (t *tree) updateParentsSize(size int) {
	ref := t.ref
	for t.ref != nil {
		t.ref.size = t.ref.size + size
		t.ref = t.ref.parent
	}
	t.ref = ref
}

func (t *tree) getDirByName(name string) *dir {
	for _, dir := range t.ref.dirs {
		if dir.name == name {
			return dir
		}
	}
	return &dir{
		name:   name,
		parent: t.ref,
		dirs:   []*dir{},
		files:  []*file{},
		size:   0,
	}
}
