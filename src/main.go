package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var (
	w    fyne.Window
	ds   *DataSet
	Raw2 string

	ohead2Text *widget.Label
	ohead3Text *widget.Label

	Screen1Page    *Pagination
	txtDSName      *widget.Entry
	rtbOutput      *widget.Entry
	rtbBall        *widget.Entry
	chkTop3        *widget.Check
	chkStarter     *widget.Check
	chkConsolation *widget.Check
	top3           bool
	starter        bool
	consolation    bool

	Screen2Page       *Pagination
	cmbDataSetText    string
	cmbDataSet        *widget.Select
	cmbAlgoText       string
	cmbAlgo           *widget.Select
	cmbFilterTypeText string
	cmbFilterType     *widget.Select
	txtFilterText     *widget.Entry
	rtbOutput2        *widget.Entry
	rtbBall2          *widget.Entry

	cmbDrawText  string
	cmbDraw      *widget.Select
	cmbAlgo2     *widget.Select
	cmbAlgoText2 string
	rtbOutput3   *widget.Entry
	rtbBall3     *widget.Entry
)

func PopulateOutput(lines []string) {
	Screen1Page.Data = lines
	Screen1Page.txtText = rtbOutput
	Screen1Page.Refresh()

	mp := ds.GetBallStats(lines)
	c := &Convert{}
	str := ""
	for _, o := range mp {
		str = str + c.IntToString(o.Key) + AlgoSeparator + c.IntToString(o.Value) + "\n"
	}
	rtbBall.SetText(str)
}

func Download() {
	
}

func InfoScreen(a fyne.App) fyne.CanvasObject {

	btnCreate := widget.NewButtonWithIcon("Update", theme.ViewRefreshIcon(), func() {
		go Download()
	})

	bottom := widget.NewVBox(
		btnCreate,
	)

	//ohead := fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer()), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), widget.NewLabel("4D Predictor"))
	ohead := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewLabel("4D Predictor"), layout.NewSpacer())

	ohead2 := widget.NewLabel("Last Draw No")
	ohead3 := widget.NewLabel("Next Draw No")

	ohead2Text = widget.NewLabel("")
	ohead3Text = widget.NewLabel("")

	lastno := ds.GetLastDrawNo()
	c := &Convert{}
	ohead2Text.SetText(c.IntToString(lastno))
	ohead3Text.SetText(c.IntToString(lastno + 1))

	middle := widget.NewVBox(
		fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), ohead2, nil), layout.NewSpacer(), layout.NewSpacer(), ohead2, ohead2Text),
		fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), ohead3, nil), layout.NewSpacer(), layout.NewSpacer(), ohead3, ohead3Text),
	)

	outputheader := fyne.NewContainerWithLayout(layout.NewBorderLayout(ohead, bottom, nil, nil), ohead, bottom, middle)
	return outputheader
}

