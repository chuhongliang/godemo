<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="植物名称">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="植物名称_英文">
        <el-input v-model="form.name_1" />
      </el-form-item>
      <el-form-item label="植物名称_俄文">
        <el-input v-model="form.name_2" />
      </el-form-item>
      <el-form-item label="卖出单价">
        <el-input-number v-model="form.sell_gold" />
      </el-form-item>
      <el-form-item label="卖出单价">
        <el-input-number v-model="form.sell_gold" />
      </el-form-item>
      <el-form-item label="上传图片">
        <el-upload class="avatar-uploader" :action="uplodurl" :show-file-list="false"
          :on-success="handleAvatarSuccess" 
          :before-upload="beforeAvatarUpload">
          <img v-if="imageUrl" :src="imageUrl" class="avatar">
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>
      <el-form-item label="上传果实图片">
        <el-upload class="avatar-uploader" :action="uplodurl" :show-file-list="false"
          :on-success="handleAvatarSuccess2" 
          :before-upload="beforeAvatarUpload">
          <img v-if="imgurl" :src="imgurl" class="avatar">
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
  max-width: 178px;
  max-height: 178px;
  line-height: 178px;
  text-align: center;
}

.avatar {
  width: 178px;
  height: 178px;
  max-width: 178px;
  max-height: 178px;
  display: block;
}
</style>


<script>
import { createItem } from '@/api/plant'
import { baseurl } from '@/api/url'
export default {
  data() {
    return {
      form: {
        name: '',
        name_1: '',
        name_2: '',
        sell_gold: 0,
        res: '',
        imgurl: '',
      },
      uplodurl: baseurl() + 'admins/upload',
      imageUrl: '',
      imgurl: '',
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        name: this.form.name,
        name_1: this.form.name_1,
        name_2: this.form.name_2,
        sell_gold: this.form.sell_gold,
        res: this.form.res,
        imgurl: this.form.imgurl,
      }
      createItem(data).then(response => {
        // this.list = response.data.items
        // this.listLoading = false
        console.log("seed", response.data.item);
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
      this.form.res = res.data.imgurl;
    },
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
      console.log('res', res.data)
      this.form.imgurl = res.data.imgurl;
    },
    beforeAvatarUpload(file) {
      // const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      // if (!isJPG) {
      //   this.$message.error('上传头像图片只能是 JPG 格式!');
      // }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      // return isJPG && isLt2M;
      return  isLt2M;
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>
