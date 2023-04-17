![teamenu](https://github.com/jassummisko/teamenu/blob/master/img/logo.png)

# teamenu: a simple terminal menu utility

## What does it do?

What teamenu does is simple:
1. it reads stdin
2. presents a menu containing each line of stdin (delimited by newline)
3. it outputs the selected option to stdout

Use the up and down arrows to move between menu options and enter to finalize your selection.

![demo](https://github.com/jassummisko/teamenu/blob/master/img/demo.gif)

## Usage:
Just pipe in input from stdin to teamenu.
1. From another shell command: `echo -e "option1\noption2\noption3" | teamenu`
2. From a file: `teamenu < options.txt` 

## Flags:
- `-t <title>` -- adds a title at the top of the menu
- `-c` -- centers the menu
