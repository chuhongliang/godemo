import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/prop/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/prop/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/prop/edit',
    method: 'post',
    data
  })
}