export const mixin_showName =  {
    methods: {
        showName() {
            alert(this.name);
        }
    },
    mounted() {
        console.log('mounted in mixin.')
    },
}

export const hunhe2 = {
    data() {
        return {
            x: 100,
            y: 200,
        }
    }
}

