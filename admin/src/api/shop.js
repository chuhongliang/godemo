import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/admins/shop/list',
    method: 'post',
    data
  })
}

export function createItem(data){
  return request({
    url: '/admins/shop/create',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/shop/edit',
    method: 'post',
    data
  })
}

export function deleteItem(data){
  return request({
    url: '/admins/shop/delete',
    method: 'post',
    data
  })
}

