package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
)

var str string
var mode int8

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转换为大写",
	"2：全部单词转换为小写",
}, "\n")


var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = ToUpper(str)
		case ModeLower:
			content = ToLower(str)
		default:
			fmt.Printf("暂不支持该模式，请执行help word查看帮助文档")
			os.Exit(-1)
		}

		fmt.Printf("输出结果为：%s\n",content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}
