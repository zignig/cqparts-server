<template>
    <div>
        <sui-segment v-show="modelList.length" basic>
            <sui-menu vertical fluid >
                <sui-menu-item :active="isActive(obj)" v-bind:model="modelList" :key="obj.name" v-for="obj in modelList"
                    v-on:click="load(obj)">
                    {{ obj.name }}
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
        load: function(obj){
            clear();
            load(obj);
            console.log(obj);
            EventBus.$emit('select',obj);
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
