<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="ID">
        <el-input v-model="form.id" />
      </el-form-item>
      <el-form-item label="品种">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="品种_英文">
        <el-input v-model="form.title_1" />
      </el-form-item>
      <el-form-item label="品种_俄文">
        <el-input v-model="form.title_2" />
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="form.desc" />
      </el-form-item>
      <el-form-item label="描述_英文">
        <el-input v-model="form.desc_1" />
      </el-form-item>
      <el-form-item label="描述_俄文">
        <el-input v-model="form.desc_2" />
      </el-form-item>
      <el-form-item label="上传图片">
        <el-upload class="avatar-uploader" action="http://localhost:3003/admins/upload" :show-file-list="false"
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


<script>
import { editItem } from '@/api/pet'
export default {
  data() {
    return {
      form: {
        id: 0,
        title: '',
        title_1: '',
        title_2: '',
        desc: '',
        desc_1: '',
        desc_2: '',
        res: ''
      },
      imageUrl: '',
    }
  },
  created() {
    const plant = this.$route.query;
    console.log('query', this.$route.query)
    if(plant){
      // this.form.name = plant.name;
      // this.form.sell_gold = plant.sell_gold;
      this.form = plant
      this.imageUrl = plant.res
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
        desc_1: this.form.desc_1,
        desc_2: this.form.desc_2,
        desc: this.form.desc,
        res: this.form.res,
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
