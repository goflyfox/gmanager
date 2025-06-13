package excel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"io"
	"net/url"
	"reflect"
	"unicode"
)

var (
	// 单元格表头
	char = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	// 默认行样式
	defaultRowStyle = &excelize.Style{Font: &excelize.Font{Color: "#666666", Size: 13, Family: "arial"}, Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"}}
)

// ExportByStructs 导出切片结构体到excel表格
func ExportByStructs(ctx context.Context, tags []string, list interface{}, fileName string, sheetName string) (err error) {
	f := excelize.NewFile()
	err = f.SetSheetName("Sheet1", sheetName)
	if err != nil {
		return
	}
	err = f.SetRowHeight(sheetName, 1, 30)
	if err != nil {
		return
	}

	rowStyleID, err := f.NewStyle(defaultRowStyle)
	if err != nil {
		return
	}
	err = f.SetSheetRow(sheetName, "A1", &tags)
	if err != nil {
		return
	}

	var (
		length    = len(tags)
		headStyle = letter(length)
		lastRow   string
		widthRow  string
	)

	for k, v := range headStyle {
		if k == length-1 {
			lastRow = fmt.Sprintf("%s1", v)
			widthRow = v
		}
	}

	if err = f.SetColWidth(sheetName, "A", widthRow, 30); err != nil {
		return err
	}

	var rowNum = 1
	for _, v := range gconv.Interfaces(list) {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		if err = f.SetSheetRow(sheetName, "A"+gconv.String(rowNum), &row); err != nil {
			return
		}
		if err = f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), lastRow, rowStyleID); err != nil {
			return
		}
	}

	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		err = gerror.New("ctx not http request")
		return
	}

	writer := r.Response.Writer
	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName)))
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	if err = f.Write(writer); err != nil {
		return
	}

	return
}

// letter 生成完整的表头
func letter(length int) []string {
	var str []string
	for i := 1; i <= length; i++ {
		str = append(str, numToChars(i))
	}
	return str
}

// numToChars 将数字转换为具体的表格表头名称
func numToChars(num int) string {
	var cols string
	v := num
	for v > 0 {
		k := v % 26
		if k == 0 {
			k = 26
		}
		v = (v - k) / 26
		cols = char[k] + cols
	}
	return cols
}

// NextLetter 传入一个字母，获取下一个字母
func NextLetter(input string) string {
	if len(input) == 0 {
		return ""
	}
	upperInput := unicode.ToUpper(rune(input[0]))
	if upperInput >= 'A' && upperInput < 'Z' {
		return string(upperInput + 1)
	}
	return "A"
}

type ImportExcelSheet struct {
	Sheet string     `json:"sheet"`
	Rows  [][]string `json:"rows"`
}

func ParseExcel(ctx context.Context, file io.Reader, sheetName string) (res *ImportExcelSheet, err error) {
	excelFile, err := excelize.OpenReader(file)
	if err != nil {
		return
	}
	defer func() {
		if err := excelFile.Close(); err != nil {
			g.Log().Error(ctx, err)
		}
	}()

	res = new(ImportExcelSheet)
	if sheetName == "" {
		sheetName = excelFile.GetSheetName(0)
	}
	sheetList := excelFile.GetSheetList()
	for _, sheet := range sheetList {
		if sheet != sheetName {
			continue
		}
		res.Sheet = sheet
		res.Rows, err = excelFile.GetRows(sheet)
		if err != nil {
			return nil, err
		}
	}
	return
}
