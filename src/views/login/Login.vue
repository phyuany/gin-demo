<template>
  <div class="login">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="登录">
          <b-form>

            <b-form-group label="电话">
              <b-form-input
                id="input-1"
                v-model="$v.user.telephone.$model"
                type="number"
                required
                placeholder="输入手机号码"
                :state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                手机号不符合要求
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group label="密码">
              <b-form-input
                id="input-1"
                v-model="$v.user.password.$model"
                type="password"
                required
                placeholder="输入用户密码"
                :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码长度必须在6-18之间
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group>
              <b-button
                @click="login"
                variant="outline-primary"
                block=""
              >登录</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';

import customValidator from '@/helper/validator';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
        maxLength: maxLength(18),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userlogin: 'login' }),
    validateState(name) {
      // 这里是es6的 析构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        console.log('API拒绝请求');
        return;
      }

      // 请求api
      this.userlogin(this.user).then(() => {
        // 跳转首页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.data.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });

      console.log('login');
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
