export default {
    state: {
        isCollapse: false,
    },
    actions: {

    },
    mutations: {
        // 修改菜单展开收起
        collapseMenu(state) {
            state.isCollapse = ! state.isCollapse;
        }
    }
}