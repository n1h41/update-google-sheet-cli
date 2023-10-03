package controller

import (
	"context"
	"fmt"

	"google.golang.org/api/sheets/v4"
)

type SheetController struct {
  s sheets.Service
}

func (c *SheetController) Init()  {
  ctx := context.Background()
  s, err := sheets.NewService(ctx)
  if err != nil {
    panic(err)
  }

  fmt.Println(s.BasePath)
}

func NewSheetController() *SheetController {
  var sheetController SheetController
  sheetController = SheetController{}
  return &sheetController
}
