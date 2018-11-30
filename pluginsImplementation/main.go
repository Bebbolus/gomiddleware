package main
import (
       "os"
       "fmt"
       "plugin"
)

func main() {
    //open plugin file
    plug, err := plugin.Open("plugins/first.so")
    if err != nil {
      fmt.Println(err)
      os.Exit(-1)
    }

    //searc for an exported symbol
    symbol, err := plug.Lookup("Talk")
    if err != nil {
      fmt.Println(err)
      os.Exit(-1)
    }

    //call the function
	symbol.(func())()

}
