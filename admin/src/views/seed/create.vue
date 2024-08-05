<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="种子等级">
        <el-input-number v-model="form.level" />
      </el-form-item>
      <el-form-item label="种子名称">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="说明">
        <el-input v-model="form.desc" />
      </el-form-item>
      <el-form-item label="成长时长(m)">
        <el-input-number v-model="form.grow_time" />
      </el-form-item>
      <el-form-item label="收获植物ID">
        <el-input-number v-model="form.plant_id" />
      </el-form-item>
      <el-form-item label="收获植物数量">
        <el-input-number v-model="form.plant_number" />
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
import { createItem } from '@/api/seed'
import { baseurl } from '@/api/url'
export default {
  data() {
    return {
      form: {
        level: 1,
        title: '',
        desc: '',
        grow_time: 0,
        imgurl: 'http',
        plant_id: 0,
        plant_number: 0,
      },
      uplodurl: baseurl() + 'admins/upload',
      imageUrl: '',
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        level: this.form.level,
        title: this.form.title,
        desc: this.form.desc,
        grow_time: this.form.grow_time,
        imgurl: this.form.imgurl,
        plant_id: this.form.plant_id,
        plant_number: this.form.plant_number,
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
      this.form.imgurl = res.data.imgurl;
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
