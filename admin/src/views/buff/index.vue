<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
      <el-table-column label="ID" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="名称" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="名称_英文" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.title_1 }}
        </template>
      </el-table-column>
      <el-table-column label="名称_俄文" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.title_2 }}
        </template>
      </el-table-column>
      <el-table-column label="说明" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.desc }}
        </template>
      </el-table-column>
      <el-table-column label="说明_英文" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.desc_1 }}
        </template>
      </el-table-column>
      <el-table-column label="说明_俄文" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.desc_2 }}
        </template>
      </el-table-column>
      <el-table-column label="增加产值" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.add_value }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="类型" width="150" align="center">
        <template slot-scope="scope" label="Status" >
          {{ getItemTypeText(scope.row.type) }}
        </template>
      </el-table-column>
      <el-table-column label="金币"  align="center">
        <template slot-scope="scope">
          {{ scope.row.gold_price }}
        </template>
      </el-table-column>
      <el-table-column label="钻石"  align="center">
        <template slot-scope="scope">
          {{ scope.row.diamond }}
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="商品图片" align="center">
        <template slot-scope="scope">
          <el-image style="width: 100px; height: 100px" :src="scope.row.img" :fit="fit"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" align="center">
        <template slot-scope="scope">
          <router-link :to="{ path: '/buff/edit', query: scope.row }">
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
import { getList } from '@/api/buff'

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
        console.log("list", this.list);

        this.listLoading = false
      })
    },
    getItemTypeText(type){
      const map = {
        1: '限时',
        2: '永久',
      }
      return map[type]
    },
  }
}
</script>
