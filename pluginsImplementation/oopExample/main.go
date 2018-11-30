package main
import (
       "os"
       "fmt"
       "plugin"
)

//define a local interface of what you want to get from plugin symbol
type MyPlug interface {
	Talk()
}


func main() {
    //open plugin file
    plug, err := plugin.Open("plugins/first.so")
    if err != nil {
      fmt.Println(err)
      os.Exit(-1)
    }

    //searc for an exported symbol
    symbol, err := plug.Lookup("MyPlugin")
    if err != nil {
      fmt.Println(err)
      os.Exit(-1)
    }

    // check that loaded symbol is type Controller
	var myPlugin MyPlug
	myPlugin, ok := symbol.(MyPlug)
	if !ok {
        fmt.Println("The module have wrong type")
        os.Exit(-1)
	}

    //call the function
	myPlugin.Talk()

}
