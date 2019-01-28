<template>
    <sui-segment basic>
        <sui-button @click="viz" :inverted="visible" basic compact icon="bars">List</sui-button>
        <sui-card-group v-show="visible" >
            <sui-card v-bind:model="modelList" :key="model.name" v-for="model in modelList">
                <sui-button class="nogap" basic v-on:click="load(model)">
                    <sui-image centered size="standard"  :src="model.img"></sui-image>
                </sui-button>
                <sui-card-content>
                    {{ model.name }}
                </sui-card-content>
            </sui-card>
        </sui-card-group>
    </sui-segment>
</template>

<script>
export default {
    data(){ return {
        visible: true 
    }
    },
    props: {
        current: "",
        modelList: {
            type: Array,
        },
    },
    methods: {
        viz: function(){
            this.visible = !this.visible;
        },
        load: function(obj){
            this.viz();
            EventBus.$emit('select',obj);
        },
    },
    computed: {
        imgref: function(){
            return "/pic/"+model+"/.png"
        }
    }
};
</script>

<style>
.nogap {
    padding-top: 0.0 !important;
    padding-right: 0.0 !important;
    padding-left: 0.0 !important;
    padding-bottom: 0.0 !important;
    border: 0px ;
}
</style>
