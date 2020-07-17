package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jlulxy/go-programming-tour-book/tour/internal/timer"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "get now time",
	Long:  "get now time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTime string
var duration string

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "calculate time",
	Long:  "calculate time",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculatetime err: %v", err)
		}
		log.Printf("输出结果：%s, %d", calculateTime.Format(layout), calculateTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已经格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效单位"ns"，"us"，"ms"，"m"，"h"`)
}
