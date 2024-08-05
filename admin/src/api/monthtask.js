import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/monthtask/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/monthtask/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/monthtask/edit',
    method: 'post',
    data
  })
}