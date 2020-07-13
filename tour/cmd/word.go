package cmd

import (
	"log"
	"strings"

	"github.com/jlulxy/go-programming-tour-book/tree/master/tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	// 全部大写
	MODE_UPPER = iota + 1
	// 全部转为小写
	MODE_LOWER
	// 下划线转为大写驼峰
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	// 下划线转诶小写驼峰
	MODE_UNDERSCORE_TO_LOWWER_CAMELCASE
	// 驼峰转为下划线单词
	MODE_UNDERSCORE_TO_UNDERSOCRE
)

var desc = strings.Join([]string{
	"该命令支持各种单词格式的转换，模式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为消息",
	"3：下划线单词转为大写驼峰单词",
	"4：驼峰单词转为下划线单词",
}, "\n")

var mode int8
var str string

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelcase(str)
		case MODE_UNDERSCORE_TO_UNDERSOCRE:
			content = word.CamelCaseToUndersocre(str)
		default:
			log.Fatalf("暂时不支持该模式转换，请执行help word 查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
