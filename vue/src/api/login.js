import request from '@/utils/request'

export function login(data) {
    return request({
        url: '/gin/user/login',
        method: 'post',
        data
    });
}

export function netLogin(data) {
    return request({
        url: '/gin/user/netLogin',
        method: 'post',
        data
    });
}

export function register(data) {
    return request({
        url: '/gin/user/register',
        method: 'post',
        data
    });
}