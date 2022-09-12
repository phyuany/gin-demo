<template>
  <div>
    <b-navbar
      toggleable="lg"
      type="dark"
      variant="info"
    >
      <b-container>
        <b-navbar-brand @click="$router.push({name: 'Home'})">Jkdev.cn</b-navbar-brand>

        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse
          id="nav-collapse"
          is-nav
        >
          <b-navbar-nav class="ml-auto">
            <b-nav-item-dropdown
             right
              v-if="userInfo"
            >
              <!-- Using 'button-content' slot -->
              <template v-slot:button-content>
                <em>{{userInfo.name}}</em>
              </template>
              <b-dropdown-item @click="$router.push({name:'profile'})">个人主页</b-dropdown-item>
              <b-dropdown-item @click="logout">退出</b-dropdown-item>
            </b-nav-item-dropdown>
            <div v-if="!userInfo">
              <b-nav-item
               v-if="$route.name != 'login'"
                @click="$router.replace({name: 'login'})"
              >登录</b-nav-item>
              <b-nav-item
               v-if="$route.name != 'register'"
                @click="$router.replace({name: 'register'})"
              >注册</b-nav-item>
            </div>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),

  methods: mapActions('userModule', ['logout']),
};
</script>

<style scoped>
</style>
