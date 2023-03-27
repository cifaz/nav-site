<template>
  <el-popover
      placement="top-start"
      :title="webInfo.name"
      :width="200"
      trigger="hover"
      :show-after="1000"
      :content="tipContent"
  >
    <template #reference>
      <div class="web-item web-item-claz">
        <div v-if="webInfo.pic"
             class="web-item-avatar"
             @click="openWeb"
             @mouseover="showBackgroundEvent(true)"
             @mouseout="showBackgroundEvent(false)"
        >
          <img
              v-if="webInfo.pic"
              :src="webInfo.pic"
              class="web-item-avatar-img"
          />
          <span class="web-item-avatar-img" v-else>{{ webInfo.name }}</span>
        </div>
        <div class="web-item-title" @click="openWeb">
          {{ webInfo.name }}
        </div>

        <el-dropdown trigger="click" class="web-item-tool" v-if="canEdit">
          <!--          <i class="el-icon-more web-item-tool-icon"></i>-->
          <el-icon>
            <MoreFilled/>
          </el-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item
                  @click="editWebItemEvent"
              >
                <el-icon>
                  <EditPen/>
                </el-icon>
                编辑
              </el-dropdown-item>
              <el-dropdown-item
                  @click="deleteWebItemEvent"
              >
                <el-icon>
                  <DeleteFilled/>
                </el-icon>
                删除
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <div class="web-item-background" :class="showBackgroundClass"></div>
        <web-form
            :modeShow="modeShow"
            :webItem="webItem"
            :webGroups="webGroups"
            @showModel="showModel"
        />
      </div>
    </template>
    <div>
      <p><span>描述：{{ webInfo.desc || webInfo.name }}</span></p>
      <p><span>地址：{{ webInfo.host }}</span></p>
    </div>
  </el-popover>
</template>

<script>
import WebForm from "@/components/Web/WebForm.vue";
import {WebDelete} from "@/api/website.js";

export default {
  name: "WebItem",
  components: {
    WebForm,
  },
  props: {
    webInfo: {
      type: Object,
      default() {
        return {
          id: "",
          group: "default",
          name: "简言",
          pic: "https://www.jianean.com/favicon.ico",
          host: "https://www.jianean.com",
          desc: "简言",
        };
      },
    },
    canEdit: {
      type: Boolean,
      default: false,
    },
    webGroups: {
      type: Array,
      default() {
        return [];
      },
    },
  },
  data() {
    return {
      showBlack: false,
      modeShow: false,
      webItem: this.webInfo,
    };
  },
  computed: {
    showBackgroundClass() {
      return this.showBlack
          ? "web-item-background-show"
          : "web-item-background-hide";
    },
    tipContent() {
      return this.webItem.desc || this.webItem.name
    }
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      // this.webItem = JSON.parse(JSON.stringify(this.webInfo));
    },
    openWeb() {
      window.open(this.webItem.host);
    },
    showBackgroundEvent(show) {
      this.showBlack = !!show;
    },
    editWebItemEvent() {
      this.modeShow = true;
    },
    deleteWebItemEvent() {
      this.$confirm("确认要删除这个站点记录吗？")
          .then(() => {
            const data = {ids: this.webItem.id};
            WebDelete(data)
                .then((res) => {
                  res = res.data;
                  if (res.code) {
                    this.$message.error(res.message);
                  } else {
                    this.$message.success({
                      message: res.message,
                      type: "success",
                      onClose() {
                        window.location.reload();
                      },
                    });
                  }
                })
                .catch((err) => {
                  console.log(err);
                  this.$message.error("删除失败");
                });
          })
          .catch((err) => {
            console.log(err);
          });
    },
    showModel(data) {
      this.modeShow = data.show;
      this.webItem = data.webItem || {};
    },
  },
  watch: {
    webItem: {
      deep: true,
    },
  },
};
</script>

<style>
.web-item {
  position: relative;
  display: block;
  border-radius: 4px;
  width: 100%;
  min-height: 10px;
  min-width: 10px;
  cursor: pointer;
}

.web-item-background {
  /*position: absolute;*/
  top: 0;
  left: 0;
  display: block;
  border-radius: 4px;
  /*height: 100px;*/
  /*width: 100px;*/
  cursor: pointer;
  background-color: gray;
  z-index: 1;
}

.web-item-background:hover {
  opacity: 0.6;
}

.web-item-background-show {
  opacity: 0.6;
}

.web-item-background-hide {
  opacity: 0;
}

.web-item-avatar {
  /*position: absolute;*/
  margin: 0 auto;
  /*padding: auto;*/
  display: block;
  border-radius: 4px;
  height: 50px;
  width: 50px;
  cursor: pointer;
  background-color: white;
  z-index: 999;
  box-shadow: 0 2px 12px 0 rgb(0 0 0 / 60%);
}

.web-item-avatar-img {
  margin: 7px;
  height: 36px;
  width: 36px;
  font-size: 12px;
  text-align: center;
  /*display: block;*/
  overflow: hidden;
}

.web-item-title {
  /*position: absolute;*/
  /*bottom: 8px;*/
  /*right: 6px;*/
  text-align: center;
  /*width: 94px;*/
  margin: 5px 0;
  color: rgb(19, 17, 17);
  font-size: 12px;
  z-index: 999;
  overflow: hidden;
}

.web-item-tool {
  position: absolute;
  /*top: 6px;*/
  top: 0;
  right: 10px;
  width: 10px;
  color: rgb(185, 150, 42);
  display: none;
  z-index: 999;
  padding: 2px;
}

.web-item:hover .web-item-tool {
  display: block;
}

.web-item-tool-icon {
  color: rgb(119, 78, 216);
}
</style>
