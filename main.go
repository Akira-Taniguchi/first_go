package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
)
func main() {
  url := "https://api.zaif.jp/api/1/last_price/btc_jpy"
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray)) // {"last_price": 817700.0}と出力される
}
