package bar

import "fmt"

//进度条
type Bar struct {
	Percent int64
	Current int64
	Total   int64
	Rate    string
	Graph   string
}

//创建func
func NewBar() *Bar {
	return &Bar{}
}

func (bar *Bar) NewOption(start, total int64) *Bar {
	bar.Total = total
	bar.Current = start
	if bar.Graph == "" {
		bar.Graph = "*"
	}
	bar.Percent = percent(bar.Current, bar.Total)
	for i := 0; i < int(bar.Percent); i++ {
		bar.Rate = bar.Graph //初始化进度条
	}
	return bar
}

//设置进度条样式
func (bar *Bar) WithGraph(graph string) *Bar {
	bar.Graph = graph
	return bar
}

func (bar *Bar) Play(current int64) {
	bar.Current = current
	last := bar.Percent
	bar.Percent = percent(bar.Current, bar.Total)
	if bar.Percent != last && bar.Percent%2 == 0 {
		bar.Rate += bar.Graph
	}
	fmt.Printf("\r [%-50s]%3d%%  %8d/%d", bar.Rate, bar.Percent, bar.Current, bar.Total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}
