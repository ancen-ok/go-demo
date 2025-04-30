package utils

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
)

var statics = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

type ExcelField struct {
	Index      int    // 下标索引
	Field      string // 属性名称
	Title      string // 名称
	LineHeight int    // 行高
}

// IExcelItem Excel写数据实体需要实现的方法
type IExcelItem interface {
	Fields() (fields []*ExcelField, err error) // 表头
}

type IExcelTarget interface {
	To() (map[string]any, error)
}

type ExcelFactory[T IExcelItem, E IExcelTarget] struct {
	Name   string         // sheet名称
	Class  T              // excel列对应的结构体
	Items  []E            // 将要写入的结构体
	fields []*ExcelField  // 属性
	excel  *excelize.File // excel操作句柄
}

// NewFactory 构建
func NewFactory[T IExcelItem, E IExcelTarget](name string, target T) *ExcelFactory[T, E] {
	return &ExcelFactory[T, E]{
		Name:  name,
		Class: target,
	}
}
func (f *ExcelFactory[T, E]) Header() (err error) {
	f.excel = excelize.NewFile()
	defer func() {
		if e := recover(); e != nil {
			_ = f.excel.Close()
		}
	}()
	// 将Sheet1改为自己想要的名字
	_ = f.excel.SetSheetName("Sheet1", f.Name)
	if f.fields, err = f.Class.Fields(); err != nil {
		return
	}
	for _, field := range f.fields {
		_ = f.excel.SetCellValue(f.Name, fmt.Sprintf("%s1", statics[field.Index]), field.Title)
	}
	return
}

func (f *ExcelFactory[T, E]) Content(items []E) error {
	var (
		item     E
		index    = 0
		rowIndex = 2
		data     map[string]any
		err      error
	)
	defer func() {
		if e := recover(); e != nil {
			_ = f.excel.Close()
		}
	}()
	for index < len(items) {
		item = items[index]
		if data, err = item.To(); err != nil {
			index++
			continue
		}
		for _, field := range f.fields {
			_ = f.excel.SetCellValue(f.Name, fmt.Sprintf("%s%d", statics[field.Index], rowIndex), data[field.Field])
		}
		rowIndex++
		index++
	}
	return nil
}

func (f *ExcelFactory[T, E]) Build() (content *bytes.Reader, err error) {
	var buffer bytes.Buffer
	defer func() {
		_ = f.excel.Close()
	}()
	if err = f.excel.Write(&buffer); err != nil {
		return
	}
	content = bytes.NewReader(buffer.Bytes())
	return
}
