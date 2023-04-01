<template>
  <el-dialog
    :title="saveTitle"
    :width="600"
    :close-on-click-modal="false"
    @closed="closeModel"
    v-model="dialogFormVisible"
  >
    <el-form :model="formData" :rules="rules" ref="formData">
      <el-form-item label="账号" :label-width="formLabelWidth" prop="name">
        <el-input
          class="web-input-width"
          v-model="formData.name"
          autocomplete="on"
        ></el-input>
      </el-form-item>
      <el-form-item label="密码" :label-width="formLabelWidth" prop="pwd">
        <el-input
          class="web-input-width"
          type="password"
          v-model="formData.pwd"
          autocomplete="on"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-form-item class="dialog-footer">
        <el-button type="primary" @click="submit('formData')">登录</el-button>
        <el-button @click="resetForm('formData')">重置</el-button>
      </el-form-item>
    </template>
  </el-dialog>
</template>

<script>
import { AuthLogin } from "@/api/auth.js";
import { CookieSet } from "@/api/cookie.js";

export default {
  emits: ['show-model'],
  name: "LoginForm",
  props: {
    modeShow: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      saveTitle: "登录",
      dialogFormVisible: false,
      formLabelWidth: "120px",
      formData: {
        name: "",
        pwd: ""
      },
      rules: {
        name: [{ required: true, message: "用户名不能为空", trigger: "blur" }],
        pwd: [{ required: true, message: "密码不能为空", trigger: "blur" }],
      },
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.dialogFormVisible = this.modeShow;
    },
    submit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.login();
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    closeModel() {
      this.emitCloseModel({ show: false });
    },
    emitCloseModel() {
      this.$emit("show-model", { show: false });
    },
    login() {
      const _this = this;
      AuthLogin(this.formData)
        .then((res) => {
          res = res.data;
          if (res.code) {
            this.$message.error(res.message || "登录失败");
          } else {
            // 获取token 和过期时间
            let token = res.data.token;
            let expire = res.data.expire;
            CookieSet(token, expire);
            this.$message.success({
              message: "登录成功",
              onClose() {
                _this.emitCloseModel();
              },
            });
          }
        })
        .catch((err) => {
          console.log(err);
          this.$message.error("保存失败");
        });
    },
  },
  watch: {
    modeShow(n) {
      this.dialogFormVisible = n;
    },
  },
};
</script>

<style>
.el-dialog {
  z-index: 999999;
}
.web-input-width input {
  max-width: 372px;
}
</style>
