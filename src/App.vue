<template>
  <div id="app">
    <navbar />
    <b-container>
      <router-view />
    </b-container>
  </div>
</template>

<script>
import Navbar from './views/layout/Navbar.vue';

export default {
  components: { Navbar },
  data() {
    return {};
  },
};
</script>

<style lang="scss">
</style>
