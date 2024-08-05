<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="等级">
        <el-input-number v-model="form.level" />
      </el-form-item>
      <el-form-item label="升级经验">
        <el-input-number v-model="form.exp" />
      </el-form-item>
      <el-form-item label="奖励类型">
        <el-select v-model="form.reward.type">
          <el-option label="种子" value="1" />
          <el-option label="肥料" value="2" />
          <el-option label="狗粮" value="4" />
          <el-option label="金币" value="5" />
          <el-option label="钻石" value="6" />
        </el-select>
      </el-form-item>
      <el-form-item label="奖励ID">
        <el-input-number v-model="form.reward.id" />
      </el-form-item>
      <el-form-item label="奖励数量">
        <el-input-number v-model="form.reward.amount" />
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
import { createItem } from '@/api/monthtask'
import { baseurl } from '@/api/url'
export default {
  data() {
    return {
      form: {
        level: '',
        exp: 0,
        reward: {
          type: '',
          id: '',
          amount: 0,
          img: '',
        },
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
        exp: this.form.exp,
        reward: {
          type: this.form.reward.type,
          id: this.form.reward.id,
          amount: this.form.reward.amount,
          img: this.form.reward.img
        }
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
      this.form.reward.img = res.data.imgurl;
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
