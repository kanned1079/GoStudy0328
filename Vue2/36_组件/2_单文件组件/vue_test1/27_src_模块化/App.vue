<template>
<div>
  <MyCount></MyCount>
  <hr>
  <MyPerson></MyPerson>
</div>
</template>

<script>
import MyCount from "@/components/Count.vue";
import MyPerson  from "@/components/Person.vue";

export default {
  name: 'App',
  components: {MyCount, MyPerson},
  // components: {MyCount},
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
