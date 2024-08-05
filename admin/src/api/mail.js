import request from '@/utils/request'

export function createItem(data) {
    return request({
        url: '/admins/system/mail',
        method: 'post',
        data
    })
}