package main
import (
	"business-incqupt/api"
	"business-incqupt/dao"
)

func main(){
	api.InitEngine()
	dao.InitDB()
}