func CreateScreen(a fyne.App) fyne.CanvasObject {
	// if fyne.CurrentDevice().IsMobile() {
	// 	logo.SetMinSize(fyne.NewSize(171, 125))
	// } else {
	// 	logo.SetMinSize(fyne.NewSize(228, 167))
	// }

	txtDSName = widget.NewEntry()
	txtDSName.SetPlaceHolder("DataSet Name")

	chkTop3 = widget.NewCheck("Top3", func(on bool) { top3 = on })
	chkStarter = widget.NewCheck("Starter", func(on bool) { starter = on })
	chkConsolation = widget.NewCheck("Consolation", func(on bool) { consolation = on })
	rtbOutput = widget.NewMultiLineEntry()
	rtbBall = widget.NewMultiLineEntry()
	//rtbOutput.Resize(fyne.NewSize(rtbOutput.Size().Width, 600))

	Screen1Page = NewPagination()
	btnCreate := widget.NewButtonWithIcon("Create", theme.ContentAddIcon(), func() {
		// b.setIcon(b.current - 1)
		result := ds.CreateDataSet(top3, starter, consolation)
		PopulateOutput(result)

		if txtDSName.Text == "" {
			return
		}

		c := &Convert{}
		last := c.IntToString(ds.GetLastDrawNo())
		next := c.IntToString(ds.GetLastDrawNo() + 1)
		header := ""
		if top3 {
			header = header + "Top3|"
		}
		if starter {
			header = header + "Starter|"
		}
		if consolation {
			header = header + "Consolation|"
		}
		if strings.HasPrefix(txtDSName.Text, last+"_") == false {
			ds.WriteDataSet(next, header, result, txtDSName.Text)
			txtDSName.SetText(next + "_" + txtDSName.Text)
		} else {
			ds.WriteDataSet("", header, result, txtDSName.Text)
		}
		RefreshDataSet()
	})
	vbox := widget.NewVBox(
		txtDSName,
		widget.NewHBox(chkTop3, chkStarter, chkConsolation),
		btnCreate,
	)

	ohead := widget.NewLabel("Numbers")
	ohead2 := widget.NewLabel("Ball Statistics")
	outputheader := fyne.NewContainerWithLayout(layout.NewBorderLayout(ohead, layout.NewSpacer(), nil, nil), ohead, layout.NewSpacer(), widget.NewVScrollContainer(rtbOutput))
	outputheader2 := fyne.NewContainerWithLayout(layout.NewBorderLayout(ohead2, layout.NewSpacer(), nil, nil), ohead2, layout.NewSpacer(), widget.NewVScrollContainer(rtbBall))
	scroll := fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), outputheader, layout.NewSpacer()), outputheader, layout.NewSpacer(), outputheader2)
	bottom := Screen1Page.Container
	//vbox.Resize(fyne.NewSize(400, 300))
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(vbox, bottom, nil, nil), vbox, bottom, scroll)
}

func RefreshDataSet() {
	datasets := ds.GetDataSetName()
	cmbDataSet.Options = datasets

	lastno := ds.GetLastDrawNo()
	c := &Convert{}
	ohead2Text.SetText(c.IntToString(lastno))
	ohead3Text.SetText(c.IntToString(lastno + 1))
}

