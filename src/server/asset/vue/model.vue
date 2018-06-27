<template>
    <div class="m">
    <sui-menu vertical>
        <sui-menu-item v-bind:model="modelList" :key="obj" v-for="obj in modelList"
            v-on:click="load(obj)">
            {{ obj }}
            <sui-icon name="delete" v-on:click="remove(obj)"/>
        </sui-menu-item>
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
        remove : function(index){
            this.modelList.splice(this.modelList.indexOf(index),1);
            EventBus.$emit('delete item',index);
        }
    }
}
</script>

<style>
</style>
