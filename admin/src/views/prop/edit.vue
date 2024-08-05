<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="狗粮ID">
        <el-input-number v-model="form.id" />
      </el-form-item>
      <el-form-item label="狗粮名称">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="狗粮名称_英文">
        <el-input v-model="form.title_1" />
      </el-form-item>
      <el-form-item label="狗粮名称_俄文">
        <el-input v-model="form.title_2" />
      </el-form-item>
      <el-form-item label="狗粮说明">
        <el-input v-model="form.desc" />
      </el-form-item>
      <el-form-item label="狗粮说明_英文">
        <el-input v-model="form.desc_1" />
      </el-form-item>
      <el-form-item label="狗粮说明_俄文">
        <el-input v-model="form.desc_2" />
      </el-form-item>
      <el-form-item label="持续时长">
        <el-input-number v-model="form.type" />
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
        <el-button type="primary" @click="onSubmit">提交</el-button>
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
import { editItem } from '@/api/prop'
import { baseurl } from '@/api/url'
export default {
  data() {
    return {
      form: {
        id: 1,
        title: '',
        title_1: '',
        title_2: '',
        desc: '',
        desc_1: '',
        desc_2: '',
        type: 1,
        img: '',
        extra: {},
      },
      uplodurl: baseurl() + 'admins/upload',
      imageUrl: '',
    }
  },
  created() {
    const seed = this.$route.query;
    console.log('seed', this.$route.query)
    if(seed){
      this.form = seed;
      this.imageUrl = seed.img;
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        id: this.form.id,
        title: this.form.title,
        title_1: this.form.title_1,
        title_2: this.form.title_2,
        desc: this.form.desc,
        desc_1: this.form.desc_1,
        desc_2: this.form.desc_2,
        type: Number(this.form.type),
        img: this.form.img,
        extra: JSON.stringify(this.form.extra),
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
