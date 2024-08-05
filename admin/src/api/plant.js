import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/plant/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/plant/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/plant/edit',
    method: 'post',
    data
  })
}