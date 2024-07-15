import axios from "axios";
import {nanoid} from "nanoid";

// 人员管理相关配置
const personOptions = {
    namespaced: true,
    actions: {
        // 业务逻辑
        addPersonWang(context, value) {
            if (value.name.indexOf('王') === 0)
                context.commit('ADD_PERSON', value)
            else
                alert('必须姓王')
        },
        addPersonServer(context) {
            axios.get('http://localhost:9090/api').then(
                response => {
                    console.log(response.data)
                    context.commit('ADD_PERSON', {id: nanoid(), name: response.data})
                },
                error => {
                    console.warn(error.message)
                }
            )
        }
    },
    mutations: {
        ADD_PERSON(state, value) {
            state.personList.unshift(value);
        }
    },
    state: {
        personList: [
            {id: '001', name: '张三'}
        ]
    },
    getters: {
        // 获取列表中第一个人
        firstPersonName(state) {
            return state.personList[0].name;
        }
    },
}

export default personOptions