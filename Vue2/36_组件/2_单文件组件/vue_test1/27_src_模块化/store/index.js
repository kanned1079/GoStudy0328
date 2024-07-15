import Vue from 'vue'
// 该文件用于创建vuex中的核心store
// 引入vuex
import Vuex from 'vuex'

import countOptions from "@/store/count";
import personOptions from "@/store/perosn";

Vue.use(Vuex);
// 创建Store
// 创建并导出Store
export default new Vuex.Store({
    modules: {
        countAbout: countOptions,
        personAbout: personOptions,
    }
})

// // 暴露/导出store
// export default store;