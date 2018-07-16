<template>
    <sui-segment v-show="modelList.length">
        <sui-input icon="search" placeholder="Search..."></sui-input>
        <sui-item-group divided >
            <sui-item v-bind:model="modelList" :key="model.name" v-for="model in modelList">
                <sui-button compact basic v-on:click="load(model)">
                    <sui-item-image wrapped v-on:click="load(model)" size="tiny" v-show="model.img" :src="model.img">
                    </sui-item-image>
                </sui-button>
                <sui-item-content>
                    <sui-header :color="isActive(model)">
                        {{ model.name }}
                        <span><sui-icon size="small" name="delete" v-on:click="remove(model)"/></sui-icon></span>
                    </sui-header>
                    <sui-item-extra>
                        <sui-button-group>
                            <sui-button size="mini" icon="newspaper" v-on:click=""></sui-button>
                            <sui-button size="mini" icon="camera" v-on:click=""></sui-button>
                            <sui-button size="mini" icon="thumbtack" v-on:click=""></sui-button>
                        </sui-button-group>
                    </sui-item-extra>
                </sui-item-content>
            </sui-item>
        </sui-item-group>
        <sui-button-group>
            <sui-button>Prev</sui-button>
            <sui-button>Next</sui-button>
        </sui-button-group>
    </sui-segment>
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
        isActive: function(obj){
            if (this.current == obj.name){
                return "green"
            }
            else {
                return "black"
            }
        },
        load: function(obj){
            clear();
            console.log(obj);
            load(obj.name);
            EventBus.$emit('select',obj);
        },
        remove : function(index){
            this.modelList.splice(this.modelList.indexOf(index),1);
            EventBus.$emit('delete item',index);
        }
    }
};

</script>

<style>
</style>
