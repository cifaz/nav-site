<template>
  <div class="fixed-nav">
    <div class="fixed-nav-item" v-for="(item,index) in webGroups" :key="index" @click="toGroup(item)">
      {{ item }}
    </div>
  </div>
</template>

<script>
export default {
  name: "WebFixedNav",
  props: {
    webGroups: {
      type: Array,
      default() {
        return [];
      },
    },
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll)
  },
  data() {
    return {};
  },
  methods: {
    // 应该是通用的锚点方法
    toGroup(item) {
      let itemDom = document.querySelector("#" + item);
      window.scrollTo({
        top: itemDom.offsetTop,
        behavior: "smooth"
      })
    },

    // 通用方法, 随拉动移动
    handleScroll() {
      const fixedNav = document.querySelector(".fixed-nav")
      const navBody = document.querySelector(".nav-body")
      let navBodyTop = navBody.getBoundingClientRect().top;

      if (navBodyTop < 0) {
        fixedNav.style.top = Math.abs(navBodyTop) + 1 + "px";
      } else {
        fixedNav.style.top = 0;
      }
    }
  }
}

</script>

<style>
.fixed-nav {
  position: absolute;
  max-width: 150px;
  min-width: 80px;
  border-radius: 5px;
  background: #FFF;
  top: 1px;
  left: -10px;
  transform: translate(-100%, 0);
}

.fixed-nav-item {
  margin: 12px 0;
  padding: 5px 10px 5px 10px;
  text-overflow: ellipsis;
  overflow: hidden;
}

.fixed-nav-item:hover {
  color: #FFFFFF;
  cursor: pointer;
  background: #3a8ee6;
}

</style>