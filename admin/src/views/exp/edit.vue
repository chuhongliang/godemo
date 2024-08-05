<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="ID">
        <el-input-number v-model="form.id" />
      </el-form-item>
      <el-form-item label="等级">
        <el-input-number v-model="form.level" />
      </el-form-item>
      <el-form-item label="升级经验">
        <el-input-number v-model="form.exp" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>


<script>
import { editItem } from '@/api/exp'
export default {
  data() {
    return {
      form: {
        id: 1,
        level: 1,
        exp: 0,
      },
    }
  },
  created() {
    const seed = this.$route.query;
    console.log('seed', this.$route.query)
    if(seed){
      this.form = seed;
      this.imageUrl = seed.res;
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        id: this.form.id,
        level: this.form.level,
        exp: this.form.exp,
      }
      editItem(data).then(response => {
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
