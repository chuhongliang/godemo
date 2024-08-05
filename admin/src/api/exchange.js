import request from '@/utils/request'



export function getList(params) {
  return request({
    url: '/exchanges/list',
    method: 'get',
    params
  })
}


export function createItem(data) {
  return request({
    url: '/exchanges/create',
    method: 'post',
    data
  })
}


export function deleteItem(data) {
  return request({
    url: '/exchanges/delete',
    method: 'post',
    data
  })
}