<template>
  <div>

    <div class="box-tool">


      <el-button size="small" type="info" v-if="canEdit" @click="groupOrder">
        分组排序
      </el-button>
      <el-button
          title="添加站点"
          class="button"
          type="danger"
          size="small"
          @click="addWeb"
      >
        <el-icon>
          <Plus/>
        </el-icon>
      </el-button>
      <el-button
          circle
          class="button"
          type="primary"
          size="small"
          :title="canEdit ? '退出登录' : '登录'"
          @click="login"
      >
        <el-icon>
          <User/>
        </el-icon>
      </el-button>
    </div>

    <el-dialog v-model="groupOrderVisible" title="分组排序" width="30%" center>
      <div class="group-order-item" v-for="(item,index) in webGroups" :key="index" @click="toGroup(item)"
           :draggable="true" @dragstart="dragstart(item)" @dragenter="dragenter(item, $event)"
           @dragend="dragend(item, $event)" @dragover="dragover($event)">
        {{ item }}
      </div>

      <template #footer>
      <span class="dialog-footer">
        <el-button @click="groupOrderVisible = false">取消</el-button>
        <el-button type="primary" @click="groupOrderSave">保存分组</el-button>
      </span>
      </template>
    </el-dialog>

    <div class="nav-body">
      <web-fixed-nav :webGroups="webGroups"/>
      <el-card
          class="box-card"
          v-for="(webList, group, index) in webListObject"
          shadow="hover"
          :key="index"
      >
        <template #header>
          <div class="card-header">
          <span class="group-info">
            <span :id="group" class="group-text">
              {{ group }}
            </span>
            <el-button ref="webSiteOrderBtn" class="webSiteOrderBtn" link size="small" v-if="canEdit" @click="orderWebsite($this, index)">
                 {{ webSiteOrderBtnText }}
            </el-button>
          </span>
          </div>
        </template>
        <web-group ref="webSiteGroup" :webList="webList" :webGroups="webGroups" :canEdit="canEdit"/>
      </el-card>
    </div>

    <web-form
        :modeShow="modeShow"
        :webGroups="webGroups"
        @showModel="showModel"
    />
    <login-form :modeShow="modeLoginForm" @showModel="showLoginModel"/>
  </div>
</template>

<script>
import WebFixedNav from "@/components/Web/WebFixedNav.vue";
import WebGroup from "@/components/Web/WebGroup.vue";
import WebForm from "@/components/Web/WebForm.vue";
import LoginForm from "@/components/Auth/LoginForm.vue";
import {WebGroups, WebList, WebSiteGroupOrder} from "@/api/website.js";
import {CookieGetToken, CookieRemoveToken} from "@/api/cookie.js";

