<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="名称">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="名称_英文">
        <el-input v-model="form.title_1" />
      </el-form-item>
      <el-form-item label="名称_俄文">
        <el-input v-model="form.title_2" />
      </el-form-item>
      <el-form-item label="说明">
        <el-input v-model="form.desc" />
      </el-form-item>
      <el-form-item label="说明_英文">
        <el-input v-model="form.desc_1" />
      </el-form-item>
      <el-form-item label="说明_俄文">
        <el-input v-model="form.desc_2" />
      </el-form-item>
      <el-form-item label="类型">
        <el-select v-model="form.type">
          <el-option label="限时" value="1" />
          <el-option label="永久" value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="增加产值">
        <el-input-number v-model="form.add_value" />
      </el-form-item>
      <el-form-item label="花费金币">
        <el-input-number v-model="form.gold_price" />
      </el-form-item>
      <el-form-item label="花费钻石">
        <el-input-number v-model="form.diamond" />
      </el-form-item>
      <el-form-item label="上传图片">
        <el-upload class="avatar-uploader" :action="uplodurl" :show-file-list="false"
          :on-success="handleAvatarSuccess" 
          :before-upload="beforeAvatarUpload">
          <img v-if="imageUrl" :src="imageUrl" class="avatar">
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">创建</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}

.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>


<script>
import { createItem } from '@/api/buff'
import { baseurl } from '@/api/url'
export default {
  data() {
    return {
      form: {
        title: '',
        title_1: '',
        title_2: '',
        desc: '',
        desc_1: '',
        desc_2: '',
        gold_price: 0,
        diamond: 0,
        img: '',
        type: 1,
        add_value: 0,

      },
      uplodurl: baseurl() + 'admins/upload',
      imageUrl: '',
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        title: this.form.title,
        title_1: this.form.title_1,
        title_2: this.form.title_2,
        desc_1: this.form.desc_1,
        desc_2: this.form.desc_2,
        desc: this.form.desc,
        gold_price: this.form.gold_price,
        diamond: this.form.diamond,
        img: this.form.img,
        type: Number(this.form.type),
        add_value: this.form.add_value,
      }
      console.log("data", data);

      createItem(data).then(response => {
        // this.list = response.data.items
        // this.listLoading = false
        console.log("task", response.data.item);
      })
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    },
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
      console.log('res', res.data)
      this.form.img = res.data.imgurl;
    },
    beforeAvatarUpload(file) {
      // const isJPG = file.type === 'image/jpeg/';
      const isLt2M = file.size / 1024 / 1024 < 2;
      // if (!isJPG) {
      //   this.$message.error('上传头像图片只能是 JPG 格式!');
      // }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      // return isJPG && isLt2M;
      return isLt2M
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>
