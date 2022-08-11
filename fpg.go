package main

import (
  "fmt"
  "flag"
  "os"
  "os/exec"
)

func main() {
  flag.Parse()
  arg1 := flag.Arg(0)
  arg2 := flag.Arg(1)

  switch arg1 {
    case "init":
      if arg2 == "" {
        fmt.Println("flask init <PROJECT_NAME> を実行してください。")
      } else {
        flaskInit(arg2)
      }

    case "help":
      fmt.Println("--- flask init がうまく動作しない場合 ---")
      fmt.Println("Python3, Python3-pip, Python3-venv のどれかに問題がある可能性があります。")
      fmt.Println("もう一度ご確認の上お試しください。")

    default: 
      fmt.Println("コマンド引数を指定してください。")
      fmt.Println("---")
      fmt.Println("flask init ・・・ プロジェクトを作成")
      fmt.Println("flask help ・・・ flask コマンドについてのヒント")
      fmt.Println("---")
  }
}
func flaskInit(str string) {
  fmt.Println("プロジェクト", str, "の作成を開始します。")
  out, err := exec.Command("sh", "./init.sh").Output()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    fmt.Println(string(out))
    fmt.Println("mv newProject", str, "で名前を変更しましょう。")
}
