import axios from "@/util/request.js";

const website = {
  add: "website/add",
  list: "website/list",
  delete: "website/delete",
  update: "website/update",
  groups: "website/groups",
  imageUpload: "upload/image",
  orderWebSite: "/website/order/list",
  orderGroup: "/website/order/group",
};

export function WebList(data) {
  return axios({
    method: "get",
    data: data,
    url: website.list,
  });
}

export function WebSiteOrder(data) {
  return axios({
    method: "put",
    data: data,
    url: website.orderWebSite,
  });
}

export function WebSiteGroupOrder(data) {
  return axios({
    method: "put",
    data: data,
    url: website.orderGroup,
  });
}

export function WebAdd(data) {
  return axios({
    method: "post",
    data: data,
    url: website.add,
  });
}


export function WebEdit(data) {
  return axios({
    method: "post",
    data: data,
    url: website.update,
  });
}

export function WebDelete(data) {
  return axios({
    method: "post",
    data: data,
    url: website.delete,
  });
}

export function WebGroups(data) {
  return axios({
    method: "get",
    data: data,
    url: website.groups,
  });
}

export const WebLogoUploadUrl = website.imageUpload
