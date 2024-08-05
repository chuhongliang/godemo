import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/buff/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/buff/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/buff/edit',
    method: 'post',
    data
  })
}