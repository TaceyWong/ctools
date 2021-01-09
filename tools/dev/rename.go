package dev

import "github.com/urfave/cli/v2"

// RenameCMD 多文件重命名
var RenameCMD = cli.Command{
	Name: "rename",
	UsageText: `ctools rename [ -h|-m|-V ] [ -v ] [ -0 ] [ -n ] [ -f ] [ -d ]
    [ -e|-E perlexpr]*|perlexpr [ files ]`,
	Usage: "renames multiple files",
	Description: `"rename" renames the filenames supplied according to the rule specified as the first argument.  The perlexpr argument is a Perl
	expression which is expected to modify the $_ string in Perl for at least some of the filenames specified.  If a given filename is not
	modified by the expression, it will not be renamed.  If no filenames are given on the command line, filenames will be read via standard
	input.

	For example, to rename all files matching "*.bak" to strip the extension, you might say

			rename 's/\.bak$//' *.bak

	To translate uppercase names to lower, you'd use

			rename 'y/A-Z/a-z/' *`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "输出成功命名文件的名称",
		},
		&cli.BoolFlag{
			Name:    "null",
			Aliases: []string{"0"},
			Usage:   `Use \0 as record separator when reading from STDIN.`,
		},
		&cli.BoolFlag{
			Name:    "nono",
			Aliases: []string{"n"},
			Usage:   "无动作：打印将要被重命名的文件名称，但不真正进行重命名",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "覆盖：允许对已经存在的文件进行覆盖",
		},
		&cli.BoolFlag{
			Name:  "fullpath",
			Usage: "Rename full path: including any directory component",
			Value: true,
		}, &cli.BoolFlag{
			Name:    "filename",
			Aliases: []string{"d"},
			Usage:   "Do not rename directory: only rename filename component of path.",
		},
		&cli.BoolFlag{
			Name:    "expression",
			Aliases: []string{"e"},
			Usage: `Expression: code to act on files name.

            May be repeated to build up code (like "perl -e"). If no -e, the
            first argument is used as code.`,
		}, &cli.BoolFlag{
			Name:    "statement",
			Aliases: []string{"E"},
			Usage:   "Statement: code to act on files name, as -e but terminated by ';'.",
		},
	},
}
