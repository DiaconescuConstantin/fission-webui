<template>
  <div class="apinode">
    <h1>This is an APINode page</h1>
    <p v-if="error" style="color:#f00">Error: {{error}}</p>
    <p>Hostname: {{hostname}}</p>
    <p>Number: {{number}}</p>
    <p><button @click="getApiRand">Get Random Number</button></p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'apinode',
  data() {
    return {
      number: 0,
      hostname: '',
      error: null,
    };
  },
  methods: {
    getApiRand() {
      const self = this;

      self.error = null;
      self.hostname = '';
      self.number = 0;

      axios.get('/nodeapi/rand')
        .then((response) => {
          if(response && response.data) {
            self.number = response.data.random;
            self.hostname = response.data.hostname;
          }
        })
        .catch((error) => {
          self.error = error;
        });
    },
  },
};
</script>

<style scoped>

</style>