func CreateScreen2(a fyne.App) fyne.CanvasObject {

	cmbDataSet = widget.NewSelect([]string{}, func(str string) {
		cmbDataSetText = str
	})
	cmbDataSet.PlaceHolder = "Select DataSet"

	cmbAlgo = widget.NewSelect([]string{"ABCD", "Occurence", "Sum", "Identity"}, func(str string) {
		cmbAlgoText = str
	})
	cmbAlgo.PlaceHolder = "Select Algo"

	cmbFilterType = widget.NewSelect([]string{"Algo", "No", "Algo Contains", "No Contains"}, func(str string) {
		cmbFilterTypeText = str
	})
	cmbFilterType.PlaceHolder = "Optional Filter Type"

	txtFilterText = widget.NewEntry()
	txtFilterText.PlaceHolder = "Optional Filter... eg. AAB?"

	rtbOutput2 = widget.NewMultiLineEntry()
	rtbBall2 = widget.NewMultiLineEntry()

	Screen2Page = NewPagination()
	Screen2Page.txtText = rtbOutput2
	btnCreate := widget.NewButtonWithIcon("Calculate", theme.ContentAddIcon(), func() {
		if cmbDataSetText == "" {
			return
		}

		nos := ds.GetDataSetNo(cmbDataSetText)
		if nos == nil {
			return
		}

		if cmbAlgoText == "" {
			return
		}

		var result []string
		for _, no := range nos {
			algo := GetPluginsNo(cmbAlgoText, no)
			result = append(result, no+AlgoSeparator+algo)
		}
		resultall := strings.Join(result, "\n")
		Raw2 = resultall

		Screen2Page.Data = result
		Screen2Page.txtText = rtbOutput2
		Screen2Page.Refresh()

		c := &Convert{}
		var result2 []string
		kvs := ds.GetAlgoCountByRaw2(resultall)
		for _, k := range kvs {
			result2 = append(result2, k.Key+AlgoSeparator+c.IntToString(k.Value))
		}

		end := len(result2)
		if end > 100 {
			end = 100
		}
		rtbBall2.SetText(strings.Join(result2[0:end], "\n"))
	})

	btnCreate2 := widget.NewButtonWithIcon("Filter", theme.ContentAddIcon(), func() {
		if cmbFilterTypeText == "" || txtFilterText.Text == "" || Raw2 == "" {
			return
		}

		var newOutput []string
		if cmbFilterTypeText == "Algo" {
			newOutput = ds.GetNoByRaw2AlgoFilter(Raw2, txtFilterText.Text)
		} else if cmbFilterTypeText == "No" {
			newOutput = ds.GetNoByRaw2NoFilter(Raw2, txtFilterText.Text)
		} else if cmbFilterTypeText == "Algo Contains" {
			newOutput = ds.GetNoByRaw2AlgoContainsFilter(Raw2, txtFilterText.Text)
		} else if cmbFilterTypeText == "No Contains" {
			newOutput = ds.GetNoByRaw2NoContainsFilter(Raw2, txtFilterText.Text)
		} else if cmbFilterTypeText == "Statistic" {
			newOutput = ds.GetNoByRaw2StatisticFilter(Raw2, txtFilterText.Text)
		}

		var result []string
		for _, no := range newOutput {
			algo := GetPluginsNo(cmbAlgoText, no)
			result = append(result, no+AlgoSeparator+algo)
		}
		Raw2 = strings.Join(result, "\n")

		Screen2Page.Data = result
		Screen2Page.txtText = rtbOutput2
		Screen2Page.Refresh()

		c := &Convert{}
		var result2 []string
		kvs := ds.GetAlgoCountByRaw2(Raw2)
		for _, k := range kvs {
			result2 = append(result2, k.Key+AlgoSeparator+c.IntToString(k.Value))
		}

		end := len(result2)
		if end > 100 {
			end = 100
		}
		rtbBall2.SetText(strings.Join(result2[0:end], "\n"))
	})

	btnDelete := widget.NewButton("Delete", func() {
		file := ds.GetDataSetPath(cmbDataSetText)
		f := &File2{}
		f.Delete(file + ".txt")
		RefreshDataSet()
		cmbDataSet.SetSelected("")
		cmbDataSet.Refresh()
	})
	dataset := fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), cmbDataSet, nil), layout.NewSpacer(), layout.NewSpacer(), cmbDataSet, btnDelete)
	algo := fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), cmbAlgo, nil), layout.NewSpacer(), layout.NewSpacer(), cmbAlgo, btnCreate)
	vbox := widget.NewVBox(
		dataset,
		algo,
		cmbFilterType,
		txtFilterText,
		btnCreate2,
	)
	ohead := widget.NewLabel("Numbers")
	ohead2 := widget.NewLabel("Statistics")
	outputheader := fyne.NewContainerWithLayout(layout.NewBorderLayout(ohead, layout.NewSpacer(), nil, nil), ohead, layout.NewSpacer(), widget.NewVScrollContainer(rtbOutput2))
	outputheader2 := fyne.NewContainerWithLayout(layout.NewBorderLayout(ohead2, layout.NewSpacer(), nil, nil), ohead2, layout.NewSpacer(), widget.NewVScrollContainer(rtbBall2))
	scroll := fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), outputheader, layout.NewSpacer()), outputheader, layout.NewSpacer(), outputheader2)
	bottom := Screen2Page.Container
	//vbox.Resize(fyne.NewSize(400, 300))
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(vbox, bottom, nil, nil), vbox, bottom, scroll)
}

