import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admins/landlevel/list',
    method: 'get',
    params
  })
}

export function createItem(data){
  return request({
    url: '/admins/landlevel/item',
    method: 'post',
    data
  })
}

export function editItem(data){
  return request({
    url: '/admins/landlevel/edit',
    method: 'post',
    data
  })
}