<template>
  <el-dialog
    :title="saveTitle"
    :width="600"
    :close-on-click-modal="false"
    v-model="dialogFormVisible"
    @closed="closeModel"
  >
    <el-form :model="formData" :rules="rules" ref="formData">
      <el-form-item label="上传图片" :label-width="formLabelWidth" prop="pic">
        <el-upload
          :action="upload.url"
          list-type="picture-card"
          :on-preview="uploadPreview"
          :on-remove="uploadRemove"
          :on-success="uploadSuccess"
          :on-error="uploadError"
          :accept="accept"
          :limit="1"
        >
<!--          <i class="el-icon-plus"></i>-->
          <el-icon><Plus /></el-icon>
        </el-upload>
        <el-dialog v-model="dialogVisible">
          <img width="100%" :src="dialogImageUrl" alt="" />
        </el-dialog>
      </el-form-item>
      <el-form-item
        label="站点图片地址"
        :label-width="formLabelWidth"
        prop="pic"
      >
        <el-input
          class="web-input-width"
          v-model="formData.pic"
          autocomplete="on"
        ></el-input>
      </el-form-item>
      <el-form-item label="站点名称" :label-width="formLabelWidth" prop="name">
        <el-input
          class="web-input-width"
          v-model="formData.name"
          autocomplete="on"
        ></el-input>
      </el-form-item>
      <el-form-item label="站点链接" :label-width="formLabelWidth" prop="host">
        <el-input
          class="web-input-width"
          v-model="formData.host"
          autocomplete="on"
        ></el-input>
      </el-form-item>
      <el-form-item label="站点分组" :label-width="formLabelWidth" prop="group">
        <el-col :span="13">
          <el-input
            class="web-input-width web-item-group-add"
            v-if="isEditGroup"
            v-model="formData.group"
            autocomplete="on"
          ></el-input>
          <el-select
            v-else
            v-model="formData.group"
            placeholder="请选择站点分组"
          >
            <el-option
              v-for="(item, index) in webGroups"
              :label="item"
              :value="item"
              :key="index"
            ></el-option>
          </el-select>
        </el-col>
        <el-col :span="11">
          <el-button
            class="edit-group"
            type="warning"
            icon="el-icon-edit"
            @click="editGroup"
            ><el-icon><Edit /></el-icon>{{ isEditGroup ? "选择分组" : "编辑分组" }}</el-button
          >
        </el-col>
      </el-form-item>
      <el-form-item label="站点描述" :label-width="formLabelWidth" prop="desc">
        <el-input
          class="web-input-width"
          v-model="formData.desc"
          autocomplete="on"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-form-item class="dialog-footer">
        <el-button type="primary" @click="submit('formData')">保存</el-button>
        <el-button @click="resetForm('formData')">重置</el-button>
      </el-form-item>
    </template>
  </el-dialog>
</template>

<script>
import { WebAdd, WebEdit } from "@/api/website.js";
import { WebLogoUploadUrl } from "@/api/upload.js";

export default {
  emits: ['show-model'],
  name: "WebForm",
  props: {
    modeShow: {
      type: Boolean,
      default: false,
    },
    webItem: {
      type: Object,
      default() {
        return {
          id: "",
          group: "",
          pic: "",
          name: "",
          host: "",
          desc: "",
        };
      },
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
      saveTitle: "",
      dialogFormVisible: false,
      formLabelWidth: "120px",
      formData: null,
      upload: {
        url: WebLogoUploadUrl,
        accept: "image/*",
        dialogImageUrl: "",
      },
      rules: {
        name: [
          { required: true, message: "站点名称不能为空", trigger: "blur" },
        ],
        host: [
          { required: true, message: "站点地址不能为空", trigger: "blur" },
          { type: "url", message: "不是有效的地址", trigger: "blur" },
        ],
      },
      isEditGroup: false,
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.dialogFormVisible = this.modeShow;
      this.formData = JSON.parse(JSON.stringify(this.webItem));
      this.saveTitle = this.formData.id ? "编辑" : "新建";
    },
    submit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.saveWeb();
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    uploadRemove(file, fileList) {
      console.log(file, fileList);
      this.formData.pic = "";
    },
    uploadPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    uploadSuccess(res) {
      if (res.code) {
        this.$message.error(res.message || "上传失败，请稍后再试！");
      } else {
        this.formData.pic = res.data.url;
        this.dialogImageUrl = this.formData.pic;
      }
    },
    uploadError(err){
      console.log(err)
    },
    closeModel() {
      this.emitCloseModel({ show: false, webItem: this.formData });
    },
    emitCloseModel() {
      this.$emit("show-model", { show: false, webItem: this.formData });
      window.location.reload();
    },
    saveWeb() {
      const _this = this;
      const saveWeb = this.formData.id ? WebEdit : WebAdd;
      const saveData = {};
      saveData.id = this.formData.id;
      saveData.group = this.formData.group;
      saveData.pic = this.formData.pic;
      saveData.name = this.formData.name;
      saveData.host = this.formData.host;
      saveData.desc = this.formData.desc;

      // 变更为统一的http头//
      /* 1. [BUG] 恢复http限制, 后期再完善
      let newHost = saveData.host;
      if (newHost.startsWith("https://")) {
        newHost = newHost.substring("https://", "//");
      } else if (newHost.startsWith("http://")) {
        newHost = newHost.substring("http://", "//");
      } else {
      //   没有开头时, 补充
        newHost = "//" + newHost;
      }
      saveData.host = newHost;
      */

      saveWeb(saveData)
        .then((res) => {
          res = res.data;
          if (res.code) {
            this.$message.error(res.message || "保存失败");
          } else {
            this.$message.success({
              message: res.message || "保存成功",
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
    editGroup() {
      this.isEditGroup = !this.isEditGroup;
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
.avatar-uploader .el-upload {
  border: 1px dashed #524e4e;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.el-upload--picture-card {
  width: 80px;
  height: 80px;
  line-height: 80px;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 18px;
  color: #8c939d;
  width: 30px;
  height: 30px;
  text-align: center;
}
.avatar {
  width: 30px;
  height: 30px;
  display: block;
}
.edit-group {
  margin-left: 2rem;
}
.web-input-width input {
  max-width: 372px !important;
}
.web-item-group-add input {
  max-width: 222px !important;
  display: inline-block;
}
</style>
