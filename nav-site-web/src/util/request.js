import axios from "axios";
import {CookieGetToken } from "@/api/cookie.js";

const instance = axios.create({
  baseURL: process.env.VUE_APP_API_BASE,
  timeout: 6000,
});

instance.interceptors.request.use(
  function(config) {
    // 在发送请求之前做些什么
    const token = CookieGetToken()
    config.headers["Authorization"] = "Bearer " + token;
    return config;
  },
  function(error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
instance.interceptors.response.use(function (response) {
  // 对响应数据做点什么
  return response;
}, function (error) {
  // 对响应错误做点什么
  return Promise.reject(error);
});

export default instance;
