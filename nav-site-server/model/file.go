package model

import (
	"encoding/json"
	"fmt"
	conf "nav-site-server/config"
	"nav-site-server/extend/util"
	"strconv"
	"strings"
	"time"
)

type WebsitesModel struct {
	Model
}

// WebsitesStoreItem 存储的站点信息
type WebsitesStoreItem struct {
	ID     string `json:"id"`     // 网站ID
	Group  string `json:"group"`  // 网站分组
	Order  int    `json:"order"`  // 网站分组
	Name   string `json:"name"`   // 网站名称
	Pic    string `json:"pic"`    // 网站图标
	Host   string `json:"host"`   // 网站地址
	Desc   string `json:"desc"`   // 网站描述
	Create string `json:"create"` // 添加时间
	Update string `json:"update"` // 更新时间
}

const (
	// WebsitesGroupDefault 默认分组名称
	WebsitesGroupDefault = "default"
)

// List 获取站点列表
func (w *WebsitesModel) List(fileSync *util.FileSync) ([]WebsitesStoreItem, error) {
	content, err := fileSync.ReadJSON()
	if err != nil {
		return nil, err
	}
	if content == nil {
		return nil, nil
	}
	list := make([]WebsitesStoreItem, 0)
	if err = json.Unmarshal(content, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// List 获取所有分组信息
func (w *WebsitesModel) ListGroupOrder(fileSync *util.FileSync) ([]string, error) {
	content, err := fileSync.ReadJSON()
	if err != nil {
		return nil, err
	}
	if content == nil {
		return nil, nil
	}
	list := make([]string, 0)
	if err = json.Unmarshal(content, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// Add 添加站点
func (w *WebsitesModel) Add(fileSync *util.FileSync, data WebsitesStoreItem, backupsDir string) (int, error) {
	// 读取
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return 0, err
	}

	data.Create = time.Now().Format(time.RFC3339)
	data.Update = data.Create
	data.Group = strings.TrimSpace(data.Group)
	if data.Group == "" {
		data.Group = WebsitesGroupDefault
	}
	data.ID = util.CreateMD5(data.Group+data.Host, true)
	for _, item := range list {
		if item.ID == data.ID {
			return 0, nil
		}
		if item.Group == "" {
			item.Group = WebsitesGroupDefault
		}
	}

	// 获取当前组有多少个, 并返回合适的order
	lastOrder := w.GetLastOrderInGroup(data.Group, list) + 1
	data.Order = lastOrder
	fmt.Println("lastOrder", lastOrder)

	list = append(list, data)
	if err := w.save(list, fileSync); err != nil {
		return 0, err
	}
	return 1, nil
}

// Add 添加站点
func (w *WebsitesModel) AddGroupOrder(groupSync *util.FileSync, data []string, backupsDir string) (int, error) {
	if err := w.saveString(data, groupSync); err != nil {
		return 0, err
	}
	return 1, nil
}

// Update 更新站点
func (w *WebsitesModel) Update(fileSync *util.FileSync, data WebsitesStoreItem, backupsDir string) (int, error) {
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return 0, err
	}
	// 遍历更新
	oldID := data.ID
	data.Update = time.Now().Format(time.RFC3339)
	data.ID = util.CreateMD5(data.Host, true)
	if data.Group == "" {
		data.Group = WebsitesGroupDefault
	}
	for index, item := range list {
		if item.ID == oldID {
			list[index] = data
		}
	}
	if err := w.save(list, fileSync); err != nil {
		return 0, err
	}
	return 1, nil
}

// Update 更新顺序
func (w *WebsitesModel) UpdateWebSiteOrder(fileSync *util.FileSync, data []WebsitesStoreItem,
	backupsDir string) (int, error) {
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return 0, err
	}

	for _, dataItem := range data {
		// 遍历更新
		oldID := dataItem.ID
		dataItem.Update = time.Now().Format(time.RFC3339)
		dataItem.ID = util.CreateMD5(dataItem.Host, true)
		if dataItem.Group == "" {
			dataItem.Group = WebsitesGroupDefault
		}
		for index, item := range list {
			if item.ID == oldID {
				list[index] = dataItem
			}
		}
	}

	if err := w.save(list, fileSync); err != nil {
		return 0, err
	}

	return 1, nil
}

// RequestWebsitesDelete 删除请求参数
type RequestWebsitesDelete struct {
	IDS string `json:"ids"`
}

// Delete 删除站点
func (w *WebsitesModel) Delete(fileSync *util.FileSync, ids []string, backupsDir string) (int, error) {
	// 读取
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return 0, err
	}
	// 遍历删除
	newList := make([]WebsitesStoreItem, 0)
	success := 0
	for _, item := range list {
		if util.StringInArray(item.ID, ids) == false {
			newList = append(newList, item)
		} else {
			success++
		}
	}
	if err := w.save(newList, fileSync); err != nil {
		return 0, err
	}
	return success, nil
}

