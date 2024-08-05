<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column  label="ID" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="商品ID" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.item_id }}</span>
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="商品类型" width="150" align="center">
        <template slot-scope="scope" label="Status" >
          {{ getItemTypeText(scope.row.item_type) }}
        </template>
      </el-table-column>
      <el-table-column label="商品名称" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_title }}
        </template>
      </el-table-column>
      <el-table-column label="商品单价" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_price }}
        </template>
      </el-table-column>
      <el-table-column label="商品描述" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_desc }}
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="商品图片" align="center">
        <template slot-scope="scope">
          <el-image style="width: 100px; height: 100px" :src="scope.row.item_imgurl" :fit="fit"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <router-link :to="{ path: '/shop/edit', query: row }">
            <el-button type="primary" size="mini" icon="el-icon-edit">
              Edit
            </el-button>
          </router-link>
          <el-button  size="mini" type="danger" style="margin-left: 10px;" @click="handleDelete(row,$index)">
            Delete
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-row>
      <el-pagination background layout="prev, pager, next" :total="total_page" :page-size="page_size" :current-page="page" @current-change="handleCurrentChange"
        @prev-click="handlePrevClick" @next-click="handleNextClick">
      </el-pagination>
    </el-row>
  </div>
</template>

<script>
import { getList, deleteItem } from '@/api/shop'

export default {
  filters: {},
  data() {
    return {
      list: null,
      listLoading: true,
      page: 1,
      page_size: 10,
      total_page: 5,
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      const data = {
        page: this.page,
        page_size: this.page_size
      }
      getList(data).then(response => {
        this.list = response.data.items
        this.listLoading = false
        this.total_page = response.data.total_page
        console.log("items", response.data.items);

      })
    },
    getItemTypeText(type){
      const map = {
        1: '种子',
        2: '肥料',
        4: '狗粮',
      }
      return map[type]
    },
    handleDelete(row, index) {
      this.$notify({
        title: 'Success',
        message: 'Delete Successfully',
        type: 'success',
        duration: 2000
      })
      console.log('row', row)
      const data = {
        id: row.id,
      }
      
      deleteItem(data).then(response => {
        // this.list = response.data.items
        // this.listLoading = false
        console.log("item", response.data.item);
      })
      this.list.splice(index, 1)
    },

    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
      this.page = val
      this.fetchData()
    },
    handlePrevClick(val) {
      console.log('handlePrevClick=>', val);
      this.page = val
      this.fetchData()
    },
    handleNextClick(val) {
      console.log('handleNextClick=>', val);
      this.page = val
      this.fetchData()
    }
  }
}
</script>
