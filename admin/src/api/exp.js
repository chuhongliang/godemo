import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/userexp/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/userexp/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/userexp/edit',
    method: 'post',
    data
  })
}