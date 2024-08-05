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
      <el-table-column label="肥料等级" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
        </template>
      </el-table-column>
      <el-table-column label="肥料名称" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="肥料名称_英文" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.name_1 }}
        </template>
      </el-table-column>
      <el-table-column label="肥料名称_俄文" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.name_2 }}
        </template>
      </el-table-column>
      <el-table-column label="减少时长(m)" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.decr_time }}
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="商品图片" align="center">
        <template slot-scope="scope">
          <el-image style="width: 100px; height: 100px" :src="scope.row.res" :fit="fit"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" align="center">
        <template slot-scope="scope">
          <router-link :to="{ path: '/muck/edit', query: scope.row }">
            <el-button type="primary" size="small" icon="el-icon-edit">
              Edit
            </el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getList } from '@/api/muck'

export default {
  filters: {},
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getList().then(response => {
        this.list = response.data.items
        this.listLoading = false;
      })
    },
    getItemTypeText(type){
      const map = {
        1: '种子',
        2: '肥料',
      }
      return map[type]
    }
  }
}
</script>
