<template>
    <div class="pure-menu">
    <ul class="pure-menu-list">
        <li class="pure-menu-item" v-bind:model="modelList" v-for="obj in modelList">
            <a class="pure-menu-link" v-on:click="load(obj)">{{ obj }}</a>
        </li>
    </ul>
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
