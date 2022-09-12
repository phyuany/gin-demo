<template>
  <div class="home">
    <div>
      <b-button pill>Button</b-button>
      <b-button pill variant="primary">Button</b-button>
      <b-button pill variant="outline-secondary">Button</b-button>
      <b-button pill variant="success">Button</b-button>
      <b-button pill variant="outline-danger">Button</b-button>
      <b-button pill variant="info">Button</b-button>
    </div>
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue';

export default {
  name: 'Home',
  components: {
    HelloWorld,
  },
};
</script>
