package ponzi

import (
	"image"
	"math"
	"strconv"

	"github.com/go-gl/gl/v4.5-core/gl"
)

// Colors used by the chart.
var (
	green  = [3]float32{0.25, 1, 0}
	red    = [3]float32{1, 0.3, 0}
	yellow = [3]float32{1, 1, 0}
	purple = [3]float32{0.5, 0, 1}
	white  = [3]float32{1, 1, 1}
)

type chart struct {
	stock             *modelStock
	volume            *chartVolume
	dailyStochastics  *chartStochastics
	weeklyStochastics *chartStochastics

	symbolQuoteText *dynamicText
	labelText       *dynamicText
	frameBorder     *vao
	frameDivider    *vao
	stickLines      *vao
	stickRects      *vao

	minPrice float32
	maxPrice float32
}

func createChart(stock *modelStock, symbolQuoteText, labelText *dynamicText) *chart {
	return &chart{
		stock:             stock,
		volume:            createChartVolume(stock, labelText),
		dailyStochastics:  createChartStochastics(stock, dailyInterval, labelText),
		weeklyStochastics: createChartStochastics(stock, weeklyInterval, labelText),

		symbolQuoteText: symbolQuoteText,
		labelText:       labelText,
		frameBorder:     createStrokedRectVAO(white, white, white, white),
		frameDivider:    createLineVAO(white, white),
	}
}

func (ch *chart) update() {
	if ch == nil {
		return
	}
	ch.updatePrices()
	ch.volume.update()
	ch.dailyStochastics.update()
	ch.weeklyStochastics.update()
}

func (ch *chart) updatePrices() {
	if ch.stock.dailySessions == nil {
		return
	}

	if ch.stickLines != nil {
		return
	}

	// Find the min and max prices for the y-axis.
	ch.minPrice, ch.maxPrice = math.MaxFloat32, 0
	for _, s := range ch.stock.dailySessions {
		if ch.minPrice > s.low {
			ch.minPrice = s.low
		}
		if ch.maxPrice < s.high {
			ch.maxPrice = s.high
		}
	}

	// Calculate vertices and indices for the candlesticks.
	var vertices []float32
	var colors []float32
	var lineIndices []uint16
	var triangleIndices []uint16

	stickWidth := 2.0 / float32(len(ch.stock.dailySessions)) // (-1 to 1) on X-axis
	leftX := -1.0 + stickWidth*0.1
	midX := -1.0 + stickWidth*0.5
	rightX := -1.0 + stickWidth*0.9

	calcY := func(value float32) float32 {
		return 2*(value-ch.minPrice)/(ch.maxPrice-ch.minPrice) - 1
	}

	for i, s := range ch.stock.dailySessions {
		// Figure out Y coordinates of the key levels.
		lowY, highY, openY, closeY := calcY(s.low), calcY(s.high), calcY(s.open), calcY(s.close)

		// Figure out the top and bottom of the candlestick.
		topY, botY := openY, closeY
		if openY < closeY {
			topY, botY = closeY, openY
		}

		// Add the vertices needed to create the candlestick.
		vertices = append(vertices,
			midX, highY, // 0
			midX, topY, // 1
			midX, lowY, // 2
			midX, botY, // 3
			leftX, topY, // 4 - Upper left of box
			rightX, topY, // 5 - Upper right of box
			leftX, botY, // 6 - Bottom left of box
			rightX, botY, // 7 - Bottom right of box
		)

		// Add the colors corresponding to the vertices.
		var c [3]float32
		switch {
		case s.close > s.open:
			c = green
		case s.close < s.open:
			c = red
		default:
			c = yellow
		}

		colors = append(colors,
			c[0], c[1], c[2],
			c[0], c[1], c[2],
			c[0], c[1], c[2],
			c[0], c[1], c[2],
			c[0], c[1], c[2],
			c[0], c[1], c[2],
			c[0], c[1], c[2],
			c[0], c[1], c[2],
		)

		// idx is function to refer to the vertices above.
		idx := func(j uint16) uint16 {
			return uint16(i)*8 + j
		}

		// Add the vertex indices to render the candlestick.
		lineIndices = append(lineIndices,
			// Top and bottom lines around the box.
			idx(0), idx(1),
			idx(2), idx(3),
		)

		if s.close > s.open {
			// Use lines for open candlestick on higher closes.
			lineIndices = append(lineIndices,
				idx(4), idx(5),
				idx(6), idx(7),
				idx(4), idx(6),
				idx(5), idx(7),
			)
		} else {
			// Use triangles for filled candlestick on lower closes.
			triangleIndices = append(triangleIndices,
				idx(4), idx(6), idx(5),
				idx(5), idx(6), idx(7),
			)
		}

		// Move the X coordinates one stick over.
		leftX += stickWidth
		midX += stickWidth
		rightX += stickWidth
	}

	ch.stickLines = createVAO(gl.LINES, vertices, colors, lineIndices)
	ch.stickRects = createVAO(gl.TRIANGLES, vertices, colors, triangleIndices)
}

