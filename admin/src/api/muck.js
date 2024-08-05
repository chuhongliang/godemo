import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/muck/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/muck/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/muck/edit',
    method: 'post',
    data
  })
}