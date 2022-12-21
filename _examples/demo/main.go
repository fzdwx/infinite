package main

import (
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/confirm"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/components/spinner"
	"github.com/fzdwx/infinite/style"
)

func main() {

	options := []string{
		"周杰伦-稻香",
		"周杰伦-晴天",
		"海明威-老人与海",
		"许巍-我相风一样自由",
		"伍佰-白鸽",
		"伍佰-枫叶",
		"陈奕迅-最佳损友",
		"陈奕迅-富士山下",
		"陈奕迅-淘汰",
		"陈奕迅-一丝不挂",
		"新裤子乐队-花火",
		"新裤子乐队-你要跳舞吗",
		"新裤子乐队-生活因你而火热",
		"新裤子乐队-没有理想的人不伤心",
		"哪吒乐队-闹海",
	}

	inf.NewSpinner(
		spinner.WithPrompt(" Loading..."),
		spinner.WithDisableOutputResult(),
	).Display(func(spinner *spinner.Spinner) {
		time.Sleep(time.Millisecond * 100 * 12)
		spinner.Info("共找到 %d 首歌", len(options))
	})

	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)

	selected, _ := inf.NewMultiSelect(options,
		multiselect.WithFilterInput(input),
	).Display("请选择你要下载的歌曲")

	yes, _ := inf.NewConfirmWithSelection(
		confirm.WithPrompt(fmt.Sprintf("请问你是否要下载这 %d 首歌", len(selected))),
	).Display()

	if !yes {
		return
	}

	inf.NewProgressGroup(len(selected)).
		AppendRunner(func(progress *components.Progress) func() {
			title := strutil.After(options[selected[progress.Id-1]], "-")
			total := random.RandInt(10, 20)
			progress.WithTotal(int64(total))
			progress.WithDefaultGradient()
			progress.WithWidth(80)
			progress.WithTitleView(func(done bool) string {
				if done {
					return fmt.Sprintf("下载 %s 成功", title)
				}
				return fmt.Sprintf("下载 %s ...", title)
			})
			return func() {
				for i := 0; i < total+1; i++ {
					progress.IncrOne()
					sleep()
				}
			}
		}).Display()

}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
