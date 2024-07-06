export default {
    install(Vue, x, y, z) {
        console.log('@@install');
        console.log(Vue);
        console.log(x, y, z);

        // 全局过滤器
        Vue.filter('mySlice', function (value) {
            return value.slice(0, 4);
        });

        // 全局自定义指令 声明全局指令
        Vue.directive('fbind', {
            fbind: {
                // 常用函数
                // 成功绑定时
                bind(element, binding) {
                    console.log('bind');
                    element.value = binding.value;
                },
                // 元素被插入页面时
                inserted(element, binding) {
                    element.focus();
                },
                // 指令所在的模版被重新解析时
                update(element, binding) {
                    element.value = binding.value;
                }
            },
        });

        // 定义混入mixin
        Vue.mixin({
            data() {
                return {
                    x: 100,
                    y: 200,
                }
            }
        });

        // 给原型对象上添加一个方法
        Vue.prototype.$hello = () => {
            alert('Hello.');
        }


    }
}

