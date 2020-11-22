package components

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"time"
)

var (
	defaultTmpDir   = ".go_echarts_tmp"
	defaultLoadTime = 2 * time.Second
)

type ScreenshotType int

const (
	PNG ScreenshotType = iota
	JPEG
)

type ScreenshotOption struct {
	IDs         []string
	Type        ScreenshotType
	TmpDir      string
	LoadTime    time.Duration
	DeviceScale float64
	Mobile      bool
}

// ScreenshotAll screenshot the charts you need
func (page *Page) Screenshot(ctx context.Context, opt ScreenshotOption) ([]Image, error) {
	tmpDir, err := createTmpDir(opt.TmpDir)
	if err != nil {
		return nil, err
	}
	opt.TmpDir = tmpDir

	filename := fmt.Sprintf("%s/go-echarts-%d.html", opt.TmpDir, time.Now().UnixNano())
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer os.Remove(filename)

	err = page.Render(file)
	if err != nil {
		return nil, err
	}

	images := make([]Image, 0, len(opt.IDs))
	for _, id := range opt.IDs {
		buf := make([]byte, 0)
		images = append(images, Image{
			ID:   id,
			Type: opt.Type,
			Buff: &buf,
		})
	}

	tasks, err := elementsScreenshot(filename, images, opt.LoadTime, opt.DeviceScale, opt.Mobile)
	if err != nil {
		return nil, err
	}

	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	err = chromedp.Run(ctx, tasks)
	defer chromedp.Stop()
	if err != nil {
		return nil, err
	}

	return images, nil
}

// ScreenshotAll screenshot all charts in page
func (page *Page) ScreenshotAll(ctx context.Context, opt ScreenshotOption) ([]Image, error) {
	tmpDir, err := createTmpDir(opt.TmpDir)
	if err != nil {
		return nil, err
	}
	opt.TmpDir = tmpDir

	filename := fmt.Sprintf("%s/go-echarts-%d.html", tmpDir, time.Now().UnixNano())
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer os.Remove(filename)

	err = page.Render(file)
	if err != nil {
		return nil, err
	}

	images := make([]Image, 0, len(page.Charts))
	for _, chart := range page.Charts {
		id := chart.ID()
		buf := make([]byte, 0)
		images = append(images, Image{
			ID:   id,
			Type: opt.Type,
			Buff: &buf,
		})
	}

	tasks, err := elementsScreenshot(filename, images, opt.LoadTime, opt.DeviceScale, opt.Mobile)
	if err != nil {
		return nil, err
	}

	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	err = chromedp.Run(ctx, tasks)
	defer chromedp.Stop()
	if err != nil {
		return nil, err
	}

	return images, nil
}

type Image struct {
	ID   string
	Type ScreenshotType
	Buff *[]byte
}

func (img *Image) WriteToFile(filename string) error {
	switch img.Type {
	case JPEG:
		filename = filename + ".jpeg"
	default:
		filename = filename + ".png"
	}
	return ioutil.WriteFile(filename, *img.Buff, 0644)
}

func elementsScreenshot(url string, images []Image, wait time.Duration, scale float64, mobile bool) (chromedp.Tasks, error) {
	//get absolute path
	path, err := filepath.Abs(url)
	if err != nil {
		return nil, err
	}

	tasks := []chromedp.Action{
		chromedp.Navigate("file://" + path),
		chromedp.Sleep(wait),
	}

	for i := range images {
		img := images[i]
		sel := fmt.Sprintf("#%s", img.ID)
		tasks = append(tasks, chromedp.QueryAfter(sel, func(ctx context.Context, nodes ...*cdp.Node) error {
			if len(nodes) < 1 {
				return fmt.Errorf("selector %q did not return any nodes", sel)
			}

			if err := deviceMetricsOverride(ctx, scale, mobile); err != nil {
				return err
			}

			return screenshot(ctx, img.Type, scale, img.Buff, nodes...)
		}, []chromedp.QueryOption{chromedp.ByID, chromedp.NodeVisible}...))
	}
	return tasks, nil
}

func screenshot(ctx context.Context, sType ScreenshotType, scale float64, res *[]byte, nodes ...*cdp.Node) error {
	// get box model
	box, err := dom.GetBoxModel().WithNodeID(nodes[0].NodeID).Do(ctx)
	if err != nil {
		return err
	}
	if len(box.Margin) != 8 {
		return chromedp.ErrInvalidBoxModel
	}

	format := page.CaptureScreenshotFormatPng
	if sType == JPEG {
		format = page.CaptureScreenshotFormatJpeg
	}

	// take screenshot of the box
	buf, err := page.CaptureScreenshot().
		WithFormat(format).
		WithClip(&page.Viewport{
			// Round the dimensions, as otherwise we might lose one pixel in either dimension.
			X:      math.Round(box.Margin[0]),
			Y:      math.Round(box.Margin[1]),
			Width:  math.Round(box.Margin[4] - box.Margin[0]),
			Height: math.Round(box.Margin[5] - box.Margin[1]),
			Scale:  scale,
		}).Do(ctx)
	if err != nil {
		return err
	}

	*res = buf
	return nil
}

func deviceMetricsOverride(ctx context.Context, scale float64, mobile bool) error {
	// get layout metrics
	_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
	if err != nil {
		return err
	}

	width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

	// force viewport emulation
	err = emulation.SetDeviceMetricsOverride(width, height, scale, mobile).
		WithScreenOrientation(&emulation.ScreenOrientation{
			Type:  emulation.OrientationTypePortraitPrimary,
			Angle: 0,
		}).
		Do(ctx)
	return err
}

func createTmpDir(tmpDir string) (string, error) {
	if tmpDir == "" {
		tmpDir = defaultTmpDir
	}

	// create tmp dir if not exist
	if _, err := os.Stat(tmpDir); err != nil {
		err = os.Mkdir(tmpDir, 0755)
		if err != nil {
			return "", err
		}
	}
	return tmpDir, nil
}
