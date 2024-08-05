<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="邮件标题">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="邮件内容">
        <el-input v-model="form.content" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">创建</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>


<script>
import { createItem } from '@/api/mail'
export default {
  data() {
    return {
      form: {
        title: '系统邮件',
        content: '',
      },
      imageUrl: '',
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        title: this.form.title,
        content: this.form.content,
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
