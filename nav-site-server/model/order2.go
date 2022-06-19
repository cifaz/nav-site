package model

import (
	"sort"
)

type MapsSort struct {
	Key     string
	MapList []WebsitesStoreItem
}

func (m *MapsSort) Len() int {
	return len(m.MapList)
}

func (m *MapsSort) Less(i, j int) bool {
	return m.MapList[i].Order < m.MapList[j].Order
}

func (m *MapsSort) Swap(i, j int) {
	m.MapList[i], m.MapList[j] = m.MapList[j], m.MapList[i]
}

func Sort(key string, maps []WebsitesStoreItem) []WebsitesStoreItem {
	mapsSort := MapsSort{}
	mapsSort.Key = key
	mapsSort.MapList = maps
	sort.Sort(&mapsSort)

	return mapsSort.MapList
}
