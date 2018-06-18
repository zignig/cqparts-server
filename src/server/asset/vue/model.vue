<template>
    <div>
    <sui-menu vertical>
        <a is="sui-menu-item" v-bind:model="modelList" v-for="obj in modelList" v-on:click="load(obj)">{{ obj }}</a>
    </sui-menu>
    </div>
</template>

<script>
export default {
    props: {
        modelList: {
            type: Array,
        },
    },
    methods:{
        load: function(event){
            clear();
            load(event);
            //console.log(event);
        },
    }
}
</script>

<style>
</style>
