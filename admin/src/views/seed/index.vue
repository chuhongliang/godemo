<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
      <el-table-column label="ID" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="种子等级" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
        </template>
      </el-table-column>
      <el-table-column label="种子名称" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="成长时长(h)" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.grow_time }}
        </template>
      </el-table-column>
      <el-table-column label="收获植物ID" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.plant_id }}
        </template>
      </el-table-column>
      <el-table-column label="收获植物数量" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.plant_number }}
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="成熟植物图片" align="center">
        <template slot-scope="scope">
          <el-image style="width: 100px; height: 100px" :src="scope.row.imgurl" :fit="fit"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" align="center">
        <template slot-scope="scope">
          <router-link :to="{ path: '/seed/edit', query: scope.row }">
            <el-button type="primary" size="small" icon="el-icon-edit">
              Edit
            </el-button>
          </router-link>
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
import { getList } from '@/api/seed'

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
        console.log("this.total_page", this.total_page)
      })
    },
    getItemTypeText(type) {
      const map = {
        1: '种子',
        2: '肥料',
      }
      return map[type]
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
