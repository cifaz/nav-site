<template>
  <div class="container">

    <web-list />
    <div class="footer">
      <p>
        <a href="http://hibao.work/" target="blank">海宝工作组 hibao.work</a>
      </p>
    </div>
  </div>
</template>

<script>
import WebList from "@/components/Web/WebList.vue";

export default {
  name: "App",
  components: {
    WebList,
  },
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

}

.footer {
  text-align: center;
}
.footer a {
  text-decoration: none;
}
</style>
