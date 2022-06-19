package model

import (
	"fmt"
	conf "nav-site-server/config"
	"testing"
)

func TestUpdateOrder(t *testing.T) {
	abc := new(WebsitesModel)

	s, _ := conf.InitConfig()
	conf.App = *s
	app := conf.App
	_ = abc.InitOrder(app.Store.FileSync)
	fmt.Println("==========")
}

func TestFileOrder(t *testing.T) {
	webSite := new(WebsitesModel)

	conf, _ := conf.InitConfig()

	list, _ := webSite.List(conf.Store.FileSync)
	//fmt.Println(list)

	mapList := make(map[string][]WebsitesStoreItem)
	for _, item := range list {
		if item.Group == "" {
			item.Group = WebsitesGroupDefault
		}
		if _, ok := mapList[item.Group]; ok == false {
			mapList[item.Group] = make([]WebsitesStoreItem, 0)
		}
		mapList[item.Group] = append(mapList[item.Group], item)
	}

	//fmt.Println(mapList)

	// 排序
	for mapKey, mapValList := range mapList {
		fmt.Println(mapKey)
		//fmt.Println(mapKey, mapValList)

		//for index, mapArr := range mapValList {
		//	fmt.Println(index, mapArr)
		//}

		sortList := Sort("order", mapValList)
		//sortList := util.SortBodyByField(mapValList, "order")
		fmt.Println(sortList)
	}
}
