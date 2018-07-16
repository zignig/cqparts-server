<template>
    <sui-segment basic>
    <sui-card-group stackable>
        <sui-card v-bind:model="modelList" :key="model.name" v-for="model in modelList">
            <sui-card-content>
                <sui-card-header>{{ model.name }}</sui-card-header>
                <sui-card-description>
                </sui-card-description>
            </sui-card-content>
            <sui-image small v-show="model.img" :src="model.img"></sui-image>
            <sui-button-group>
                <sui-button >
                    <sui-icon name="truck"></sui-icon>
                    Add 
                </sui-button>
                <sui-button >
                    <sui-icon name="eye"></sui-icon>
                    View 
                </sui-button>
            </sui-button-group>
        </sui-card>
    </sui-card-group>
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
    computed: {
        imgref: function(){
            return "/pic/"+model+"/.png" 
        }
    }
};
</script>

<style>
</style>
