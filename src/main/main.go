package main

import (
	"bufio"
	"fmt"
	"github.com/goxt/dog2"
	"github.com/goxt/dog2/util"
	"os"
	"ws"
)

/**
 * 设置定时任务，写在 util.StartSchedules() 之前
 * 示例：util.NewSchedule(执行函数名).EveryMinute()
 * 周期函数：
 *		EveryMinute			每分钟
 *		EveryFiveMinutes	每五分钟
 *		EveryTenMinutes		每十分钟
 *		EveryThirtyMinutes	每半小时
 *		HourlyAt			每小时的x分
 *		DailyAt				每天的x时x分
 *		WeeklyAt			每周的星期几x时x分
 *		MonthlyAt			每月的x日x时x分
 *		YearlyAt			每年的x月x日x时x分
 * 重复和冲突：
 * 		SetRepeat(bool)		设置此定时器是否重复执行
 *								如果true，表示可以重复执行，比如1分钟执行1次函数A
 *									在下个1分钟，本该触发函数A的时候，由于上一个函数A运行量很大，依旧没有执行完毕
 *									那么，设置了可重复的话，两次函数A互不影响，各执行各的。
 *								但如果为false，表示不可以重复执行，同上情况时，第2次的函数A就不再执行
 *									除非之前的函数A在触发时间点之前已经执行完毕
 */
func mySchedules() {
	util.StartSchedules()
}

// 命令行帮助菜单
var consoleHelp = `
	命令行帮助菜单：
	`

// 命令行执行
var consoleFunc = map[string]func(){
	// 命令行映射数组 命令行名称 => 执行程序名
}

func main() {

	// 加载框架
	dog2.Load()

	var args = os.Args

	// Http服务
	if len(args) == 1 {

		// 运行定时任务
		mySchedules()

		// 启动websocket
		ws.Start()

		// 启动HTTP监听服务
		dog2.Run()
		return
	}

	// 进入命令行模式
	if args[1] == "console" {
		myConsole()
		return
	}
}

func myConsole() {
	var input = bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("当前处于主程序，请输入子程序名，进入相应子程序命令行（`h`帮助，`q`退出）")
		input.Scan()
		var key = input.Text()
		if key == "h" {
			fmt.Println(consoleHelp)
			continue
		}
		if key == "q" {
			return
		}

		// 异常捕捉
		myConsoleSub(key)
	}
}

func myConsoleSub(key string) {
	defer func() {
		if e := recover(); e != nil {
			switch v := e.(type) {
			case util.BizException:
				fmt.Println(v.Msg)
			default:
				var str = e.(error).Error()
				util.LogException("命令行执行失败：" + str)
				fmt.Println(str)
			}
		}
	}()

	// 执行脚本
	if consoleFunc[key] == nil {
		fmt.Println("没有【" + key + "】命令行程序，目前支持以下程序：")
		fmt.Println(consoleHelp)
	} else {
		consoleFunc[key]()
	}
}