func CreateScreen3(a fyne.App) fyne.CanvasObject {
	c := &Convert{}
	lastno := ds.GetLastDrawNo()
	var draw []string
	for i := ds.GetLastDrawNo(); i >= lastno-100; i-- {
		draw = append(draw, c.IntToString(i))
	}

	cmbDraw = widget.NewSelect(draw, func(str string) {
		cmbDrawText = str
	})
	cmbDraw.PlaceHolder = "Select Draw"

	cmbAlgo2 = widget.NewSelect([]string{"ABCD", "Occurence", "Sum", "Identity"}, func(str string) {
		cmbAlgoText2 = str
	})
	cmbAlgo2.PlaceHolder = "Select Algo"

	rtbOutput3 = widget.NewMultiLineEntry()
	rtbBall3 = widget.NewMultiLineEntry()

	btnCreate := widget.NewButtonWithIcon("Calculate", theme.ContentAddIcon(), func() {
		if cmbDrawText == "" {
			return
		}

		nos := ds.GetNoByDrawNo(c.ToInt32(cmbDrawText))
		if nos == nil {
			return
		}

		if cmbAlgoText2 == "" {
			return
		}

		var result []string
		for _, no := range nos {
			algo := GetPluginsNo(cmbAlgoText2, no)
			result = append(result, no+AlgoSeparator+algo)
		}
		resultall := strings.Join(result, "\n")
		s := &Strings2{}

		resultx := s.InsertAt(result, "Top3", 0)
		resultx = s.InsertAt(resultx, "Starter", 4)
		resultx = s.InsertAt(resultx, "Consolation", 15)
		rtbOutput3.SetText(strings.Join(resultx, "\n"))

		c := &Convert{}
		var result2 []string
		kvs := ds.GetAlgoCountByRaw2(resultall)
		for _, k := range kvs {
			result2 = append(result2, k.Key+AlgoSeparator+c.IntToString(k.Value))
		}

		end := len(result2)
		if end > 100 {
			end = 100
		}
		rtbBall3.SetText(strings.Join(result2[0:end], "\n"))
	})

	vbox := widget.NewVBox(
		cmbDraw,
		cmbAlgo2,
		btnCreate,
	)
	ohead := widget.NewLabel("Numbers")
	ohead2 := widget.NewLabel("Statistics")
	outputheader := fyne.NewContainerWithLayout(layout.NewBorderLayout(ohead, layout.NewSpacer(), nil, nil), ohead, layout.NewSpacer(), widget.NewVScrollContainer(rtbOutput3))
	outputheader2 := fyne.NewContainerWithLayout(layout.NewBorderLayout(ohead2, layout.NewSpacer(), nil, nil), ohead2, layout.NewSpacer(), widget.NewVScrollContainer(rtbBall3))
	scroll := fyne.NewContainerWithLayout(layout.NewBorderLayout(layout.NewSpacer(), layout.NewSpacer(), outputheader, layout.NewSpacer()), outputheader, layout.NewSpacer(), outputheader2)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(vbox, layout.NewSpacer(), nil, nil), vbox, layout.NewSpacer(), scroll)
}

func main() {
	ds = &DataSet{}
	a := app.NewWithID("com.johnkenedy.4DAnalysis")
	w = a.NewWindow("4D Predictor")

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Info", theme.HomeIcon(), InfoScreen(a)),
		widget.NewTabItemWithIcon("DataSet", theme.ContentAddIcon(), CreateScreen(a)),
		widget.NewTabItemWithIcon("Analysis", theme.SearchIcon(), CreateScreen2(a)),
		widget.NewTabItemWithIcon("View", theme.InfoIcon(), CreateScreen3(a)),
	)

	tabs.SetTabLocation(widget.TabLocationLeading)
	// tabs.SelectTabIndex(a.Preferences().Int())
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(400, 500))
	RefreshDataSet()
	w.ShowAndRun()

	// a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
