package tools

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

// InteractiveAction 交互式命令
func InteractiveAction(c *cli.Context) (err error) {
	action := ""
	cateActions := map[string][]string{}
	cates := []string{}
	cateCount := 0
	newCateCount := 0

	for _, cmd := range c.App.Commands {
		cate := ""
		if cmd.Category == "" {
			cate = "主要"
		} else {
			cate = cmd.Category
		}
		cateActions[cate] = append(cateActions[cmd.Category], cmd.Usage)
		newCateCount = len(cateActions)
		if newCateCount > cateCount {
			cateCount = newCateCount
			cates = append(cates, cate)
		}
	}
	for cate := range cateActions {
		cates = append(cates, cate)
	}
	prompt := &survey.Select{
		Message:  "你想进行什么操作:",
		Options:  cates,
		PageSize: 10,
	}

	survey.AskOne(prompt, &action)
	prompt = &survey.Select{
		Message:  "你想进行什么操作:",
		Options:  cateActions[action],
		PageSize: 10,
	}
	survey.AskOne(prompt, &action)

	return err

}

func InteractiveL1Action(c *cli.Context) (err error) {
	return err
}
