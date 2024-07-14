import Vue from 'vue'
// 该文件用于创建vuex中的核心store
// 引入vuex
import Vuex from 'vuex'

// 准备Actions 用于响应组件中的作用
const actions = {
    // jia(context, value) {
    //     console.log('Action中的jia被调用了');
    //     // console.log(context, value);
    //     context.commit('JIA', value)
    // },
    // jian(context, value) {
    //     context.commit('JIAN', value);
    // },
    jiaOdd(context, value) {
        if (context.state.sum % 2 !== 0)
            context.commit('JIA', value)
    },
    jiaWait(context, value) {
        setTimeout(() => {
            context.commit('JIA', value)
        }, 500)
    }
}
// 准备Mutations 用于操作数据
const mutations = {
    JIA(state, value) {
        console.log('Mutations中的JIA被调用了');
        state.sum += value;
    },
    JIAN(state, value) {
        state.sum -= value;
    },

}
// 准备State 用于存储数据
const state = {
    sum: 0, // 求和
    school: 'NNU',
    subject: '前端',

}
// 准备getters用于将state中的数值进行加工
const getters = {
    bigSum(state) {
        return state.sum * 10;
    }
}

Vue.use(Vuex);
// 创建Store
// 创建并导出Store
export default new Vuex.Store({
    actions,
    mutations,
    state,
    getters,
})

// // 暴露/导出store
// export default store;