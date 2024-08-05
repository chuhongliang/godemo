import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/admins/seed/list',
    method: 'post',
    data
  })
}

export function createItem(data){
  return request({
    url: '/admins/seed/create',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/seed/edit',
    method: 'post',
    data
  })
}