<template>
    <div :xl="6" :lg="7" class="bg-login">
        <!-- <el-mage :src="require('@/assets/logo.png')" class="logo"/> -->
        <el-row type="flex" class="row-gb row-two" justify="center" align="middle">
            <el-col :span="6"></el-col>
            <el-col :span="6">
                <h1 class="title">PASSWORD MANAGER</h1>
            </el-col>
            <el-col :span="6"></el-col>
        </el-row>
        <el-row type="flex" class="row-gb card" justify="center" align="bottom">
            <el-col :span="7" class="login-card">
                <el-form :model="loginForm" :rules="rules" ref="loginForm" label-width="25%" class="loginForm">
                    <el-form-item label="ACCOUNT" prop="username" style="width: 380px">
                        <el-input v-model="loginForm.username"></el-input>
                    </el-form-item>
                    <el-form-item label="PASSWORD" prop="password" style="width: 380px">
                        <el-input type="password" v-model="loginForm.password"></el-input>
                    </el-form-item>
                    <el-form-item class="btn-ground">
                        <el-button type="primary" @click="submitForm('loginForm')">Login</el-button>
                        <el-button @click="resetForm('loginForm')">Reset</el-button>
                        <el-button @click="dialog = true">Register</el-button>

                        <el-drawer 
                            title="USER REGISTER"
                            :before-close="handleClose" 
                            :visible.sync="dialog"
                            :modal-append-to-body="false"
                            direction="rtl" 
                            customer-class="register-drawer" 
                            ref="drawer"
                            >
                            <div class="register-drawer_content">
                                <el-form :model="registerForm" :rules="rules" ref="registerForm" label-width="25%" class="registerForm">
                                    <el-form-item label="ACCOUNT"  prop="username" style="width: 380px">
                                        <el-input v-model="registerForm.username" autocomplete="off"></el-input>
                                    </el-form-item>
                                    <el-form-item label="PASSWORD" prop="password" style="width: 380px">
                                        <el-input v-model="registerForm.password" autocomplete="off"></el-input>
                                    </el-form-item>
                                    <el-form-item label="PHONE" prop="phone" style="width: 380px">
                                        <el-input v-model="registerForm.phone" autocomplete="off"></el-input>
                                    </el-form-item>
                                </el-form>
                                <div class="register-drawer_foolter">
                                    <el-button @click="cancelForm">Cancel</el-button>
                                    <el-button type="primary" @click="$ref.drawer.closeDrawer()" :loading="loading">{{loading ? 'submiting...' : 'confirm'}}</el-button>
                                </div>
                            </div>
                        </el-drawer>
                    </el-form-item>
                    <el-form-item label="net login" prop="netLogin">
                        <el-switch v-model="loginForm.netLogin"></el-switch>
                    </el-form-item>
                </el-form>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import ElementUI from 'element-ui';
import router from "@/router"
import {login, netLogin, register} from '@/api/login'

export default {
    name: "Login",
    data() {
        return {
            dialog: false,
            loading: false,
            timer: null,
            registerForm: {
                username: '',
                password: '',
                phone: ''
            },
            loginForm: {
                username: '',
                password: '',
                netLogin: false
            },
            rules: {
                username: [
                    {required: true, message: 'Please enter your account', trigger: "blur"},
                    {min: 3, max: 10, message: 'The length is between 3 and 10', trigger: "blur"}
                ],
                password: [
                    {required: true, message: 'Please enter your password', trigger: "blur"},
                    {min: 1, max: 100, message: 'The length is between 1 and 100', trigger: "blur"}
                ],
                phone:[
                    {required: true, message: 'Please enter your phone', trigger: "blur"}
                ]
            }
        };
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    console.log("form", this.loginForm)
                    if (!this.loginForm.netLogin) {
                        login({username:this.loginForm.username, password:this.loginForm.password}).then(res => {
                            let success = JSON.parse(res.data.success);
                            let message = res.data.msg;
                            if (success) {
                                
                                // 登录成功
                                ElementUI.Message.success(message);
                                // 跳转页面
                                router.push('/hello')
                            } else {
                                // 打印错误信息
                                ElementUI.Message.error(message);
                            }
                        })
                    } else {
                        netLogin({username:this.loginForm.username, password:this.loginForm.password}).then(res => {
                            let success = JSON.parse(res.data.success);
                            let message = res.data.msg;
                            if (success) {
                                
                                // 登录成功
                                ElementUI.Message.success(message);
                                // 跳转页面
                                router.push('/hello')
                            } else {
                                // 打印错误信息
                                ElementUI.Message.error(message);
                            }
                        })
                    }
                } else {
                    console.log('error submit');
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        handleClose(done) {
            if (this.loading) {
                return;
            }
            this.$confirm('submit it ?').then(_ => {
                this.loading = true;
                this.timer = setTimeout(() => {
                    done();
                    setTimeout(() => {
                        this.loading = false;
                    }, 400);
                }, 2000);
            })
            .catch(_ => {})
        }
    },
    cancelForm() {
        this.loading = false;
        this.dialog = false;
        clearTimeout(this.timer);
    }
}
</script>

<style scoped>
.bg-login {
    height: 100%;
    /* background-image: url("../assets/logo.png"); */
    background-size: 200%;
}

.btn-ground {
    text-align: center;
}

.logo {
    margin: 30px;
    height: 70px;
    width: 70px;
    position: fixed;
}

.title {
    text-shadow: -3px 3px 1px #5f565e;
    text-align: center;
    margin-top: 60%;
    color: #41b9a6;
    font-size: 40px;
}

.login-card {
    background-color: #ffffff;
    opacity: 0.9;
    box-shadow: 0 0 20px #ffffff;
    border-radius: 10px;
    padding: 40px 40px 30px 15px;
    width: auto;
}
</style>