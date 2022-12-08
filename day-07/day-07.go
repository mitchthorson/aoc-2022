package day07

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
	name   string
	files  []*File
	dirs   map[string]*Dir
	parent *Dir
}

func (d *Dir) changeDir(dirName string) *Dir {
	switch dirName {
	case "..":
		return d.parent
	case "/":
		newDir := d
		for newDir.parent != nil {
			newDir = d.parent
		}
		return newDir
	default:
		subDir, ok := d.dirs[dirName]
		if !ok {
			panic(fmt.Sprintf("Subdir %s does not exist in %s, %v\n", dirName, d.name, d.dirs))
		}
		return subDir
	}
}

func (d *Dir) setChildren(output []string) *Dir {
	dirs := map[string]*Dir{}
	files := []*File{}
	for _, line := range output {
		outputFields := strings.Fields(line)
		switch outputFields[0] {
		case "dir":
			subd := new(Dir)
			subd.name = outputFields[1]
			subd.parent = d
			dirs[subd.name] =  subd
		default:
			f := new(File)
			f.fromOutput(outputFields)
			files = append(files, f)
		}

	}
	d.files = files
	d.dirs = dirs
	return d
}

func (d *Dir) getDirSize() int {
	size := 0
	// get file sizes first
	for _, file := range d.files {
		size += file.size
	}
	for _, dir := range d.dirs {
		size += dir.getDirSize()
	}
	return size
}

func (d *Dir) getAllDirSizes() []int {
	dirSizes := []int{d.getDirSize()}
	for _, dir := range d.dirs {
		dirSizes = append(dirSizes, dir.getAllDirSizes()...)
	}
	return dirSizes

}

type File struct {
	size int
	name string
}

func (f *File) fromOutput(outputFields []string) *File {
	filesize, err := strconv.Atoi(outputFields[0])
	utils.Check(err)
	f.size = filesize
	f.name = outputFields[1]
	return f
}

type Command struct {
	cmd    string
	arg    string
	output []string
}

func (c *Command) fromString(input string) *Command {
	lines := strings.Split(input, "\n")
	commandArgs := strings.Fields(lines[0])
	c.cmd = commandArgs[0]
	if len(commandArgs) > 1 {
		c.arg = commandArgs[1]
	}
	if len(lines) > 1 {
		commandOutput := []string{}
		for _, output := range lines[1:] {
			if output == "" {
				continue
			}
			commandOutput = append(commandOutput, output)
		}
		c.output = commandOutput
	}
	return c
}

func getFileSystemFromCommands(input string) *Dir {
	commandInput := strings.Split(input, "$")
	// first we need to parse the commands into a useful structure
	var commands []*Command
	for _, c := range commandInput {
		cleanCommand := strings.Trim(c, " ")
		if cleanCommand == "" {
			continue
		}
		command := new(Command)
		commands = append(commands, command.fromString(cleanCommand))
	}
	cwd := new(Dir)
	cwd.name = "/"
	for _, c := range commands {
		switch c.cmd {
		case "cd":
			cwd = cwd.changeDir(c.arg)

		case "ls":
			cwd.setChildren(c.output)
		}
	}

	for cwd.parent != nil {
		cwd = cwd.parent
	}
	return cwd
}

func GetResult1(input string) int {
	fs := getFileSystemFromCommands(input)
	limit := 100000
	dirSizes := fs.getAllDirSizes()
	smallDirs := []int{}
	for _, d := range dirSizes {
		if d < limit {
			smallDirs = append(smallDirs, d)
		}
	}
	result := 0
	for _, d := range smallDirs {
		result += d
	}
	return result
}
func GetResult2(input string) int {
	fs := getFileSystemFromCommands(input)
	TOTAL_SPACE := 70000000
	REQUIRED_SPACE := 30000000
	AVAILABLE_SPACE := TOTAL_SPACE - fs.getDirSize()
	NEED_TO_DELETE := REQUIRED_SPACE - AVAILABLE_SPACE
	allDirs := fs.getAllDirSizes()
	sort.Ints(allDirs)
	result := 0
	for _, d := range allDirs {
		if d > NEED_TO_DELETE {
			result = d
			break
		}
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-07/input.txt")
	fmt.Printf("Day 07 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 07 part 2 result is:\n%d\n", GetResult2(input))
}
