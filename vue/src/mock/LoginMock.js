const Mock = require('mockjs');
const Random = Mock.Random;

let result = {
    msg: '',
    data: ''
}

let username = 'admin';
let password = '123456';

Mock.mock('/login', 'post', (req) => {
    let form = JSON.parse(req.body);
    if (username == form.username) {
        if (password == form.password) {
            result.msg = 'login success'
            result.data = 'true'
        } else {
            result.msg = 'password is wrong'
            result.data = 'false'
        }
    } else {
        result.msg = 'user is not exist'
        result.data = 'false'
    }

    return result;
})