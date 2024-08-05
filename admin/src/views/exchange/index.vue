<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
      <el-table-column label="ID" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column  label="兑换码" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.code }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="奖励1类型" width="150" align="center">
        <template slot-scope="scope" label="Status">
          {{ getItemTypeText(scope.row.item_type) }}
        </template>
      </el-table-column>
      <el-table-column label="奖励1ID">
        <template slot-scope="scope">
          {{ scope.row.item_id }}
        </template>
      </el-table-column>
      <el-table-column  label="奖励1数量" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_number }}
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="奖励2类型" width="150" align="center">
        <template slot-scope="scope" label="Status">
          {{ getItemTypeText(scope.row.item_type_2) }}
        </template>
      </el-table-column>
      <el-table-column label="奖励2ID">
        <template slot-scope="scope">
          {{ scope.row.item_id_2 }}
        </template>
      </el-table-column>
      <el-table-column  label="奖励2数量" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_number_2 }}
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="奖励2类型" width="150" align="center">
        <template slot-scope="scope" label="Status">
          {{ getItemTypeText(scope.row.item_type_3) }}
        </template>
      </el-table-column>
      <el-table-column label="奖励3ID">
        <template slot-scope="scope">
          {{ scope.row.item_id_3 }}
        </template>
      </el-table-column>
      <el-table-column  label="奖励3数量" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_number_3 }}
        </template>
      </el-table-column>

      <el-table-column  label="有效期" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.expire_at }}
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button size="mini" type="danger" style="margin-left: 10px;" @click="handleDelete(row, $index)">
            Delete
          </el-button>
        </template>
      </el-table-column>

    </el-table>
  </div>
</template>

<script>
import { getList, deleteItem } from '@/api/exchange'

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
        console.log("items", response.data.items);
      })
    },
    getItemTypeText(type) {
      const map = {
        1: '种子',
        2: '肥料',
        5: '金币',
        6: '钻石',
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
  }
}
</script>
