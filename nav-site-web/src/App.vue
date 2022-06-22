<template>
  <div class="container">
    <div class="header">
      <img src="logo.png" class="logo" v-if="siteInfo.HasLogo" />
      <div class="site-name">{{ siteInfo.Title }}</div>
    </div>

    <web-list/>
    <div class="footer">
      <div>
        <div class="copyright1" target="blank"  v-for="(item,index) in copyrightArr" :key="index">
          {{ item }}
        </div>
        <a class="copyright2" :href="siteInfo.Url" target="blank">{{ copyright }} 版权所有&copy;2022-{{ currYear }}</a>
      </div>
    </div>
  </div>
</template>

<script>
import WebList from "@/components/Web/WebList.vue";
import {WebInfo} from "@/api/website.js";

export default {
  name: "App",
  components: {
    WebList,
  },
  data() {
    return {
      title: process.env.VUE_APP_TITLE,
      siteInfo: {},
      copyrightArr: [],
      copyright: null,
      currYear: new Date().getFullYear()
    }
  },
  created() {
    this.getSiteInfo()
  },
  mounted() {
    this.title = process.env.VUE_APP_TITLE ? process.env.VUE_APP_TITLE : "迷你网址导航NavSite"
    document.title = this.title

  },
  computed: {},
  methods: {
    getSiteInfo() {
      const that = this
      WebInfo().then(res => {
        let data = res.data.data
        that.siteInfo = data

        if (that.siteInfo.Title) document.title = that.siteInfo.Title
        that.initFooterData()
      })
    },

    initFooterData() {
      const copyright = this.siteInfo.Copyright
      if (copyright) this.copyrightArr = copyright.split("|")
      if (this.copyrightArr.length > 1) this.copyright = this.copyrightArr.pop()
    },

    initHeaderFooterData() {
      // this.siteInfo
      // this.title = process.env.VUE_APP_TITLE ? process.env.VUE_APP_TITLE : "迷你网址导航NavSite"
      // document.title = this.title
    }
  }
};
</script>

<style>
body {
  background-color: #e9e9eb;
}

.container {
  width: 100%;
  min-height: 100%;
  position: relative;
  display: block;
}

.header {
  max-width: 1100px;
  margin: 10px auto 0 auto;
  position: relative;
}

.logo {
  width: 120px;
  height: 40px;
  margin-right: 12px;
}

.site-name {
  /*position: absolute;*/
  /*left: 0;*/
  display: inline-block;
  font-size: 28px;
  text-align: left;
  font-weight: bolder;
  color: #3a8ee6;
  height: 40px;
  line-height: 40px;
}

.footer {
  text-align: center;
}

.footer a {
  text-decoration: none;
}

.copyright1 {
  font-size: 14px;
  color: #605f5f;
  padding: 10px 0;
}

.copyright2 {
  color: #555;

}
.copyright2:active {
  color: #555;
}
</style>
