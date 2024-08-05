<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="兑换码">
        <el-input v-model="form.code" />
      </el-form-item>
      <el-form-item label="奖励类型">
        <el-select v-model="form.item_type">
          <el-option label="种子" value="1" />
          <el-option label="肥料" value="2" />
          <el-option label="金币" value="5" />
          <el-option label="钻石" value="6" />
        </el-select>
      </el-form-item>
      <el-form-item label="奖励ID">
        <el-input-number v-model="form.item_id" />
      </el-form-item>
      <el-form-item label="奖励数量">
        <el-input-number v-model="form.amount" />
      </el-form-item>

      <el-form-item label="奖励类型2">
        <el-select v-model="form.item_type2">
          <el-option label="种子" value="1" />
          <el-option label="肥料" value="2" />
          <el-option label="金币" value="5" />
          <el-option label="钻石" value="6" />
        </el-select>
      </el-form-item>
      <el-form-item label="奖励ID2">
        <el-input-number v-model="form.item_id2" />
      </el-form-item>
      <el-form-item label="奖励数量2">
        <el-input-number v-model="form.item_number2" />
      </el-form-item>

      <el-form-item label="奖励类型3">
        <el-select v-model="form.item_type3">
          <el-option label="种子" value="1" />
          <el-option label="肥料" value="2" />
          <el-option label="金币" value="5" />
          <el-option label="钻石" value="6" />
        </el-select>
      </el-form-item>
      <el-form-item label="奖励ID3">
        <el-input-number v-model="form.item_id3" />
      </el-form-item>
      <el-form-item label="奖励数量3">
        <el-input-number v-model="form.item_number3" />
      </el-form-item>

      <el-form-item label="有效期">
        <div class="block">
          <span class="demonstration"></span>
          <el-date-picker v-model="form.expire_at" type="datetime" placeholder="选择日期时间" align="right"
            :picker-options="pickerOptions">
          </el-date-picker>
        </div>
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
import { createItem, getListByType } from '@/api/exchange'
import { baseurl } from '@/api/url'
export default {
  data() {
    return {
      form: {
        code: '',
        item_id: '',
        item_type: '',
        amount: '',
        expire_at: '',

        itme_type2: '',
        item_id2: '',
        item_number2: '',

        item_type3: '',
        item_id3: '',
        item_number3: '',
      },
      options: [],
      uplodurl: baseurl() + 'admins/upload',
      imageUrl: '',
    }
  },
  created() {
    const data = {
      code: this.form.code,
      item_type: this.form.item_type,
      item_id: this.form.item_id,
      amount: this.form.amount,
      expire_at: this.form.expire_at,

      item_type2: this.form.item_type2,
      item_id2: this.form.item_id2,
      item_number2: this.form.item_number2,

      item_type3: this.form.item_type3,
      item_id3: this.form.item_id3,
      item_number3: this.form.item_number3,

    }
    getListByType(data).then(response => {
      this.options = response.data.items
    })
  },
  watch: {
    'form.item_type': function (newVal, oldVal) {
      const data = {
        item_type: this.form.item_type
      }
      getListByType(data).then(response => {
        console.log("response", response);
        this.options = response.data.items
      })
    }
  },

  methods: {
    onSubmit() {
      this.$message('submit!')
      const data = {
        item_id: this.form.item_id,
        item_type: Number(this.form.item_type),
        amount: Number(this.form.amount),
        code: this.form.code,
        expire_at: this.form.expire_at,

        item_type2: Number(this.form.item_type2),
        item_id2: this.form.item_id2,
        item_number2: this.form.item_number2,

        item_type3: Number(this.form.item_type3),
        item_id3: this.form.item_id3,
        item_number3: this.form.item_number3,
      }
      createItem(data).then(response => {
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
    getItems() {
      getListByType(data).then(response => {
        this.options = response.data.items
      })
    },
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
      console.log('res', res.data)
      this.form.item_res = res.data.imgurl;
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
      return isLt2M;
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>