// Groups 获取站点分组列表
func (w *WebsitesModel) Groups(fileSync *util.FileSync) ([]string, error) {
	// 读取
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return nil, err
	}
	allGroups := make([]string, 0)
	mapGroups := make(map[string]string)
	for _, item := range list {
		if item.Group == "" {
			continue
		}
		if _, ok := mapGroups[item.Group]; ok {
			continue
		}
		mapGroups[item.Group] = item.Group
		allGroups = append(allGroups, item.Group)
	}

	app := conf.App
	currGroupsInGroupFile, _ := w.ListGroupOrder(app.GroupStore.FileSync)
	newGroups := currGroupsInGroupFile
	if len(currGroupsInGroupFile) == len(allGroups) {
		newGroups = currGroupsInGroupFile
	} else {
		isAppendNew := false

		if currGroupsInGroupFile == nil {
			newGroups = allGroups
			isAppendNew = true
		} else {
			for _, groupName := range allGroups {
				isContains := false
				for _, groupName2 := range currGroupsInGroupFile {
					if groupName == groupName2 {
						isContains = true
					}
				}

				if !isContains {
					newGroups = append(newGroups, groupName)
					isAppendNew = true
				}
			}
		}

		if isAppendNew {
			w.AddGroupOrder(app.GroupStore.FileSync, newGroups, "")
		}
	}

	return newGroups, nil
}

func (w *WebsitesModel) save(list []WebsitesStoreItem, fileSync *util.FileSync) error {
	if len(list) == 0 {
		return nil
	}
	content, err := json.Marshal(list)
	if err != nil {
		return err
	}
	if err := fileSync.CoverJSON(content); err != nil {
		return err
	}
	return nil
}

func (w *WebsitesModel) saveString(data []string, groupSync *util.FileSync) error {
	if len(data) == 0 {
		return nil
	}
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := groupSync.CoverJSON(content); err != nil {
		return err
	}
	return nil
}

// InitOrder /** 初始化顺序, 如果没有的情况下, 如果有顺序将会按读取顺序添加 */
func (w *WebsitesModel) InitOrder(fileSync *util.FileSync) error {
	// 读取
	srcList := make([]WebsitesStoreItem, 0)
	srcList, _ = w.List(fileSync)

	// group list by group
	mapList := make(map[string][]WebsitesStoreItem)
	for _, item := range srcList {
		if item.Group == "" {
			item.Group = WebsitesGroupDefault
		}
		if _, ok := mapList[item.Group]; ok == false {
			mapList[item.Group] = make([]WebsitesStoreItem, 0)
		}
		mapList[item.Group] = append(mapList[item.Group], item)
	}

	// sort by group, add to orderList
	orderMapList := make(map[string][]WebsitesStoreItem)
	orderList := make([]WebsitesStoreItem, 0)
	for key, val := range mapList {
		for index, item := range val {
			item.Order = index + 1
			orderMapList[key] = append(orderMapList[key], item)
			orderList = append(orderList, item)
		}
	}

	// save
	w.save(orderList, fileSync)
	return nil
}

// 获取当前组有多少个, 并返回合适的order
func (w *WebsitesModel) GetLastOrderInGroup(groupKey string, webSiteList []WebsitesStoreItem) int {
	order := 1
	if len(webSiteList) > 0 {
		listMap := w.OrderWebSiteByOrder(webSiteList)

		var newGroupKey = groupKey
		newGroupKeyExistsCount := 0
		for key, _ := range listMap {
			if strings.HasSuffix(key, groupKey) {
				newGroupKey = key
				newGroupKeyExistsCount++
			}
		}

		println(listMap)
		if newGroupKeyExistsCount > 0 {
			order = listMap[newGroupKey][len(listMap[newGroupKey])-1].Order
		} else {
			order = len(listMap)
		}

	}

	return order
}

func (w *WebsitesModel) List2listMap(webSiteList []WebsitesStoreItem) map[string][]WebsitesStoreItem {
	mapList := make(map[string][]WebsitesStoreItem)
	for _, item := range webSiteList {
		if item.Group == "" {
			item.Group = WebsitesGroupDefault
		}
		if _, ok := mapList[item.Group]; ok == false {
			mapList[item.Group] = make([]WebsitesStoreItem, 0)
		}
		mapList[item.Group] = append(mapList[item.Group], item)
	}
	return mapList
}

// 将website list 转为分组并排序的map
func (w *WebsitesModel) OrderWebSiteByOrder(webSiteList []WebsitesStoreItem) map[string][]WebsitesStoreItem {
	mapList := w.List2listMap(webSiteList)

	// 排序
	orderMapList := make(map[string][]WebsitesStoreItem)
	//groupArr := make([]string, 0)
	for mapKey, mapValList := range mapList {
		sortList := Sort("order", mapValList)
		orderMapList[mapKey] = sortList
		//append(groupArr, mapKey)
	}

	// 分组排序
	//orderMapList
	app := conf.App

	groupArr, _ := w.ListGroupOrder(app.GroupStore.FileSync)

	// 编号-分组名称, 用于前端排序不规则
	newWebList := make(map[string][]WebsitesStoreItem)
	if len(groupArr) > 0 && len(orderMapList) > 0 {
		for index, group := range groupArr {
			newWebList[strconv.Itoa(index)+"-"+group] = orderMapList[group]
		}
	} else {
		newWebList = orderMapList
	}

	return newWebList
}
