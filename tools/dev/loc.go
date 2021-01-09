package dev

import (
	"github.com/TaceyWong/ctools/pkg/linecount"
	"github.com/urfave/cli/v2"
)

var lineCountPathDenyList = cli.NewStringSlice()
var lineCountAllowListExtensions = cli.NewStringSlice()
var lineCountGeneratedMarkers = cli.NewStringSlice()
var lineCountExclude = cli.NewStringSlice()

// LineCountCMD linecount命令用户界面
var LineCountCMD = cli.Command{
	Name:    "linecount",
	Aliases: []string{"loc", "scc"},
	Usage:   "统计项目代码量并进行复杂度估计",
	Action: func(c *cli.Context) error {
		parseStringSliceFlag2StringList()
		linecount.DirFilePaths = c.Args().Slice()
		if linecount.ConfigureLimits != nil {
			linecount.ConfigureLimits()
		}
		linecount.ConfigureGc()
		linecount.ConfigureLazy(true)
		linecount.Process()
		return nil
	},
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:        "avg-wage",
			Value:       56286,
			Usage:       "基本COCOMO计算的平均工资值",
			Destination: &linecount.AverageWage,
		},
		&cli.BoolFlag{
			Name:        "binary",
			Value:       false,
			Usage:       "禁用二进制文件检测",
			Destination: &linecount.DisableCheckBinary,
		},
		&cli.BoolFlag{
			Name:        "by-file",
			Value:       false,
			Usage:       "",
			Destination: &linecount.Files,
		},
		&cli.BoolFlag{
			Name:        "ci",
			Value:       false,
			Usage:       "",
			Destination: &linecount.Ci,
		},
		&cli.BoolFlag{
			Name:        "no-ignore",
			Value:       false,
			Usage:       "",
			Destination: &linecount.Ignore,
		},
		&cli.BoolFlag{
			Name:        "no-gitignore",
			Value:       false,
			Usage:       "",
			Destination: &linecount.GitIgnore,
		},
		&cli.BoolFlag{
			Name:        "debug",
			Value:       false,
			Usage:       "开启调试输出",
			Destination: &linecount.Debug,
		},
		&cli.StringSliceFlag{
			Name:        "exclude-dir",
			Value:       cli.NewStringSlice(".git", ".hg", ".svn"),
			Usage:       "要排除的目录",
			Destination: lineCountPathDenyList,
		},
		&cli.IntFlag{
			Name:        "file-gc-count",
			Value:       10000,
			Usage:       "开启GC的解析文件数量",
			Destination: &linecount.GcFileCount,
		},
		&cli.StringFlag{
			Name:        "format",
			Aliases:     []string{"f"},
			Value:       "tabular",
			Usage:       "设置输出格式[tabular, wide, json, csv, cloc-yaml, html, html-table]",
			Destination: &linecount.Format,
		},
		&cli.StringSliceFlag{
			Name:        "include-ext",
			Aliases:     []string{"i"},
			Value:       cli.NewStringSlice(),
			Usage:       "限制文件后缀 [逗号分割列表: e.g. go,java,js]",
			Destination: lineCountAllowListExtensions,
		},
		&cli.BoolFlag{
			Name:        "languages",
			Aliases:     []string{"l"},
			Value:       false,
			Usage:       "打印支持的语言和文件后缀",
			Destination: &linecount.Languages,
		},
		&cli.BoolFlag{
			Name:        "no-cocomo",
			Value:       false,
			Usage:       "移除COCOMO计算的输出",
			Destination: &linecount.Cocomo,
		},
		&cli.BoolFlag{
			Name:        "no-size",
			Value:       false,
			Usage:       "移除文件大小计算的输出",
			Destination: &linecount.Size,
		},
		&cli.StringFlag{
			Name:        "size-unit",
			Value:       "si",
			Usage:       "设置文件大小的单位[si, binary, mixed, xkcd-kb, xkcd-kelly, xkcd-imaginary, xkcd-intel, xkcd-drive, xkcd-bakers]",
			Destination: &linecount.SizeUnit,
		},
		&cli.BoolFlag{
			Name:        "no-complexity",
			Aliases:     []string{"c"},
			Value:       false,
			Usage:       "跳过代码复杂度计算",
			Destination: &linecount.Complexity,
		},
		&cli.BoolFlag{
			Name:        "no-duplicates",
			Aliases:     []string{"d"},
			Value:       false,
			Usage:       "从统计和输出中移除重复文件",
			Destination: &linecount.Duplicates,
		},
		&cli.BoolFlag{
			Name:        "min-gen",
			Aliases:     []string{"z"},
			Value:       false,
			Usage:       "识别压缩和自动生成的文件",
			Destination: &linecount.MinifiedGenerated,
		},
		&cli.BoolFlag{
			Name:        "min",
			Value:       false,
			Usage:       "识别压缩文件",
			Destination: &linecount.Minified,
		},
		&cli.BoolFlag{
			Name:        "gen",
			Value:       false,
			Usage:       "识别自动生成的文件",
			Destination: &linecount.Generated,
		},
		&cli.StringSliceFlag{
			Name:        "generated-markers",
			Value:       cli.NewStringSlice("do not edit", "不要编辑"),
			Usage:       "自动生成文件的文件头标记",
			Destination: lineCountGeneratedMarkers,
		},
		&cli.BoolFlag{
			Name:        "no-min-gen",
			Value:       false,
			Usage:       "在输出中忽略压缩和生成的文件(隐含 --min-gen)",
			Destination: &linecount.IgnoreMinifiedGenerate,
		},
		&cli.BoolFlag{
			Name:        "no-min",
			Value:       false,
			Usage:       "在输出中忽略压缩文件(隐含 --min)",
			Destination: &linecount.IgnoreMinified,
		},
		&cli.BoolFlag{
			Name:        "no-gen",
			Value:       false,
			Usage:       "在输出中会略自动生成的文件(隐含 --gen)",
			Destination: &linecount.IgnoreGenerated,
		},
		&cli.IntFlag{
			Name:        "min-gen-line-length",
			Value:       255,
			Usage:       "number of bytes per average line for file to be considered minified or generated",
			Destination: &linecount.MinifiedGeneratedLineByteLength,
		},
		&cli.StringSliceFlag{
			Name:        "not-match",
			Aliases:     []string{"M"},
			Value:       cli.NewStringSlice(),
			Usage:       "忽略匹配正则表达式的文件和目录",
			Destination: lineCountExclude,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Value:       "",
			Usage:       "输出文件名（缺省为标准输出）",
			Destination: &linecount.FileOutput,
		},
		&cli.StringFlag{
			Name:        "sort",
			Aliases:     []string{"s"},
			Value:       "files",
			Usage:       "column to sort by [files, name, lines, blanks, code, comments, complexity]",
			Destination: &linecount.SortBy,
		},
		&cli.BoolFlag{
			Name:        "trace",
			Aliases:     []string{"t"},
			Value:       false,
			Usage:       "enable trace output (not recommended when processing multiple files)",
			Destination: &linecount.Trace,
		},
		&cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"v"},
			Value:       false,
			Usage:       "verbose output",
			Destination: &linecount.Verbose,
		},
		&cli.BoolFlag{
			Name:        "wide",
			Aliases:     []string{"w"},
			Value:       false,
			Usage:       "包含更多统计信息的宽输出 (隐含 --complexity)",
			Destination: &linecount.More,
		},
		&cli.BoolFlag{
			Name:        "no-large",
			Value:       false,
			Usage:       "ignore files over certain byte and line size set by max-line-count and max-byte-count",
			Destination: &linecount.NoLarge,
		},
		&cli.BoolFlag{
			Name:        "include-synlinks",
			Value:       false,
			Usage:       "if set will count symlink files",
			Destination: &linecount.IncludeSymLinks,
		},
		&cli.Int64Flag{
			Name:        "large-line-count",
			Value:       40000,
			Usage:       "number of lines a file can contain before being removed from output",
			Destination: &linecount.LargeLineCount,
		},
		&cli.StringFlag{
			Name:        "count-as",
			Value:       "",
			Usage:       "count extension as language [e.g. jsp:htm,chead:\"C Header\" maps extension jsp to html and chead to C Header]",
			Destination: &linecount.CountAs,
		},
		&cli.StringFlag{
			Name:        "format-multi",
			Value:       "",
			Usage:       "have multiple format output overriding --format [e.g. tabular:stdout,csv:file.csv,json:file.json]",
			Destination: &linecount.FormatMulti,
		},
		&cli.StringFlag{
			Name:        "remap-unknown",
			Value:       "",
			Usage:       "inspect files of unknown type and remap by checking for a string and remapping the language [e.g. \"-*- C++ -*-\":\"C Header\"]",
			Destination: &linecount.RemapUnknown,
		},
		&cli.StringFlag{
			Name:        "remap-all",
			Value:       "",
			Usage:       "inspect every file and remap by checking for a string and remapping the language [e.g. \"-*- C++ -*-\":\"C Header\"]",
			Destination: &linecount.RemapAll,
		},
	},
}

func parseStringSliceFlag2StringList() {
	linecount.PathDenyList = lineCountPathDenyList.Value()
	linecount.AllowListExtensions = lineCountAllowListExtensions.Value()
	linecount.GeneratedMarkers = lineCountGeneratedMarkers.Value()
	linecount.Exclude = lineCountExclude.Value()
}
