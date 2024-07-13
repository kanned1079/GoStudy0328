<script>
export default {
  name: 'MyList',
  data () {
    return {
      users: [],
    }
  },
  methods: {

  }
}
</script>

<template>
  <div class="row">
    <div class="card">
      <a href="https://github.com/xxxxxx" target="_blank">
        <img src="https://v2.cn.vuejs.org/images/logo.svg" alt="" style="width: 100px">
      </a>
      <p class="card-text">xxxxxx</p>
    </div>

    <div class="card">
      <a href="https://github.com/xxxxxx" target="_blank">
        <img src="https://v2.cn.vuejs.org/images/logo.svg" alt="" style="width: 100px">
      </a>
      <p class="card-text">xxxxxx</p>
    </div>

    <div class="card">
      <a href="https://github.com/xxxxxx" target="_blank">
        <img src="https://v2.cn.vuejs.org/images/logo.svg" alt="" style="width: 100px">
      </a>
      <p class="card-text">xxxxxx</p>
    </div>

    <div class="card">
      <a href="https://github.com/xxxxxx" target="_blank">
        <img src="https://v2.cn.vuejs.org/images/logo.svg" alt="" style="width: 100px">
      </a>
      <p class="card-text">xxxxxx</p>
    </div>

    <div class="card">
      <a href="https://github.com/xxxxxx" target="_blank">
        <img src="https://v2.cn.vuejs.org/images/logo.svg" alt="" style="width: 100px">
      </a>
      <p class="card-text">xxxxxx</p>
    </div>

    <div class="card">
      <a href="https://github.com/xxxxxx" target="_blank">
        <img src="https://v2.cn.vuejs.org/images/logo.svg" alt="" style="width: 100px">
      </a>
      <p class="card-text">xxxxxx</p>
    </div>

  </div>
</template>

<style scoped>
.album {
  min-height: 50rem; /* Can be removed; just added for demo purposes */
  padding-top: 3rem;
  padding-bottom: 3rem;
  background-color: #f7f7f7;
}

.card {
  float: left;
  width: 33.333%;
  padding: .75rem;
  margin-bottom: 2rem;
  border: 1px solid #efefef;
  text-align: center;
}

.card > img {
  margin-bottom: .75rem;
  border-radius: 100px;
}

.card-text {
  font-size: 85%;
}
</style>