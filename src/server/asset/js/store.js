// Make a vuex store for all stuff
const store = new Vuex.Store({
    strict: true,
    state: { 
        count:0,
        models: [] 
    },
    mutations:{
        increment (state){
            state.count++
        }
    }
});
