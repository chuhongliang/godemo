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
      <el-table-column label="等级" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
        </template>
      </el-table-column>
      <el-table-column label="减少时间" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.decr_time }}
        </template>
      </el-table-column>
      <el-table-column label="增加产量" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.add_result }}
        </template>
      </el-table-column>
      <el-table-column label="升级消耗(金币)" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.consume }}
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="土地图片" align="center">
        <template slot-scope="scope">
          <el-image style="width: 100px; height: 100px" :src="scope.row.img" :fit="fit"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" align="center">
        <template slot-scope="scope">
          <router-link :to="{ path: '/land/edit', query: scope.row }">
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
import { getList } from '@/api/land'

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
        this.listLoading = false
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
