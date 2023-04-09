![tmenu](https://github.com/jassummisko/tmenu/blob/master/img/logo.png)

# tmenu: a terminal menu inspired by dmenu

## What does it do?

What tmenu does is simple:
1. it reads stdin
2. presents a menu containing each line of stdin (delimited by newline)
3. it outputs the selected option to stdout

Use the up and down arrows to move between menu options and enter to finalize your selection.

![demo](https://github.com/jassummisko/tmenu/blob/master/img/demo.gif)

## Usage:
Just pipe in input from stdin to tmenu.
1. From another shell command: `echo -e "option1\noption2\noption3" | tmenu`
2. From a file: `tmenu < options.txt` 