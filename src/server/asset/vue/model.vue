<template>
    <div>
    <sui-segment v-show="modelList.length" basic>
    <sui-menu vertical fluid >
        <sui-menu-item :active="isActive(obj)" v-bind:model="modelList" :key="obj" v-for="obj in modelList"
            v-on:click="load(obj)">
            {{ obj }}
            <sui-icon name="delete" v-on:click="remove(obj)"/>
        </sui-menu-item>
    </sui-menu>
    </sui-segment>
    </div>
</template>

<script>
export default {
    props: {
        current: "",
        modelList: {
            type: Array,
        },
    },
    methods:{
        isActive: function(name){
            return this.current ===  name
        },
        load: function(event){
            clear();
            load(event);
            //console.log(event);
            EventBus.$emit('active',event);
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
