<template>
    <div class="pure-menu">
    <ul class="pure-menu-list">
        <li class="pure-menu-item" v-bind:model="modelList" v-for="obj in modelList">
            <a class="pure-menu-link" v-on:click="load">{{ obj }}</a>
        </li>
    </ul>
    </div>
</template>

<script>
export default { 
    props: {
        modelList: [], 
    },
    methods:{
        load: function(event){
            console.log(this);
        },
        add: function(item){
            this.models.push(item);
        }
    }
}
</script>

<style>
</style>
