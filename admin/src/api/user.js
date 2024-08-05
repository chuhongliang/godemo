import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/admins/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/admins/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/admins/logout',
    method: 'post'
  })
}
