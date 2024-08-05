<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
      <el-table-column label="ID" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="等级" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.level }}
        </template>
      </el-table-column>
      <el-table-column label="升级经验" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.exp }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="奖励类型" width="150" align="center">
        <template slot-scope="scope" label="Status" >
          {{ getItemTypeText(scope.row.reward.type) }}
        </template>
      </el-table-column>
      <el-table-column label="奖励ID"  align="center">
        <template slot-scope="scope">
          {{ scope.row.reward.id }}
        </template>
      </el-table-column>
      <el-table-column label="奖励数量"  align="center">
        <template slot-scope="scope">
          {{ scope.row.reward.amount }}
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="奖励图片" align="center">
        <template slot-scope="scope">
          <el-image style="width: 100px; height: 100px" :src="scope.row.reward.img" :fit="fit"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" align="center">
        <template slot-scope="scope">
          <router-link :to="{ path: '/monthtask/edit', query: scope.row }">
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
import { getList } from '@/api/monthtask'

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
        4: '狗粮',
        5: '金币',
        6: '钻石',
      }
      return map[type]
    },
  }
}
</script>
