<template>
  <div class="group-card">
    <el-row>
      <el-col class="web-site-col" :xs="8" :sm="6" :md="4" :lg="3" :xl="3"
              v-for="(webInfo, index) in sortMapList" :key="index" :draggable="canDrag"
              @dragstart="dragstart(webInfo)" @dragenter="dragenter(webInfo, $event)"
              @dragend="dragend(webInfo, $event)" @dragover="dragover($event)"
      >
        <web-item ref="webItem1" :webInfo="webInfo" :webGroups="webGroups" :canEdit="canEdit"/>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import WebItem from "@/components/Web/WebItem.vue";
import {WebSiteOrder} from "@/api/website.js";

export default {
  name: "WebGroup",
  components: {
    // eslint-disable-next-line vue/no-unused-components
    WebItem,
  },
  props: {
    webList: {
      type: Array,
      default() {
        return [];
      },
    },
    webGroups: {
      type: Array,
      default() {
        return [];
      },
    },
    canEdit: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
      sortMapList: this.webList,
      webList2: null,
      oldVal: null,
      newVal: null,
      // 是否能拖动
      canDrag: false,
    }
  },
  mounted() {

  },
  methods: {
    // eslint-disable-next-line no-unused-vars
    endDrag(index) {
      this.canDrag = false;
      let groupCar = document.querySelectorAll(".group-card")[index];

      groupCar.children[0].children.forEach(function (curr) {
        if (curr.children.length > 1) {
          curr.lastChild.remove();
          curr.firstElementChild.style.pointerEvents = "auto";
        }
      })
    },
    startDrag(index) {
      this.canDrag = true;
      let groupCar = document.querySelectorAll(".group-card")[index];

      groupCar.children[0].children.forEach(function (curr) {
        if (curr.children.length === 1) {
          let tagName = "div";
          let htmlDivElement = document.createElement(tagName);
          htmlDivElement.style.position = "absolute";
          htmlDivElement.style.top = "0";
          htmlDivElement.style.left = "0";
          htmlDivElement.style.width = "100%";
          htmlDivElement.style.height = "100%";
          htmlDivElement.style.opacity = "0.1";

          htmlDivElement.onmouseover = () => {
            htmlDivElement.style.background = "#333";
            htmlDivElement.style.cursor = "pointer";
          }

          htmlDivElement.onmouseout = () => {
            htmlDivElement.style.background = "#fff";
          }

          curr.style.position = "relative";
          curr.appendChild(htmlDivElement);

          curr.firstElementChild.style.pointerEvents = "none";
        } else {
          curr.lastChild.remove();
          curr.firstElementChild.style.pointerEvents = "auto";
        }

      })
    },

    dragstart(val) {
      if (!this.canDrag) return false
      this.oldVal = val;
    },

    dragenter(val, event) {
      this.newVal = val;
      if (!this.canDrag) return false
      event.preventDefault()
    },

    // eslint-disable-next-line no-unused-vars
    dragend(val, event) {
      if (!this.canDrag) return false
      if (this.newVal != this.oldVal) {
        let newList = [...this.sortMapList]
        let oldIndex = this.sortMapList.indexOf(this.oldVal);
        let newIndex = this.sortMapList.indexOf(this.newVal);

        newList.splice(oldIndex, 1, this.newVal)
        newList.splice(newIndex, 1, this.oldVal)
        newList = newList.map((item, index) => {
          item.order = (index + 1)
          return item
        })

        this.sortMapList = newList

        let _this = this;
        // 网站列表排序
        WebSiteOrder(this.sortMapList).then((res) => {
          res = res.data;
          if (res.code) {
            this.$message.error(res.message || "保存失败");
          } else {
            this.$message.success({
              message: res.message || "保存成功",
              onClose() {
                try {
                  _this.emitCloseModel();
                  // eslint-disable-next-line no-empty
                } catch (e) {}
              },
            });
          }
        }).catch((err) => {
          console.log(err);
          this.$message.error("保存失败");
        });
      }
    },

    // eslint-disable-next-line no-unused-vars
    dragover(event) {
      if (!this.canDrag) return false
      event.preventDefault()
    },
  }
};
</script>

<style scoped>
.group-card {
  display: block;
  max-width: 1068px;
  height: 100%;
  z-index: 1;
  /*background-color: #FAFAFA;*/
}

.web-site-col {
  margin: 4px 0 0 0;
  flex-direction: column;
  display: flex;
  align-items: center;
  justify-items: center;
}
.web-site-col:hover {
  background: #ebebeb;
  cursor: pointer;
}
</style>
