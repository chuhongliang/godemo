<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="ID">
        <el-input-number v-model="form.id" />
      </el-form-item>
      <el-form-item label="商品类型">
        <el-select v-model="form.item_type">
          <el-option label="种子" value="1" />
          <el-option label="肥料" value="2" />
          <el-option label="狗粮" value="4" />
        </el-select>
      </el-form-item>
      <el-form-item label="商品ID">
        <el-input-number v-model="form.item_id" />
      </el-form-item>
      <el-form-item label="商品名称">
        <el-input v-model="form.item_title" />
      </el-form-item>
      <el-form-item label="商品单价">
        <el-input-number v-model="form.item_price" />
      </el-form-item>
      <el-form-item label="商品描述">
        <el-input v-model="form.item_desc" />
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
import { editItem } from '@/api/shop'
import { baseurl } from '@/api/url'
export default {
  data() {
    return {
      form: {
        id: 1,
        item_id: '',
        item_type: '',
        item_title: '',
        item_price: '',
        item_desc: '',
        item_imgurl: '',
      },
      uplodurl: baseurl() + 'admins/upload',
      imageUrl: '',
    }
  },
  created() {
    const item = this.$route.query;
    console.log('item', this.$route.query)
    if(item){
      this.form = item;
      this.imageUrl = item.item_res;
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        id: this.form.id,
        item_id: this.form.item_id,
        item_type: Number(this.form.item_type),
        item_title: this.form.item_title,
        item_price: this.form.item_price,
        item_desc: this.form.item_desc,
        item_imgurl: this.form.item_imgurl,
      }
      editItem(data).then(response => {
        // this.list = response.data.items
        // this.listLoading = false
        console.log("item", response.data.item);
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
      this.form.item_imgurl = res.data.imgurl;
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
