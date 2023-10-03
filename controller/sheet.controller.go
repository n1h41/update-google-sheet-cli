package controller

import (
	"context"
	"fmt"
	"os"

	// "fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetController struct {
  s *sheets.Service
}

func (c *SheetController) Init()  {
  ctx := context.Background()

  creds, err:= os.ReadFile("/home/n1h41/.config/gcloud/application_default_credentials.json")
  if err != nil {
    panic(err)
  }

  sheetService, err := sheets.NewService(ctx, option.WithCredentialsJSON(creds))
  if err != nil {
    panic(err)
  }
  c.s = sheetService
}

func (c *SheetController) GetSheetHandler()  {
  data, err := c.s.Spreadsheets.Get("1AEmYe3wpxub9oj98hY1kV3i1mBeIG0A0zUTY-AIDlYg").Do()
  if err != nil {
    panic(err)
  }
  fmt.Println(data)
}

func NewSheetController() *SheetController {
  var sheetController SheetController
  sheetController = SheetController{}
  return &sheetController
}
