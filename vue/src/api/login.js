import request from '@/utils/request'

export function login(data) {
    return request({
        url: '/gin/user/login',
        method: 'post',
        data
    });
}