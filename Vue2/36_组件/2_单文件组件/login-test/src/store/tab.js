export default {
    namespaced: true,
    state: {
        isCollapse: false,  // 菜单折叠
        tabList: [  // 面包屑
            {
                path: '/',
                name: 'home',
                label: '首页',
                icon: 's-home',
                url: 'Home/Home'
            },
        ],
    },
    actions: {

    },
    mutations: {
        // 修改菜单展开收起
        collapseMenu(state) {
            state.isCollapse = ! state.isCollapse;
        },
        // 修改面包屑
        selectMenu(state, val) {
            console.log('selectMenu ', val)
            // 判断添加的数据是否是首页
            if (val.name !== 'home') {  // 不是首页 那就添加
                let index =  state.tabList.findIndex(item => item.name===val.name)
                if (index === -1)
                    state.tabList.push(val)
            }
            console.log('added ', state.tabList)
        },
        closeTag(state, item) {   // 用于删除指定的Tag数据
            console.log('关闭Tag')
            console.log(state)
            let index = state.tabList.findIndex(val => val.name === item.name)
            state.tabList.splice(index, 1);
        }
    }
}