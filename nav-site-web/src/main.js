import { createApp } from "vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import App from "./App.vue";
import axios from "axios";
import "dayjs/locale/zh-cn";
import locale from "element-plus/lib/locale/lang/zh-cn";
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const app = createApp(App);


for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

document.title = process.env.VUE_APP_TITLE | "网址导航"

app.config.globalProperties.$axios = axios;
app.use(ElementPlus, { locale });
app.mount("#app");
