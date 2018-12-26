package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

// Mapping
var shogunNameMap = map[string]string{
  "1": "源頼朝",
  "2": "源頼家",
  "3": "源実朝",
  "4": "九条頼経",
  "5": "九条頼嗣",
  "6": "宗尊親王",
  "7": "惟康親王",
  "8": "久明親王",
  "9": "守邦親王",
}

var shikkenNameMap = map[string]string{
  "1": "北条時政",
  "2": "北条義時",
  "3": "北条泰時",
  "4": "北条経時",
  "5": "北条時頼",
  "6": "北条長時",
  "7": "北条政村",
  "8": "北条時宗",
  "9": "北条貞時",
  "10": "北条師時",
  "11": "北条宗宣",
  "12": "北条煕時",
  "13": "北条基時",
  "14": "北条高時",
  "15": "北条貞顕",
  "16": "北条守時",
}

func main() {
  app := cli.NewApp()

  app.Name = "go-kamakura"
  app.Usage = "Display shogun or shikken Name in Kamakura Bakufu"
  app.Version = "0.0.1"

  app.Action = func (c *cli.Context) error {
    str := c.Args().Get(0)
    if c.Bool("shikken") {
      fmt.Println(shikkenNameMap[str])
    } else {
      fmt.Println(shogunNameMap[str])
    }
    return nil
  }

  app.Flags = []cli.Flag{
    cli.BoolFlag {
      Name: "shikken, s",
      Usage: "Display shikken Name in Kamakura Bakufu",
    },
  }

  app.Run(os.Args)
}
