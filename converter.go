package main
import (
	"fmt"
	"github.com/gohouse/converter"
)
func main() {
	err := converter.NewTable2Struct().
		SavePath("E:/opt/htdocs/go/test/model/model.go").
		Dsn("debin:debin10086@tcp(qq10086.zhidaohu.com:3307)/test?charset=utf8").
		EnableJsonTag(true).
		TagKey("db").
		Table("douban_top_movie").
		RealNameMethod("TableName").
		Run()
	fmt.Println(err)
}