export default {
  name: "WebList",
  components: {
    WebGroup,
    WebForm,
    LoginForm,
    WebFixedNav,
  },
  data() {
    return {
      // 网址对象, 未分组
      webListObject: {},
      // 分组
      webGroups: [],
      modeShow: false,
      modeLoginForm: false,
      canEdit: false,
      token: "",
      loginIcon: "",
      // 是否允许排序
      webSiteOrderEnable: false,
      webSiteOrderBtnText: "排序",
      // 分组排序参数
      groupOrderVisible: false,
      groupOldVal: null,
      groupNewVal: null,
    };
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      this.token = CookieGetToken();
      this.isCanEdit();
      this.getWebGroups();
      this.setLoginIcon();
      this.getWebList();
    },
    getWebList() {
      let that = this;
      WebList()
          .then((res) => {
            res = res.data;
            if (res.code) {
              that.$message.error(
                  res.message || "获取站点失败，请联系系统管理员！"
              );
            } else {
              let rows = res.data.rows

              let newRows = {}
              for (let groupKey in rows) {
                let groupKeyArr = groupKey.split("-")
                newRows[groupKeyArr[1]] = rows[groupKey];
                // rows[groupKeyArr[0]] = rows[groupKeyArr[0]].sort(that.sortId);
              }

              that.webListObject = newRows;
            }
          }).catch((err) => {
            console.log(err);
          });
    },
    sortId(a, b) {
      return a.order - b.order;
    },
    getWebGroups() {
      WebGroups().then((res) => {
            res = res.data;
            if (res.code) {
              this.$message.error(
                  res.message || "获取站点失败，请联系系统管理员！"
              );
            } else {
              this.webGroups = res.data || [];
            }
          })
          .catch((err) => {
            console.log(err);
          });
    },
    addWeb() {
      if (this.token) {
        this.modeShow = true;
      } else {
        this.modeLoginForm = true;
      }
    },
    login() {
      if (this.canEdit) {
        const _this = this;
        this.$confirm("确定要退出登录吗？").then(() => {
          CookieRemoveToken();
          CookieGetToken();
          _this.isCanEdit();
          _this.setLoginIcon();
          window.location.reload();
        });
      } else {
        this.modeLoginForm = true;
      }
    },
    showModel(data) {
      this.modeShow = data.show;
      window.location.reload();
    },
    showLoginModel(data) {
      this.modeLoginForm = data.show;
      this.token = CookieGetToken();
      this.setLoginIcon();
      this.isCanEdit();
    },
    isCanEdit() {
      return (this.canEdit = !!CookieGetToken());
    },
    setLoginIcon() {
      return (this.loginIcon = CookieGetToken()
          ? "el-icon-switch-button"
          : "el-icon-user-solid")
    },

    orderWebsite(_this, index) {
      if (!this.webSiteOrderEnable) {
        this.$refs.webSiteGroup[index].startDrag(index)
        this.webSiteOrderEnable = true
        // this.webSiteOrderBtnText = "取消排序"
        this.$refs.webSiteOrderBtn[index].$el.innerText="取消排序"
      } else {
        this.webSiteOrderEnable = false;
        this.$refs.webSiteGroup[index].endDrag(index)
        this.$refs.webSiteOrderBtn[index].$el.innerText="排序"
      }
    },

    groupOrder() {
      this.groupOrderVisible = true
    },

    dragstart(val) {
      this.groupOldVal = val;
    },

    dragenter(val, event) {
      this.groupNewVal = val;
      event.preventDefault()
    },

    // eslint-disable-next-line no-unused-vars
    dragend(val, event) {
      if (this.groupNewVal != this.groupOldVal) {
        let oldIndex = this.webGroups.indexOf(this.groupOldVal);
        let newIndex = this.webGroups.indexOf(this.groupNewVal);

        this.webGroups.splice(oldIndex, 1, this.groupNewVal)
        this.webGroups.splice(newIndex, 1, this.groupOldVal)
      }
    },

    dragover(event) {
      event.preventDefault()
    },

    groupOrderSave() {
      let _this = this;
      // 网站列表排序
      WebSiteGroupOrder(this.webGroups).then((res) => {
        res = res.data;
        if (res.code) {
          this.$message.error(res.message || "保存失败");
        } else {
          this.$message.success({
            message: res.message || "保存成功",
            onClose() {
              try {
                _this.emitCloseModel();
              } catch (e) {
                _this.groupOrderVisible = false
              }
            },
          });
        }
      }).catch((err) => {
        console.log(err);
        this.$message.error("保存失败");
      });
    }
  },
  watch: {
    token() {
      this.setLoginIcon();
      this.isCanEdit();
    },
    canEdit() {
    },
    loginIcon() {
    },
  },
};
</script>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.nav-body {
  max-width: 1100px;
  margin: 15px auto;
  position: relative;
}

.group-order-item {
  margin: 5px 0;
  padding: 3px 0;
}

.group-order-item:hover {
  background: #8c939d;
  cursor: pointer;
}

.box-card {
  margin-top: 10px;
}

.el-card__header {
  padding: 4px 20px;
}

.el-icon-s-tools {
  color: #409eff;
}

.group-info {
  color: #409eff;
}

.group-text {
  /*padding: 0 0.2rem;*/
  font-size: 0.9rem;
  text-decoration: none;
  color: #000;
}

.group-text:active {
  color: #000;
}

.webSiteOrderBtn {
  margin-left: 15px;
}

.box-tool {
  max-width: 1100px;
  text-align: right;
  margin: 0 auto;
  position: relative;
  /*margin-right: 4rem;*/
  /*margin-top: 2rem;*/
}


</style>