func (ch *chart) render(r image.Rectangle) {
	if ch == nil {
		return
	}
	const pad = 3
	subRects := ch.renderFrame(r)
	ch.renderPrices(subRects[3].Inset(pad))
	ch.volume.render(subRects[2].Inset(pad))
	ch.dailyStochastics.render(subRects[1].Inset(pad))
	ch.weeklyStochastics.render(subRects[0].Inset(pad))
}

func (ch *chart) renderFrame(r image.Rectangle) []image.Rectangle {
	if ch == nil {
		return nil
	}

	// Start rendering from the top left. Track position with point.
	pt := image.Pt(r.Min.X, r.Max.Y)

	//
	// Render the frame around the chart.
	//

	gl.Uniform1f(colorMixAmountLocation, 1)
	setModelMatrixRectangle(r)
	ch.frameBorder.render()

	//
	// Render the symbol and its quote.
	//

	const pad = 10
	s := ch.symbolQuoteText.measure(ch.stock.symbol)
	pt.Y -= pad + s.Y
	{
		c := pt
		c.X += pad
		c = c.Add(ch.symbolQuoteText.render(ch.stock.symbol, c))
		c = c.Add(ch.symbolQuoteText.render(formatQuote(ch.stock.quote), c))
	}
	pt.Y -= pad

	//
	// Render the dividers between the sections.
	//

	r.Max.Y = pt.Y
	gl.Uniform1f(colorMixAmountLocation, 1)

	rects := sliceRectangle(r, 0.13, 0.13, 0.13, 0.6)
	for _, r := range rects {
		setModelMatrixRectangle(image.Rect(r.Min.X, r.Max.Y, r.Max.X, r.Max.Y))
		ch.frameDivider.render()
	}
	return rects
}

func (ch *chart) renderPrices(r image.Rectangle) {
	if ch == nil {
		return
	}

	if ch.stock.dailySessions != nil {
		makeLabel := func(v float32) string {
			return strconv.FormatFloat(float64(v), 'f', 2, 32)
		}

		labelSize := ch.labelText.measure(makeLabel(ch.maxPrice))
		labelPaddingX, labelPaddingY := 4, labelSize.Y/2
		pricePerPixel := (ch.maxPrice - ch.minPrice) / float32(r.Dy())

		// Start at top and decrement one label with top and bottom padding.
		c := r.Max
		dc := image.Pt(0, labelPaddingY+labelSize.Y+labelPaddingY)

		// Start at top with max price and decrement change in price of a label with padding.
		v := ch.maxPrice
		dv := pricePerPixel * float32(dc.Y)

		// Offets to the cursor and price value when drawing.
		dcy := labelPaddingY + labelSize.Y   // Puts cursor at the baseline of the text.
		dvy := labelPaddingY + labelSize.Y/2 // Uses value in the middle of the label.

		for {
			{
				c := image.Pt(c.X, c.Y-dcy)
				if c.Y < r.Min.Y {
					break
				}

				v := v - pricePerPixel*float32(dvy)
				l := makeLabel(v)
				s := ch.labelText.measure(l)
				c.X -= s.X + labelPaddingX
				ch.labelText.render(l, c)
			}

			c = c.Sub(dc)
			v -= dv
		}

		r.Max.X -= labelSize.X + labelPaddingX*2
	}

	gl.Uniform1f(colorMixAmountLocation, 1)
	setModelMatrixRectangle(r)
	ch.stickLines.render()
	ch.stickRects.render()
}

func (ch *chart) close() {
	if ch == nil {
		return
	}
	ch.volume.close()
	ch.dailyStochastics.close()
	ch.weeklyStochastics.close()

	ch.frameDivider.close()
	ch.frameBorder.close()
	ch.stickLines.close()
	ch.stickRects.close()
}
