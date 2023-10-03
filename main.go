package main

import "n1h41/update-google-sheet-cli/controller"


func main() {
  // features.GetCommitMessages("/mnt/d/nihal/Development/Flutter/works/raf-pharmacy/", true)

  var sheetController *controller.SheetController
  sheetController = controller.NewSheetController()

  sheetController.Init()

  sheetController.GetSheetHandler() 
}
