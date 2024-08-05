import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/pet_kind/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/pet_kind/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/pet_kind/edit',
    method: 'post',
    data
  })
}