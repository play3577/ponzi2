package ponzi

import (
	"fmt"

	"golang.org/x/image/font/gofont/goregular"

	"github.com/btmura/ponzi2/internal/gfx"
)

const (
	chartRounding = 10
	chartPadding  = 5
)

var (
	chartSymbolQuoteTextRenderer = gfx.NewTextRenderer(goregular.TTF, 24)
	chartFormatQuote             = func(st *ModelStock) string {
		if st.Price() != 0 {
			return fmt.Sprintf("%.2f %+5.2f %+5.2f%%", st.Price(), st.Change(), st.PercentChange()*100.0)
		}
		return ""
	}
	chartLoadingText = NewCenteredText(chartSymbolQuoteTextRenderer, "LOADING...")
)

// Shared variables used by multiple chart components.
var (
	chartAxisLabelTextRenderer = gfx.NewTextRenderer(goregular.TTF, 12)
	chartGridHorizLine         = gfx.HorizColoredLineVAO(gray, gray)
)

// Chart shows a stock chart for a single stock.
type Chart struct {
	// header renders the header with the symbol, quote, and buttons.
	header *ChartHeader

	// lines renders the vertical weekly lines.
	lines *ChartLines

	// prices renders the candlesticks.
	prices *ChartPrices

	// avgLines renders the moving average lines.
	avgLines *ChartAvgLines

	// volume renders the volume bars.
	volume *ChartVolume

	// dailyStochastics renders the daily stochastics.
	dailyStochastics *ChartStochastics

	// weeklyStochastics renders the weekly stochastics.
	weeklyStochastics *ChartStochastics

	// loading is true when data for the stock is being retrieved.
	loading bool

	// hasStockUpdated is true if the stock has reported data.
	hasStockUpdated bool
}

// NewChart creates a new chart.
func NewChart() *Chart {
	return &Chart{
		header: NewChartHeader(&ChartHeaderArgs{
			SymbolQuoteTextRenderer: chartSymbolQuoteTextRenderer,
			QuoteFormatter:          chartFormatQuote,
			ShowRefreshButton:       true,
			ShowAddButton:           true,
			Rounding:                chartRounding,
			Padding:                 chartPadding,
		}),
		lines:             &ChartLines{},
		prices:            &ChartPrices{},
		avgLines:          NewChartAvgLines(),
		volume:            &ChartVolume{},
		dailyStochastics:  &ChartStochastics{Interval: DailyInterval},
		weeklyStochastics: &ChartStochastics{Interval: WeeklyInterval},
		loading:           true,
	}
}

// SetLoading sets the Chart's loading state.
func (ch *Chart) SetLoading(loading bool) {
	ch.loading = loading
	ch.header.SetLoading(loading)
}

// SetStock sets the Chart's stock.
func (ch *Chart) SetStock(st *ModelStock) {
	ch.hasStockUpdated = !st.LastUpdateTime.IsZero()
	ch.header.SetStock(st)
	ch.lines.SetStock(st)
	ch.prices.SetStock(st)
	ch.avgLines.SetStock(st)
	ch.volume.SetStock(st)
	ch.dailyStochastics.SetStock(st)
	ch.weeklyStochastics.SetStock(st)
}

// Update updates the Chart.
func (ch *Chart) Update() {
	ch.header.Update()
}

// Render renders the Chart.
func (ch *Chart) Render(vc ViewContext) {
	// Render the border around the chart.
	renderRoundedRect(vc.Bounds, chartRounding)

	// Render the header and the line below it.
	r, _ := ch.header.Render(vc)
	rects := sliceRect(r, 0.13, 0.13, 0.13, 0.61)
	renderHorizDivider(rects[3], horizLine)

	if ch.loading && !ch.hasStockUpdated {
		chartLoadingText.Render(r)
		return
	}

	for i := 0; i < 3; i++ {
		renderHorizDivider(rects[i], horizLine)
	}

	pr, vr, dr, wr := rects[3], rects[2], rects[1], rects[0]

	pr = pr.Inset(chartPadding)
	vr = vr.Inset(chartPadding)
	dr = dr.Inset(chartPadding)
	wr = wr.Inset(chartPadding)

	maxWidth := ch.prices.RenderLabels(pr)
	if w := ch.volume.RenderLabels(vr); w > maxWidth {
		maxWidth = w
	}
	if w := ch.dailyStochastics.RenderLabels(dr); w > maxWidth {
		maxWidth = w
	}
	if w := ch.weeklyStochastics.RenderLabels(wr); w > maxWidth {
		maxWidth = w
	}

	pr.Max.X -= maxWidth + chartPadding
	vr.Max.X -= maxWidth + chartPadding
	dr.Max.X -= maxWidth + chartPadding
	wr.Max.X -= maxWidth + chartPadding

	ch.lines.Render(pr)
	ch.lines.Render(vr)
	ch.lines.Render(dr)
	ch.lines.Render(wr)

	ch.prices.Render(pr)
	ch.avgLines.Render(pr)
	ch.volume.Render(vr)
	ch.dailyStochastics.Render(dr)
	ch.weeklyStochastics.Render(wr)
}

// SetRefreshButtonClickCallback sets the callback for refresh button clicks.
func (ch *Chart) SetRefreshButtonClickCallback(cb func()) {
	ch.header.SetRefreshButtonClickCallback(cb)
}

// SetAddButtonClickCallback sets the callback for add button clicks.
func (ch *Chart) SetAddButtonClickCallback(cb func()) {
	ch.header.SetAddButtonClickCallback(cb)
}

// Close frees the resources backing the chart.
func (ch *Chart) Close() {
	if ch.header != nil {
		ch.header.Close()
		ch.header = nil
	}
	if ch.lines != nil {
		ch.lines.Close()
		ch.lines = nil
	}
	if ch.prices != nil {
		ch.prices.Close()
		ch.prices = nil
	}
	if ch.volume != nil {
		ch.volume.Close()
		ch.volume = nil
	}
	if ch.dailyStochastics != nil {
		ch.dailyStochastics.Close()
		ch.dailyStochastics = nil
	}
	if ch.weeklyStochastics != nil {
		ch.weeklyStochastics.Close()
		ch.weeklyStochastics = nil
	}
}
