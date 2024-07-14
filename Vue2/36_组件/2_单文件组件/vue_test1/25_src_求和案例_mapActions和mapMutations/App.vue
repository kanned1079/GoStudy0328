<template>
  <MyCount></MyCount>
</template>>

<script>
import MyCount from "@/components/Count.vue";

export default {
  name: 'App',
  components: {MyCount},
  data() {
    return {
    }
  },
  mounted() {
    console.log('App', this);
  }
}
</script>

<style scoped>
</style>
