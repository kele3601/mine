<template>
  <div class="login-container">
    <h2>Login</h2>
    <form @submit.prevent="handleLogin">
      <div class="form-group">
<!--        用户名-->
        <label for="name">Username:</label>
        <input type="text" id="name" v-model="name" required>
      </div>
      <div class="form-group">
<!--        密码-->
        <label for="pass">Pass:</label>
        <input type="password" id="pass" v-model="pass" required>
      </div>
<!--      按钮-->
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Login',
  data(){
    return{
      name:'',
      pass:''
    };
  },
  methods:{
    async handleLogin(){
      try{
        // 组装数据
        console.log('Login:',this.name,this.pass);
        const user = {
          name:this.name,
          pass:this.pass,
        }
        // 发送post请求
        const response = await axios.post('localhost:5219/public/user/login',user);
        console.log(response);
      }catch (error){
        console.error(error);
      }
    }
  }
}
</script>

<style scoped>
.login-container {
  background-color: gray;
  margin:auto;
  padding:20px;
  display: flex;
  justify-content: center;
  align-items: center;
  height:30vh;
}
</style>